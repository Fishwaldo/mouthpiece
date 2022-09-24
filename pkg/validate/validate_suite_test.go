package validate_test

import (
	"testing"

	"github.com/Fishwaldo/mouthpiece/pkg/validate"

	//mptest "github.com/Fishwaldo/mouthpiece/pkg/mocks"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestValidate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Validate Suite")
}

var _ = Describe("Validate", func() {
	Context("Validate", func() {
		It("Should Return the Validate Struct", func() {
			v := validate.Get()
			Expect(v).ToNot(BeNil())
		})
	})
	Context("Ent Validation Functions", func() {
		It("Should Return a String Validator", func() {
			v := validate.EntStringValidator("required")
			Expect(v).ToNot(BeNil())
			err := v("test")
			Expect(err).To(BeNil())			
		})
		It("Should return a error for a invalid validation", func() {
			v := validate.EntStringValidator("email")
			Expect(v).ToNot(BeNil())
			err := v("test")
			Expect(err).ToNot(BeNil())
		})
	})
})