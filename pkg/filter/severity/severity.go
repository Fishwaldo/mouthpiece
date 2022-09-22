package severity

import (
	"context"
	"encoding/json"

	"github.com/Fishwaldo/mouthpiece/pkg/filter"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"
)

func init() {
	filter.RegisterFilterImpl("SeverityFilter", SevFilterFactory{})
}

type SeverityFilterOp int

const (
	SeverityFilterOpEQ SeverityFilterOp = iota
	SeverityFilterOpNE
	SeverityFilterOpGT
	SeverityFilterOpLT
	SeverityFilterOpGE
	SeverityFilterOpLE
)

type SeverityFilter struct {
	config *SeverityFilterConfig
}

type SeverityFilterConfig struct {
	Op       SeverityFilterOp
	Severity int
}

func (c *SeverityFilterConfig) AsJSON() (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *SeverityFilterConfig) FromJSON(data string) error {
	return json.Unmarshal([]byte(data), c)
}

type SevFilterFactory struct {
}

func (sff SevFilterFactory) FilterFactory(ctx context.Context, config string) (interfaces.FilterImplI, error) {
	var cfg SeverityFilterConfig
	if err := cfg.FromJSON(config); err != nil {
		return nil, mperror.ErrFilterConfigInvalid
	}
	return &SeverityFilter{config: &cfg}, nil
}

func (sff SevFilterFactory) DefaultConfig(ctx context.Context) interfaces.MarshableConfigI {
	return &SeverityFilterConfig{
		Op:       SeverityFilterOpGE,
		Severity: 0,
	}
}

func (sf *SeverityFilter) Init(ctx context.Context) error {
	return nil
}

func (sf *SeverityFilter) Process(ctx context.Context, msg interfaces.MessageI) (interfaces.FilterAction, error) {
	switch sf.config.Op {
	case SeverityFilterOpEQ:
		if msg.GetSeverity() == sf.config.Severity {
			return interfaces.FilterMatch, nil
		}
	case SeverityFilterOpNE:
		if msg.GetSeverity() != sf.config.Severity {
			return interfaces.FilterMatch, nil
		}
	case SeverityFilterOpGT:
		if msg.GetSeverity() > sf.config.Severity {
			return interfaces.FilterMatch, nil
		}
	case SeverityFilterOpLT:
		if  msg.GetSeverity() < sf.config.Severity {
			return interfaces.FilterMatch, nil
		}
	case SeverityFilterOpGE:
		if msg.GetSeverity() >= sf.config.Severity {
			return interfaces.FilterMatch, nil
		}
	case SeverityFilterOpLE:
		if msg.GetSeverity() <= sf.config.Severity {
			return interfaces.FilterMatch, nil
		}
	default:
		return interfaces.FilterPass, mperror.ErrFilterConfigInvalid
	}
	return interfaces.FilterNoMatch, nil
}

func (sf *SeverityFilter) FilterName() string {
	return "SeverityFilter"
}

func (sf *SeverityFilter) SetConfig(ctx context.Context, config interfaces.MarshableConfigI) error {
	var ok bool
	if sf.config, ok = config.(*SeverityFilterConfig); !ok {
		return mperror.ErrFilterConfigInvalid
	}
	if sf.config.Op > SeverityFilterOpLE {
		return mperror.ErrFilterConfigInvalid
	}
	return nil
}
func (sf *SeverityFilter) GetConfig(ctx context.Context) (interfaces.MarshableConfigI, error) {
	return sf.config, nil
}

var _ interfaces.FilterImplI = (*SeverityFilter)(nil)

// func trimFileExtension(fileName string) string {
// 	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
// }

// func filterFiles(root, ext string) []string {
// 	var a []string
// 	fs.WalkDir(ScriptFiles, root, func(s string, d fs.DirEntry, e error) error {
// 		if e != nil {
// 			return e
// 		}
// 		if filepath.Ext(d.Name()) == ext {
// 			a = append(a, s)
// 		}

// 		return nil
// 	})
// 	return a
// }

// func (myfs *FilterService) loadEvalScriptFiles(files []string, scripttype interfaces.FilterType) {
// 	for _, script := range files {
// 		flt := &Filter{}
// 		if tx := db.Db.Find(flt, "name = ? and type = ?", trimFileExtension(filepath.Base(script)), scripttype); tx.RowsAffected == 0 {
// 			llog.V(1).Info("Reading Filter Script from Filesystem", "type", scripttype, "filter", trimFileExtension(filepath.Base(script)))
// 			content, err := fs.ReadFile(ScriptFiles, script)
// 			if err != nil {
// 				llog.Error(err, "Failed to read Filter Script File", "filename", script)
// 				continue
// 			}
// 			//
// 			// Create an evalfilter, with the script inside it.
// 			//
// 			config := []Filterconfig{
// 				{
// 					Name:  "script",
// 					Value: string(content),
// 				},
// 				{
// 					Name:  "name",
// 					Value: trimFileExtension(filepath.Base(script)),
// 				},
// 			}
// 			if fltimpl, err := GetNewFilterImpl(context.Background(), "EvalFilter", config); err != nil {
// 				llog.Error(err, "Failed to create EvalFilter", "filename", script)
// 				continue
// 			} else {
// 				if err := fltimpl.Init(); err == nil {
// 					llog.Info("Loaded Filter Script ", "type", scripttype, "filter", fltimpl.FilterName())
// 					flt := &Filter{
// 						Type:           scripttype,
// 						filterImpl:     fltimpl,
// 						FilterImplType: "EvalFilter",
// 						Config:         config,
// 						lastUsed:       time.Now(),
// 						Enabled:        true,
// 					}
// 					flt.Name = flt.GetConfig("name").(string)
// 					myfs.fltMutex.Lock()
// 					myfs.Filters[flt.Name] = flt
// 					myfs.fltMutex.Unlock()
// 					if tx := db.Db.Save(&flt); tx.Error != nil {
// 						llog.Error(tx.Error, "Failed to save Filter into database", "name", flt.Name)
// 					}
// 				} else {
// 					llog.Error(err, "Failed to load Filter Script into database", "type", scripttype, "filter", fltimpl.FilterName())
// 				}
// 			}
// 		} else {
// 			llog.Info("Filter Script already loaded into Database", "type", scripttype, "filter", trimFileExtension(filepath.Base(script)))
// 		}
// 	}
// }

// scripts := filterFiles("scripts/apps", ".scp")
// fs.loadEvalScriptFiles(scripts, interfaces.AppFilter)

// scripts = filterFiles("scripts/transports", ".scp")
// fs.loadEvalScriptFiles(scripts, interfaces.TransportFilter)
