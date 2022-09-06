package filter

import (
	"context"
	"embed"
	"io/fs"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/msg"

	"github.com/go-logr/logr"
	"gorm.io/gorm"
	//	"gorm.io/datatypes"
)

//go:embed scripts
var ScriptFiles embed.FS

var llog logr.Logger

type FilterImplI interface {
	FilterName() string
	Init() error
	Process(context.Context, *msg.Message) (bool, error)
}

type Filter struct {
	ID             uint `gorm:"primary_key"`
	Name           string
	Type           interfaces.FilterType
	Enabled        bool
	FilterImplType string
	Config         []Filterconfig
	lastUsed       time.Time   `gorm:"-"`
	filterImpl     FilterImplI `gorm:"-"`
}

type Filterconfig struct {
	ID       uint `gorm:"primary_key"`
	FilterID uint
	Name     string
	Value    string
}

type FilterService struct {
	ctx      *interfaces.MPContext
	fltMutex sync.Mutex
	Filters  map[string]*Filter
}

func NewFilterService() *FilterService {
	llog = log.Log.WithName("filter")
	llog.Info("Loading Filters")


	fs := &FilterService{
		Filters: make(map[string]*Filter),
	}

	return fs
}

func (fs *FilterService) Start(ctx *interfaces.MPContext) error {
	fs.ctx = ctx
	db.Db.AutoMigrate(&Filter{}, &Filterconfig{})
	
	scripts := filterFiles("scripts/apps", ".scp")
	fs.loadEvalScriptFiles(scripts, interfaces.AppFilter)

	scripts = filterFiles("scripts/transports", ".scp")
	fs.loadEvalScriptFiles(scripts, interfaces.TransportFilter)
	go fs.expireFilters()
	return nil
}

func (fs *FilterService) expireFilters() {
	for {
		select {
		case <-fs.ctx.Done():
			return
		case <-time.After(time.Second * 60):
			fs.fltMutex.Lock()
			defer fs.fltMutex.Unlock()
			for _, flt := range fs.Filters {
				if flt.lastUsed.Add(time.Second * 60).Before(time.Now()) {
					llog.Info("Expiring Filter", "name", flt.Name, "lastUsed", flt.lastUsed)
					delete(fs.Filters, flt.Name)
				}
			}
		}
	}
}

func (fs *FilterService) Get(ctx context.Context, name string, scripttype interfaces.FilterType) interfaces.FilterI {
	var flt Filter
	fs.fltMutex.Lock()
	defer fs.fltMutex.Unlock()
	if flt, ok := fs.Filters[name]; ok {
		if flt.Type == scripttype {
			flt.lastUsed = time.Now()
			return flt
		}
	}
	tx := db.Db.WithContext(ctx).Preload("Config").Find(&flt, "name = ? and type = ?", name, scripttype)
	if tx != nil && tx.Error != gorm.ErrRecordNotFound {
		llog.Error(tx.Error, "Finding filter Error", "name", name)
		return nil
	} else if tx.Error == gorm.ErrRecordNotFound {
		return nil
	}
	var err error
	if flt.filterImpl, err = GetNewFilterImpl(ctx, flt.FilterImplType, flt.Config); err != nil {
		llog.Error(err, "Failed to create Filter", "name", name)
		return nil
	}
	flt.lastUsed = time.Now()
	fs.Filters[name] = &flt
	return flt
}

func (fs *FilterService) GetByID(ctx context.Context, id uint, scripttype interfaces.FilterType) interfaces.FilterI {
	var flt Filter
	fs.fltMutex.Lock()
	defer fs.fltMutex.Unlock()
	for _, v := range fs.Filters {
		if v.ID == id {
			v.lastUsed = time.Now()
			return v
		}
	}
	tx := db.Db.WithContext(ctx).Preload("Config").Find(&flt, "id = ? and type = ?", id, scripttype)
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		llog.Error(tx.Error, "Finding filter Error", "id", id)
		return nil
	} else if tx.Error == gorm.ErrRecordNotFound {
		return nil
	}
	var err error
	if flt.filterImpl, err = GetNewFilterImpl(ctx, flt.FilterImplType, flt.Config); err != nil {
		llog.Error(err, "Failed to create Filter", "name", flt.Name)
		return nil
	}
	flt.lastUsed = time.Now()
	fs.Filters[flt.Name] = &flt
	return flt
}

func (myfs *FilterService) loadEvalScriptFiles(files []string, scripttype interfaces.FilterType) {
	for _, script := range files {
		flt := &Filter{}
		if tx := db.Db.Find(flt, "name = ? and type = ?", trimFileExtension(filepath.Base(script)), scripttype); tx.RowsAffected == 0 {
			llog.V(1).Info("Reading Filter Script from Filesystem", "type", scripttype, "filter", trimFileExtension(filepath.Base(script)))
			content, err := fs.ReadFile(ScriptFiles, script)
			if err != nil {
				llog.Error(err, "Failed to read Filter Script File", "filename", script)
				continue
			}
			//
			// Create an evalfilter, with the script inside it.
			//
			config := []Filterconfig{
				{
					Name:  "script",
					Value: string(content),
				},
				{
					Name:  "name",
					Value: trimFileExtension(filepath.Base(script)),
				},
			}
			if fltimpl, err := GetNewFilterImpl(context.Background(), "EvalFilter", config); err != nil {
				llog.Error(err, "Failed to create EvalFilter", "filename", script)
				continue
			} else {
				if err := fltimpl.Init(); err == nil {
					llog.Info("Loaded Filter Script ", "type", scripttype, "filter", fltimpl.FilterName())
					flt := &Filter{
						Type:           scripttype,
						filterImpl:     fltimpl,
						FilterImplType: "EvalFilter",
						Config:         config,
						lastUsed:       time.Now(),
						Enabled:        true,
					}
					flt.Name = flt.GetConfig("name").(string)
					myfs.fltMutex.Lock()
					myfs.Filters[flt.Name] = flt
					myfs.fltMutex.Unlock()
					if tx := db.Db.Save(&flt); tx.Error != nil {
						llog.Error(tx.Error, "Failed to save Filter into database", "name", flt.Name)
					}
				} else {
					llog.Error(err, "Failed to load Filter Script into database", "type", scripttype, "filter", fltimpl.FilterName())
				}
			}
		} else {
			llog.Info("Filter Script already loaded into Database", "type", scripttype, "filter", trimFileExtension(filepath.Base(script)))
		}
	}
}

func (f Filter) GetID() uint {
	return f.ID
}

func (f Filter) GetName() string {
	return f.Name
}

func (f Filter) ProcessMessage(ctx context.Context, msg *msg.Message) (bool, error) {
	if f.Config == nil {
		if tx := db.Db.WithContext(ctx).Find(&f.Config, "filter_id = ? ", f.ID); tx.Error != nil {
			return false, tx.Error
		}
	}
	var err error
	if f.filterImpl == nil {
		if f.filterImpl, err = GetNewFilterImpl(ctx, f.FilterImplType, f.Config); err != nil {
			return true, err
		}
	}
	return f.filterImpl.Process(ctx, msg)
}

func (fs Filter) GetConfig(item string) interface{} {
	for _, v := range fs.Config {
		if strings.EqualFold(v.Name, item) {
			return v.Value
		}
	}
	return nil
}

func trimFileExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

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
