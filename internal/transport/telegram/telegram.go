package telegram

import (
	"fmt"
	//	"os"

	. "github.com/Fishwaldo/mouthpiece/internal/log"
	"github.com/Fishwaldo/mouthpiece/internal/message"
	"github.com/Fishwaldo/mouthpiece/internal/transport"

	"github.com/spf13/viper"

	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	"github.com/mymmrac/telego/telegoutil"
)

type TelegramTransport struct {
}

func init() {
	viper.SetDefault("transport.telegram.enabled", false)
	tp := NewTGTransport()
	transport.RegisterTransport(tp)
}

func NewTGTransport() transport.ITransport {
	return &TelegramTransport{}
}

func (t TelegramTransport) GetName() string {
	return "telegram"
}

func (t TelegramTransport) Start() {
	bot, err := telego.NewBot(viper.GetString("transport.telegram.token"), telego.WithDefaultLogger(false, true))
	if err != nil {
		Log.Error(err, "Telegram Transport Failed to Start")
	}

	// Get updates channel
	updates, _ := bot.UpdatesViaLongPulling(nil)
	//defer bot.StopLongPulling()

	// Create bot handler and specify from where to get updates
	bh, _ := telegohandler.NewBotHandler(bot, updates)

	// Register new handler with match on command `/start`
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// Send message
		_, _ = bot.SendMessage(telegoutil.Message(
			telegoutil.ID(update.Message.Chat.ID),
			fmt.Sprintf("Hello %s!", update.Message.From.FirstName),
		))
	}, telegohandler.CommandEqual("start"))

	// Register new handler with match on any command
	// Handlers will match only once and in order of registration, so this handler will be called on any command
	// except `/start` command
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// Send message
		_, _ = bot.SendMessage(telegoutil.Message(
			telegoutil.ID(update.Message.Chat.ID),
			"Unknown command, use /start",
		))
	}, telegohandler.AnyCommand())

	// Start handling updates
	go bh.Start()

	// Stop handling updates
	//defer bh.Stop()
	Log.Info("Transport Started", "name", t.GetName())
}

func (t TelegramTransport) NewTransportConfig() {
	//	user.TransportConfigs = append(user.TransportConfigs, mouthpiece.TransportConfig{
	//		Transport: t.GetName(),
	//		Config: user.Username,
	//	})
}

func (t TelegramTransport) SendMessage(config transport.TransportConfig, msg msg.Message) (err error) {
	fmt.Println("=========================================================")
	fmt.Printf("Message: %s\n", msg.Body.Message)
	fmt.Println("=========================================================")
	transport.UpdateTransportStatus(t, msg, "sent")
	return nil
}
