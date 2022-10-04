package mouthpiece_test

import (
	"context"
	"testing"
	"database/sql"

	mouthpiece "github.com/Fishwaldo/mouthpiece/pkg"
	"github.com/Fishwaldo/mouthpiece/pkg/msg"

	//	"github.com/Fishwaldo/mouthpiece/pkg/log"

	"github.com/go-logr/logr/testr"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var tst *testing.T

func TestMouthPiece(t *testing.T) {
	tst = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pkg Suite")
}

var _ = Describe("MouthPiece", func() {
	Context("Initialize", func() {
		It("should initialize", func() {
			logger := testr.NewWithOptions(tst, testr.Options{Verbosity: 10})

			db, err := sql.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			Expect(err).To(BeNil())
			Expect(db).ToNot(BeNil())

			mp := mouthpiece.NewMouthPiece(context.Background(), "sqlite3", db, logger)
			Expect(mp).ToNot(BeNil())
			ctx := mp.SetAdminTenant(context.Background())
			err = mp.Start(ctx)
			Expect(err).To(BeNil())

			app, err := mp.GetAppService().Get(ctx, "MouthPiece")
			Expect(err).To(BeNil())
			Expect(app).ToNot(BeNil())

			msg1 := msg.NewMessage(ctx, "Hello World", app)
			Expect(msg1.SetTopic(ctx, "Test")).To(BeNil())
			Expect(msg1.Save(ctx, app)).To(BeNil())

			err = mp.RouteMessage(ctx, msg1)
			Expect(err).To(BeNil())

			msg2 := msg.NewMessage(ctx, "Hello Again World", app)
			Expect(msg2.SetTopic(ctx, "MyTopic")).To(BeNil())
			Expect(msg2.SetSeverity(ctx, 4)).To(BeNil())
			Expect(msg2.SetShortMsg(ctx, "Short Message")).To(BeNil())
			Expect(msg2.Save(ctx, app)).To(BeNil())
			Expect(mp.RouteMessage(ctx, msg2)).To(BeNil())
		})
	})
})
