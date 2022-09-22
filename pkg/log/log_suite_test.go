package log_test

import (
	"testing"

	"github.com/Fishwaldo/mouthpiece/pkg/log"

	"github.com/go-logr/logr"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Log Suite")
}

var _ = Describe("Log", func() {
	Context("Initialize", func() {
		It("Should Complete", func() {
			log := log.InitLogger(logr.Discard())
			Expect(log).ToNot(BeNil())
		})
	})
})
