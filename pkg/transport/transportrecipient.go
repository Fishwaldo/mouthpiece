package transport

import (
	"context"
	"fmt"
	"sync"

	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/ent"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbtransportrecipients"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"

	"github.com/go-logr/logr"
)

type TransportRecipient struct {
	lock   sync.RWMutex
	dbTr   *ent.DbTransportRecipients
	config interfaces.MarshableConfigI
	tpi    interfaces.TransportInstance
	log    logr.Logger
}

func newTransportRecipient(ctx context.Context, logger logr.Logger, tpi interfaces.TransportInstance, name string, config interfaces.MarshableConfigI) (interfaces.TransportRecipient, error) {
	tpr := &TransportRecipient{
		tpi: tpi,
		log: logger.WithName("TransportRecipient").WithValues("name", name),
	}
	var err error
	jsoncfg, err := config.AsJSON()
	if err != nil {
		tpr.log.Error(err, "Error marshalling config")
		return nil, mperror.ErrTransportConfigInvalid
	}
	if tpr.dbTr, err = db.DbClient.DbTransportRecipients.Create().SetName(name).SetConfig(jsoncfg).SetTransportInstanceID(tpi.GetID()).Save(ctx); err != nil {
		tpr.log.Error(err, "Failed to Create Transport Recipient")
		return nil, mperror.FilterErrors(err)
	}
	if err = tpr.init(); err != nil {
		tpr.log.Error(err, "Failed to Initialize Transport Recipient")
		return nil, mperror.FilterErrors(err)
	}
	return tpr, nil
}

func (tr *TransportRecipient) init() error {
	return nil
}

func (tr *TransportRecipient) Load(ctx context.Context, logger logr.Logger, dbtr any) error {
	tr.lock.Lock()
	defer tr.lock.Unlock()

	var ok bool
	tr.dbTr, ok = dbtr.(*ent.DbTransportRecipients)
	if !ok {
		logger.Error(mperror.ErrInvalidType, "Failed to load Transport Recipient", "Type", fmt.Sprintf("%T", dbtr))
		return mperror.ErrInvalidType
	}
	tr.log = logger.WithName("TransportRecipient").WithValues("name", tr.dbTr.Name)

	var err error
	if tr.dbTr.Edges.TransportInstance, err = tr.dbTr.Edges.TransportInstanceOrErr(); err != nil {
		if tr.dbTr.Edges.TransportInstance, err = tr.dbTr.QueryTransportInstance().Only(ctx); err != nil {
			tr.log.Error(err, "Failed to load Transport Instance")
			return mperror.FilterErrors(err)
		}
	}
	tr.tpi, err = interfaces.GetTransportService(ctx).GetTransportInstanceByID(ctx, tr.dbTr.Edges.TransportInstance.ID)
	if err != nil {
		tr.log.Error(err, "Failed to load Transport Instance")
		return mperror.FilterErrors(err)
	}

	tr.config, err = tr.tpi.LoadTransportReciepientConfig(ctx, tr.dbTr.Config)
	if err != nil {
		tr.log.Error(err, "Failed to load Transport Recipient Config")
		return mperror.FilterErrors(err)
	}

	return tr.init()
}

func (tr *TransportRecipient) GetID() int {
	tr.lock.RLock()
	defer tr.lock.RUnlock()
	return tr.dbTr.ID
}

func (tr *TransportRecipient) GetName() string {
	tr.lock.RLock()
	defer tr.lock.RUnlock()
	return tr.dbTr.Name
}

func (tr *TransportRecipient) SetName(ctx context.Context, name string) error {
	tr.lock.Lock()
	defer tr.lock.Unlock()

	dbtmp, err := tr.dbTr.Update().SetName(name).Save(ctx)
	if err != nil {
		tr.log.Error(err, "Failed to Update Transport Recipient Name")
		return mperror.FilterErrors(err)
	}
	tr.dbTr = dbtmp
	return nil
}
func (tr *TransportRecipient) GetDescription() string {
	tr.lock.RLock()
	defer tr.lock.RUnlock()
	return tr.dbTr.Description
}

func (tr *TransportRecipient) SetDescription(ctx context.Context, description string) error {
	tr.lock.Lock()
	defer tr.lock.Unlock()
	dbtmp, err := tr.dbTr.Update().SetDescription(description).Save(ctx)
	if err != nil {
		tr.log.Error(err, "Failed to Update Transport Recipient Description")
		return mperror.FilterErrors(err)
	}
	tr.dbTr = dbtmp
	return nil
}

func (tr *TransportRecipient) SetConfig(ctx context.Context, config interfaces.MarshableConfigI) error {
	tr.lock.Lock()
	defer tr.lock.Unlock()
	if err := tr.tpi.ValidateTransportRecipientConfig(ctx, config); err != nil {
		tr.log.Error(err, "Failed to Validate Config")
		return mperror.FilterErrors(err)
	}

	if configstr, err := config.AsJSON(); err != nil {
		tr.log.Error(err, "Error marshalling config")
		return mperror.ErrFilterConfigInvalid
	} else {
		if err := tr.dbTr.Update().SetConfig(string(configstr)).Exec(ctx); err != nil {
			tr.log.Error(err, "Error updating config")
			return mperror.FilterErrors(err)
		}
		tr.config = config
	}
	return nil
}

func (tr *TransportRecipient) GetConfig() (interfaces.MarshableConfigI, error) {
	tr.lock.RLock()
	defer tr.lock.RUnlock()
	return tr.config, nil
}

func (tr *TransportRecipient) loadEdges(ctx context.Context) error {
	var err error
	if tr.dbTr.Edges.UserRecipient, err = tr.dbTr.Edges.UserRecipientOrErr(); err != nil {
		if tr.dbTr, err = db.DbClient.DbTransportRecipients.Query().WithGroupRecipient().WithUserRecipient().Where(dbtransportrecipients.IDEQ(tr.dbTr.ID)).First(ctx); err != nil {
			tr.log.Error(err, "Failed to load Transport Instance")
			return mperror.ErrInternalError
		}
		return nil
	}
	if tr.dbTr.Edges.GroupRecipient, err = tr.dbTr.Edges.GroupRecipientOrErr(); err != nil {
		if tr.dbTr, err = db.DbClient.DbTransportRecipients.Query().WithGroupRecipient().WithUserRecipient().Where(dbtransportrecipients.IDEQ(tr.dbTr.ID)).First(ctx); err != nil {
			tr.log.Error(err, "Failed to load Transport Instance")
			return mperror.ErrInternalError
		}
	}
	return nil
}

func (tr *TransportRecipient) SetUser(ctx context.Context, user interfaces.UserI) error {
	if tr.GetRecipientType(ctx) != interfaces.TransportRecipientTypeNotSet {
		tr.log.Error(mperror.ErrInvalidType, "Transport Recipient already has a recipient type set")
		return mperror.ErrTransportRecipientSet
	}
	tr.lock.Lock()
	defer tr.lock.Unlock()

	dbtmp, err := tr.dbTr.Update().SetUserRecipientID(user.GetID()).Save(ctx)
	if err != nil {
		tr.log.Error(err, "Failed to Update Transport Recipient User")
		return mperror.FilterErrors(err)
	}
	tr.dbTr = dbtmp
	return nil
}
func (tr *TransportRecipient) GetUser(ctx context.Context) (interfaces.UserI, error) {

	if tr.GetRecipientType(ctx) != interfaces.TransportRecipientTypeUser {
		tr.log.Error(mperror.ErrInvalidType, "Transport Recipient is not a user")
		return nil, mperror.ErrTransportRecipientSet
	}
	tr.lock.RLock()
	defer tr.lock.RUnlock()
	if tr.dbTr.Edges.UserRecipient.ID > 0 {
		return interfaces.GetUserService(ctx).GetByID(ctx, tr.dbTr.Edges.UserRecipient.ID)
	}
	return nil, mperror.ErrTransportRecipientGroupOrUserNotSet
}
func (tr *TransportRecipient) SetGroup(ctx context.Context, group interfaces.GroupI) error {

	if tr.GetRecipientType(ctx) != interfaces.TransportRecipientTypeNotSet {
		tr.log.Error(mperror.ErrInvalidType, "Transport Recipient already has a recipient type set")
		return mperror.ErrTransportRecipientSet
	}
	tr.lock.Lock()
	defer tr.lock.Unlock()
	dbtmp, err := tr.dbTr.Update().SetGroupRecipientID(group.GetID()).Save(ctx)
	if err != nil {
		tr.log.Error(err, "Failed to Update Transport Recipient Group")
		return mperror.FilterErrors(err)
	}
	tr.dbTr = dbtmp
	return nil
}
func (tr *TransportRecipient) GetGroup(ctx context.Context) (interfaces.GroupI, error) {

	if tr.GetRecipientType(ctx) != interfaces.TransportRecipientTypeGroup {
		tr.log.Error(mperror.ErrInvalidType, "Transport Recipient is not a Group")
		return nil, mperror.ErrTransportRecipientSet
	}
	tr.lock.RLock()
	defer tr.lock.RUnlock()

	if tr.dbTr.Edges.GroupRecipient.ID > 0 {
		return interfaces.GetGroupService(ctx).GetByID(ctx, tr.dbTr.Edges.GroupRecipient.ID)
	}
	return nil, mperror.ErrTransportRecipientGroupOrUserNotSet
}
func (tr *TransportRecipient) GetRecipientType(ctx context.Context) interfaces.TransportRecipientType {
	tr.lock.RLock()
	defer tr.lock.RUnlock()
	tr.loadEdges(ctx)
	if (tr.dbTr.Edges.UserRecipient != nil) && (tr.dbTr.Edges.UserRecipient.ID > 0) {
		return interfaces.TransportRecipientTypeUser
	}
	if (tr.dbTr.Edges.GroupRecipient != nil) && (tr.dbTr.Edges.GroupRecipient.ID > 0) {
		return interfaces.TransportRecipientTypeGroup
	}
	return interfaces.TransportRecipientTypeNotSet
}

func (tr *TransportRecipient) ProcessMessage(ctx context.Context, msg interfaces.MessageI) error {
	return mperror.FilterErrors(tr.tpi.Send(ctx, tr, msg))
}
