package console

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"
	"github.com/Fishwaldo/mouthpiece/pkg/transport"

	"github.com/go-logr/logr"
)

type ConsoleTransportProvider struct {

}

type ConsoleTransportInstance struct {
	cfg *ConsoleConfig
	log logr.Logger
}

type ConsoleConfig struct {

}

func (c *ConsoleConfig) AsJSON() (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *ConsoleConfig) FromJSON(data string) error {
	return json.Unmarshal([]byte(data), c)
}

type ConsoleRecipientConfig struct {

}

func (c *ConsoleRecipientConfig) AsJSON() (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *ConsoleRecipientConfig) FromJSON(data string) error {
	return json.Unmarshal([]byte(data), c)
}



func init() {
	tp := NewConsoleTransportProvider()
	transport.RegisterTransportProvider(tp)
}

func NewConsoleTransportProvider() interfaces.TransportProvider {
	return &ConsoleTransportProvider{}
}

func (t *ConsoleTransportProvider) GetName() string {
	return "console"
}

func (t *ConsoleTransportProvider) CreateInstance(ctx context.Context, logger logr.Logger, name string, config interfaces.MarshableConfigI) (interfaces.TransportInstanceImpl, error) {
	log.Log.Info("Creating Console Transport Instance", "name", t.GetName())
	cticonfig, ok := config.(*ConsoleConfig)
	if !ok {
		log.Log.Error(mperror.ErrTransportConfigInvalid, "Invalid Config", "Type", fmt.Sprintf("%T", config))
		return nil, mperror.ErrTransportConfigInvalid
	}

	tpi := &ConsoleTransportInstance{
		cfg: cticonfig,
		log: logger.WithName("ConsoleTransportInstance").WithValues("name", name),
	}
	tpi.log.Info("Creating Console Transport Instance")
	return tpi, nil
}

func (t *ConsoleTransportProvider) LoadConfigFromJSON(ctx context.Context, data string) (interfaces.MarshableConfigI, error) {
	var cfg ConsoleConfig
	err := cfg.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (t *ConsoleTransportInstance) Init(context.Context) error {
	return nil
}

func (t *ConsoleTransportInstance) SetConfig(ctx context.Context, config interfaces.MarshableConfigI) error {
	cfg, ok := config.(*ConsoleConfig)
	if !ok {
		t.log.Error(mperror.ErrTransportConfigInvalid, "Invalid Config", "Type", fmt.Sprintf("%T", config))
		return mperror.ErrTransportConfigInvalid
	}
	t.cfg = cfg
	return nil
}

func (t *ConsoleTransportInstance) GetConfig(ctx context.Context) interfaces.MarshableConfigI {
	return t.cfg
}

func (tpi *ConsoleTransportInstance) Start(context.Context) error {
	tpi.log.Info("Starting Transport Instance")
	return nil
}

func (tpi *ConsoleTransportInstance) Stop(context.Context) error {
	tpi.log.Info("Stopping Transport Instance")
	return nil
}

func (tpi *ConsoleTransportInstance) ValidateTransportRecipientConfig(ctx context.Context, config interfaces.MarshableConfigI) error {
	if _, ok := config.(*ConsoleRecipientConfig); !ok {
		tpi.log.Error(mperror.ErrTransportConfigInvalid, "Invalid Config", "Type", fmt.Sprintf("%T", config))
		return mperror.ErrTransportConfigInvalid
	}
	return nil
}

func (tpi *ConsoleTransportInstance) LoadTransportReciepientConfig(ctx context.Context, config string) (interfaces.MarshableConfigI, error) {
	cfg := &ConsoleRecipientConfig{}
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


func (tpi *ConsoleTransportInstance) Send(ctx context.Context, tpr interfaces.TransportRecipient, msg interfaces.MessageI) error {
	fmt.Println("=========================================================")
	fmt.Printf("Message: %s\n", msg.String())
	if tpr.GetRecipientType(ctx) == interfaces.TransportRecipientTypeUser {
		user, err := tpr.GetUser(ctx)
		if err != nil {
			tpi.log.Error(err, "Error Getting User")
			return err
		}
		fmt.Printf("\tFor User: %s\n", user.GetEmail())
	} else if tpr.GetRecipientType(ctx) == interfaces.TransportRecipientTypeGroup {
		group, err := tpr.GetGroup(ctx)
		if err != nil {
			tpi.log.Error(err, "Error Getting Group")
			return err
		}
		fmt.Printf("\tFor Group: %s\n", group.GetName())
	} else {
		fmt.Printf("\tFor: NOT SET!\n")
	}

	flds, err := msg.GetFields(ctx)
	if err == nil {
		fmt.Printf("\tFields:\n")
		for k, v := range flds {
			fmt.Printf("\t\tField: %s = %s\n", k, v)
		}
	}
	metadata, err := msg.GetMetadataFields(ctx)
	if err == nil {
		fmt.Printf("\tMetaData Fields:\n")
		for k, v := range metadata {
			fmt.Printf("\t\tMetaData: %s = %+v\n", k, v)
		}
	}
	fmt.Printf("\tTransport Recipient: %s\n", tpr.GetName())
	fmt.Println("=========================================================")
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

var _ interfaces.TransportInstanceImpl = (*ConsoleTransportInstance)(nil)
