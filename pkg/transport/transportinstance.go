package transport

import (
	"context"
	"fmt"
	"sync"

	"github.com/Fishwaldo/mouthpiece/pkg/dbdriver"
	"github.com/Fishwaldo/mouthpiece/pkg/ent"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"

	"github.com/mitchellh/mapstructure"
	"github.com/go-logr/logr"
)

type TransportInstance struct {
	dbTi              *ent.DbTransportInstances
	transportProvider interfaces.TransportProvider
	lock              sync.RWMutex
	log               logr.Logger
	impl              interfaces.TransportInstanceImpl
}

func newTransportInstance(ctx context.Context, log logr.Logger, trp interfaces.TransportProvider, name string, config interfaces.MarshableConfigI) (interfaces.TransportInstance, error) {
	var err error
	var db_ti *ent.DbTransportInstances

	newlog := log.WithName("TransportInstance").WithValues("TransportInstance", name)

	jsonconfig, err := config.AsJSON()
	if err != nil {
		newlog.Error(err, "Failed to Marshal Config to save to DB")
		return nil, mperror.ErrInvalidType
	}

	if db_ti, err = dbdriver.DbClient.DbTransportInstances.Create().
		SetName(name).
		SetTransportProvider(trp.GetName()).
		SetConfig(string(jsonconfig)).
		Save(ctx); err != nil {
		newlog.Error(err, "Failed to Save Transport Instance")
		return nil, mperror.FilterErrors(err)
	}

	ti := &TransportInstance{
		dbTi:              db_ti,
		transportProvider: trp,
		log:               newlog,
	}

	ti.impl, err = trp.CreateInstance(ctx, newlog, name, config)
	if err != nil {
		ti.log.Error(err, "Failed to Create Transport Instance")
		return nil, mperror.FilterErrors(err)
	}

	if err := ti.SetConfig(ctx, config); err != nil {
		ti.log.Error(err, "Failed to Set Config on Transport Instance")
		return nil, mperror.ErrTransportConfigInvalid
	}

	if err := ti.init(ctx); err != nil {
		ti.log.Error(err, "Failed to Init Transport Instance")
		return nil, mperror.FilterErrors(err)
	}

	if err := ti.Start(ctx); err != nil {
		newlog.Error(err, "Failed to Start Transport Instance")
		return nil, mperror.FilterErrors(err)
	}
	ti.log.V(1).Info("New Transport Instance Created")
	return ti, nil
}

func (ti *TransportInstance) Load(ctx context.Context, logger logr.Logger, dbTpi any) error {
	var ok bool
	if ti.dbTi, ok = dbTpi.(*ent.DbTransportInstances); !ok {
		logger.Error(mperror.ErrInvalidType, "Invalid Config Type", "Type", fmt.Sprintf("%T", dbTpi))
		return mperror.ErrInvalidType
	}
	ti.log = logger.WithName("TransportInstance").WithValues("TransportInstance", ti.dbTi.Name)

	var trp interfaces.TransportProvider
	var err error
	if trp, err = interfaces.GetTransportService(ctx).GetTransportProvider(ctx, ti.dbTi.TransportProvider); err != nil {
		ti.log.Error(err, "Failed to get Transport Provider")
		return mperror.FilterErrors(err)
	}

	ti.transportProvider = trp
	
	config, err := ti.transportProvider.LoadConfigFromJSON(ctx, ti.dbTi.Config)
	if err != nil {
		ti.log.Error(err, "Failed to Load Config")
		return mperror.FilterErrors(err)
	}

	if ti.impl, err = ti.transportProvider.CreateInstance(ctx, ti.log, ti.dbTi.Name, config); err != nil {
		ti.log.Error(err, "Failed to Create Transport Instance Implemntation")
		return mperror.FilterErrors(err)
	}

	if err := ti.init(ctx); err != nil {
		ti.log.Error(err, "Error Initializing Transport Instance")
		return mperror.FilterErrors(err)
	}

	if err := ti.Start(ctx); err != nil {
		ti.log.Error(err, "Error Starting Transport Instance")
		return mperror.FilterErrors(err)
	}
	ti.log.Info("Loading Transport Instance")
	return nil
}

func (ti *TransportInstance) init(ctx context.Context) error {
	return ti.impl.Init(ctx)
}

func (ti *TransportInstance) GetID() int {
	ti.lock.RLock()
	defer ti.lock.RUnlock()
	return ti.dbTi.ID
}

func (ti *TransportInstance) GetName() string {
	ti.lock.RLock()
	defer ti.lock.RUnlock()
	return ti.dbTi.Name
}

func (ti *TransportInstance) SetName(ctx context.Context, name string) error {
	ti.lock.Lock()
	defer ti.lock.Unlock()
	
	dbtmp, err := ti.dbTi.Update().SetName(name).Save(ctx)
	if err != nil {
		ti.log.Error(err, "Failed to Update Name")
		return mperror.FilterErrors(err)
	}
	ti.dbTi = dbtmp
	return nil
}

func (ti *TransportInstance) GetDescription() string {
	ti.lock.RLock()
	defer ti.lock.RUnlock()
	return ti.dbTi.Description
}

func (ti *TransportInstance) SetDescription(ctx context.Context, description string) error {
	ti.lock.Lock()
	defer ti.lock.Unlock()
	dbtmp, err := ti.dbTi.Update().SetDescription(description).Save(ctx)
	if err != nil {
		ti.log.Error(err, "Failed to Set Description")
		return mperror.FilterErrors(err)
	}
	ti.dbTi = dbtmp
	return nil
}

func (ti *TransportInstance) SetConfig(ctx context.Context, config interfaces.MarshableConfigI) error {
	ti.lock.Lock()
	defer ti.lock.Unlock()
	if err := ti.impl.SetConfig(ctx, config); err != nil {
		ti.log.Error(err, "Failed to Validate Config")
		return mperror.FilterErrors(err)
	}

	dbconfig, err := config.AsJSON()
	if err != nil {
		ti.log.Error(err, "Failed to Marshal Config to save to DB")
		return mperror.ErrTransportConfigInvalid
	}

	dbtmp, err := ti.dbTi.Update().SetConfig(string(dbconfig)).Save(ctx)
	if err != nil {
		ti.log.Error(err, "Failed to save config to DB")
		return mperror.FilterErrors(err)
	}
	ti.dbTi = dbtmp
	return nil
}

func (ti *TransportInstance) GetConfig(ctx context.Context) (interfaces.MarshableConfigI, error) {
	ti.lock.RLock()
	defer ti.lock.RUnlock()
	return ti.impl.GetConfig(ctx), nil

}

func (ti *TransportInstance) Start(ctx context.Context) error {
	ti.lock.Lock()
	defer ti.lock.Unlock()
	ti.log.Info("Starting Transport Instance")
	return mperror.FilterErrors(ti.impl.Start(ctx))
}

func (ti *TransportInstance) Stop(ctx context.Context) error {
	ti.lock.Lock()
	defer ti.lock.Unlock()
	ti.log.Info("Stopping Transport Instance")
	return mperror.FilterErrors(ti.impl.Stop(ctx))
}

func (ti *TransportInstance) Send(ctx context.Context, tpr interfaces.TransportRecipient, msg interfaces.MessageI) error {
	ti.lock.Lock()
	defer ti.lock.Unlock()
	ti.log.Info("Sending Message", "Message", msg.GetID())
	return mperror.FilterErrors(ti.impl.Send(ctx, tpr, msg))
}

func (ti *TransportInstance) ValidateTransportRecipientConfig(ctx context.Context, config interfaces.MarshableConfigI) error {
	ti.lock.Lock()
	defer ti.lock.Unlock()
	return mperror.FilterErrors(ti.impl.ValidateTransportRecipientConfig(ctx, config))
}

func (ti *TransportInstance) LoadTransportReciepientConfig(ctx context.Context, config string) (interfaces.MarshableConfigI, error) {
	ti.lock.Lock()
	defer ti.lock.Unlock()
	cfg, err := ti.impl.LoadTransportReciepientConfig(ctx, config)
	return cfg, mperror.FilterErrors(err)
}

func (ti *TransportInstance) GetAppData(ctx context.Context, name string, data any) (err error) {
	ti.lock.RLock()
	defer ti.lock.RUnlock()
	newdata, ok := ti.dbTi.AppData.Data[name]
	if !ok {
		return mperror.ErrAppDataNotFound
	}
	err = mapstructure.Decode(newdata, &data)
	if err != nil {
		ti.log.Error(err, "Error decoding AppData", "name", name)
	}
	return  nil
}

func (ti *TransportInstance) SetAppData(ctx context.Context, name string, data any) (err error) {
	ti.lock.Lock()
	defer ti.lock.Unlock()
	appdata := ti.dbTi.AppData
	appdata.Data[name] = data
	dbtmp, err := ti.dbTi.Update().SetAppData(appdata).Save(ctx)
	if err != nil {
		ti.log.Error(err, "Error setting app data on TransportInstance", "name", name)
		return mperror.FilterErrors(err)
	}
	ti.dbTi = dbtmp
	return nil
}

var _ interfaces.TransportInstance = (*TransportInstance)(nil)