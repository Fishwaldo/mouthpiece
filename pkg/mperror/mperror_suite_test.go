package mperror_test

import (
	"testing"

	"github.com/Fishwaldo/mouthpiece/pkg/mperror"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMperror(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mperror Suite")
}

var _ = Describe("Mperror", func() {
	Context("Error Codes", func() {
		It("Error Filter Should Work", func() {
			err := mperror.FilterErrors(mperror.ErrValidationError)
			Expect(err).ToNot(BeNil())
			Expect(err).To(MatchError(mperror.ErrInternalError))
		})
	})
})