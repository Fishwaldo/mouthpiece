package field_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/Fishwaldo/mouthpiece/pkg/filter/field"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/mocks"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestField(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Field Suite")
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

		fltfactory := &field.FieldFilterFactory{}
		var err error
		flt, err = fltfactory.FilterFactory(context.Background(), "{}")
		Expect(err).To(BeNil())
	})
	Context("Config", func() {
		It("Should Return False When Procesing with Invalid Config", func() {
			eqflt := &field.FieldFilterConfig{
				Field: "test",
				Value: "blah",
				Op: 666,
			}
			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).ToNot(BeNil())
		})
		It("Return a Valid Config", func() {
			eqflt := &field.FieldFilterConfig{
				Field: "test",
				Value: "blah",
				Op: field.FieldFilterOpEQ,
			}
			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			cfg, err := flt.GetConfig(context.Background())
			Expect(err).To(BeNil())
			Expect(cfg).To(Equal(eqflt))
		})
		It("Return a Valid Config as JSON", func() {
			eqflt := &field.FieldFilterConfig{
				Field: "test",
				Value: "blah",
				Op: field.FieldFilterOpEQ,
			}
			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			cfg, err := flt.GetConfig(context.Background())
			Expect(err).To(BeNil())
			Expect(cfg).To(Equal(eqflt))
			jscfg, err := cfg.AsJSON()
			Expect(err).To(BeNil())
			var blah field.FieldFilterConfig
			err = json.Unmarshal([]byte(jscfg), &blah)
			Expect(err).To(BeNil())
		})
	})
	Context("Filter Filed Equal", func() {
		It("Should Return FilterMatch", func() {
			eqflt := &field.FieldFilterConfig{
				Field: "test",
				Value: "blah",
				Op: field.FieldFilterOpEQ,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetField(gomock.Any(), gomock.Eq("test")).Return("blah", nil)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterMatch))
		})
		It("Should Return FilterNoMatch for a different filed", func() {
			eqflt := &field.FieldFilterConfig{
				Field: "test",
				Value: "blah2",
				Op: field.FieldFilterOpEQ,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetField(gomock.Any(), gomock.Eq("test")).Return("blah", nil)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterNoMatch))
		})
		It("Should Return FilterNoMatch for a non-existant Field", func() {
			eqflt := &field.FieldFilterConfig{
				Field: "test",
				Value: "blah2",
				Op: field.FieldFilterOpEQ,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetField(gomock.Any(), gomock.Eq("test")).Return("", mperror.ErrMsgFieldNotFound)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).ToNot(BeNil())
			Expect(ok).To(Equal(interfaces.FilterNoMatch))
		})
	})
	Context("Filter Filed Contains", func() {
		It("Should Return FilterMatch", func() {
			eqflt := &field.FieldFilterConfig{
				Field: "test",
				Value: "blah",
				Op: field.FieldFilterOpContains,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetField(gomock.Any(), gomock.Eq("test")).Return("hahablahhaha", nil)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterMatch))
		})
		It("Should Return FilterNoMatch for a different filed", func() {
			eqflt := &field.FieldFilterConfig{
				Field: "test",
				Value: "blah2",
				Op: field.FieldFilterOpContains,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetField(gomock.Any(), gomock.Eq("test")).Return("ooops", nil)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterNoMatch))
		})
		It("Should Return FilterNoMatch for a non-existant Field", func() {
			eqflt := &field.FieldFilterConfig{
				Field: "test",
				Value: "blah2",
				Op: field.FieldFilterOpContains,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetField(gomock.Any(), gomock.Eq("test")).Return("", mperror.ErrMsgFieldNotFound)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).ToNot(BeNil())
			Expect(ok).To(Equal(interfaces.FilterNoMatch))
		})
	})
	Context("Filter Field Present", func() {
		It("Should Return FilterMatch", func() {
			eqflt := &field.FieldFilterConfig{
				Field: "test",
				Op: field.FieldFilterOpPresent,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetField(gomock.Any(), gomock.Eq("test")).Return("hahablahhaha", nil)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterMatch))
		})
		It("Should Return FilterNoMatch for a non-existant Field", func() {
			eqflt := &field.FieldFilterConfig{
				Field: "test",
				Op: field.FieldFilterOpPresent,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetField(gomock.Any(), gomock.Eq("test")).Return("", mperror.ErrMsgFieldNotFound)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterNoMatch))
		})
	})
	Context("Filter Field Missing", func() {
		It("Should Return FilterNoMatch", func() {
			eqflt := &field.FieldFilterConfig{
				Field: "test",
				Op: field.FieldFilterOpMissing,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetField(gomock.Any(), gomock.Eq("test")).Return("hahablahhaha", nil)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterNoMatch))
		})
		It("Should Return Filter for a non-existant Field", func() {
			eqflt := &field.FieldFilterConfig{
				Field: "test",
				Op: field.FieldFilterOpMissing,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetField(gomock.Any(), gomock.Eq("test")).Return("", mperror.ErrMsgFieldNotFound)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterMatch))
		})
	})
	Context("Filter Field Regular Expression", func() {
		It("Should Return FilterMatch", func() {
			eqflt := &field.FieldFilterConfig{
				Field: "test",
				Value: ".*blah.*",
				Op: field.FieldFilterOpRegex,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetField(gomock.Any(), gomock.Eq("test")).Return("hahablahhaha", nil)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterMatch))
		})
		It("Should Return FilterNoMatch for a different filed", func() {
			eqflt := &field.FieldFilterConfig{
				Field: "test",
				Value: ".*blah2.*",
				Op: field.FieldFilterOpRegex,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetField(gomock.Any(), gomock.Eq("test")).Return("ooops", nil)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterNoMatch))
		})
		It("Should Return FilterNoMatch for a non-existant Field", func() {
			eqflt := &field.FieldFilterConfig{
				Field: "test",
				Value: ".*blah2.*",
				Op: field.FieldFilterOpRegex,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetField(gomock.Any(), gomock.Eq("test")).Return("", mperror.ErrMsgFieldNotFound)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).ToNot(BeNil())
			Expect(ok).To(Equal(interfaces.FilterNoMatch))
		})
		It("Should Return FilterPass for a invalid regexp", func() {
			eqflt := &field.FieldFilterConfig{
				Field: "test",
				Value: "[",
				Op: field.FieldFilterOpRegex,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			mockMessage.EXPECT().GetField(gomock.Any(), gomock.Eq("test")).Return("ooops", nil)

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).ToNot(BeNil())
			Expect(ok).To(Equal(interfaces.FilterPass))
		})
	})
	Context("Filter Field NoOp", func() {
		It("Should Return FilterPass", func() {
			eqflt := &field.FieldFilterConfig{
				Field: "test",
				Op: field.FieldFilterOpNoOp,
			}

			err := flt.SetConfig(context.Background(), eqflt)
			Expect(err).To(BeNil())

			ok, err := flt.Process(context.Background(), mockMessage)
			Expect(err).To(BeNil())
			Expect(ok).To(Equal(interfaces.FilterPass))
		})
	})

})

