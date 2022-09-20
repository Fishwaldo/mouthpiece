package severity_test

import (
	"context"
	"testing"
	"encoding/json"

	"github.com/Fishwaldo/mouthpiece/pkg/filter/severity"
	"github.com/Fishwaldo/mouthpiece/pkg/mocks"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)


func TestSeverity(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Severity Suite")
}

var _ = Describe("Severity", func() {
	var (
		mockMessage *mock_interfaces.MockMessageI
		flt         interfaces.FilterImplI
		ctrl 	  *gomock.Controller
	)
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockMessage = mock_interfaces.NewMockMessageI(ctrl)

		fltfactory := &severity.SevFilterFactory{}
		var err error
		flt, err = fltfactory.FilterFactory(context.Background(), "{}")
		Expect(err).To(BeNil())
	})
	Context("Config", func() {
		It("Should Return False When Procesing with Invalid Config", func() {
			eqflt := &severity.SeverityFilterConfig{
				Severity: 3,
				Op:       666,
			}
			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).ToNot(BeNil())
		})
		It("Return a Valid Config", func() {
			eqflt := &severity.SeverityFilterConfig{
				Severity: 3,
				Op:       severity.SeverityFilterOpEQ,
			}
			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			cfg, err := flt.GetConfig(context.Background())
			Expect(err).To(BeNil())
			Expect(cfg).To(Equal(eqflt))
		})
		It("Return a Valid Config as JSON", func() {
			eqflt := &severity.SeverityFilterConfig{
				Severity: 3,
				Op:       severity.SeverityFilterOpEQ,
			}
			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			cfg, err := flt.GetConfig(context.Background())
			Expect(err).To(BeNil())
			Expect(cfg).To(Equal(eqflt))
			jscfg, err := cfg.AsJSON()
			Expect(err).To(BeNil())
			var blah severity.SeverityFilterConfig
			err = json.Unmarshal([]byte(jscfg), &blah)
			Expect(err).To(BeNil())
		})

	})

	Context("Filter Severity Equal", func() {
		It("Should Return FilterMatch", func() {
			eqflt := &severity.SeverityFilterConfig{
				Severity: 3,
				Op:       severity.SeverityFilterOpEQ,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetSeverity().Return(3)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterMatch))
		})
		It("Should Return FilterNoMatch when not equal", func() {
			eqflt := &severity.SeverityFilterConfig{
				Severity: 2,
				Op:       severity.SeverityFilterOpEQ,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetSeverity().Return(3)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterNoMatch))
		})
	})
	Context("Filter Severity Not Equal", func() {
		It("Should Return FilterMatch", func() {
			eqflt := &severity.SeverityFilterConfig{
				Severity: 3,
				Op:       severity.SeverityFilterOpNE,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetSeverity().Return(3)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterNoMatch))
		})
		It("Should Return FilterNoMatch when not equal", func() {
			eqflt := &severity.SeverityFilterConfig{
				Severity: 2,
				Op:       severity.SeverityFilterOpNE,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetSeverity().Return(3)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterMatch))
		})
	})
	Context("Filter Severity Greater Than", func() {
		It("Should Return FilterMatch", func() {
			eqflt := &severity.SeverityFilterConfig{
				Severity: 2,
				Op:       severity.SeverityFilterOpGT,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetSeverity().Return(3)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterMatch))
		})
		It("Should Return FilterNoMatch when Less than", func() {
			eqflt := &severity.SeverityFilterConfig{
				Severity: 4,
				Op:       severity.SeverityFilterOpGT,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetSeverity().Return(3)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterNoMatch))
		})
		It("Should Return FilterNoMatch when Equal", func() {
			eqflt := &severity.SeverityFilterConfig{
				Severity: 3,
				Op:       severity.SeverityFilterOpGT,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetSeverity().Return(3)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterNoMatch))
		})
	})
	Context("Filter Severity Less Than", func() {
		It("Should Return FilterMatch", func() {
			eqflt := &severity.SeverityFilterConfig{
				Severity: 4,
				Op:       severity.SeverityFilterOpLT,
			}

			mockMessage.EXPECT().GetSeverity().Return(3)

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterMatch))
		})
	})
	Context("Filter Severity Greater Than or Equal", func() {
		It("Should Return FilterMatch when equal", func() {
			eqflt := &severity.SeverityFilterConfig{
				Severity: 3,
				Op:       severity.SeverityFilterOpGE,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetSeverity().Return(3)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterMatch))
		})
		It("Should return FilterMatch when greater than", func() {
			eqflt := &severity.SeverityFilterConfig{
				Severity: 2,
				Op:       severity.SeverityFilterOpGE,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetSeverity().Return(3)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterMatch))
		})
		It("Should return FilterNoMatch when less than when checking greater than", func() {
			eqflt := &severity.SeverityFilterConfig{
				Severity: 4,
				Op:       severity.SeverityFilterOpGE,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetSeverity().Return(3)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterNoMatch))
		})
	})
	Context("Filter Severity Less Than or Equal", func() {
		It("Should Return FilterMatch", func() {
			eqflt := &severity.SeverityFilterConfig{
				Severity: 3,
				Op:       severity.SeverityFilterOpLE,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetSeverity().Return(3)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterMatch))
		})
		It("Should return FilterMatch when less than", func() {
			eqflt := &severity.SeverityFilterConfig{
				Severity: 4,
				Op:       severity.SeverityFilterOpLE,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetSeverity().Return(3)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterMatch))
		})
		It("Should return FilterNoMatch when greater than when checking less than", func() {
			eqflt := &severity.SeverityFilterConfig{
				Severity: 2,
				Op:       severity.SeverityFilterOpLE,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetSeverity().Return(3)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterNoMatch))
		})
	})
})
