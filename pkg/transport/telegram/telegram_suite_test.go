package telegram_test

import (
	"os"
	"testing"
	"strconv"

	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	mptest "github.com/Fishwaldo/mouthpiece/pkg/mocks"

	"github.com/Fishwaldo/mouthpiece/pkg/transport"
	"github.com/Fishwaldo/mouthpiece/pkg/transport/telegram"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTelegram(t *testing.T) {
	mptest.Tst = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Telegram Suite")
}

var _ = mptest.DBBeforeSuite
var _ = mptest.DBBeforeSuite

var globalTPP interfaces.TransportInstanceImpl
var globalTCfg *telegram.TelegramConfig

var _ = Describe("Telegram", func() {
	Context("Provider", func() {
		It("SHould Register with the Provider", func() {
			provider, err := transport.GetTransportProvider(mptest.Ctx, "telegram")
			Expect(err).To(BeNil())
			Expect(provider).ToNot(BeNil())
			Expect(provider.GetName()).To(Equal("telegram"))
		})
		It("Should Create a new Telegram instance", func() {
			var err error
			provider, err := transport.GetTransportProvider(mptest.Ctx, "telegram")
			Expect(err).To(BeNil())
			Expect(provider).ToNot(BeNil())
			Expect(provider.GetName()).To(Equal("telegram"))
			var cfg interfaces.MarshableConfigI
			cfg, err = provider.LoadConfigFromJSON(mptest.Ctx, "{}")
			Expect(err).To(BeNil())
			Expect(cfg).ToNot(BeNil())
			Expect(cfg).To(BeAssignableToTypeOf(&telegram.TelegramConfig{}))
			globalTCfg = cfg.(*telegram.TelegramConfig)
			globalTPP, err = provider.CreateInstance(mptest.Ctx, mptest.GlobalLogger, "telegraminstance", cfg)
			Expect(err).To(BeNil())
			Expect(globalTPP).ToNot(BeNil())
			Expect(globalTPP).To(BeAssignableToTypeOf(&telegram.TelegramTransportInstance{}))
		})
	})
	Context("Run", func() {
		It("Should Run", func() {
			ttoken := os.Getenv("TELEGRAM_TOKEN")
			if len(ttoken) == 0 {
				Skip("No Telegram Chat Enviroment Vars")
			} else {
				globalTCfg.TelegramToken = ttoken
			}
			err := globalTPP.SetConfig(mptest.Ctx, globalTCfg)
			Expect(err).To(BeNil())
			err = globalTPP.Start(mptest.Ctx)		
			Expect(err).To(BeNil())
//			time.Sleep(10 * time.Second)	
		})
	})
	Context("Send", func() {
		It("It Should Send a Message", func() {
			tchat := os.Getenv("TELEGRAM_CHAT")
			ttoken := os.Getenv("TELEGRAM_TOKEN")
			if (len(tchat) == 0) || (len(ttoken) == 0) {
				Skip("No Telegram Chat Enviroment Vars")
			} 
			ctrl := gomock.NewController(GinkgoT())
			msg := mptest.NewMockMessageI(ctrl)
			tpr := mptest.NewMockTransportRecipient(ctrl)
			id, _ := strconv.Atoi(tchat)
			trcfg := &telegram.TelegramRecipientConfig{
				ChatID: int64(id),
			}

			tpr.EXPECT().GetConfig().Return(trcfg, nil).AnyTimes()

			msg.EXPECT().GetMessage().Return("Hello World - This is my birthday").AnyTimes()

			err := globalTPP.Send(mptest.Ctx, tpr, msg)
			Expect(err).To(BeNil())
			//time.Sleep(11 * time.Second)

		})
	})
	Context("Stop", func() {
		It("Should Stop", func() {
			tchat := os.Getenv("TELEGRAM_CHAT")
			ttoken := os.Getenv("TELEGRAM_TOKEN")
			if (len(tchat) == 0) || (len(ttoken) == 0) {
				Skip("No Telegram Chat Enviroment Vars")
			} 
			err := globalTPP.Stop(mptest.Ctx)
			Expect(err).To(BeNil())
		})
	})
})