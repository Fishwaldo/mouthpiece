package ent_test

import (
//	"context"
	"time"

	"github.com/Fishwaldo/mouthpiece/pkg/ent"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/app"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/message"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var DemoApp *ent.App
var DemoMsg *ent.Message

var _ = Describe("Messages: ", func() {
	It("Setup Demo App", func() {
		var err error
		DemoApp, err = eClient.App.Create().
			SetName("DemoApp").
			SetStatus(app.StatusEnabled).
			SetDescription("Demo Application").
			SetIcon("https://example.com/icon.png").
			SetURL("https://example.com").
			Save(ctx)
		Expect(err).To(BeNil())
		Expect(DemoApp).ToNot(BeNil())
	})
	When("Creating", func() {
		DescribeTable("Some Messages",
			func(message string, shortmsg string, topic string, severity int, ts time.Time, expectedmessage string, expectedshortmsg string, expectedtopic string, expectedseverity int, expectedts time.Time) {
				newmsg, err := eClient.Message.Create().
					SetMessage(message).
					SetShortMsg(shortmsg).
					SetTopic(topic).
					SetSeverity(severity).
					SetTimestamp(ts).
					SetApp(DemoApp).
					Save(ctx)
				Expect(err).To(BeNil())
				Expect(newmsg).ToNot(BeNil())
				Expect(newmsg.Message).To(Equal(expectedmessage))
				Expect(newmsg.ShortMsg).To(HaveValue(Equal(expectedshortmsg)))
				Expect(newmsg.Topic).To(HaveValue(Equal(expectedtopic)))
				Expect(newmsg.Severity).To(Equal(expectedseverity))
				Expect(newmsg.Timestamp).To(Equal(expectedts))
			},
			Entry("Normal Message", "message", "shortmsg",  "topic", 3, time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), "message", "shortmsg", "topic", 3, time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)),
			Entry("Another Normal Message", "anothermessage", "anothershortmsg", "anothertopic", 1, time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC), "anothermessage", "anothershortmsg", "anothertopic", 1, time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)),
		)
	})
	When("Creating Messages with Fields", func() {
		It("Should be ok", func() {
			var err error
			DemoMsg, err = eClient.Message.Create().
				SetMessage("message").
				SetShortMsg("shortmsg").
				SetTopic("topic").
				SetSeverity(3).
				SetTimestamp(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)).
				SetApp(DemoApp).
				Save(ctx)

			var1, err1 := eClient.MsgVar.Create().SetName("key1").SetValue("value1").SetOwner(DemoMsg).Save(ctx)
			var2, err2 := eClient.MsgVar.Create().SetName("key2").SetValue("value2").SetOwner(DemoMsg).Save(ctx)

			Expect(err).To(BeNil())
			Expect(DemoMsg).ToNot(BeNil())
			Expect(DemoMsg.Message).To(Equal("message"))
			Expect(DemoMsg.ShortMsg).To(HaveValue(Equal("shortmsg")))
			Expect(DemoMsg.Topic).To(HaveValue(Equal("topic")))
			Expect(DemoMsg.Severity).To(Equal(3))
			Expect(DemoMsg.Timestamp).To(Equal(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)))

			Expect(err1).To(BeNil())
			Expect(var1).ToNot(BeNil())
			Expect(var1.Name).To(Equal("key1"))
			Expect(var1.Value).To(Equal("value1"))
			Expect(err2).To(BeNil())
			Expect(var2).ToNot(BeNil())
			Expect(var2.Name).To(Equal("key2"))
			Expect(var2.Value).To(Equal("value2"))


			dbmsg, err := eClient.Message.Query().Where(message.ID(DemoMsg.ID)).WithVars().WithApp().Only(ctx)

			Expect(err).To(BeNil())
			Expect(dbmsg).ToNot(BeNil())
			Expect(dbmsg.Message).To(Equal("message"))
			Expect(dbmsg.ShortMsg).To(HaveValue(Equal("shortmsg")))
			Expect(dbmsg.Topic).To(HaveValue(Equal("topic")))
			Expect(dbmsg.Severity).To(Equal(3))
			Expect(dbmsg.Timestamp).To(Equal(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)))
			Expect(dbmsg.Edges.App.ID).To(Equal(DemoApp.ID))
			Expect(dbmsg.Edges.Vars).To(HaveLen(2))
			Expect(dbmsg.Edges.Vars[0].Name).To(Equal("key1"))
			Expect(dbmsg.Edges.Vars[0].Value).To(Equal("value1"))
			Expect(dbmsg.Edges.Vars[1].Name).To(Equal("key2"))
			Expect(dbmsg.Edges.Vars[1].Value).To(Equal("value2"))
		})	
		It("Cannot Have Duplicate Fields", func() {
			newmsg, err := eClient.Message.Query().WithVars().Where(message.ID(DemoMsg.ID)).First(ctx)
			Expect(err).To(BeNil())
			Expect(newmsg).ToNot(BeNil())
			var3, err := eClient.MsgVar.Create().SetName("key1").SetValue("value1").SetOwner(newmsg).Save(ctx)
			Expect(err).ToNot(BeNil())
			Expect(var3).To(BeNil())
		})
	})
	When("Retrieving Messages", func() {
		It("Should have 3 messages", func() {
			messages, err := eClient.Message.Query().All(ctx)
			Expect(err).To(BeNil())
			Expect(messages).ToNot(BeNil())
			Expect(len(messages)).To(Equal(3))
		})
	})
	When("Message Must be Associated with a App", func() {
		It("Should fail to create a Message without an App", func() {
			_, err := eClient.Message.Create().
				SetMessage("message").
				SetShortMsg("shortmsg").
				SetTopic("topic").
				SetSeverity(3).
				SetTimestamp(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)).
				Save(ctx)
			Expect(err).ToNot(BeNil())
		})
		It("Should have the Apps", func() {
			msg, err := eClient.Message.Query().WithApp().First(ctx)
			Expect(err).To(BeNil())
			Expect(msg).ToNot(BeNil())
			Expect(msg.Edges.App).ToNot(BeNil())
			Expect(msg.Edges.App.Name).To(Equal(DemoApp.Name))
		})
	})
})