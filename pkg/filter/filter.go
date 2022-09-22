package filter

import (
	"context"
	"embed"
	"fmt"
	"sync"

	"github.com/Fishwaldo/mouthpiece/pkg/mperror"

	"github.com/Fishwaldo/mouthpiece/pkg/db"

	"github.com/Fishwaldo/mouthpiece/pkg/ent"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"

	"github.com/go-logr/logr"
)

//go:embed scripts
var ScriptFiles embed.FS

type Filter struct {
	interfaces.CacheAble
	dbFilter   *ent.DbFilter
	filterImpl interfaces.FilterImplI
	lock       sync.RWMutex
	log        logr.Logger
}

func newFilter(ctx context.Context, log logr.Logger, fltimpl string, name string, flttype interfaces.FilterType) interfaces.FilterI {
	newlog := log.WithName("Filter")
	newlog.Info("Creating new filter")
	flt, err := GetNewFilterImpl(ctx, fltimpl, "{}")
	if err != nil {
		newlog.Error(err, "Error creating new filter")
		return nil
	}
	dbFilter, err := db.DbClient.DbFilter.Create().SetName(name).SetFilterImpl(fltimpl).SetType(flttype).SetConfig("{}").Save(ctx)
	if err != nil {
		newlog.Error(err, "Error creating new filter")
		return nil
	}
	return &Filter{
		log: log.WithName(dbFilter.Name),
		filterImpl: flt,
		dbFilter: dbFilter,
	}
}

func (f *Filter) Load(ctx context.Context, logger logr.Logger, dbFilter any) (err error) {
	var ok bool
	if f.dbFilter, ok = dbFilter.(*ent.DbFilter); ok {
		f.log = logger.WithName("Filter").WithValues("Filter", f.dbFilter.Name)
		return f.loadFilterImpl(ctx)
	}
	logger.Error(mperror.ErrInvalidType, "Invalid type passed to filter Loader", "Type", fmt.Sprintf("%T", dbFilter))
	return mperror.ErrInternalError
}


func (f *Filter) GetID() int {
	f.lock.RLock()
	defer f.lock.RUnlock()
	return f.dbFilter.ID
}

func (f *Filter) GetName() string {
	f.lock.RLock()
	defer f.lock.RUnlock()
	return f.dbFilter.Name
}

func (f *Filter) SetName(ctx context.Context, name string) (err error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	dbtmp, err := f.dbFilter.Update().SetName(name).Save(ctx)
	if err != nil {
		f.log.Error(err, "Failed to update name")
		return mperror.FilterErrors(err)
	}
	f.dbFilter = dbtmp
	
	
	return nil
}

func (f *Filter) GetDescription() string {
	f.lock.RLock()
	defer f.lock.RUnlock()
	return f.dbFilter.Description
}

func (f *Filter) SetDescription(ctx context.Context, desc string) (err error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	dbtmp, err := f.dbFilter.Update().SetDescription(desc).Save(ctx)
	if err != nil {
		f.log.Error(err, "Failed to update description")
		return mperror.FilterErrors(err)
	}
	f.dbFilter = dbtmp
	return nil
}

func (f *Filter) GetType() string {
	f.lock.RLock()
	defer f.lock.RUnlock()
	return f.dbFilter.Type.String()
}

func (f *Filter) GetFilterImplementation() interfaces.FilterImplI {
	f.lock.RLock()
	defer f.lock.RUnlock()
	return f.filterImpl
}

func (f *Filter) ProcessMessage(ctx context.Context, msg interfaces.MessageI) (ok interfaces.FilterAction, err error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.log.V(1).Info("Processing message", "msg", msg.GetID())
	if err := f.loadFilterImpl(ctx); err != nil {
		f.log.Error(err, "Error loading filter implementation")
		return interfaces.FilterPass, mperror.ErrInternalError
	}
	ok, err = f.filterImpl.Process(ctx, msg)
	return ok, mperror.FilterErrors(err)
}

func (f *Filter) GetConfig(ctx context.Context) (config interfaces.MarshableConfigI, err error) {
	f.lock.RLock()
	defer f.lock.RUnlock()
	if err := f.loadFilterImpl(ctx); err != nil {
		f.log.Error(err, "Error loading filter implementation")
		return nil, mperror.ErrInternalError
	}
	cfg, err := f.filterImpl.GetConfig(ctx)
	return cfg, mperror.FilterErrors(err)
}

func (f *Filter) SetConfig(ctx context.Context, config interfaces.MarshableConfigI) (err error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	cfg, err  := config.AsJSON();
	if err != nil {
		f.log.Error(err, "Error marshalling config")
		return mperror.ErrInternalError
	}
	dbtmp, err := f.dbFilter.Update().SetConfig(cfg).Save(ctx)
	if err != nil {
		f.log.Error(err, "Error saving filter config")
		return mperror.FilterErrors(err)
	}
	f.dbFilter = dbtmp

	if err := f.loadFilterImpl(ctx); err != nil {
		f.log.Error(err, "Error loading filter implementation")
		return mperror.FilterErrors(err)
	}
	if err := f.filterImpl.SetConfig(ctx, config); err != nil {
		f.log.Error(err, "Error updating filter implementation config")
		return mperror.FilterErrors(err)
	}

	return nil
}

func (f *Filter) loadFilterImpl(ctx context.Context) (err error) {
	if f.filterImpl == nil {
		if f.filterImpl, err = GetNewFilterImpl(ctx, f.dbFilter.FilterImpl, f.dbFilter.Config); err != nil {
			f.log.Error(err, "Error loading filter implementation", "filter", f.dbFilter.Name)
			return mperror.FilterErrors(err)
		}
	}
	return nil
}


func (f *Filter) Save(ctx context.Context) (err error) {
	dbtmp, err := f.dbFilter.Update().Save(ctx)
	if err != nil {
		f.log.Error(err, "Error saving filter", "filter", f.dbFilter.Name)
		return mperror.FilterErrors(err)
	}
	f.dbFilter = dbtmp
	return nil
}

var _ interfaces.FilterI = (*Filter)(nil)
