package console

import (
	"context"
	"fmt"
	"sync"

	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/msg"
	"github.com/Fishwaldo/mouthpiece/pkg/transport"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"

	"github.com/go-logr/logr"
)

type ConsoleTransportProvider struct {
	svcctx *interfaces.MPContext
	instancelock sync.Mutex
	instances    map[string]*ConsoleTransportInstance
	log          logr.Logger
}

type ConsoleInstanceConfig struct {
	transport.TransportInstanceStdConfigFields
}

type ConsoleTransportInstance struct {
	Config ConsoleInstanceConfig `gorm:"embedded"` 
	log  logr.Logger
	recptlock  sync.Mutex
	recipients map[uint]*ConsoleTransportRecipient

}

type ConsoleRecipientConfig struct {
	transport.TransportRecipientConfigFields
}

type ConsoleTransportRecipient struct {
	Config ConsoleRecipientConfig `gorm:"embedded"`
	instance   *ConsoleTransportInstance
}


func init() {
	tp := NewStdoutTransportProvider()
	transport.RegisterTransportProvider(tp)
}

func NewStdoutTransportProvider() interfaces.TransportProvider {
	return &ConsoleTransportProvider{
		instances: make(map[string]*ConsoleTransportInstance),
	}
}

func (t *ConsoleTransportProvider) GetName() string {
	return "console"
}

func (t *ConsoleTransportProvider) Start(ctx *interfaces.MPContext, log logr.Logger) error {
	t.instancelock.Lock()
	defer t.instancelock.Unlock()

	t.svcctx = ctx
	t.log = log.WithName("ConsoleTransportProvider")


	db.Db.AutoMigrate(&ConsoleTransportInstance{}, &ConsoleTransportRecipient{})

	var tpi []ConsoleTransportInstance
	if tx := db.Db.Find(&tpi); tx.Error != nil {
		t.log.Error(tx.Error, "Error Loading Transport Instances", "name", t.GetName())
		return tx.Error
	}
	for _, v := range tpi {
		v.log = t.log.WithName(v.GetName())
		v.recipients = make(map[uint]*ConsoleTransportRecipient)
		if err := v.Start(t.svcctx); err != nil {
			t.log.Error(err, "Error Starting Transport Instance", "name", v.GetName())
		} else { 
			t.instances[v.GetName()] = &v
		}
	}
	t.log.Info("Transport Started", "name", t.GetName())
	return nil
}

func (t *ConsoleTransportProvider) CreateInstance(ctx context.Context, config any) (interfaces.TransportInstance, error) {
	t.log.Info("Creating Transport Instance", "name", t.GetName())
	cticonfig, ok := config.(ConsoleInstanceConfig); 
	if ok == false {
		return nil, fmt.Errorf("Invalid Config - Cannot Cast to ConsoleInstanceConfig")
	}

	tpi := &ConsoleTransportInstance{
		Config: cticonfig,
		log:  t.log.WithName(cticonfig.Name),
		recipients: make(map[uint]*ConsoleTransportRecipient),
	}
	if tx := db.Db.WithContext(ctx).Create(tpi); tx.Error != nil {
		t.log.Error(tx.Error, "Error Creating Transport Instance", "name", tpi.GetName())
		return nil, tx.Error
	}
	if err := tpi.Start(t.svcctx); err != nil {
		t.log.Error(err, "Error Starting Transport Instance", "name", tpi.GetName())
		return nil, err
	}
	t.instancelock.Lock()
	defer t.instancelock.Unlock()
	t.instances[tpi.GetName()] = tpi
	return tpi, nil
}

func (t *ConsoleTransportProvider) DeleteInstance(ctx context.Context, tpi interfaces.TransportInstance) error {
	t.log.Info("Deleting Transport Instance", "name", tpi.GetName())
	if tx := db.Db.WithContext(ctx).Delete(tpi); tx.Error != nil {
		t.log.Error(tx.Error, "Error Deleting Transport Instance", "name", tpi.GetName())
		return tx.Error
	}
	if err := tpi.Stop(ctx); err != nil {
		t.log.Error(err, "Error Stopping Transport Instance", "name", tpi.GetName())
	}
	t.instancelock.Lock()
	defer t.instancelock.Unlock()
	delete(t.instances, tpi.GetName())
	return nil
}

func (t *ConsoleTransportProvider) GetInstance(ctx context.Context, ID uint) (interfaces.TransportInstance, error) {
	t.instancelock.Lock()
	defer t.instancelock.Unlock()
	for _, v := range t.instances {
		if v.GetID() == ID {
			return v, nil
		}
	}
	return nil, fmt.Errorf("transport Instance Not Found")
}

func (t *ConsoleTransportProvider) GetInstanceByName(ctx context.Context, name string) (interfaces.TransportInstance, error) {
	t.instancelock.Lock()
	defer t.instancelock.Unlock()
	if v, ok := t.instances[name]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("transport Instance Not Found")
}

func (t *ConsoleTransportProvider) GetInstances(ctx context.Context) ([]uint, error) {
	t.instancelock.Lock()
	defer t.instancelock.Unlock()
	var IDs []uint
	for _, v := range t.instances {
		IDs = append(IDs, v.GetID())
	}
	return IDs, nil
}



func (tpi *ConsoleTransportInstance) GetName() string {
	return tpi.Config.Name
}

func (tpi *ConsoleTransportInstance) GetID() uint {
	return tpi.Config.ID
}

func (tpi *ConsoleTransportInstance) Start(*interfaces.MPContext) error {
	tpi.log.Info("Starting Transport Instance", "name", tpi.GetName())
	var tprs []ConsoleTransportRecipient
	if tx := db.Db.Find(&tprs); tx.Error != nil {
		tpi.log.Error(tx.Error, "Error Loading Transport Recipients", "name", tpi.GetName())
		return tx.Error
	}
	for _, v := range tprs {
		v.instance = tpi
		tpi.recipients[v.GetID()] = &v
	}
	return nil
}

func (tpi *ConsoleTransportInstance) Stop (context.Context) error {
	return nil
}

func (tpi *ConsoleTransportInstance) CreateGroupTransportRecipient(ctx context.Context, assignedID uint, grp interfaces.GroupI, config any) (interfaces.TransportRecipient, error) {
	tpi.log.Info("Creating Transport Reciepient", "name", tpi.GetName(), "group", grp.GetID())
	return tpi.createTransportReciepient(ctx, assignedID, nil, grp, config)
}

func (tpi *ConsoleTransportInstance) CreateUserTransportRecipient(ctx context.Context, assignedID uint, user interfaces.UserI, config any) (interfaces.TransportRecipient, error) {
	tpi.log.Info("Creating Transport Reciepient", "name", tpi.GetName(), "user", user.GetID())
	return tpi.createTransportReciepient(ctx, assignedID, user, nil, config)
}

func (tpi *ConsoleTransportInstance) createTransportReciepient(ctx context.Context, assignedID uint, user interfaces.UserI, grp interfaces.GroupI, config any) (interfaces.TransportRecipient, error) {
	if assignedID == 0 {
		return nil, fmt.Errorf("assignedID Cannot Be Zero")
	}

	tpr := ConsoleTransportRecipient{
		instance: tpi,
	}
	tpr.Config.ID = assignedID
	if user != nil {
		tpr.Config.UserID = user.GetID()
	}
	if grp != nil {
		tpr.Config.GroupID = grp.GetID()
	}

	if tx := db.Db.WithContext(ctx).Create(&tpr); tx.Error != nil {
		tpi.log.Error(tx.Error, "Error Creating Transport Recipient", "ID", tpr.GetID())
		return nil, tx.Error
	}
	tpi.recptlock.Lock()
	defer tpi.recptlock.Unlock()
	tpi.recipients[tpr.GetID()] = &tpr
	return tpr, nil
}

func (tpi *ConsoleTransportInstance) GetTransportReciepients(ctx context.Context)  ([]uint, error) {
	tpi.recptlock.Lock()
	defer tpi.recptlock.Unlock()
	var IDs []uint
	for _, v := range tpi.recipients {
		IDs = append(IDs, v.GetID())
	}
	return IDs, nil
}

func (tpi *ConsoleTransportInstance) GetTransportRecipientByTransportID(ctx context.Context, ID uint) (interfaces.TransportRecipient, error) {
	tpi.recptlock.Lock()
	defer tpi.recptlock.Unlock()
	if tpr, ok := tpi.recipients[ID]; ok {
		return tpr, nil
	}
	return nil, nil
}


func (tpi *ConsoleTransportInstance) DeleteTransportRecipient(ctx context.Context, tpr interfaces.TransportRecipient) error {
	tpi.recptlock.Lock()
	defer tpi.recptlock.Unlock()
	for k, v := range tpi.recipients {
		if v.GetID() == tpr.GetID() {
			delete(tpi.recipients, k)
			if tx := db.Db.WithContext(ctx).Delete(tpr); tx.Error != nil {
				tpi.log.Error(tx.Error, "Error Deleting Transport Recipient", "ID", tpr.GetID())
				return tx.Error
			}
			return nil
		}
	}
	return mperror.ErrTransportReciepiantNotFound
}

func (tpr ConsoleTransportRecipient) GetID() uint {
	return tpr.Config.ID
}

func (tpr ConsoleTransportRecipient) GetGroupID() uint {
	return tpr.Config.GroupID
}

func (tpr ConsoleTransportRecipient) GetUserID() uint {
	return tpr.Config.UserID
}

func (tpr ConsoleTransportRecipient) ProcessGroupMessage(ctx context.Context, msg msg.Message) error {
	fmt.Println("=========================================================")
	fmt.Printf("Group Message: %s\n", msg.Body.Message)
	fmt.Println("=========================================================")
	return nil
}

func (tpr ConsoleTransportRecipient) ProcessMessage(ctx context.Context, msg msg.Message) error {
	fmt.Println("=========================================================")
	fmt.Printf("Message: %s\n", msg.Body.Message)
	fmt.Println("=========================================================")
	return nil
}
