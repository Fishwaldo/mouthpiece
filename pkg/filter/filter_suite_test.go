package filter_test

import (
	"testing"
	"time"

	"github.com/Fishwaldo/mouthpiece/pkg/filter"
	_ "github.com/Fishwaldo/mouthpiece/pkg/filter/noop"
	noopfilter "github.com/Fishwaldo/mouthpiece/pkg/filter/noop"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"
	"github.com/Fishwaldo/mouthpiece/pkg/msg"

	mptest "github.com/Fishwaldo/mouthpiece/pkg/mocks"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFilter(t *testing.T) {
	mptest.Tst = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Filter Suite")
}

var _ = mptest.DBBeforeSuite
var _ = mptest.DBAfterSuite
var globalFilter interfaces.FilterI

var _ = Describe("Filter", func() {
	Context("Provider", func() {
		It("Should List Filter Providers", func() {
			providers := filter.GetFilterImpls(mptest.Ctx)
			Expect(providers).To(HaveLen(2))
//			Expect(providers[0]).To(Equal("NoOpFilter"))
		})
		It("Should Get a New Filter Implementation", func() {
			provider, err := filter.GetNewFilterImpl(mptest.Ctx, "NoOpFilter", "{}")
			Expect(err).To(BeNil())
			Expect(provider).ToNot(BeNil())
			Expect(provider.FilterName()).To(Equal("NoOpFilter"))
		})
		It("Should Get a New Filter Config Implementation", func() {
			provider, err := filter.GetFilterImplDefaultConfig(mptest.Ctx, "NoOpFilter")
			Expect(err).To(BeNil())
			Expect(provider).ToNot(BeNil())
			Expect(provider).To(BeAssignableToTypeOf(&noopfilter.NoOpFilterConfig{}))
		})
		It("Should Return a Error for a non-existant FilterImpl", func() {
			_, err := filter.GetNewFilterImpl(mptest.Ctx, "NoOpFilter2", "{}")
			Expect(err).ToNot(BeNil())
			Expect(err).To(MatchError(mperror.ErrFilterImplNotFound))
		})
		It("Should return Not Found for a non-existant Filter Config Implementation", func() {
			_, err := filter.GetFilterImplDefaultConfig(mptest.Ctx, "NoOpFilter2")
			Expect(err).ToNot(BeNil())
			Expect(err).To(MatchError(mperror.ErrFilterImplNotFound))
		})
	})
	Context("Filter", func() {
		It("Should Create a New Filter", func() {
			var err error
			globalFilter, err = mptest.Mp.GetFilterService().Create(mptest.Ctx, "NoOpFilter", "newfilter", interfaces.AppFilter)
			Expect(err).To(BeNil())
			Expect(globalFilter).ToNot(BeNil())
			Expect(globalFilter.GetName()).To(Equal("newfilter"))
			Expect(globalFilter.GetType()).To(Equal(interfaces.AppFilter.String()))
		})
		It("Should Modify Name of a Filter", func() {
			err := globalFilter.SetName(mptest.Ctx, "newfilter2")
			Expect(err).To(BeNil())
			Expect(globalFilter.GetName()).To(Equal("newfilter2"))
		})
		It("Should Modify Description of a Filter", func() {
			err := globalFilter.SetDescription(mptest.Ctx, "newfilter2")
			Expect(err).To(BeNil())
			Expect(globalFilter.GetDescription()).To(Equal("newfilter2"))
		})
		It("Should Get the ID of a Filter", func() {
			Expect(globalFilter.GetID()).ToNot(BeNil())
		})
		It("Should Get the Filter Implemenation", func() {
			impl := globalFilter.GetFilterImplementation()
			Expect(impl).ToNot(BeNil())
			Expect(impl.FilterName()).To(Equal("NoOpFilter"))
		})
		It("Should be able to Set the COnfig", func() {
			cfg, err := filter.GetFilterImplDefaultConfig(mptest.Ctx, "NoOpFilter")
			Expect(err).To(BeNil())
			Expect(cfg).ToNot(BeNil())
			Expect(cfg).To(BeAssignableToTypeOf(&noopfilter.NoOpFilterConfig{}))
			err = globalFilter.SetConfig(mptest.Ctx, cfg)
			Expect(err).To(BeNil())
		})
		It("Should be able to Get the COnfig", func() {
			cfg, err := globalFilter.GetConfig(mptest.Ctx)
			Expect(err).To(BeNil())
			Expect(cfg).ToNot(BeNil())
			Expect(cfg).To(BeAssignableToTypeOf(&noopfilter.NoOpFilterConfig{}))
		})
		It("Should Process a Message", func() {
			mymsg:= msg.NewMessage(mptest.Ctx, "Message", mptest.App)
			action, err := globalFilter.ProcessMessage(mptest.Ctx, mymsg)
			Expect(err).To(BeNil())
			Expect(action).To(Equal(interfaces.FilterPass))
		})
	})
	Context("Filter Service", func() {
		It("SHould Get a Filter by Name", func() {
			f, err := mptest.Mp.GetFilterService().Get(mptest.Ctx, "newfilter2", interfaces.AppFilter)
			Expect(err).To(BeNil())
			Expect(f).ToNot(BeNil())
			Expect(f.GetName()).To(Equal("newfilter2"))
		})
		It("SHould Get a Filter by ID", func() {
			f, err := mptest.Mp.GetFilterService().GetByID(mptest.Ctx, globalFilter.GetID(), interfaces.AppFilter)
			Expect(err).To(BeNil())
			Expect(f).ToNot(BeNil())
			Expect(f.GetName()).To(Equal("newfilter2"))
		})
		It("SHould Start a Filter", func() {
			err := mptest.Mp.GetFilterService().Start(mptest.Ctx)
			Expect(err).To(BeNil())
		})
		It("Should get Expired Filter by Name", func() {
			time.Sleep(5 * time.Second)
			f, err := mptest.Mp.GetFilterService().Get(mptest.Ctx, "newfilter2", interfaces.AppFilter)
			Expect(err).To(BeNil())
			Expect(f).ToNot(BeNil())
			Expect(f.GetName()).To(Equal("newfilter2"))
		})
		It("Should get Expired Filters by ID", func() {
			time.Sleep(5 * time.Second)
			f, err := mptest.Mp.GetFilterService().GetByID(mptest.Ctx, globalFilter.GetID(), interfaces.AppFilter)
			Expect(err).To(BeNil())
			Expect(f).ToNot(BeNil())
			Expect(f.GetName()).To(Equal("newfilter2"))
		})
	})
})