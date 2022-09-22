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
	FilterCache  map[int]interfaces.FilterI
	log logr.Logger
}

func NewFilterService(ctx context.Context, logger logr.Logger) *FilterService {
	fs := &FilterService{
		FilterCache: make(map[int]interfaces.FilterI),
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
		case <-time.After(1 * time.Second):
			fs.fltMutex.Lock()
			for name, flt := range fs.FilterCache {
				if flt.GetLastUsed().Add(interfaces.Config.ExpireFilters).Before(time.Now()) {
					fs.log.Info("Expiring Filter From Cache", "name", name, "lastUsed", flt.GetLastUsed())
					delete(fs.FilterCache, flt.GetID())
				}
			}
			fs.fltMutex.Unlock()
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
	fs.FilterCache[flt.GetID()] = flt
	return flt, nil
}

func (fs *FilterService) Get(ctx context.Context, name string, scripttype interfaces.FilterType) (interfaces.FilterI, error) {
	fs.fltMutex.Lock()
	defer fs.fltMutex.Unlock()
	for _, v := range fs.FilterCache {
		if v.GetName() == name && v.GetType() == scripttype.String() {
			v.SetLastUsed()
			return v, nil
		}
	}

	fs.log.V(1).Info("Filter not found in cache", "name", name)
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
	fs.FilterCache[flt.GetID()] = flt
	return flt, nil
}

func (fs *FilterService) GetByID(ctx context.Context, id int, scripttype interfaces.FilterType) (interfaces.FilterI, error) {
	var flt interfaces.FilterI
	fs.fltMutex.Lock()
	defer fs.fltMutex.Unlock()
	if flt, ok := fs.FilterCache[id]; ok {
		if flt.GetType() == scripttype.String() {
			flt.SetLastUsed()
			return flt, nil
		}
	}

	fs.log.V(1).Info("Filter not found in cache", "id", id)
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
	fs.FilterCache[flt.GetID()] = flt

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
		flt.SetLastUsed()
		fs.FilterCache[flt.GetID()] = flt
		return flt, nil
	}
	
}
