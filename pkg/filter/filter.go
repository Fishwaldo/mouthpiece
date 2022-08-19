package filter

import (
	"context"
	"io/fs"
	"path/filepath"
	"strings"

	//"io/ioutil"
	"embed"
	"fmt"

	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/message"
	"github.com/go-logr/logr"
	"github.com/skx/evalfilter/v2"
	"github.com/skx/evalfilter/v2/object"
	"gorm.io/gorm"
)

//go:embed scripts
var ScriptFiles embed.FS

type FilterType int

var llog logr.Logger

const (
	AppFilter = iota
	UserFilter
	TransportFilter
)

func (ft FilterType) String() string {
	return [...]string{"AppFilter", "UserFilter", "TransportFilter"}[ft]
}

type Filter struct {
	gorm.Model       `json:"-"`
	Name             string
	Content          string
	Type             FilterType
	Enabled          bool
	script           *evalfilter.Eval `gorm:"-" json:"-"`
	ok               bool             `gorm:"-"`
	processedMessage *msg.Message     `gorm:"-" json:"-"`
}

var Filters []*Filter

func filterFiles(root, ext string) []string {
	var a []string
	fs.WalkDir(ScriptFiles, root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}

		return nil
	})
	return a
}

func trimFileExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func loadScriptFiles(files []string, scripttype FilterType) {
	for _, script := range files {
		/* first see if this script exists in the Database first */
		var flt *Filter
		tx := db.Db.Where("name = ? and type = ?", trimFileExtension(filepath.Base(script)), scripttype).First(&flt)
		if tx.RowsAffected == 0 {
			llog.V(1).Info("Reading Filter Script from Filesystem", "type", scripttype, "filter", trimFileExtension(filepath.Base(script)))
			content, err := fs.ReadFile(ScriptFiles, script)
			if err != nil {
				llog.Error(err, "Failed to read Filter Script File", "filename", script)
				continue
			}
			//
			// Create an evalfilter, with the script inside it.
			//
			flt = &Filter{
				Name:    trimFileExtension(filepath.Base(script)),
				Content: string(content),
				Type:    scripttype,
			}
		} else {
			llog.V(1).Info("Loading Filter Script from Databse", "type", scripttype, "filter", flt.Name)
		}
		if err := flt.SetupEvalFilter(); err == nil {
			llog.Info("Loaded Filter Script ", "type", scripttype, "filter", flt.Name)
			Filters = append(Filters, flt)
			db.Db.Save(flt)
		} else {
			llog.Error(err, "Failed to load Filter Script", "type", scripttype, "filter", flt.Name)
		}
	}
}

func InitFilter() {
	llog = log.Log.WithName("filter")
	db.Db.AutoMigrate(&Filter{})
	Filters = make([]*Filter, 0)
	scripts := filterFiles("scripts/apps", ".scp")
	loadScriptFiles(scripts, AppFilter)
	scripts = filterFiles("scripts/transports", ".scp")
	loadScriptFiles(scripts, TransportFilter)
}

func FindFilter(name string) *Filter {
	for _, flt := range Filters {
		if flt.Name == name {
			return flt
		}
	}
	return nil
}

func (ev *Filter) fnPrintf(args []object.Object) object.Object {
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

	llog.Info("Filter Script Output", "filter", ev.Name, "output", out)
	return &object.Void{}
}

func (ev *Filter) fnPrint(args []object.Object) object.Object {
	for _, e := range args {
		llog.Info("Filter Script Output", "filter", ev.Name, "Output", e.Inspect())
	}
	return &object.Void{}
}

func (ev *Filter) fnSetField(args []object.Object) object.Object {
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

	llog.Info("Setting Field Value", "filter", ev.Name, "field", fld, "value", arg)
	ev.processedMessage.Body.Fields[fld] = arg
	return &object.Void{}
}

func (ev *Filter) fnClearField(args []object.Object) object.Object {
	if len(args) != 1 {
		return &object.Null{}
	}
	// Type-check
	if args[0].Type() != object.STRING {
		return &object.Null{}
	}
	fld := args[0].(*object.String).Value
	llog.Info("Clearing Field Value", "filter", ev.Name, "field", fld)
	if _, ok := ev.processedMessage.Body.Fields[fld]; ok {
		delete(ev.processedMessage.Body.Fields, fld)
	} else {
		llog.Info("Field Not Found", "filter", ev.Name, "field", fld)
	}
	return &object.Void{}
}

func (ev *Filter) fnSetShortMessage(arg []object.Object) object.Object {
	if len(arg) != 1 {
		return &object.Null{}
	}
	// Type-check
	if arg[0].Type() != object.STRING {
		return &object.Null{}
	}
	msg := arg[0].(*object.String).Value
	llog.Info("Setting Short Message", "filter", ev.Name, "message", msg)
	ev.processedMessage.Body.ShortMsg = msg
	return &object.Void{}
}

func (ev *Filter) fnSetSeverity(arg []object.Object) object.Object {
	if len(arg) != 1 {
		return &object.Null{}
	}
	// Type-check
	if arg[0].Type() != object.STRING {
		return &object.Null{}
	}
	msg := arg[0].(*object.String).Value
	llog.Info("Setting Severity", "filter", ev.Name, "Severity", msg)
	ev.processedMessage.Body.Severity = msg
	return &object.Void{}
}

func (ev *Filter) ProcessMessage(ctx context.Context, msg *msg.Message) (bool, error) {
	if !ev.Enabled {
		return true, nil
	}
	defer func() {
		if err := recover(); err != nil {
			llog.Error(err.(error), "Filter Script Error", "filter", ev.Name)
		}
	}()
	if !ev.ok {
		llog.Info("Filter Script Not ready", "filter", ev.Name)
		return true, nil
	}
	ev.processedMessage = msg
	ev.script.SetContext(ctx)
	ok, err := ev.script.Run(msg.Body)
	ev.processedMessage = nil
	if err != nil {
		llog.Info("Filter Run Failed", "filter", ev.Name, "result", ok, "Error", err)
		return true, err
	}
	llog.V(1).Info("Filter Run Success", "filter", ev.Name, "result", ok)
	return ok, nil
}

func (ev *Filter) SetupEvalFilter() error {
	//
	// Create an evaluator, with the script inside it.
	//
	ev.script = evalfilter.New(ev.Content)
	llog.V(1).Info("Filter Script Content", "filter", ev.Name, "content", ev.Content)
	ev.script.AddFunction("printf", ev.fnPrintf)
	ev.script.AddFunction("print", ev.fnPrint)
	ev.script.AddFunction("setfield", ev.fnSetField)
	ev.script.AddFunction("clearfield", ev.fnClearField)
	ev.script.AddFunction("setshortmessage", ev.fnSetShortMessage)
	ev.script.AddFunction("setseverity", ev.fnSetSeverity)
	if err := ev.script.Prepare(); err != nil {
		llog.Error(err, "Compile Filter Script Failed", "filter", ev.Name)
		ev.ok = false
		return err
	}
	llog.V(1).Info("Compile Filter Script Success", "filter", ev.Name)
	ev.ok = true
	return nil
}
