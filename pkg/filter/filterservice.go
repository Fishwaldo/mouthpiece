package filter

import (
	"context"
	"sync"
	"time"

	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/ent"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbfilter"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"

	"github.com/go-logr/logr"
)



type FilterService struct {
	fltMutex sync.Mutex
	Filters  map[string]interfaces.FilterI
	log logr.Logger
}

func NewFilterService(ctx context.Context, logger logr.Logger) *FilterService {
	fs := &FilterService{
		Filters: make(map[string]interfaces.FilterI),
		log: logger.WithName("FilterService"),
	}
	fs.log.V(1).Info("New Filter Service")
	fltimpls := GetFilterImpls(ctx)
	for _, fltimpl := range fltimpls {
		fs.log.V(1).Info("Filter Registered", "name", fltimpl)
	}
	return fs
}

func (fs *FilterService) Start(ctx context.Context) error {
	go fs.expireFilters()
	return nil
}

func (fs *FilterService) expireFilters() {
	for {
		select {
//		case <-fs.ctx.Done():
//			return
		case <-time.After(time.Second * 60):
			fs.fltMutex.Lock()
			defer fs.fltMutex.Unlock()
			for name, flt := range fs.Filters {
				if flt.GetLastUsed().Add(time.Second * 60).Before(time.Now()) {
					fs.log.Info("Expiring Filter", "name", name, "lastUsed", flt.GetLastUsed())
					delete(fs.Filters, name)
				}
			}
		}
	}
}

func (fs *FilterService) Load(ctx context.Context, dbflt any) (interfaces.FilterI, error) {
	mydbflt, ok := dbflt.(*ent.DbFilter)
	if !ok {
		fs.log.Error(nil, "Invalid filter type")
		return nil, mperror.ErrInvalidType
	}
	flt := &Filter{}
	if err := flt.Load(ctx, fs.log, mydbflt); err != nil {
		fs.log.Error(err, "Failed to load filter", "name", mydbflt.Name)
		return nil, mperror.ErrInternalError
	}
	flt.SetLastUsed()
	fs.Filters[flt.GetName()] = flt
	return flt, nil
}

func (fs *FilterService) Get(ctx context.Context, name string, scripttype interfaces.FilterType) (interfaces.FilterI, error) {
	fs.fltMutex.Lock()
	defer fs.fltMutex.Unlock()
	if v, ok := fs.Filters[name]; ok {
		v.SetLastUsed()
		return v, nil
	}
	dbflt, err := db.DbClient.DbFilter.Query().Where(dbfilter.Name(name), dbfilter.TypeEQ(scripttype)).Only(ctx)
	if err != nil {
		fs.log.Error(err, "Error getting filter", "name", name)
		return nil, mperror.ErrFilterNotFound
	}
	flt, err := fs.Load(ctx, dbflt)
	if err != nil {
		fs.log.Error(err, "Error loading filter", "name", name)
		return nil, mperror.ErrInternalError
	}
	flt.SetLastUsed()
	fs.Filters[name] = flt
	return flt, nil
}

func (fs *FilterService) GetByID(ctx context.Context, id int, scripttype interfaces.FilterType) (interfaces.FilterI, error) {
	var flt interfaces.FilterI
	fs.fltMutex.Lock()
	defer fs.fltMutex.Unlock()
	for _, v := range fs.Filters {
		if v.GetID() == id {
			v.SetLastUsed()
			return v, nil
		}
	}
	
	dbflt, err := db.DbClient.DbFilter.Query().Where(dbfilter.ID(id), dbfilter.TypeEQ(scripttype)).Only(ctx)
	if err != nil {
		fs.log.Error(err, "Error getting filter", "id", id)
		return nil, mperror.ErrFilterNotFound
	}
	flt, err = fs.Load(ctx, dbflt)
	if err != nil {
		fs.log.Error(err, "Error loading filter", "id", id)
		return nil, mperror.ErrInternalError
	}
	flt.SetLastUsed()
	fs.Filters[flt.GetName()] = flt

	return flt, nil
}

func (fs *FilterService) Create(ctx context.Context, fltimpl string, name string, flttype interfaces.FilterType) (interfaces.FilterI, error) {
	fs.fltMutex.Lock()
	defer fs.fltMutex.Unlock()

	//TODO Check if a filter exists

	if flt := newFilter(ctx, fs.log, fltimpl, name, flttype); flt == nil {
		fs.log.Error(nil, "Error creating filter", "name", name)
		return nil, mperror.ErrInternalError
	} else {
		fs.Filters[name] = flt
		return flt, nil
	}
	
}
