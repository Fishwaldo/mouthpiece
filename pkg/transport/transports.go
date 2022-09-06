package transport

import (
	"context"
	"errors"
	"sync"

	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"

	//	"github.com/Fishwaldo/mouthpiece/pkg/msg"
	//"gorm.io/gorm"
	"github.com/go-logr/logr"
)

type TransportInstanceStdConfigFields struct {
	ID                  uint `gorm:"primary_key"`
	Name                string
	Description         string
	TransportProviderID uint
}

type TransportRecipientConfigFields struct {
	ID                  uint `gorm:"primary_key"`
	Name                string
	Description         string
	UserID              uint
	GroupID             uint
	TransportInstanceID uint
	TransportProviderID uint
}

type TransportService struct {
	serviceContext          *interfaces.MPContext
	log                     logr.Logger
	transportInstancesMutex sync.Mutex
	transportInstances      map[uint]interfaces.TransportInstance
}

type transportServiceRecieptientMap struct {
	ID                   uint `gorm:"primary_key"`
	TransportInstance    uint
	TransportRecipientID uint
}

// Map of the Available Transports
var transports map[string]interfaces.TransportProvider

// Mutex to protect the transports map
var transportsMutex sync.Mutex

func RegisterTransportProvider(transport interfaces.TransportProvider) {
	transportsMutex.Lock()
	defer transportsMutex.Unlock()
	if transports == nil {
		transports = make(map[string]interfaces.TransportProvider)
	}
	transports[transport.GetName()] = transport
}

func GetTransportProvider(ctx context.Context, name string) (interfaces.TransportProvider, error) {
	transportsMutex.Lock()
	defer transportsMutex.Unlock()
	if t, ok := transports[name]; ok {
		return t, nil
	}
	return nil, errors.New("Transport Not Found")
}

func GetTransportProviders(ctx context.Context) []string {
	transportsMutex.Lock()
	defer transportsMutex.Unlock()
	var a []string
	for k := range transports {
		a = append(a, k)
	}
	return a
}

func NewTransportService() *TransportService {
	return &TransportService{
		log: log.Log.WithName("transports"),
		transportInstances: make(map[uint]interfaces.TransportInstance),
	}
}

func (tps *TransportService) Start(svcctx *interfaces.MPContext) error {

	db.Db.AutoMigrate(&transportServiceRecieptientMap{})

	transportsMutex.Lock()
	defer transportsMutex.Unlock()

	tps.serviceContext = svcctx
	for k, t := range transports {
		tps.log.Info("Starting Transport", "transport", k)
		if err := t.Start(svcctx, tps.log); err != nil {
			tps.log.Error(err, "Error Starting Transport", "transport", k)
		}
	}
	return nil
}

func (tps *TransportService) CreateTransportInstance(ctx context.Context, provider string, config any) (interfaces.TransportInstance, error) {
	if t, ok := transports[provider]; ok {
		if tpi, err := t.CreateInstance(ctx, config); err != nil {
			tps.log.Error(err, "Error Creating Transport Instance", "transport", provider)
			return nil, err
		} else {
			tps.transportInstancesMutex.Lock()
			defer tps.transportInstancesMutex.Unlock()
			tps.transportInstances[tpi.GetID()] = tpi
			return tpi, nil
		}
	}
	return nil, errors.New("transport Provider Not Found")
}
func (tps *TransportService) GetTransportInstances(context.Context) ([]uint, error) {
	var tpis []uint
	tps.transportInstancesMutex.Lock()
	defer tps.transportInstancesMutex.Unlock()
	for _, v := range tps.transportInstances {
		tpis = append(tpis, v.GetID())
	}
	return tpis, nil
}
func (tps *TransportService) GetTransportInstance(ctx context.Context, id uint) (interfaces.TransportInstance, error) {
	tps.transportInstancesMutex.Lock()
	defer tps.transportInstancesMutex.Unlock()
	if tpi, ok := tps.transportInstances[id]; ok {
		return tpi, nil
	}
	return nil, mperror.ErrTransportInstanceNotFound
}

func (tps *TransportService) GetTransportInstanceByName(ctx context.Context, name string) (interfaces.TransportInstance, error) {
	tps.transportInstancesMutex.Lock()
	defer tps.transportInstancesMutex.Unlock()
	for _, v := range tps.transportInstances {
		if v.GetName() == name {
			return v, nil
		}
	}
	return nil, mperror.ErrTransportInstanceNotFound
}
func (tps *TransportService) DeleteTransportInstance(ctx context.Context, tpi interfaces.TransportInstance) error {
	tps.transportInstancesMutex.Lock()
	defer tps.transportInstancesMutex.Unlock()
	delete(tps.transportInstances, tpi.GetID())

	var tprmap []transportServiceRecieptientMap
	if tx := db.Db.WithContext(ctx).Find(&tprmap, "transport_instance_id = ?", tpi.GetID()); tx.Error != nil {
		log.Log.Error(tx.Error, "Error Deleting Transport Instance", "transport", tpi.GetName())
		return tx.Error
	} else {
		for _, v := range tprmap {
			if tpr, err := tpi.GetTransportRecipientByTransportID(ctx, v.TransportRecipientID); err != nil {
				log.Log.Error(err, "Error Getting Transport Recieptients", "transport", tpi.GetName())
			} else {
				if err := tpi.DeleteTransportRecipient(ctx, tpr); err != nil {
					log.Log.Error(err, "Error Deleting Transport Recipient", "transport", tpi.GetName())
				}
				if tx := db.Db.WithContext(ctx).Delete(&v); tx.Error != nil {
					log.Log.Error(tx.Error, "Error Deleting Transport Recipient from Map", "transport", tpi.GetName())
				}
			}
		}
	}
	return nil
}

func (tps *TransportService) CreateGroupTransportRecipient(ctx context.Context, tpi interfaces.TransportInstance, grp interfaces.GroupI, config any) (interfaces.TransportRecipient, error) {
	if tpi == nil {
		return nil, mperror.ErrTransportInstanceNotFound
	}
	tx := db.Db.Begin()
	tprmap := &transportServiceRecieptientMap{
		TransportInstance: tpi.GetID(),
	}
	if res := tx.Create(tprmap); res.Error != nil {
		tx.Rollback()
		return nil, res.Error
	}
	if tpr, err := tpi.CreateGroupTransportRecipient(ctx, tprmap.ID, grp, config); err != nil {
		tx.Rollback()
		return nil, err
	} else {
		tprmap.TransportRecipientID = tpr.GetID()
		if res2 := tx.Save(tprmap); res2.Error != nil {
			tx.Rollback()
			return nil, res2.Error
		}
		tx.Commit()
		return tpr, nil
	}
}

func (tps *TransportService) CreateUserTransportRecipient(ctx context.Context, tpi interfaces.TransportInstance, user interfaces.UserI, config any) (interfaces.TransportRecipient, error) {
	if tpi == nil {
		return nil, mperror.ErrTransportInstanceNotFound
	}
	tx := db.Db.Begin()
	tprmap := &transportServiceRecieptientMap{
		TransportInstance: tpi.GetID(),
	}
	if res := tx.Create(tprmap); res.Error != nil {
		tx.Rollback()
		return nil, res.Error
	}
	if tpr, err := tpi.CreateUserTransportRecipient(ctx, tprmap.ID, user, config); err != nil {
		tx.Rollback()
		return nil, err
	} else {
		tprmap.TransportRecipientID = tpr.GetID()
		if res2 := tx.Save(tprmap); res2.Error != nil {
			tx.Rollback()
			return nil, res2.Error
		}
		if res := tx.Commit(); res.Error != nil {
			tx.Rollback()
			return nil, res.Error
		}
		return tpr, nil
	}
}
func (tps *TransportService) DeleteTransportRecipient(ctx context.Context, tpr interfaces.TransportRecipient) error {
	var tprmap transportServiceRecieptientMap
	if res := db.Db.WithContext(ctx).Find(&tprmap, "transport_recipient_id = ?", tpr.GetID()); res.Error != nil {
		return res.Error
	}
	tx := db.Db.Begin()
	if tpi, err := tps.GetTransportInstance(ctx, tprmap.TransportInstance); err != nil {
		log.Log.Error(err, "Error Getting Transport Instance", "transport_instance", tprmap.TransportInstance)
		tx.Rollback()
		return err
	} else {
		if err := tpi.DeleteTransportRecipient(ctx, tpr); err != nil {
			tx.Rollback()
			return err
		}
		if res := tx.WithContext(ctx).Delete(&tprmap); res != nil {
			tx.Rollback()
			return res.Error
		}
		if res := tx.Commit(); res.Error != nil {
			tx.Rollback()
			return res.Error
		}
		return nil
	}
}
func (tps *TransportService) GetTransportReciepientsForGroup(ctx context.Context, grp interfaces.GroupI) ([]uint, error) {
	var res []uint
	for _, v := range tps.transportInstances {
		if tprs, err := v.GetTransportReciepients(ctx); err != nil {
			log.Log.Error(err, "Error Getting Transport Recipients", "transport_instance", v.GetID())
		} else {
			for _, tpid := range tprs {
				if tpr, err := v.GetTransportRecipientByTransportID(ctx, tpid); err != nil {
					log.Log.Error(err, "Error Getting Transport Recipient", "transport_recipient", tpid)
				} else {
					if tpr.GetGroupID() == grp.GetID() {
						res = append(res, tpr.GetID())
					}
				}
			}
		}
	}
	return res, nil
}
func (tps *TransportService) GetTransportReciepientsForUser(ctx context.Context, user interfaces.UserI) ([]uint, error) {
	var res []uint
	for _, v := range tps.transportInstances {
		if tprs, err := v.GetTransportReciepients(ctx); err != nil {
			log.Log.Error(err, "Error Getting Transport Recipients", "transport_instance", v.GetID())
		} else {
			for _, tpid := range tprs {
				if tpr, err := v.GetTransportRecipientByTransportID(ctx, tpid); err != nil {
					log.Log.Error(err, "Error Getting Transport Recipient", "transport_recipient", tpid)
				} else {
					if tpr.GetUserID() == user.GetID() {
						res = append(res, tpr.GetID())
					}
				}
			}
		}
	}
	return res, nil
}
func (tps *TransportService) GetTransportReciepients(ctx context.Context) ([]uint, error) {
	var res []uint
	for _, v := range tps.transportInstances {
		if tprs, err := v.GetTransportReciepients(ctx); err != nil {
			log.Log.Error(err, "Error Getting Transport", "transport_instance", v.GetID())
		} else {
			res = append(res, tprs...)
		}
	}
	return res, nil
}

func (tps *TransportService) GetTransportReciepient(ctx context.Context, tid uint) (interfaces.TransportRecipient, error) {
	var tpr transportServiceRecieptientMap
	if res := db.Db.WithContext(ctx).Find(&tpr, "id = ?", tid); res.Error != nil {
		return nil, res.Error
	}
	if tpi, err := tps.GetTransportInstance(ctx, tpr.TransportInstance); err != nil {
		log.Log.Error(err, "Error Getting Transport Instance", "transport_instance", tpr.TransportInstance)
		return nil, err
	} else {
		return tpi.GetTransportRecipientByTransportID(ctx, tpr.TransportRecipientID)
	}
}
