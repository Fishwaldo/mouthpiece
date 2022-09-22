package noop

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"

	"github.com/Fishwaldo/mouthpiece/pkg/transport"
	"github.com/go-logr/logr"
	"github.com/google/uuid"

)

var Messages map[uuid.UUID]bool


type NoOpTransportProvider struct {

}

type NoOpTransportInstance struct {
	cfg *NoOpConfig
	log logr.Logger
}

type NoOpConfig struct {

}

func (c *NoOpConfig) AsJSON() (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *NoOpConfig) FromJSON(data string) error {
	return json.Unmarshal([]byte(data), c)
}

type NoOpRecipientConfig struct {

}

func (c *NoOpRecipientConfig) AsJSON() (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *NoOpRecipientConfig) FromJSON(data string) error {
	return json.Unmarshal([]byte(data), c)
}



func init() {
	tp := NewConsoleTransportProvider()
	transport.RegisterTransportProvider(tp)
}

func NewConsoleTransportProvider() interfaces.TransportProvider {
	return &NoOpTransportProvider{}
}

func (t *NoOpTransportProvider) GetName() string {
	return "noop"
}

func (t *NoOpTransportProvider) CreateInstance(ctx context.Context, logger logr.Logger, name string, config interfaces.MarshableConfigI) (interfaces.TransportInstanceImpl, error) {
	log.Log.Info("Creating NoOp Transport Instance", "name", t.GetName())
	cticonfig, ok := config.(*NoOpConfig)
	if !ok {
		log.Log.Error(mperror.ErrTransportConfigInvalid, "Invalid Config", "Type", fmt.Sprintf("%T", config))
		return nil, mperror.ErrTransportConfigInvalid
	}

	tpi := &NoOpTransportInstance{
		cfg: cticonfig,
		log: logger.WithName("NoOpTransportInstance").WithValues("name", name),
		
	}
	Messages = make(map[uuid.UUID]bool)
	tpi.log.Info("Creating NoOp Transport Instance")
	return tpi, nil
}

func (t *NoOpTransportProvider) LoadConfigFromJSON(ctx context.Context, data string) (interfaces.MarshableConfigI, error) {
	var cfg NoOpConfig
	err := cfg.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (t *NoOpTransportInstance) Init(context.Context) error {
	return nil
}

func (t *NoOpTransportInstance) SetConfig(ctx context.Context, config interfaces.MarshableConfigI) error {
	cfg, ok := config.(*NoOpConfig)
	if !ok {
		t.log.Error(mperror.ErrTransportConfigInvalid, "Invalid Config", "Type", fmt.Sprintf("%T", config))
		return mperror.ErrTransportConfigInvalid
	}
	t.cfg = cfg
	return nil
}

func (t *NoOpTransportInstance) GetConfig(ctx context.Context) interfaces.MarshableConfigI {
	return t.cfg
}

func (tpi *NoOpTransportInstance) Start(context.Context) error {
	tpi.log.Info("Starting Transport Instance")
	return nil
}

func (tpi *NoOpTransportInstance) Stop(context.Context) error {
	tpi.log.Info("Stopping Transport Instance")
	return nil
}

func (tpi *NoOpTransportInstance) ValidateTransportRecipientConfig(ctx context.Context, config interfaces.MarshableConfigI) error {
	if _, ok := config.(*NoOpRecipientConfig); !ok  {
		tpi.log.Error(mperror.ErrTransportConfigInvalid, "Invalid Config", "Type", fmt.Sprintf("%T", config))
		return mperror.ErrTransportConfigInvalid
	}
	return nil
}

func (tpi *NoOpTransportInstance) LoadTransportReciepientConfig(ctx context.Context, config string) (interfaces.MarshableConfigI, error) {
	cfg := &NoOpRecipientConfig{}
	if err := cfg.FromJSON(config); err != nil {
		tpi.log.Error(err, "Error Unmarshalling Config")
		return nil, err
	}
	if err := tpi.ValidateTransportRecipientConfig(ctx, cfg); err != nil {
		tpi.log.Error(err, "Error Validating Config")
		return nil, err
	}
	return cfg, nil
}


func (tpi *NoOpTransportInstance) Send(ctx context.Context, tpr interfaces.TransportRecipient, msg interfaces.MessageI) error {
	tpi.log.Info("Sending Message", "Message", msg)
	if _, ok := Messages[msg.GetID()]; ok {
		tpi.log.Info("Message already sent", "Message", msg)
	} else {
		Messages[msg.GetID()] = true
	}
	return nil
}

// func (tpr ConsoleTransportRecipient) ProcessGroupMessage(ctx context.Context, msg msg.Message) error {
// 	fmt.Println("=========================================================")
// 	fmt.Printf("Group Message: %s\n", msg.Body.Message)
// 	fmt.Println("=========================================================")
// 	return nil
// }

// func (tpr ConsoleTransportRecipient) ProcessMessage(ctx context.Context, msg msg.Message) error {
// 	fmt.Println("=========================================================")
// 	fmt.Printf("Message: %s\n", msg.Body.Message)
// 	fmt.Println("=========================================================")
// 	return nil
// }

var _ interfaces.TransportInstanceImpl = (*NoOpTransportInstance)(nil)
