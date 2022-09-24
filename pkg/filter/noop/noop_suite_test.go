package noopfilter_test

import (
	"testing"
	"context"
	"time"


	_ "github.com/Fishwaldo/mouthpiece/pkg/filter/noop"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/filter"
	mptest "github.com/Fishwaldo/mouthpiece/pkg/mocks"


	"github.com/google/uuid"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestNoop(t *testing.T) {
	mptest.Tst = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Noop Suite")
}

var globalFilter interfaces.FilterImplI

var _ = Describe("Evalfilter", func() {
	Context("It Should Register Filters", func() {
		It("Embeded Scripts", func() {
			filters := filter.GetFilterImpls(context.Background())
			Expect(filters).ToNot(BeNil())
		})
		It("SHould Create a Filter From a script", func() {
			var err error
			globalFilter, err = filter.GetNewFilterImpl(context.Background(), "NoOpFilter", "{}")
			Expect(err).To(BeNil())
			Expect(globalFilter).ToNot(BeNil())
			Expect(globalFilter.FilterName()).To(Equal("NoOpFilter"))
		})
	})
	Context("Should Filter Messages", func() {
		It("Should Filter a Message", func() {
			ctrl := gomock.NewController(GinkgoT())
			msg := mptest.NewMockMessageI(ctrl)
			msg.EXPECT().GetSeverity().Return(3).AnyTimes()
			msg.EXPECT().GetMessage().Return("Test warning Message").AnyTimes()
			msg.EXPECT().GetShortMsg().Return(nil).AnyTimes()
			msg.EXPECT().GetTopic().Return(nil).AnyTimes()
			msg.EXPECT().GetTimestamp().Return(time.Now()).AnyTimes()
			msg.EXPECT().SetSeverity(gomock.Any(), 4).Return(nil).AnyTimes()
			msg.EXPECT().GetID().Return(uuid.New()).AnyTimes()
			res, err := globalFilter.Process(context.Background(), msg)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(interfaces.FilterPass))
		})
	})
})
