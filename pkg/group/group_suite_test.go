package group_test

import (
	"context"
	"testing"

	//	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/ent"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/rules"
	_ "github.com/Fishwaldo/mouthpiece/pkg/ent/runtime"
	"github.com/Fishwaldo/mouthpiece/pkg/group"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/log"

	"github.com/go-logr/logr/testr"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var tst *testing.T
func TestGroup(t *testing.T) {
	tst = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Group Suite")
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

})

var _ = AfterSuite(func() {
	eClient.Close()
})

var globalGroup interfaces.GroupI
var groupService interfaces.GroupServiceI


var _ = Describe("Group", func() {
	Context("Create GroupService", func() {
		It("Should Create a new Group Service", func() {
			groupService = group.NewGroupService()
			Expect(groupService).ToNot(BeNil())
			Expect(groupService.Start(context.Background())).ToNot(HaveOccurred())
		})
	})
	Context("Create", func() {
		It("should create a new group", func() {
			var err error
			globalGroup, err = group.NewGroupService().Create(ctx,"TestGroup","TestGroupDescription");
			Expect(err).To(BeNil())
			Expect(globalGroup).ToNot(BeNil())
		})
		It("Should Not Create a Duplicate Group Name", func() {
			_, err := group.NewGroupService().Create(ctx,"TestGroup","TestGroupDescription");
			Expect(err).ToNot(BeNil())
		})
		It("Should allow more than 1 group with the same description", func() {
			_, err := group.NewGroupService().Create(ctx,"TestGroup2","TestGroupDescription");
			Expect(err).To(BeNil())
		})
	})
	Context("Get", func() {
		It("should get a group by Name", func() {
			g2, err := groupService.Get(ctx, "TestGroup")
			Expect(err).To(BeNil())
			Expect(g2).ToNot(BeNil())
			Expect(g2.GetID()).To(Equal(globalGroup.GetID()))
		})
		It("should get a group by ID", func() {
			g2, err := groupService.GetByID(ctx, globalGroup.GetID())
			Expect(err).To(BeNil())
			Expect(g2).ToNot(BeNil())
			Expect(g2.GetID()).To(Equal(globalGroup.GetID()))
		})
		It("Should Get Group by DBGroup", func() {
			dbgroup, err := eClient.DbGroup.Get(ctx, globalGroup.GetID())
			Expect(err).To(BeNil())
			Expect(dbgroup).ToNot(BeNil())
			g2, err := groupService.Load(ctx, dbgroup)
			Expect(err).To(BeNil())
			Expect(g2).ToNot(BeNil())
			Expect(g2.GetID()).To(Equal(globalGroup.GetID()))
		})
	})
	Context("Modify", func() {
		It("Should Modify Group Name", func() {
			err := globalGroup.SetName(ctx, "TestGroup3")
			Expect(err).To(BeNil())
			Expect(globalGroup.GetName()).To(Equal("TestGroup3"))
		})
		It("Should Modify Group Description", func() {
			err := globalGroup.SetDescription(ctx, "TestGroupDescription2")
			Expect(err).To(BeNil())
			Expect(globalGroup.GetDescription()).To(Equal("TestGroupDescription2"))
		})
	})
	Context("Delete", func() {
		It("Should Delete Group", func() {
			grp2, err := groupService.Get(ctx, "TestGroup2")
			Expect(err).To(BeNil())
			Expect(grp2).ToNot(BeNil())
			err = groupService.Delete(ctx, grp2)
			Expect(err).To(BeNil())
		})
	})

})