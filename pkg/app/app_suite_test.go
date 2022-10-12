package app_test

import (
	"testing"

	_ "github.com/Fishwaldo/mouthpiece/pkg/ent/runtime"
	_ "github.com/Fishwaldo/mouthpiece/pkg/filter/noop"
	noopfilter "github.com/Fishwaldo/mouthpiece/pkg/filter/noop"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	mptest "github.com/Fishwaldo/mouthpiece/pkg/mocks"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"
	"github.com/Fishwaldo/mouthpiece/pkg/msg"
	"github.com/Fishwaldo/mouthpiece/pkg/transport"
	nooptrp "github.com/Fishwaldo/mouthpiece/pkg/transport/noop"

	//	"github.com/Fishwaldo/mouthpiece/pkg/ent"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestApp(t *testing.T) {
	mptest.Tst = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "App Suite")
}

var _ = mptest.DBBeforeSuite
var _ = mptest.DBAfterSuite

var globalApp interfaces.AppI
var globalflt interfaces.FilterI

var _ = Describe("App", func() {
	Context("Create AppService", func() {
		It("should create a new app", func() {
			var err error
			globalApp, err = mptest.Mp.GetAppService().Create(mptest.Ctx, "TestApp", "Desciption")
			Expect(err).To(BeNil())
			Expect(globalApp).ToNot(BeNil())
			Expect(globalApp.GetName()).To(Equal("TestApp"))
			Expect(globalApp.GetDescription()).To(Equal("Desciption"))
		})
		It("Should Not Create a Duplicate App Name", func() {
			_, err := mptest.Mp.GetAppService().Create(mptest.Ctx, "TestApp", "Desciption")
			Expect(err).To(MatchError(mperror.ErrAppExists))
		})
		It("Should allow more than 1 app with the same description", func() {
			_, err := mptest.Mp.GetAppService().Create(mptest.Ctx, "TestApp2", "Desciption")
			Expect(err).To(BeNil())
		})
	})
	Context("Get App", func() {
		It("Should get an app by name", func() {
			app, err := mptest.Mp.GetAppService().Get(mptest.Ctx, "TestApp")
			Expect(err).To(BeNil())
			Expect(app).ToNot(BeNil())
			Expect(app.GetName()).To(Equal("TestApp"))
			Expect(app.GetDescription()).To(Equal("Desciption"))
		})
		It("Should get an app by ID", func() {
			app, err := mptest.Mp.GetAppService().GetByID(mptest.Ctx, globalApp.GetID())
			Expect(err).To(BeNil())
			Expect(app).ToNot(BeNil())
			Expect(app.GetName()).To(Equal("TestApp"))
			Expect(app.GetDescription()).To(Equal("Desciption"))
		})
		It("Should Get All Apps", func() {
			apps, err := mptest.Mp.GetAppService().GetAll(mptest.Ctx)
			Expect(err).To(BeNil())
			Expect(apps).ToNot(BeNil())
			Expect(len(apps)).To(Equal(3))
			Expect(apps[0].GetName()).To(Equal("MouthPiece"))
			Expect(apps[1].GetName()).To(Equal("TestApp"))
			Expect(apps[1].GetDescription()).To(Equal("Desciption"))
			Expect(apps[2].GetName()).To(Equal("TestApp2"))
			Expect(apps[2].GetDescription()).To(Equal("Desciption"))
		})
		It("Should Test if Exists by name", func() {
			exists, err := mptest.Mp.GetAppService().Exists(mptest.Ctx, "TestApp")
			Expect(err).To(BeNil())
			Expect(exists).To(BeTrue())
		})
		It("Should Test if Exists by ID", func() {
			exists, err := mptest.Mp.GetAppService().ExistsByID(mptest.Ctx, globalApp.GetID())
			Expect(err).To(BeNil())
			Expect(exists).To(BeTrue())
		})
	})
	Context("Modify", func() {
		It("Should Modify Name", func() {
			err := globalApp.SetName(mptest.Ctx, "NewName")
			Expect(err).To(BeNil())
			Expect(globalApp.GetName()).To(Equal("NewName"))
		})
		It("Should Modify Description", func() {
			err := globalApp.SetDescription(mptest.Ctx, "NewDescription")
			Expect(err).To(BeNil())
			Expect(globalApp.GetDescription()).To(Equal("NewDescription"))
		})
		It("Should Modify App URL", func() {
			err := globalApp.SetURL(mptest.Ctx, "https://example.com")
			Expect(err).To(BeNil())
			Expect(globalApp.GetURL()).To(Equal("https://example.com"))
		})
		It("SHould Modify App Icon", func() {
			err := globalApp.SetIcon(mptest.Ctx, "https://example.com/icon.png")
			Expect(err).To(BeNil())
			Expect(globalApp.GetIcon()).To(Equal("https://example.com/icon.png"))
		})
		It("Should Reject Invalid URL's", func() {
			err := globalApp.SetURL(mptest.Ctx, "blah")
			Expect(err).ToNot(BeNil())
		})
		It("Should Reject Invalid Icon URL's", func() {
			err := globalApp.SetIcon(mptest.Ctx, "blah")
			Expect(err).ToNot(BeNil())
		})
		It("Should Be able to Set Status", func() {
			err := globalApp.SetStatus(mptest.Ctx, interfaces.Disabled)
			Expect(err).To(BeNil())
			Expect(globalApp.GetStatus()).To(Equal(interfaces.Disabled))
		})
	})
	Context("Process Messages", func() {
		It("Should Process a Message", func() {
			newuser, err := mptest.Mp.GetUserService().Create(mptest.Ctx, "dummy@example.com", "dummy")
			Expect(err).To(BeNil())
			Expect(newuser).ToNot(BeNil())

			tpp, err := transport.GetTransportProvider(mptest.Ctx, "noop")
			Expect(err).To(BeNil())
			Expect(tpp).ToNot(BeNil())

			tpi, err := mptest.Mp.GetTransportService().CreateTransportInstance(mptest.Ctx, tpp, "NoOp Transport", &nooptrp.NoOpConfig{})
			Expect(err).To(BeNil())
			Expect(tpi).ToNot(BeNil())

			tpr, err := mptest.Mp.GetTransportService().Create(mptest.Ctx, tpi, "TestNoOp Transport Recipient", &nooptrp.NoOpRecipientConfig{})
			Expect(err).To(BeNil())
			Expect(tpr).ToNot(BeNil())

			globalGroup, err := mptest.Mp.GetGroupService().Create(mptest.Ctx, "TestGroup", "Desciption")
			Expect(err).To(BeNil())
			Expect(globalGroup).ToNot(BeNil())
			Expect(globalGroup.GetName()).To(Equal("TestGroup"))
			Expect(globalGroup.GetDescription()).To(Equal("Desciption"))

			err = globalGroup.AddTransportRecipient(mptest.Ctx, tpr)
			Expect(err).To(BeNil())

			err = globalGroup.AddUser(mptest.Ctx, newuser)
			Expect(err).To(BeNil())

			err = globalGroup.AddApp(mptest.Ctx, globalApp)
			Expect(err).To(BeNil())

			globalflt, err = mptest.Mp.GetFilterService().Create(mptest.Ctx, "NoOpFilter", "NoOp Filter", interfaces.AppFilter)
			Expect(err).To(BeNil())
			Expect(globalflt).ToNot(BeNil())

			err = globalApp.AddFilter(mptest.Ctx, globalflt)
			Expect(err).To(BeNil())

			mymsg := msg.NewMessage(mptest.Ctx, "Message", mptest.App)
			Expect(mymsg).ToNot(BeNil())

			err = globalApp.ProcessMessage(mptest.Ctx, mymsg)
			Expect(err).To(BeNil())

			metadata, err := mymsg.GetMetadata(mptest.Ctx, interfaces.MD_AppName)
			Expect(err).To(BeNil())
			Expect(metadata).To(Equal("NewName"))

			_, ok := noopfilter.Messages[mymsg.GetID()]
			Expect(ok).To(BeTrue())

		})
	})
	Context("Filters", func() {
		It("Should List Filters", func() {
			filters, err := globalApp.GetFilters(mptest.Ctx)
			Expect(err).To(BeNil())
			Expect(filters).To(HaveLen(1))
		})
		It("Should Delete Filters", func() {
			err := globalApp.DelFilter(mptest.Ctx, globalflt)
			Expect(err).To(BeNil())
			filters, err := globalApp.GetFilters(mptest.Ctx)
			Expect(err).To(BeNil())
			Expect(filters).To(HaveLen(0))
		})
	})
})
