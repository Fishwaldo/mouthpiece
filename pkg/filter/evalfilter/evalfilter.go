package evalfilter

import (
	"context"
	"fmt"

	"github.com/Fishwaldo/mouthpiece/pkg/filter"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/msg"
	"github.com/go-logr/logr"
	"github.com/skx/evalfilter/v2"
	"github.com/skx/evalfilter/v2/object"

	"golang.org/x/exp/slices"
)

var llog logr.Logger

type EvalFilter struct {
	ready            bool
	filter           *evalfilter.Eval
	processedMessage *msg.Message
	config           []filter.Filterconfig
}

var (
	FilterImplDetails = filter.FilterImplDetails{Factory: NewEvalFilter}
)

func init() {
	fmt.Println("Registering EvalFilter")
	filter.RegisterFilterImpl("EvalFilter", FilterImplDetails)
}

func NewEvalFilter(ctx context.Context, config []filter.Filterconfig) (filter.FilterImplI, error) {
	llog = log.Log.WithName("EvalFilter")

	if idx := slices.IndexFunc(config, func(config filter.Filterconfig) bool { return config.Name == "name" }); idx == -1 {
		return &EvalFilter{}, fmt.Errorf("EvalFilter: No name specified")
	}

	if idx := slices.IndexFunc(config, func(config filter.Filterconfig) bool { return config.Name == "script" }); idx == -1 {
		return &EvalFilter{}, fmt.Errorf("EvalFilter: No Script specified")
	}

	ef := &EvalFilter{
		config: config,
	}
	return ef, nil
}

func (ef *EvalFilter) FilterName() string {
	return "EvalFilter"
}

func (ev *EvalFilter) Init() error {
	idx := slices.IndexFunc(ev.config, func(config filter.Filterconfig) bool { return config.Name == "script" })

	ev.filter = evalfilter.New(ev.config[idx].Value)
	ev.filter.AddFunction("printf", ev.fnPrintf)
	ev.filter.AddFunction("print", ev.fnPrint)
	ev.filter.AddFunction("setfield", ev.fnSetField)
	ev.filter.AddFunction("clearfield", ev.fnClearField)
	ev.filter.AddFunction("setshortmessage", ev.fnSetShortMessage)
	ev.filter.AddFunction("setseverity", ev.fnSetSeverity)
	if err := ev.filter.Prepare(); err != nil {
		llog.Error(err, "Compile Filter Script Failed", "filter", ev.getName())
		ev.ready = false
		return err
	}
	llog.V(1).Info("Compile Filter Script Success", "filter", ev.getName())
	ev.ready = true
	return nil
}

func (ev *EvalFilter) getName() string {
	idx := slices.IndexFunc(ev.config, func(config filter.Filterconfig) bool { return config.Name == "name" })
	return ev.config[idx].Value
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

	llog.Info("Filter Script Output", "filter", ev.getName(), "output", out)
	return &object.Void{}
}

func (ev *EvalFilter) fnPrint(args []object.Object) object.Object {
	for _, e := range args {
		llog.Info("Filter Script Output", "filter", ev.getName(), "Output", e.Inspect())
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

	llog.Info("Setting Field Value", "filter", ev.getName(), "field", fld, "value", arg)
	ev.processedMessage.Body.Fields[fld] = arg
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
	llog.Info("Clearing Field Value", "filter", ev.getName(), "field", fld)
	if _, ok := ev.processedMessage.Body.Fields[fld]; ok {
		delete(ev.processedMessage.Body.Fields, fld)
	} else {
		llog.Info("Field Not Found", "filter", ev.getName(), "field", fld)
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
	llog.Info("Setting Short Message", "filter", ev.getName(), "message", msg)
	ev.processedMessage.Body.ShortMsg = msg
	return &object.Void{}
}

func (ev *EvalFilter) fnSetSeverity(arg []object.Object) object.Object {
	if len(arg) != 1 {
		return &object.Null{}
	}
	// Type-check
	if arg[0].Type() != object.STRING {
		return &object.Null{}
	}
	msg := arg[0].(*object.String).Value
	llog.Info("Setting Severity", "filter", ev.getName(), "Severity", msg)
	ev.processedMessage.Body.Severity = msg
	return &object.Void{}
}

func (ev *EvalFilter) Process(ctx context.Context, msg *msg.Message) (bool, error) {
	if !ev.ready {
		if err := ev.Init(); err != nil {
			llog.Error(err, "Filter Init Failed", "filter", ev.getName())
			return true, err
		}
	}
	defer func() {
		if err := recover(); err != nil {
			llog.Error(err.(error), "Filter Script Error", "filter", ev.getName())
		}
	}()
	ev.processedMessage = msg
	ev.filter.SetContext(ctx)
	ok, err := ev.filter.Run(msg.Body)
	ev.processedMessage = nil
	if err != nil {
		llog.Info("Filter Run Failed", "filter", ev.getName(), "result", ok, "Error", err)
		return true, err
	}
	llog.V(1).Info("Filter Run Success", "filter", ev.getName(), "result", ok)
	return ok, nil
}
