package msg_test

import (
	"testing"
	"time"

	"github.com/google/uuid"

	_ "github.com/Fishwaldo/mouthpiece/pkg/ent/runtime"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	mptest "github.com/Fishwaldo/mouthpiece/pkg/mocks"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"
	"github.com/Fishwaldo/mouthpiece/pkg/msg"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMsg(t *testing.T) {
	mptest.Tst = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Msg Suite")
}

var _ = mptest.DBBeforeSuite
var _ = mptest.DBAfterSuite

var globalMsg interfaces.MessageI

var _ = Describe("Msg", func() {
	var msgid uuid.UUID
	Context("Create", func() {
		It("Should Create a new Msg ", func() {
			Expect(mptest.Ctx).ToNot(BeNil())
			Expect(mptest.App).ToNot(BeNil())
			globalMsg = msg.NewMessage(mptest.Ctx, "Message", mptest.App)
			Expect(globalMsg).ToNot(BeNil())

			Expect(globalMsg.SetSeverity(mptest.Ctx, 1)).To(BeNil())
			Expect(globalMsg.SetShortMsg(mptest.Ctx, "Short Message")).To(BeNil())
			Expect(globalMsg.SetTopic(mptest.Ctx, "Topic")).To(BeNil())
			Expect(globalMsg.SetMessage(mptest.Ctx, "Message2")).To(BeNil())
			Expect(globalMsg.SetTimestamp(mptest.Ctx, time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))).To(BeNil())

			err := globalMsg.Save(mptest.Ctx, mptest.App)
			Expect(err).To(BeNil())
			msgid = globalMsg.GetID()
			Expect(msgid).ToNot(BeNil())

			Expect(globalMsg.GetSeverity()).To(Equal(1))
			Expect(globalMsg.GetTopic()).To(HaveValue(Equal("Topic")))
			Expect(globalMsg.GetShortMsg()).To(HaveValue(Equal("Short Message")))
			Expect(globalMsg.GetMessage()).To(Equal("Message2"))
			Expect(globalMsg.GetTimestamp()).To(Equal(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)))
			newapp, err := globalMsg.GetApp(mptest.Ctx)
			Expect(err).To(BeNil())
			Expect(newapp).ToNot(BeNil())
			Expect(newapp.GetName()).To(Equal("MouthPiece"))
		})
		It("Can Load a Message Based on MsgID", func() {
			Expect(mptest.Ctx).ToNot(BeNil())
			Expect(mptest.App).ToNot(BeNil())
			Expect(msgid).ToNot(BeNil())
			mymsg, err := msg.LoadFromDB(mptest.Ctx, msgid)
			Expect(err).To(BeNil())
			Expect(mymsg).ToNot(BeNil())
			Expect(mymsg.GetID()).To(Equal(msgid))
		})
	})
	Context("Message Fields", func() {
		It("We can add Fields", func() {
			flds := make(map[string]string)
			flds["test"] = "test"
			flds["test2"] = "test2"
			err := globalMsg.SetFields(mptest.Ctx, flds)
			Expect(err).To(BeNil())
		})
		It("We can Get a Fields", func() {
			val, err := globalMsg.GetField(mptest.Ctx, "test")
			Expect(err).To(BeNil())
			Expect(val).To(Equal("test"))
		})
		It("We can Get all Fields", func() {
			val, err := globalMsg.GetFields(mptest.Ctx)
			Expect(err).To(BeNil())
			Expect(val["test"]).To(Equal("test"))
			Expect(val["test2"]).To(Equal("test2"))
		})
		It("We can update a Fields", func() {
			err := globalMsg.SetField(mptest.Ctx, "test", "test3")
			Expect(err).To(BeNil())
			val, err := globalMsg.GetField(mptest.Ctx, "test")
			Expect(err).To(BeNil())
			Expect(val).To(Equal("test3"))
		})
		It("We can update Fields with a map", func() {
			flds := make(map[string]string)
			flds["test"] = "update"
			flds["test2"] = "test2"
			err := globalMsg.SetFields(mptest.Ctx, flds)
			Expect(err).To(BeNil())
			val, err := globalMsg.GetField(mptest.Ctx, "test")
			Expect(err).To(BeNil())
			Expect(val).To(Equal("update"))
		})
		It("We Get a Error with a unknown field", func() {
			_, err := globalMsg.GetField(mptest.Ctx, "test3")
			Expect(err).To(MatchError(mperror.ErrMsgFieldNotFound))
		})
	})
	Context("MetaData Fields", func() {
		It("We can add Metadata", func() {
			Expect(globalMsg.SetMetadata(mptest.Ctx, "test", "test")).To(BeNil())
		})
		It("Can Get Metadata", func() {
			val, err := globalMsg.GetMetadata(mptest.Ctx, "test")
			Expect(err).To(BeNil())
			Expect(val).To(Equal("test"))
		})
		It("Errors on Invalid Metadata", func() {
			_, err := globalMsg.GetMetadata(mptest.Ctx, "test2")
			Expect(err).To(MatchError(mperror.ErrMsgMetadataNotFound))
		})
		It("Can Replace Metadata", func() {
			Expect(globalMsg.SetMetadata(mptest.Ctx, "test", "test2")).To(BeNil())
			val, err := globalMsg.GetMetadata(mptest.Ctx, "test")
			Expect(err).To(BeNil())
			Expect(val).To(Equal("test2"))
		})
		It("Can Handle Any Interface for Metadata", func() {
			Expect(globalMsg.SetMetadata(mptest.Ctx, "testint", 1)).To(BeNil())
			val, err := globalMsg.GetMetadata(mptest.Ctx, "testint")
			Expect(err).To(BeNil())
			Expect(val).To(Equal(1))
		})
		It("Can Get All Metadata", func() {
			val, err := globalMsg.GetMetadataFields(mptest.Ctx)
			Expect(err).To(BeNil())
			Expect(val["testint"]).To(Equal(1))
			Expect(val["test"]).To(Equal("test2"))
		})
	})
	Context("Cloned Messages", func() {
		It("Can Clone a Message", func() {
			newmsg := globalMsg.Clone()
			Expect(newmsg).ToNot(BeNil())
			Expect(newmsg.GetID()).To(Equal(globalMsg.GetID()))
			Expect(newmsg.GetTopic()).To(Equal(globalMsg.GetTopic()))
			Expect(newmsg.GetShortMsg()).To(Equal(globalMsg.GetShortMsg()))
			Expect(newmsg.GetMessage()).To(Equal(globalMsg.GetMessage()))
			Expect(newmsg.GetTimestamp()).To(Equal(globalMsg.GetTimestamp()))
			Expect(newmsg.GetSeverity()).To(Equal(globalMsg.GetSeverity()))

			newapp, err := newmsg.GetApp(mptest.Ctx)
			Expect(err).To(BeNil())
			Expect(newapp).ToNot(BeNil())

			oldapp, err := globalMsg.GetApp(mptest.Ctx)
			Expect(err).To(BeNil())
			Expect(oldapp).ToNot(BeNil())
			Expect(newapp.GetName()).To(Equal(oldapp.GetName()))

			Expect(newmsg.SetMessage(mptest.Ctx, "Cloned Message")).To(MatchError(mperror.ErrMsgLocked))

			Expect(newmsg.GetMessage()).To(Equal("Message2"))
			Expect(globalMsg.GetMessage()).To(Equal("Message2"))

			/* make sure our Metadata Fields are not shared */
			Expect(newmsg.SetMetadata(mptest.Ctx, "clonetest", "Cloned Value")).To(BeNil())
			val, err := newmsg.GetMetadata(mptest.Ctx, "clonetest")
			Expect(err).To(BeNil())
			Expect(val).To(Equal("Cloned Value"))

			_, err = globalMsg.GetMetadata(mptest.Ctx, "clonetest")
			Expect(err).To(MatchError(mperror.ErrMsgMetadataNotFound))

		})
	})
	Context("Process", func() {
		It("Can Process a Message", func() {
			//Expect(globalMsg.ProcessMessage(mptest.Ctx)).To(BeNil())
		})
	})

	// 	})
	// })
	// Context("We can get Message Data", func() {
	// 	It("Can Get the Message", func() {
	// 		loadmsg, err := msg.LoadFromDB(ctx, globalMsg.GetID())
	// 		Expect(err).To(BeNil())
	// 		Expect(globalMsg).ToNot(BeNil())
	// 		Expect(globalMsg.GetID()).To(Equal(globalMsg.GetID()))
	// 		globalMsg = loadmsg
	// 	})
	// 	It("Get ID", func() {
	// 		Expect(globalMsg.GetID()).To(Equal(globalMsg.GetID()))
	// 	})
	// 	It("Get Message", func() {
	// 		Expect(globalMsg.GetMessage()).To(Equal("Hello World"))
	// 	})
	// 	It("Get Short Message", func() {
	// 		Expect(globalMsg.GetShortMsg()).To(HaveValue(Equal("string")))
	// 	})
	// 	It("Get Severity", func() {
	// 		Expect(globalMsg.GetSeverity()).To(Equal(3))
	// 	})
	// 	It("Get Topic", func() {
	// 		Expect(globalMsg.GetTopic()).To(HaveValue(Equal("topic")))
	// 	})
	// 	It("Get Timestamp", func() {
	// 		Expect(globalMsg.GetTimestamp()).To(Equal(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)))
	// 	})
	// 	It("Get Fields", func() {
	// 		flds, err := globalMsg.GetFields(ctx)
	// 		Expect(err).To(BeNil())
	// 		Expect(flds["test"]).To(Equal("test"))
	// 		Expect(flds["test2"]).To(Equal("test2"))
	// 	})
	// 	It("Get App", func() {
	// 		Skip("Need to Mock AppService")
	// 		app, err := globalMsg.GetApp(ctx)
	// 		Expect(err).To(BeNil())
	// 		Expect(app).ToNot(BeNil())
	// 		Expect(app.GetID()).To(Equal(globalApp.ID))
	// 	})
	// })
	// Context("Modify", func() {
	// 	It("We can modify the Message", func() {
	// 		err := globalMsg.SetMessage(ctx, "Hello World 2")
	// 		Expect(err).To(BeNil())
	// 		Expect(globalMsg.GetMessage()).To(Equal("Hello World 2"))
	// 	})
	// 	It("We can modify the Short Message", func() {
	// 		err := globalMsg.SetShortMsg(ctx, "string2")
	// 		Expect(err).To(BeNil())
	// 		Expect(globalMsg.GetShortMsg()).To(HaveValue(Equal("string2")))
	// 	})
	// 	It("We can modify the Severity", func() {
	// 		err := globalMsg.SetSeverity(ctx, 4)
	// 		Expect(err).To(BeNil())
	// 		Expect(globalMsg.GetSeverity()).To(Equal(4))
	// 	})
	// 	It("We can modify the Topic", func() {
	// 		err := globalMsg.SetTopic(ctx, "topic2")
	// 		Expect(err).To(BeNil())
	// 		Expect(globalMsg.GetTopic()).To(HaveValue(Equal("topic2")))
	// 	})
	// 	It("We can modify the Timestamp", func() {
	// 		err := globalMsg.SetTimestamp(ctx, time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	// 		Expect(err).To(BeNil())
	// 		Expect(globalMsg.GetTimestamp()).To(Equal(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)))
	// 	})
	// })
	// Context("Process", func() {
	// 	It("We can process the message", func() {
	// 		err := globalMsg.ProcessMessage(ctx)
	// 		Expect(err).To(BeNil())
	// 	})
	// })
})
