package transport_test

import (
	"testing"

	"github.com/Fishwaldo/mouthpiece/pkg/transport/noop"
	"github.com/Fishwaldo/mouthpiece/pkg/msg"
		"github.com/Fishwaldo/mouthpiece/pkg/mperror"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	mptest "github.com/Fishwaldo/mouthpiece/pkg/mocks"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = mptest.DBBeforeSuite
var _ = mptest.DBAfterSuite

var globalTransportProvider interfaces.TransportProvider
var globalTransportInstance interfaces.TransportInstance
var globalTransportRecipient interfaces.TransportRecipient
var globalTransportRecipient2 interfaces.TransportRecipient

func TestTransport(t *testing.T) {
	mptest.Tst = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Transport Suite")
}

var _ = Describe("Transport", func() {
	Context("Transport Provider", func() {
		It("Should List Transport Providers", func() {
			providers, err := mptest.Mp.GetTransportService().GetAllTransportProviders(mptest.Ctx)
			Expect(providers).To(HaveLen(2))
			Expect(err).To(BeNil())
		})
		It("Should Get a Transport Provider", func() {
			var err error
			globalTransportProvider, err = mptest.Mp.GetTransportService().GetTransportProvider(mptest.Ctx, "noop")
			Expect(err).To(BeNil())
			Expect(globalTransportProvider).ToNot(BeNil())
			Expect(globalTransportProvider.GetName()).To(Equal("noop"))
		})
	})
	Context("Transport Instance", func() {
		It("Should Create a New Transport Instance", func() {
			var err error
			cfg, err := globalTransportProvider.LoadConfigFromJSON(mptest.Ctx, "{}")
			Expect(err).To(BeNil())
			Expect(cfg).To(BeAssignableToTypeOf(&noop.NoOpConfig{}))
			globalTransportInstance, err = mptest.Mp.GetTransportService().CreateTransportInstance(mptest.Ctx, globalTransportProvider, "test", cfg)
			Expect(err).To(BeNil())
			Expect(globalTransportInstance).ToNot(BeNil())
			Expect(globalTransportInstance.GetName()).To(Equal("test"))
		})
		It("Should Get a Transport Instance", func() {
			tpi, err := mptest.Mp.GetTransportService().GetTransportInstance(mptest.Ctx, "test")
			Expect(err).To(BeNil())
			Expect(tpi).ToNot(BeNil())
			Expect(tpi.GetName()).To(Equal("test"))
		})
		It("Should get a transport instance by ID", func() {
			tpi, err := mptest.Mp.GetTransportService().GetTransportInstanceByID(mptest.Ctx, globalTransportInstance.GetID())
			Expect(err).To(BeNil())
			Expect(tpi).ToNot(BeNil())
			Expect(tpi.GetName()).To(Equal("test"))
		})
		It("Should Get All Transport Instances", func() {
			tpis, err := mptest.Mp.GetTransportService().GetAllTransportInstances(mptest.Ctx)
			Expect(err).To(BeNil())
			Expect(tpis).ToNot(BeNil())
			Expect(tpis).To(HaveLen(2))
		})
		It("Should Delete a Transport Instance", func() {
			var err error
			cfg, err := globalTransportProvider.LoadConfigFromJSON(mptest.Ctx, "{}")
			Expect(err).To(BeNil())
			Expect(cfg).To(BeAssignableToTypeOf(&noop.NoOpConfig{}))
			tpi, err := mptest.Mp.GetTransportService().CreateTransportInstance(mptest.Ctx, globalTransportProvider, "test2", cfg)
			Expect(err).To(BeNil())
			Expect(tpi).ToNot(BeNil())
			Expect(tpi.GetName()).To(Equal("test2"))
			err = mptest.Mp.GetTransportService().DeleteTransportInstance(mptest.Ctx, tpi)
			Expect(err).To(BeNil())
		})
		It("SHould Check if a Transport Exists by Name", func() {
			exists := mptest.Mp.GetTransportService().ExistsTransportInstance(mptest.Ctx, "test")
			Expect(exists).To(BeTrue())
		})
		It("SHould Check of a Transport Exists by ID", func() {
			exists := mptest.Mp.GetTransportService().ExistsByIDTransportInstance(mptest.Ctx, globalTransportInstance.GetID())
			Expect(exists).To(BeTrue())
		})
		It("Should Start Transport Services", func() {
			err := mptest.Mp.GetTransportService().Start(mptest.Ctx)
			Expect(err).To(BeNil())
		})
		It("Should Set the Name of a Transport Instance", func() {
			err := globalTransportInstance.SetName(mptest.Ctx, "test2")
			Expect(err).To(BeNil())
		})
		It("Should Get the Name of a Transport Instance", func() {
			name := globalTransportInstance.GetName()
			Expect(name).To(Equal("test2"))
		})
		It("Should Set the Description of a Transport Instance", func() {
			err := globalTransportInstance.SetDescription(mptest.Ctx, "test2")
			Expect(err).To(BeNil())
		})
		It("Should Get the Description of a Transport Instance", func() {
			desc := globalTransportInstance.GetDescription()
			Expect(desc).To(Equal("test2"))
		})
		It("Should Set the Config of a Transport Instance", func() {
			cfg, err := globalTransportProvider.LoadConfigFromJSON(mptest.Ctx, "{}")
			Expect(err).To(BeNil())
			Expect(cfg).To(BeAssignableToTypeOf(&noop.NoOpConfig{}))
			err = globalTransportInstance.SetConfig(mptest.Ctx, cfg)
			Expect(err).To(BeNil())
		})
		It("Should Get the Config of a Transport Instance", func() {
			cfg, err := globalTransportInstance.GetConfig(mptest.Ctx)
			Expect(err).To(BeNil())
			Expect(cfg).To(BeAssignableToTypeOf(&noop.NoOpConfig{}))
		})


	})
	Context("Transport Recipients", func() {
		It("Should Create a New Transport Recipient", func() {
			cfg, err := globalTransportProvider.LoadConfigFromJSON(mptest.Ctx, "{}")
			Expect(err).To(BeNil())
			Expect(cfg).To(BeAssignableToTypeOf(&noop.NoOpConfig{}))
			globalTransportRecipient, err = mptest.Mp.GetTransportService().Create(mptest.Ctx, globalTransportInstance, "testrecipient", cfg)
			Expect(err).To(BeNil())
			Expect(globalTransportRecipient).ToNot(BeNil())
			Expect(globalTransportRecipient.GetName()).To(Equal("testrecipient"))
		})
		It("Should Get By ID", func() {
			tr, err := mptest.Mp.GetTransportService().GetByID(mptest.Ctx, globalTransportRecipient.GetID())
			Expect(err).To(BeNil())
			Expect(tr).ToNot(BeNil())
			Expect(tr.GetName()).To(Equal("testrecipient"))
		})
		It("Should Get By Name", func() {
			tr, err := mptest.Mp.GetTransportService().Get(mptest.Ctx, "testrecipient")
			Expect(err).To(BeNil())
			Expect(tr).ToNot(BeNil())
			Expect(tr.GetName()).To(Equal("testrecipient"))
		})
		It("Should Check if Exists by Name", func() {
			exists := mptest.Mp.GetTransportService().Exists(mptest.Ctx, "testrecipient")
			Expect(exists).To(BeTrue())
		})
		It("SHould Check if exists by ID", func() {
			exists := mptest.Mp.GetTransportService().ExistsByID(mptest.Ctx, globalTransportRecipient.GetID())
			Expect(exists).To(BeTrue())
		})
		It("SHould Get All Transport Reciepients", func() {
			trs, err := mptest.Mp.GetTransportService().GetAll(mptest.Ctx)
			Expect(err).To(BeNil())
			Expect(trs).ToNot(BeNil())
			Expect(trs).To(HaveLen(3))
		})
		It("Should Set the Name of a Transport Recipient", func() {
			err := globalTransportRecipient.SetName(mptest.Ctx, "testrecipient2")
			Expect(err).To(BeNil())
		})
		It("Should Get the Name of a Transport Recipient", func() {
			name := globalTransportRecipient.GetName()
			Expect(name).To(Equal("testrecipient2"))
		})
		It("Should Set the Description of a Transport Recipient", func() {
			err := globalTransportRecipient.SetDescription(mptest.Ctx, "testrecipient2")
			Expect(err).To(BeNil())
		})
		It("Should Get the Description of a Transport Recipient", func() {
			desc := globalTransportRecipient.GetDescription()
			Expect(desc).To(Equal("testrecipient2"))
		})
		It("Should Return a invalid Recipient Type", func() {
			rtype := globalTransportRecipient.GetRecipientType(mptest.Ctx)
			Expect(rtype).To(Equal(interfaces.TransportRecipientTypeNotSet))
		})
		It("Add a User Reciptient", func() {
			admin, err := mptest.Mp.GetUserService().Get(mptest.Ctx, "admin@example.com")
			Expect(err).To(BeNil())
			Expect(admin).ToNot(BeNil())
			err = globalTransportRecipient.SetUser(mptest.Ctx, admin)
			Expect(err).To(BeNil())
		})
		It("Should Return a User Recipient Type", func() {
			rtype := globalTransportRecipient.GetRecipientType(mptest.Ctx)
			Expect(rtype).To(Equal(interfaces.TransportRecipientTypeUser))
		})
		It("Should Get the User Recipient", func() {
			user, err := globalTransportRecipient.GetUser(mptest.Ctx)
			Expect(err).To(BeNil())
			Expect(user).ToNot(BeNil())
			Expect(user.GetEmail()).To(Equal("admin@example.com"))
		})
		It("Should Fail if we try to add a Group Recipient", func() {
			group, err := mptest.Mp.GetGroupService().Get(mptest.Ctx, "MouthPiece")
			Expect(err).To(BeNil())
			Expect(group).ToNot(BeNil())
			err = globalTransportRecipient.SetGroup(mptest.Ctx, group)
			Expect(err).ToNot(BeNil())
			Expect(err).To(MatchError(mperror.ErrTransportRecipientSet))
		})
		It("Should Add A group Recipient", func() {
			cfg, err := globalTransportProvider.LoadConfigFromJSON(mptest.Ctx, "{}")
			Expect(err).To(BeNil())
			Expect(cfg).To(BeAssignableToTypeOf(&noop.NoOpConfig{}))
			globalTransportRecipient2, err = mptest.Mp.GetTransportService().Create(mptest.Ctx, globalTransportInstance, "testrecipient4", cfg)
			Expect(err).To(BeNil())
			Expect(globalTransportRecipient2).ToNot(BeNil())
			Expect(globalTransportRecipient2.GetName()).To(Equal("testrecipient4"))

			group, err := mptest.Mp.GetGroupService().Get(mptest.Ctx, "MouthPiece")
			Expect(err).To(BeNil())
			Expect(group).ToNot(BeNil())
			err = globalTransportRecipient2.SetGroup(mptest.Ctx, group)
			Expect(err).To(BeNil())
		})
		It("Should Return a Group Recipient Type", func() {
			rtype := globalTransportRecipient2.GetRecipientType(mptest.Ctx)
			Expect(rtype).To(Equal(interfaces.TransportRecipientTypeGroup))
		})
		It("Should Get the Group Recipient", func() {
			group, err := globalTransportRecipient2.GetGroup(mptest.Ctx)
			Expect(err).To(BeNil())
			Expect(group).ToNot(BeNil())
			Expect(group.GetName()).To(Equal("MouthPiece"))
		})
		It("Should Fail if we try to add a User Recipient", func() {
			admin, err := mptest.Mp.GetUserService().Get(mptest.Ctx, "admin@example.com")
			Expect(err).To(BeNil())
			Expect(admin).ToNot(BeNil())
			err = globalTransportRecipient2.SetUser(mptest.Ctx, admin)
			Expect(err).ToNot(BeNil())
			Expect(err).To(MatchError(mperror.ErrTransportRecipientSet))
		})
		It("Should Be able to set Config", func() {
			cfg, err := globalTransportInstance.LoadTransportReciepientConfig(mptest.Ctx, "{}")
			Expect(err).To(BeNil())
			Expect(cfg).To(BeAssignableToTypeOf(&noop.NoOpRecipientConfig{}))
			err = globalTransportRecipient2.SetConfig(mptest.Ctx, cfg)
			Expect(err).To(BeNil())
		})
		It("Should Be able to get Config", func() {
			cfg, err := globalTransportRecipient2.GetConfig()
			Expect(err).To(BeNil())
			Expect(cfg).To(BeAssignableToTypeOf(&noop.NoOpRecipientConfig{}))
		})
		It("Should Process a Message", func() {
			app, err := mptest.Mp.GetAppService().Get(mptest.Ctx, "MouthPiece")
			Expect(err).To(BeNil())
			Expect(app).ToNot(BeNil())
			msg := msg.NewMessage(mptest.Ctx, "Hello World", app)
			err = globalTransportRecipient2.ProcessMessage(mptest.Ctx, msg)
			Expect(err).To(BeNil())
			_, ok := noop.Messages[msg.GetID()]
			Expect(ok).To(BeTrue())
		})
	})
	Context("Stop and Delete", func() {
		It("SHould Stop the TransportInstances", func() {
			err := mptest.Mp.GetTransportService().Stop(mptest.Ctx)
			Expect(err).To(BeNil())
		})
		It("Should Delete the TransportInstances", func() {
			err := mptest.Mp.GetTransportService().Delete(mptest.Ctx, globalTransportRecipient2)
			Expect(err).To(BeNil())
		})
	})
})
