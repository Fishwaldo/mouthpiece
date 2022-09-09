package ent_test

import (
	"testing"
	"context"



	"github.com/Fishwaldo/mouthpiece/pkg/ent"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/rules"
	_ "github.com/Fishwaldo/mouthpiece/pkg/ent/runtime"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/migrate"

	"entgo.io/ent/dialect"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	_ "github.com/mattn/go-sqlite3"
)

func TestEnt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ent Suite")
}

var eClient *ent.Client
var ctx context.Context

var _ = BeforeSuite(func() {
	var err error
	eClient, err = ent.Open(dialect.SQLite, "file:test.db?mode=memory&cache=shared&_fk=1")
	Expect(err).To(BeNil())
	Expect(eClient).ToNot(BeNil())
	// Run the automatic migration tool to create all schema resources.
	Expect(eClient.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true))).To(BeNil())
	tntrole := &rules.UserViewer{
		Role: rules.GlobalAdmin,
	}
	ctx = rules.NewContext(context.Background(), tntrole)
	tnt, err := eClient.Tenant.Create().SetName("Default").Save(ctx)
	tntrole.Role = rules.Admin
	Expect(err).To(BeNil())
	Expect(tnt).ToNot(BeNil())
	tntrole.T = tnt

})

var _ = AfterSuite(func() {
	eClient.Close()
})