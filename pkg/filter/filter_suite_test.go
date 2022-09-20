package filter_test

import (
	"testing"

	_ "github.com/Fishwaldo/mouthpiece/pkg/filter"
	_ "github.com/Fishwaldo/mouthpiece/pkg/filter/severity"	
	_ "github.com/Fishwaldo/mouthpiece/pkg/filter/field"	
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFilter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Filter Suite")
}
