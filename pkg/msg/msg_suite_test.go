package msg_test

import (
	"context"
	"testing"
	"time"

	//	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/ent"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbapp"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/rules"
	_ "github.com/Fishwaldo/mouthpiece/pkg/ent/runtime"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/msg"

	"github.com/go-logr/logr/testr"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var tst *testing.T
func TestMsg(t *testing.T) {
	tst = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Msg Suite")
}

var eClient *ent.Client
var ctx context.Context

var _ = BeforeSuite(func() {
	var err error
	eClient, err = db.Initialize("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	Expect(err).To(BeNil())
	logger := testr.NewWithOptions(tst, testr.Options{Verbosity: -1})
	log.InitLogger(logger)
	tntrole := &rules.UserViewer{
		Role: rules.GlobalAdmin,
	}
	ctx = rules.NewContext(context.Background(), tntrole)
	tnt, err := eClient.Tenant.Create().SetName("Default").Save(ctx)
	tntrole.Role = rules.Admin
	Expect(err).To(BeNil())
	Expect(tnt).ToNot(BeNil())
	tntrole.T = tnt

	globalApp, err = eClient.DbApp.Create().SetName("Global").
		SetStatus(interfaces.Enabled).
		SetDescription("test").
		SetIcon("http://www.example.com/icon.png").
		SetURL("http://www.example.com").
		Save(ctx)
	Expect(err).To(BeNil())
	Expect(globalApp).ToNot(BeNil())

})

var _ = AfterSuite(func() {
	eClient.Close()
})

var globalApp *ent.DbApp
var globalMsg *msg.Message

var _ = Describe("Msg", func() {
	Context("Create", func() {
		It("Should Create a new Msg ", func() {
			var shtmsg = "string"
			var tpic = "topic"
			msgstruct := &ent.DbMessage {
				Message: "Hello World",
				ShortMsg: &shtmsg,
				Severity: 3,
				Topic: &tpic,
				Timestamp: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
			}
			globalMsg = msg.NewMessage()
			Expect(globalMsg).ToNot(BeNil())
			err := globalMsg.Load(ctx, msgstruct)
			Expect(err).To(BeNil())
			err = globalMsg.Save(ctx, globalApp)
			Expect(err).To(BeNil())
		})
		It("We can add Fields", func() {
			flds := make(map[string]string)
			flds["test"] = "test"
			flds["test2"] = "test2"
			err := globalMsg.SetFields(ctx, flds)
			Expect(err).To(BeNil())

		})
	})
	Context("We can get Message Data", func() {
		It("Can Get the Message", func() {
			loadmsg, err := msg.LoadFromDB(ctx, globalMsg.GetID())
			Expect(err).To(BeNil())
			Expect(globalMsg).ToNot(BeNil())
			Expect(globalMsg.GetID()).To(Equal(globalMsg.GetID()))
			globalMsg = loadmsg
		})
		It("Get ID", func() {
			Expect(globalMsg.GetID()).To(Equal(globalMsg.GetID()))
		})
		It("Get Message", func() {
			Expect(globalMsg.GetMessage()).To(Equal("Hello World"))
		})
		It("Get Short Message", func() {
			Expect(globalMsg.GetShortMsg()).To(HaveValue(Equal("string")))
		})
		It("Get Severity", func() {
			Expect(globalMsg.GetSeverity()).To(Equal(3))
		})
		It("Get Topic", func() {
			Expect(globalMsg.GetTopic()).To(HaveValue(Equal("topic")))
		})
		It("Get Timestamp", func() {
			Expect(globalMsg.GetTimestamp()).To(Equal(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)))
		})
		It("Get Fields", func() {
			flds, err := globalMsg.GetFields(ctx)
			Expect(err).To(BeNil())
			Expect(flds["test"]).To(Equal("test"))
			Expect(flds["test2"]).To(Equal("test2"))
		})
		It("Get App", func() {
			Skip("Need to Mock AppService")
			app, err := globalMsg.GetApp(ctx)
			Expect(err).To(BeNil())
			Expect(app).ToNot(BeNil())
			Expect(app.GetID()).To(Equal(globalApp.ID))
		})
	})
	Context("Modify", func() {
		It("We can modify the Message", func() {
			err := globalMsg.SetMessage(ctx, "Hello World 2")
			Expect(err).To(BeNil())
			Expect(globalMsg.GetMessage()).To(Equal("Hello World 2"))
		})
		It("We can modify the Short Message", func() {
			err := globalMsg.SetShortMsg(ctx, "string2")
			Expect(err).To(BeNil())
			Expect(globalMsg.GetShortMsg()).To(HaveValue(Equal("string2")))
		})
		It("We can modify the Severity", func() {
			err := globalMsg.SetSeverity(ctx, 4)
			Expect(err).To(BeNil())
			Expect(globalMsg.GetSeverity()).To(Equal(4))
		})
		It("We can modify the Topic", func() {
			err := globalMsg.SetTopic(ctx, "topic2")
			Expect(err).To(BeNil())
			Expect(globalMsg.GetTopic()).To(HaveValue(Equal("topic2")))
		})
		It("We can modify the Timestamp", func() {
			err := globalMsg.SetTimestamp(ctx, time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
			Expect(err).To(BeNil())
			Expect(globalMsg.GetTimestamp()).To(Equal(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)))
		})
	})
	Context("Process", func() {
		It("We can process the message", func() {
			err := globalMsg.ProcessMessage(ctx)
			Expect(err).To(BeNil())
		})
	})
})