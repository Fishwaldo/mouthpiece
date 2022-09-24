package noop_test

import (
	"testing"

	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	mptest "github.com/Fishwaldo/mouthpiece/pkg/mocks"
	"github.com/Fishwaldo/mouthpiece/pkg/transport"

	"github.com/Fishwaldo/mouthpiece/pkg/transport/noop"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestNoop(t *testing.T) {
	mptest.Tst = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Noop Suite")
}

var _ = mptest.DBBeforeSuite
var _ = mptest.DBBeforeSuite

var globalTPP interfaces.TransportProvider
var globalCfg interfaces.MarshableConfigI
var globalTPI interfaces.TransportInstanceImpl
var globalRcpCfg interfaces.MarshableConfigI

var _ = Describe("NoOp", func() {
	Context("Provider", func() {
		It("SHould Register with the Provider", func() {
			var err error
			globalTPP, err = transport.GetTransportProvider(mptest.Ctx, "noop")
			Expect(err).To(BeNil())
			Expect(globalTPP).ToNot(BeNil())
			Expect(globalTPP.GetName()).To(Equal("noop"))
		})
		It("Should Create a New noop Config Instance", func() {
			var err error
			globalCfg, err = globalTPP.LoadConfigFromJSON(mptest.Ctx, "{}")
			Expect(err).To(BeNil())
			Expect(globalCfg).ToNot(BeNil())
			Expect(globalCfg).To(BeAssignableToTypeOf(&noop.NoOpConfig{}))
		})
		It("Should Create a New Console Instance Implementation", func() {
			var err error
			globalTPI, err = globalTPP.CreateInstance(mptest.Ctx, mptest.GlobalLogger, "consoleinstance", globalCfg)
			Expect(err).To(BeNil())
			Expect(globalTPI).ToNot(BeNil())
			Expect(globalTPI).To(BeAssignableToTypeOf(&noop.NoOpTransportInstance{}))
		})
		It("Should Return a Valid Config", func() {
			cfg := globalTPI.GetConfig(mptest.Ctx)
			Expect(cfg).ToNot(BeNil())
			Expect(cfg).To(BeAssignableToTypeOf(&noop.NoOpConfig{}))
		})
		It("Load a Recipient Config", func() {
			var err error
			globalRcpCfg, err = globalTPI.LoadTransportReciepientConfig(mptest.Ctx, "{}")
			Expect(err).To(BeNil())
			Expect(globalRcpCfg).ToNot(BeNil())
			Expect(globalRcpCfg).To(BeAssignableToTypeOf(&noop.NoOpRecipientConfig{}))
		})
		It("Should Validate a Config", func() {
			err := globalTPI.ValidateTransportRecipientConfig(mptest.Ctx, globalRcpCfg)
			Expect(err).To(BeNil())
		})
		It("Should Process a Message", func() {
			mptest.TestTransportSend(globalTPI)
		})
		It("Should Stop", func() {
			err := globalTPI.Stop(mptest.Ctx)
			Expect(err).To(BeNil())
		})
	})
})
