package evalfilter

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/Fishwaldo/mouthpiece/pkg/filter"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/go-logr/logr"
	"github.com/skx/evalfilter/v2"
	"github.com/skx/evalfilter/v2/object"
)

//go:embed scripts
var embededScripts embed.FS

func init() {
	filter.RegisterFilterImpl("EvalFilter", EvalFilterFactory{})
}

type EvalFilter struct {
	ready            bool
	filter           *evalfilter.Eval
	processedMessage interfaces.MessageI
	config           *EvalFilterConfig
	ctx              context.Context
	filtertype       FilterType
	filtername       string
	log              logr.Logger
}

type EvalFilterConfig struct {
	Script string
}

func (c *EvalFilterConfig) AsJSON() (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *EvalFilterConfig) FromJSON(data string) error {
	return json.Unmarshal([]byte(data), c)
}

type FilterType int

const (
	FilterTypeInvalid FilterType = iota
	FilterTypeFile
	FilterTypeEmbeded
	FilterTypeStatic
)

type EvalFilterFactory struct {
	filtertype   FilterType
	filtersource string
}

func (eff EvalFilterFactory) FilterFactory(ctx context.Context, config string) (interfaces.FilterImplI, error) {
	var cfg EvalFilterConfig
	if err := cfg.FromJSON(config); err != nil {
		return nil, mperror.ErrFilterConfigInvalid
	}
	eflt := &EvalFilter{
		config:     &cfg,
		filtertype: eff.filtertype,
		log: log.Log.WithName("EvalFilter"),
	}

	var filterscript string

	switch eff.filtertype {
	case FilterTypeFile:
		if filecontext, err := os.ReadFile(eff.filtersource); err != nil {
			eflt.log.Error(err, "Filter Script File Read Failed", "script", eff.filtersource)
			return nil, err
		} else {
			filterscript = string(filecontext)
			eflt.filtername = trimFileExtension(filepath.Base(eff.filtersource))
		}
	case FilterTypeEmbeded:
		if filecontent, err := fs.ReadFile(embededScripts, eff.filtersource); err != nil {
			eflt.log.Error(err, "Filter Script Embeded Read Failed", "script", eff.filtersource)
			return nil, err
		} else {
			filterscript = string(filecontent)
			eflt.filtername = trimFileExtension(filepath.Base(eff.filtersource))
		}
	case FilterTypeStatic:
		filterscript = eflt.config.Script
		eflt.filtername = "Static"
	default:
		return nil, mperror.ErrFilterConfigInvalid
	}
	eflt.log = log.Log.WithName(eflt.filtername)

	eflt.filter = evalfilter.New(filterscript)
	eflt.filter.AddFunction("printf", eflt.fnPrintf)
	eflt.filter.AddFunction("print", eflt.fnPrint)
	eflt.filter.AddFunction("setfield", eflt.fnSetField)
	eflt.filter.AddFunction("clearfield", eflt.fnClearField)
	eflt.filter.AddFunction("setshortmessage", eflt.fnSetShortMessage)
	eflt.filter.AddFunction("setseverity", eflt.fnSetSeverity)
	if err := eflt.filter.Prepare(); err != nil {
		eflt.log.Error(err, "Compile Filter Script Failed", "filter", eflt.FilterName())
		eflt.ready = false
		return nil, err
	}
	eflt.log.V(1).Info("Compile Filter Script Success", "filter", eflt.FilterName())
	eflt.ready = true
	return eflt, nil

}

func (sff EvalFilterFactory) DefaultConfig(ctx context.Context) interfaces.MarshableConfigI {
	return &EvalFilterConfig{}
}

func (sf *EvalFilter) Init(ctx context.Context) error {
	return nil
}

func (ef *EvalFilter) Process(ctx context.Context, msg interfaces.MessageI) (interfaces.FilterAction, error) {
	ef.ctx = ctx
	ef.processedMessage = msg

	if !ef.ready {
		if err := ef.Init(ctx); err != nil {
			ef.log.Error(err, "Filter Init Failed", "filter", ef.FilterName())
			return interfaces.FilterPass, err
		}
	}
	defer func() {
		if err := recover(); err != nil {
			ef.log.Error(err.(error), "Filter Script Error", "filter", ef.FilterName())
		}
	}()
	ef.filter.SetContext(ctx)
	ok, err := ef.filter.Run(getFlatMessage(ctx, msg))
	if err != nil {
		ef.log.Info("Filter Run Failed", "filter", ef.FilterName(), "result", ok, "Error", err)
		return interfaces.FilterPass, err
	}
	ef.log.V(1).Info("Filter Run Success", "filter", ef.FilterName(), "result", ok)

	ef.ctx = nil
	ef.processedMessage = nil

	if ok {
		return interfaces.FilterMatch, nil
	} else {
		return interfaces.FilterNoMatch, nil
	}
}

func (sf *EvalFilter) FilterName() string {
	return sf.filtername
}

func (sf *EvalFilter) SetConfig(ctx context.Context, config interfaces.MarshableConfigI) error {
	var ok bool
	if sf.config, ok = config.(*EvalFilterConfig); !ok {
		return mperror.ErrFilterConfigInvalid
	}
	return nil
}
func (sf *EvalFilter) GetConfig(ctx context.Context) (interfaces.MarshableConfigI, error) {
	return sf.config, nil
}

func (ev *EvalFilter) fnPrintf(args []object.Object) object.Object {
	// We expect 1+ arguments
	if len(args) < 1 {
		return &object.Null{}
	}

	// Type-check
	if args[0].Type() != object.STRING {
		return &object.Null{}
	}

	// Get the format-string.
	fs := args[0].(*object.String).Value

	// Convert the arguments to something go's sprintf
	// code will understand.
	argLen := len(args)
	fmtArgs := make([]interface{}, argLen-1)

	// Here we convert and assign.
	for i, v := range args[1:] {
		fmtArgs[i] = v.ToInterface()
	}

	// Call the helper
	out := fmt.Sprintf(fs, fmtArgs...)

	ev.log.Info("Filter Script Output", "filter", ev.FilterName(), "output", out)
	return &object.Void{}
}

func (ev *EvalFilter) fnPrint(args []object.Object) object.Object {
	for _, e := range args {
		ev.log.Info("Filter Script Output", "filter", ev.FilterName(), "Output", e.Inspect())
	}
	return &object.Void{}
}

func (ev *EvalFilter) fnSetField(args []object.Object) object.Object {
	// We expect 2 arguments
	if len(args) != 2 {
		return &object.Null{}
	}

	// Type-check
	if args[0].Type() != object.STRING {
		return &object.Null{}
	}

	// Get the field name.
	fld := args[0].(*object.String).Value

	arg := args[0].ToInterface()

	if val, ok := arg.(string); !ok {
		ev.log.Error(nil, "Cannot Convert Argument to String", "filter", ev.FilterName(), "field", fld, "value", val)
		return &object.Null{}
	}

	ev.log.Info("Setting Field Value", "filter", ev.FilterName(), "field", fld, "value", arg)
	if err := ev.processedMessage.SetMetadata(ev.ctx, fld, arg); err != nil {
		ev.log.Error(err, "Set Field Failed", "filter", ev.FilterName(), "field", fld, "value", arg)
		return &object.Null{}
	}
	return &object.Void{}
}

func (ev *EvalFilter) fnClearField(args []object.Object) object.Object {
	if len(args) != 1 {
		return &object.Null{}
	}
	// Type-check
	if args[0].Type() != object.STRING {
		return &object.Null{}
	}
	fld := args[0].(*object.String).Value
	ev.log.Info("Clearing Field Value", "filter", ev.FilterName(), "field", fld)
	if err := ev.processedMessage.SetMetadata(ev.ctx, fld, nil); err != nil {
		ev.log.Info("Clear Field Failed", "filter", ev.FilterName(), "field", fld)
	}
	return &object.Void{}
}

func (ev *EvalFilter) fnSetShortMessage(arg []object.Object) object.Object {
	if len(arg) != 1 {
		return &object.Null{}
	}
	// Type-check
	if arg[0].Type() != object.STRING {
		return &object.Null{}
	}
	msg := arg[0].(*object.String).Value
	ev.log.Info("Setting Short Message", "filter", ev.FilterName(), "Short Message", msg)
	ev.processedMessage.SetShortMsg(ev.ctx, msg)
	return &object.Void{}
}

func (ev *EvalFilter) fnSetSeverity(arg []object.Object) object.Object {
	if len(arg) != 1 {
		return &object.Null{}
	}
	// Type-check
	if arg[0].Type() != object.INTEGER {
		return &object.Null{}
	}
	msg := arg[0].(*object.Integer).Value
	ev.log.Info("Setting Severity", "filter", ev.FilterName(), "Severity", msg)
	ev.processedMessage.SetSeverity(ev.ctx, int(msg))
	return &object.Void{}
}

type FlatMessage struct {
	Message string
	Severity int
	//ShortMsg *string
	//Topic *string
	TimeStamp time.Time
}

func getFlatMessage(ctx context.Context, mymsg interfaces.MessageI) FlatMessage {
	var fm FlatMessage
	fm.Message = mymsg.GetMessage()
	fm.Severity = mymsg.GetSeverity()
	//fm.ShortMsg = *mymsg.GetShortMsg()
	//fm.Topic = *mymsg.GetTopic()
	fm.TimeStamp = mymsg.GetTimestamp()
	return fm
}