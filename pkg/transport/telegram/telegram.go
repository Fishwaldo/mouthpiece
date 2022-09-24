package telegram

import (
	"context"
	"fmt"
	"encoding/json"

	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/transport"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"

	"github.com/go-logr/logr"
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	"github.com/mymmrac/telego/telegoutil"
)

type TelegramTransportProvider struct {

}


type TelegramTransportInstance struct {
	cfg *TelegramConfig
	log logr.Logger
	bot *telego.Bot
	updates <-chan telego.Update
	handler *telegohandler.BotHandler
}

type TelegramConfig struct {
	TelegramToken string
}

func (c *TelegramConfig) AsJSON() (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *TelegramConfig) FromJSON(data string) error {
	return json.Unmarshal([]byte(data), c)
}

type TelegramRecipientConfig struct {
	ChatID int64
}

func (c *TelegramRecipientConfig) AsJSON() (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *TelegramRecipientConfig) FromJSON(data string) error {
	return json.Unmarshal([]byte(data), c)
}


func init() {
	tp := NewTelegramTransportProvider()
	transport.RegisterTransportProvider(tp)

}


func NewTelegramTransportProvider() interfaces.TransportProvider {
	return &TelegramTransportProvider{}
}

func (t *TelegramTransportProvider) GetName() string {
	return "telegram"
}

func (t *TelegramTransportProvider) CreateInstance(ctx context.Context, logger logr.Logger, name string, config interfaces.MarshableConfigI) (interfaces.TransportInstanceImpl, error) {
	log.Log.Info("Creating Telegram Transport Instance", "name", t.GetName())
	tticonfig, ok := config.(*TelegramConfig)
	if !ok {
		log.Log.Error(mperror.ErrTransportConfigInvalid, "Invalid Config", "Type", fmt.Sprintf("%T", config))
		return nil, mperror.ErrTransportConfigInvalid
	}

	tpi := &TelegramTransportInstance{
		cfg: tticonfig,
		log: logger.WithName("TelegramTransportInstance").WithValues("name", name),
	}
	tpi.log.Info("Creating Telegram Transport Instance")
	return tpi, nil
}

func (t *TelegramTransportProvider) LoadConfigFromJSON(ctx context.Context, data string) (interfaces.MarshableConfigI, error) {
	var cfg TelegramConfig
	err := cfg.FromJSON(data)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}


func (t *TelegramTransportInstance) Init(context.Context) error {
	return nil
}

func (t *TelegramTransportInstance) SetConfig(ctx context.Context, config interfaces.MarshableConfigI) error {
	cfg, ok := config.(*TelegramConfig)
	if !ok {
		t.log.Error(mperror.ErrTransportConfigInvalid, "Invalid Config", "Type", fmt.Sprintf("%T", config))
		return mperror.ErrTransportConfigInvalid
	}
	t.cfg = cfg
	return nil
}

func (t *TelegramTransportInstance) GetConfig(ctx context.Context) interfaces.MarshableConfigI {
	return t.cfg
} 

func (tpi *TelegramTransportInstance) Start(context.Context) error {
	tpi.log.Info("Starting Telegram Transport Instance")

	tlogger := &telegramLogger {
		log: tpi.log,
		token: tpi.cfg.TelegramToken,
	}


	var err error
	tpi.bot, err = telego.NewBot(tpi.cfg.TelegramToken, telego.WithLogger(tlogger), telego.WithHealthCheck())
	if err != nil {
		log.Log.Error(err, "Telegram Transport Failed to Start")
		return err
	}

	// Get updates channel
	tpi.updates, err = tpi.bot.UpdatesViaLongPulling(nil)
	if err != nil {
		log.Log.Error(err, "Telegram Transport Failed to St Updates via Long Pull")
		return err
	}

	// Create bot handler and specify from where to get updates
	tpi.handler, err = telegohandler.NewBotHandler(tpi.bot, tpi.updates)
	if err != nil {
		log.Log.Error(err, "Telegram Transport Failed to Create Bot Handler")
		return err
	}
	// Register new handler with match on command `/start`
	tpi.handler.Handle(func(bot *telego.Bot, update telego.Update) {
		// Send message
		_, _ = bot.SendMessage(telegoutil.Message(
			telegoutil.ID(update.Message.Chat.ID),
			fmt.Sprintf("Hello %s!", update.Message.From.FirstName),
		))
	}, telegohandler.CommandEqual("start"))

	// Register new handler with match on any command
	// Handlers will match only once and in order of registration, so this handler will be called on any command
	// except `/start` command
	tpi.handler.Handle(func(bot *telego.Bot, update telego.Update) {
		// Send message
		_, _ = bot.SendMessage(telegoutil.Message(
			telegoutil.ID(update.Message.Chat.ID),
			"Unknown command, use /start",
		))
	}, telegohandler.AnyCommand())

	// Start handling updates
	go tpi.handler.Start()

	log.Log.Info("Telegram Transport Started")

	return nil
}

func (tpi *TelegramTransportInstance) Stop(context.Context) error {
	tpi.log.Info("Stopping Transport Instance")
	tpi.handler.Stop()
	tpi.bot.StopLongPulling()

	return nil
}

func (tpi *TelegramTransportInstance) ValidateTransportRecipientConfig(ctx context.Context, config interfaces.MarshableConfigI) error {
	if _, ok := config.(*TelegramRecipientConfig); !ok {
		tpi.log.Error(mperror.ErrTransportConfigInvalid, "Invalid Config", "Type", fmt.Sprintf("%T", config))
		return mperror.ErrTransportConfigInvalid
	}
	return nil
}

func (tpi *TelegramTransportInstance) LoadTransportReciepientConfig(ctx context.Context, config string) (interfaces.MarshableConfigI, error) {
	cfg := &TelegramRecipientConfig{}
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

func (tpi *TelegramTransportInstance) Send(ctx context.Context, tpr interfaces.TransportRecipient, msg interfaces.MessageI) error {
	trcfg, err := tpr.GetConfig()
	if err != nil {
		tpi.log.Error(err, "Error Getting Recipient Config")
		return err
	}
	cfg, ok := trcfg.(*TelegramRecipientConfig)
	if !ok {
		tpi.log.Error(mperror.ErrTransportConfigInvalid, "Invalid Config", "Type", fmt.Sprintf("%T", trcfg))
		return mperror.ErrTransportConfigInvalid
	}
	tmsg, err := tpi.bot.SendMessage(
		telegoutil.Message(
			telegoutil.ID(cfg.ChatID),
			msg.GetMessage(),
		),
	)
	if err != nil {
		tpi.log.Error(err, "Error Sending Message")
		return err
	}
	tpi.log.Info("Message Sent", "Message", tmsg)

	return nil
}

