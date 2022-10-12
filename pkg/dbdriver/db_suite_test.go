package dbdriver_test

import (
	"database/sql"
	"testing"
	"os"

	"github.com/Fishwaldo/mouthpiece/pkg/dbdriver"
//	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/go-logr/logr/testr"
	"github.com/go-logr/logr"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var tst *testing.T

func TestDb(t *testing.T) {
	tst = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Db Suite")
}

var _ = Describe("Db", func() {
	Context("Initialize", func() {
		It("Should Complete", func() {
			if len(os.Getenv("MP_DEBUG_LOG")) > 0 {
				log.InitLogger(testr.NewWithOptions(tst, testr.Options{Verbosity: 10}))
			} else {
				log.InitLogger(logr.Discard())
			}
			db, err := sql.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			Expect(err).To(BeNil())
			Expect(db).ToNot(BeNil())
			clnt, err := dbdriver.Initialize("sqlite3", db)
			Expect(err).To(BeNil())
			Expect(clnt).ToNot(BeNil())
		})
	})
	Context("Tenant", func() {
		It("Get Default Tenant", func() {
			if len(os.Getenv("MP_DEBUG_LOG")) > 0 {
				log.InitLogger(testr.NewWithOptions(tst, testr.Options{Verbosity: 10}))
			} else {
				log.InitLogger(logr.Discard())
			}
			db, err := sql.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			Expect(err).To(BeNil())
			Expect(db).ToNot(BeNil())
			clnt, err := dbdriver.Initialize("sqlite3", db)
			Expect(err).To(BeNil())
			Expect(clnt).ToNot(BeNil())
			tenant := dbdriver.GetDefaultTenant()
			Expect(tenant).ToNot(BeNil())
			Expect(tenant.Name).To(Equal("default"))

		})
	})
})
