package transport

import (
	"context"
	"sync"

	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/ent"

	//	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbtransportinstances"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbtransportrecipients"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"

	"github.com/go-logr/logr"
)

type TransportService struct {
	transportInstances map[string]interfaces.TransportInstance
	tilock             sync.RWMutex
	log                logr.Logger
}

func NewTransportService(ctx context.Context, log logr.Logger) *TransportService {
	ts := &TransportService{
		transportInstances: make(map[string]interfaces.TransportInstance),
		log:                log.WithName("TransportService"),
	}
	ts.log.V(1).Info("New Transport Service")

	tprs := GetTransportProviders(ctx)
	for _, tpr := range tprs {
		ts.log.V(1).Info("Transport Provider Registered", "Name", tpr.GetName())
	}
	return ts
}

func (ts *TransportService) Start(ctx context.Context) error {
	var db_tis []*ent.DbTransportInstances
	var err error
	if db_tis, err = db.DbClient.DbTransportInstances.Query().All(ctx); err != nil {
		ts.log.Error(err, "Error Bootstrapping all TransportInstances")
		return mperror.ErrInternalError
	}
	for _, db_ti := range db_tis {
		if ti, err := ts.LoadTransportInstance(ctx, db_ti); err != nil {
			ts.log.Error(err, "Error loading TransportInstance", "name", db_ti.Name)
			continue
		} else {
			if err := ti.Start(ctx); err != nil {
				ts.log.Error(err, "Error starting TransportInstance", "name", db_ti.Name)
				continue
			}
			ts.tilock.Lock()
			ts.transportInstances[db_ti.Name] = ti
			ts.tilock.Unlock()
		}
	}
	return nil
}

func (ts *TransportService) Create(ctx context.Context, transport interfaces.TransportInstance, name string, config interfaces.MarshableConfigI) (interfaces.TransportRecipient, error) {
	if ok := ts.Exists(ctx, name); ok {
		return nil, mperror.ErrTransportReciptientExists
	}
	app, err := newTransportRecipient(ctx, ts.log, transport, name, config)
	if err != nil {
		ts.log.Error(err, "Error creating TransportRecipient", "name", name)
		return nil, err
	} else {
		return app, nil
	}
}

func (ts *TransportService) Delete(ctx context.Context, tr interfaces.TransportRecipient) error {
	if ok := ts.Exists(ctx, tr.GetName()); !ok {
		return mperror.ErrTransportReciepientNotFound
	}
	if err := db.DbClient.DbTransportRecipients.DeleteOneID(tr.GetID()).Exec(ctx); err != nil {
		ts.log.Error(err, "Error deleting TransportRecipient", "name", tr.GetName())
		return err
	}
	return nil
}

func (ts *TransportService) Get(ctx context.Context, name string) (interfaces.TransportRecipient, error) {
	db_tr, err := db.DbClient.DbTransportRecipients.Query().Where(dbtransportrecipients.Name(name)).Only(ctx)
	if err != nil {
		ts.log.Error(err, "Error getting TransportRecipient", "name", name)
		return nil, err
	}
	return ts.Load(ctx, db_tr)
}

func (ts *TransportService) GetByID(ctx context.Context, id int) (interfaces.TransportRecipient, error) {
	db_tr, err := db.DbClient.DbTransportRecipients.Query().Where(dbtransportrecipients.ID(id)).Only(ctx)
	if err != nil {
		ts.log.Error(err, "Error getting TransportRecipient", "ID", id)
		return nil, err
	}
	return ts.Load(ctx, db_tr)
}

func (ts *TransportService) GetAll(ctx context.Context) (tprs []interfaces.TransportRecipient, err error) {
	var dbtrs []*ent.DbTransportRecipients
	if dbtrs, err = db.DbClient.DbTransportRecipients.Query().All(ctx); err != nil {
		ts.log.Error(err, "Error getting all TransportRecipient")
		return nil, mperror.ErrInternalError
	}

	for _, dbtr := range dbtrs {
		if tr, err := ts.Load(ctx, dbtr); err != nil {
			ts.log.Error(err, "Error loading TransportRecipient", "name", dbtr.Name)
			return nil, mperror.ErrInternalError
		} else {
			tprs = append(tprs, tr)
		}
	}
	return tprs, nil
}

func (ts *TransportService) Load(ctx context.Context, dbtr any) (interfaces.TransportRecipient, error) {
	entTr, ok := dbtr.(*ent.DbTransportRecipients)
	if !ok {
		return nil, mperror.ErrInvalidType
	}
	var err error
	if entTr.Edges.TransportInstance, err = entTr.Edges.TransportInstanceOrErr(); err != nil {
		entTr.Edges.TransportInstance = entTr.QueryTransportInstance().OnlyX(ctx)
	}

	tr := &TransportRecipient{}

	if err := tr.Load(ctx, ts.log, entTr); err != nil {
		return nil, err
	}
	return tr, nil
}

func (ts *TransportService) Exists(ctx context.Context, name string) bool {
	if ok, err := db.DbClient.DbTransportRecipients.Query().Where(dbtransportrecipients.Name(name)).Exist(ctx); err != nil {
		ts.log.Error(err, "Error checking if TransportRecipient exists", "name", name)
		return false
	} else {
		return ok
	}
}

func (ts *TransportService) ExistsByID(ctx context.Context, id int) bool {
	if ok, err := db.DbClient.DbTransportRecipients.Query().Where(dbtransportrecipients.ID(id)).Exist(ctx); err != nil {
		ts.log.Error(err, "Error checking if TransportRecipient exists", "ID", id)
		return false
	} else {
		return ok
	}
}

func (ts *TransportService) CreateTransportInstance(ctx context.Context, trp interfaces.TransportProvider, name string, config interfaces.MarshableConfigI) (interfaces.TransportInstance, error) {
	if ok := ts.ExistsTransportInstance(ctx, name); ok {
		return nil, mperror.ErrTransportInstanceExists
	}

	ts.log.Info("Creating Transport Instance", "name", name, "provider", trp.GetName(), "config", config)
	tri, err := newTransportInstance(ctx, ts.log, trp, name, config)
	if err != nil {
		ts.log.Error(err, "Error creating TransportInstance", "name", name)
		return nil, mperror.ErrInternalError
	} else {
		ts.tilock.Lock()
		ts.transportInstances[name] = tri
		ts.tilock.Unlock()
		ts.log.Info("Transport Instance Created", "name", name)
		return tri, nil
	}
}

func (ts *TransportService) GetTransportInstance(ctx context.Context, name string) (interfaces.TransportInstance, error) {
	ts.tilock.RLock()
	defer ts.tilock.RUnlock()
	if tri, ok := ts.transportInstances[name]; ok {
		return tri, nil
	}
	return nil, mperror.ErrTransportInstanceNotFound
}
func (ts *TransportService) GetTransportInstanceByID(ctx context.Context, id int) (interfaces.TransportInstance, error) {
	ts.tilock.RLock()
	defer ts.tilock.RUnlock()
	for _, tri := range ts.transportInstances {
		if tri.GetID() == id {
			return tri, nil
		}
	}
	return nil, mperror.ErrTransportInstanceNotFound
}
func (ts *TransportService) GetAllTransportInstances(ctx context.Context) (tis []interfaces.TransportInstance, err error) {
	ts.tilock.RLock()
	defer ts.tilock.RUnlock()
	for _, ti := range ts.transportInstances {
		tis = append(tis, ti)
	}
	return tis, nil
}
func (ts *TransportService) DeleteTransportInstance(ctx context.Context, ti interfaces.TransportInstance) error {

	if ok := ts.ExistsTransportInstance(ctx, ti.GetName()); !ok {
		return mperror.ErrTransportInstanceNotFound
	}
	ts.tilock.Lock()
	defer ts.tilock.Unlock()
	if err := db.DbClient.DbTransportInstances.DeleteOneID(ti.GetID()).Exec(ctx); err != nil {
		ts.log.Error(err, "Error deleting TransportInstance", "name", ti.GetName())
		return mperror.ErrInternalError
	}
	delete(ts.transportInstances, ti.GetName())
	return nil
}

func (ts *TransportService) ExistsTransportInstance(ctx context.Context, name string) bool {
	ts.tilock.RLock()
	defer ts.tilock.RUnlock()
	if _, ok := ts.transportInstances[name]; ok {
		return true
	}
	return false
}
func (ts *TransportService) ExistsByIDTransportInstance(ctx context.Context, id int) bool {
	ts.tilock.RLock()
	defer ts.tilock.RUnlock()
	for _, tri := range ts.transportInstances {
		if tri.GetID() == id {
			return true
		}
	}
	return false
}

func (ts *TransportService) LoadTransportInstance(ctx context.Context, dbti any) (interfaces.TransportInstance, error) {
	entTi, ok := dbti.(*ent.DbTransportInstances)
	if !ok {
		return nil, mperror.ErrInvalidType
	}
	var ti interfaces.TransportInstance

	if err := ti.Load(ctx, ts.log, entTi); err != nil {
		ts.log.Error(err, "Error loading TransportInstance", "name", entTi.Name)
		return nil, err
	}
	if err := ti.Start(ctx); err != nil {
		ts.log.Error(err, "Error starting TransportInstance", "name", entTi.Name)
		return nil, mperror.ErrInternalError
	}
	ts.tilock.Lock()
	ts.transportInstances[ti.GetName()] = ti
	ts.tilock.Unlock()
	return ti, nil
}

func (ts *TransportService) GetTransportProvider(ctx context.Context, name string) (interfaces.TransportProvider, error) {
	return GetTransportProvider(ctx, name)
}

func (ts *TransportService) GetAllTransportProviders(ctx context.Context) ([]interfaces.TransportProvider, error) {
	return GetTransportProviders(ctx), nil
}

var _ interfaces.TransportServiceI = (*TransportService)(nil)
