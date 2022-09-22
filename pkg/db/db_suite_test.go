package db_test

import (
	"testing"

	"github.com/Fishwaldo/mouthpiece/pkg/db"

	_ "github.com/mattn/go-sqlite3"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Db Suite")
}

var _ = Describe("Db", func() {
	Context("Initialize", func() {
		It("Should Complete", func() {
			clnt, err := db.Initialize("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			Expect(err).To(BeNil())
			Expect(clnt).ToNot(BeNil())
		})
	})
	Context("Tenant", func() {
		It("Get Default Tenant", func() {
			clnt, err := db.Initialize("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			Expect(err).To(BeNil())
			Expect(clnt).ToNot(BeNil())
			tenant := db.GetDefaultTenant()
			Expect(tenant).ToNot(BeNil())
			Expect(tenant.Name).To(Equal("default"))

		})
	})
})
