package group_test

import (
	_ "fmt"
	"testing"

	_ "github.com/Fishwaldo/mouthpiece/pkg/ent/runtime"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	mptest "github.com/Fishwaldo/mouthpiece/pkg/mocks"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"
	"github.com/Fishwaldo/mouthpiece/pkg/transport"
	"github.com/Fishwaldo/mouthpiece/pkg/msg"


	"github.com/Fishwaldo/mouthpiece/pkg/transport/noop"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGroup(t *testing.T) {
	mptest.Tst = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Group Suite")
}

var _ = mptest.DBBeforeSuite
var _ = mptest.DBAfterSuite

var globalGroup interfaces.GroupI

var _ = Describe("Group", func() {
	Context("Create GroupService", func() {
		It("should create a new group", func() {
			var err error
			globalGroup, err = mptest.Mp.GetGroupService().Create(mptest.Ctx, "TestGroup", "Desciption")
			Expect(err).To(BeNil())
			Expect(globalGroup).ToNot(BeNil())
			Expect(globalGroup.GetName()).To(Equal("TestGroup"))
			Expect(globalGroup.GetDescription()).To(Equal("Desciption"))
		})
		It("Should Not Create a Duplicate Group Name", func() {
			_, err := mptest.Mp.GetGroupService().Create(mptest.Ctx, "TestGroup", "Desciption")
			Expect(err).To(MatchError(mperror.ErrGroupExists))
		})
		It("Should allow more than 1 group with the same description", func() {
			_, err := mptest.Mp.GetGroupService().Create(mptest.Ctx, "TestGroup2", "Desciption")
			Expect(err).To(BeNil())
		})
	})
	Context("Get A Group", func() {
		It("Should get a group by name", func() {
			grp, err := mptest.Mp.GetGroupService().Get(mptest.Ctx, "TestGroup")
			Expect(err).To(BeNil())
			Expect(grp).ToNot(BeNil())
			Expect(grp.GetName()).To(Equal("TestGroup"))
			Expect(grp.GetDescription()).To(Equal("Desciption"))
		})
		It("Should get a group by ID", func() {
			grp, err := mptest.Mp.GetGroupService().GetByID(mptest.Ctx, globalGroup.GetID())
			Expect(err).To(BeNil())
			Expect(grp).ToNot(BeNil())
			Expect(grp.GetName()).To(Equal("TestGroup"))
			Expect(grp.GetDescription()).To(Equal("Desciption"))
		})
		It("Exists Should work", func() {
			exists, err := mptest.Mp.GetGroupService().Exists(mptest.Ctx, "TestGroup")
			Expect(err).To(BeNil())
			Expect(exists).To(BeTrue())
		})
		It("ExistsByID should work", func() {
			exists, err := mptest.Mp.GetGroupService().ExistsByID(mptest.Ctx, globalGroup.GetID())
			Expect(err).To(BeNil())
			Expect(exists).To(BeTrue())
		})
	})
	Context("Group Data", func() {
		It("Should be able to set and get Name", func() {
			err := globalGroup.SetName(mptest.Ctx, "NewName")
			Expect(err).To(BeNil())
			Expect(globalGroup.GetName()).To(Equal("NewName"))
		})
		It("Should be able to set and get Description", func() {
			err := globalGroup.SetDescription(mptest.Ctx, "NewDescription")
			Expect(err).To(BeNil())
			Expect(globalGroup.GetDescription()).To(Equal("NewDescription"))
		})
	})
	Context("Manage Users to a Group", func() {
		var user interfaces.UserI
		It("Add A User to a Group", func() {
			var err error
			user, err = mptest.Mp.GetUserService().Create(mptest.Ctx, "TestUser@example.com", "TestUser")
			Expect(err).To(BeNil())
			Expect(user).ToNot(BeNil())
			err = globalGroup.AddUser(mptest.Ctx, user)
			Expect(err).To(BeNil())
		})
		It("Get A User from a group", func() {
			users := globalGroup.GetUsers(mptest.Ctx)
			Expect(users).To(HaveLen(1))
			Expect(users[0].GetID()).To(Equal(user.GetID()))
		})
		It("Delete Users from a Group", func() {
			err := globalGroup.DelUser(mptest.Ctx, user)
			Expect(err).To(BeNil())
			users := globalGroup.GetUsers(mptest.Ctx)
			Expect(users).To(HaveLen(0))
		})
	})
	Context("Manage Apps to a Group", func() {
		It("Add an App to a Group", func() {
			app, err := mptest.Mp.GetAppService().Create(mptest.Ctx, "TestApp", "TestApp")
			Expect(err).To(BeNil())
			Expect(app).ToNot(BeNil())
			err = globalGroup.AddApp(mptest.Ctx, app)
			Expect(err).To(BeNil())
		})
		It("Get Apps from a Group", func() {
			apps := globalGroup.GetApps(mptest.Ctx)
			Expect(apps).To(HaveLen(1))
			Expect(apps[0].GetName()).To(Equal("TestApp"))

		})
		It("Delete an App from a Group", func() {
			app, err := mptest.Mp.GetAppService().Get(mptest.Ctx, "TestApp")
			Expect(err).To(BeNil())
			Expect(app).ToNot(BeNil())
			err = globalGroup.DelApp(mptest.Ctx, app)
			Expect(err).To(BeNil())
			apps := globalGroup.GetApps(mptest.Ctx)
			Expect(apps).To(HaveLen(0))
		})
	})
	Context("Manage Transport Recipients", func() {
		It("Add a Transport Recipient", func() {
			tpp, err := transport.GetTransportProvider(mptest.Ctx, "noop")
			Expect(err).To(BeNil())
			Expect(tpp).ToNot(BeNil())

			tpi, err := mptest.Mp.GetTransportService().CreateTransportInstance(mptest.Ctx, tpp, "NoOp Transport", &noop.NoOpConfig{})
			Expect(err).To(BeNil())
			Expect(tpi).ToNot(BeNil())

			tpr, err := mptest.Mp.GetTransportService().Create(mptest.Ctx, tpi, "TestNoOp Transport Recipient", &noop.NoOpRecipientConfig{})
			Expect(err).To(BeNil())
			Expect(tpr).ToNot(BeNil())

			err = globalGroup.AddTransportRecipient(mptest.Ctx, tpr)
			Expect(err).To(BeNil())
		})
		It("Get Transport Recipients", func() {
			tprs := globalGroup.GetTransportRecipients(mptest.Ctx)
			Expect(tprs).To(HaveLen(1))
			Expect(tprs[0].GetName()).To(Equal("TestNoOp Transport Recipient"))
		})
		It("Delete a Transport Recipient", func() {
			tpr, err := mptest.Mp.GetTransportService().Get(mptest.Ctx, "TestNoOp Transport Recipient")
			Expect(err).To(BeNil())
			Expect(tpr).ToNot(BeNil())
			err = globalGroup.DelTransportRecipient(mptest.Ctx, tpr)
			Expect(err).To(BeNil())
			tprs := globalGroup.GetTransportRecipients(mptest.Ctx)
			Expect(tprs).To(HaveLen(0))
		})
	})
	Context("Get All Groups", func() {
		It("Should Return All Groups", func() {
			groups, err := mptest.Mp.GetGroupService().GetAll(mptest.Ctx)
			Expect(err).To(BeNil())
			Expect(groups).To(HaveLen(3))
			/* remember we seeded the database */
			Expect(groups[0].GetName()).To(Equal("MouthPiece"))
			Expect(groups[1].GetName()).To(Equal("NewName"))
			Expect(groups[2].GetName()).To(Equal("TestGroup2"))

		})
	})
	Context("Process Message", func() {
		It("Should Process a Message", func() {
			newuser, err := mptest.Mp.GetUserService().Create(mptest.Ctx, "dummy@example.com", "dummy")
			Expect(err).To(BeNil())
			Expect(newuser).ToNot(BeNil())

			tpr, err := mptest.Mp.GetTransportService().Get(mptest.Ctx, "TestNoOp Transport Recipient")
			Expect(err).To(BeNil())
			Expect(tpr).ToNot(BeNil())

			err = globalGroup.AddTransportRecipient(mptest.Ctx, tpr)
			Expect(err).To(BeNil())
			err = globalGroup.AddUser(mptest.Ctx, newuser)
			Expect(err).To(BeNil())

			mymsg:= msg.NewMessage(mptest.Ctx, "Message", mptest.App)
			
			Expect(mymsg).ToNot(BeNil())
			err = globalGroup.ProcessMessage(mptest.Ctx, mymsg)
			Expect(err).To(BeNil())
			_, ok := noop.Messages[mymsg.GetID()]
			Expect(ok).To(BeTrue())
		})
	})
	Context("Delete Groups", func() {
		It("Should Work", func() {
			err := mptest.Mp.GetGroupService().Delete(mptest.Ctx, globalGroup)
			Expect(err).To(BeNil())
			ok, err := mptest.Mp.GetGroupService().Exists(mptest.Ctx, "NewName")
			Expect(err).To(BeNil())
			Expect(ok).To(BeFalse())
		})
	})

})

// 	Context("Create", func() {
// 		It("should create a new group", func() {
// 			var err error
// 			globalGroup, err = group.NewGroupService().Create(ctx,"TestGroup","TestGroupDescription");
// 			Expect(err).To(BeNil())
// 			Expect(globalGroup).ToNot(BeNil())
// 		})
// 		It("Should Not Create a Duplicate Group Name", func() {
// 			_, err := group.NewGroupService().Create(ctx,"TestGroup","TestGroupDescription");
// 			Expect(err).ToNot(BeNil())
// 		})
// 		It("Should allow more than 1 group with the same description", func() {
// 			_, err := group.NewGroupService().Create(ctx,"TestGroup2","TestGroupDescription");
// 			Expect(err).To(BeNil())
// 		})
// 	})
// 	Context("Get", func() {
// 		It("should get a group by Name", func() {
// 			g2, err := groupService.Get(ctx, "TestGroup")
// 			Expect(err).To(BeNil())
// 			Expect(g2).ToNot(BeNil())
// 			Expect(g2.GetID()).To(Equal(globalGroup.GetID()))
// 		})
// 		It("should get a group by ID", func() {
// 			g2, err := groupService.GetByID(ctx, globalGroup.GetID())
// 			Expect(err).To(BeNil())
// 			Expect(g2).ToNot(BeNil())
// 			Expect(g2.GetID()).To(Equal(globalGroup.GetID()))
// 		})
// 		It("Should Get Group by DBGroup", func() {
// 			dbgroup, err := eClient.DbGroup.Get(ctx, globalGroup.GetID())
// 			Expect(err).To(BeNil())
// 			Expect(dbgroup).ToNot(BeNil())
// 			g2, err := groupService.Load(ctx, dbgroup)
// 			Expect(err).To(BeNil())
// 			Expect(g2).ToNot(BeNil())
// 			Expect(g2.GetID()).To(Equal(globalGroup.GetID()))
// 		})
// 	})
// 	Context("Modify", func() {
// 		It("Should Modify Group Name", func() {
// 			err := globalGroup.SetName(ctx, "TestGroup3")
// 			Expect(err).To(BeNil())
// 			Expect(globalGroup.GetName()).To(Equal("TestGroup3"))
// 		})
// 		It("Should Modify Group Description", func() {
// 			err := globalGroup.SetDescription(ctx, "TestGroupDescription2")
// 			Expect(err).To(BeNil())
// 			Expect(globalGroup.GetDescription()).To(Equal("TestGroupDescription2"))
// 		})
// 	})
// 	Context("Delete", func() {
// 		It("Should Delete Group", func() {
// 			grp2, err := groupService.Get(ctx, "TestGroup2")
// 			Expect(err).To(BeNil())
// 			Expect(grp2).ToNot(BeNil())
// 			err = groupService.Delete(ctx, grp2)
// 			Expect(err).To(BeNil())
// 		})
// 	})

// })
