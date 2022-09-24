package user_test

import (
	"testing"

	noopflt "github.com/Fishwaldo/mouthpiece/pkg/filter/noop"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	mptest "github.com/Fishwaldo/mouthpiece/pkg/mocks"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"
	"github.com/Fishwaldo/mouthpiece/pkg/msg"
	noop "github.com/Fishwaldo/mouthpiece/pkg/transport/noop"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = mptest.DBBeforeSuite
var _ = mptest.DBAfterSuite

func TestUser(t *testing.T) {
	mptest.Tst = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Suite")
}

var globalUser interfaces.UserI

var _ = Describe("User", func() {
	Context("UserService", func() {
		It("should be able to create a user", func() {
			var err error
			globalUser, err = mptest.Mp.GetUserService().Create(mptest.Ctx, "test@test.com", "test user")
			Expect(err).To(BeNil())
			Expect(globalUser).ToNot(BeNil())
			Expect(globalUser.GetEmail()).To(Equal("test@test.com"))
		})
		It("Should be able to get a user by email", func() {
			user, err := mptest.Mp.GetUserService().Get(mptest.Ctx, "test@test.com")
			Expect(err).To(BeNil())
			Expect(user).ToNot(BeNil())
			Expect(user.GetEmail()).To(Equal("test@test.com"))
		})
		It("Should be able to get a user by ID", func() {
			user, err := mptest.Mp.GetUserService().GetByID(mptest.Ctx, globalUser.GetID())
			Expect(err).To(BeNil())
			Expect(user).ToNot(BeNil())
			Expect(user.GetEmail()).To(Equal("test@test.com"))
		})
		It("Should be able to get get All Users", func() {
			users, err := mptest.Mp.GetUserService().GetAll(mptest.Ctx)
			Expect(err).To(BeNil())
			Expect(users).ToNot(BeNil())
			Expect(len(users)).To(Equal(3))
		})
		It("Should be able to Test if exists by emaail", func() {
			ok, err := mptest.Mp.GetUserService().Exists(mptest.Ctx, "test@test.com")
			Expect(err).To(BeNil())
			Expect(ok).To(BeTrue())
		})
		It("SHould be able to test if exists by email", func() {
			ok, err := mptest.Mp.GetUserService().ExistsByID(mptest.Ctx, globalUser.GetID())
			Expect(err).To(BeNil())
			Expect(ok).To(BeTrue())
		})
		It("Should be able to test if exists by ID", func() {
			ok, err := mptest.Mp.GetUserService().ExistsByID(mptest.Ctx, globalUser.GetID())
			Expect(err).To(BeNil())
			Expect(ok).To(BeTrue())
		})
		It("Should be able to Delete a user", func() {
			user, err := mptest.Mp.GetUserService().Create(mptest.Ctx, "test@delete.me", "test user")
			Expect(err).To(BeNil())
			Expect(user).ToNot(BeNil())
			Expect(user.GetEmail()).To(Equal("test@delete.me"))
			err = mptest.Mp.GetUserService().Delete(mptest.Ctx, user)
			Expect(err).To(BeNil())
		})
		It("Should not allow duplicate email addresses", func() {
			_, err := mptest.Mp.GetUserService().Create(mptest.Ctx, "test@test.com", "test user")
			Expect(err).ToNot(BeNil())
			Expect(err).To(MatchError(mperror.ErrUserExists))
		})
	})
	Context("User", func() {
		It("Should be able to set and get the email", func() {
			err := globalUser.SetEmail(mptest.Ctx, "test2@test.com")
			Expect(err).To(BeNil())
			Expect(globalUser.GetEmail()).To(Equal("test2@test.com"))
		})
		It("Should be able to set and get the name", func() {
			err := globalUser.SetName(mptest.Ctx, "test user2")
			Expect(err).To(BeNil())
			Expect(globalUser.GetName()).To(Equal("test user2"))
		})
		It("Should reject invalid email addresses", func() {
			err := globalUser.SetEmail(mptest.Ctx, "test2test.com")
			Expect(err).ToNot(BeNil())
			Expect(err).To(MatchError(mperror.ErrValidationError))
		})
		It("Should be able to get and set the description", func() {
			err := globalUser.SetDescription(mptest.Ctx, "test description")
			Expect(err).To(BeNil())
			Expect(globalUser.GetDescription()).To(Equal("test description"))
		})
		It("Should be able to Set and Get a Metadata Field", func() {
			err := globalUser.SetField(mptest.Ctx, "test", "test value")
			Expect(err).To(BeNil())
			val, err := globalUser.GetField(mptest.Ctx, "test")
			Expect(err).To(BeNil())
			Expect(val).To(Equal("test value"))
		})
		It("Should be able to set and get Metadata Fields", func() {
			flds := map[string]string{
				"test":  "test value222",
				"test4": "test value2",
			}
			err := globalUser.SetFields(mptest.Ctx, flds)
			Expect(err).To(BeNil())
			vals, err := globalUser.GetFields(mptest.Ctx)
			Expect(err).To(BeNil())
			validate := map[string]string{
				"test":  "test value222",
				"test4": "test value2",
			}
			Expect(vals).To(Equal(validate))
		})
		It("Should be able to add Filters", func() {
			flt, err := mptest.Mp.GetFilterService().Create(mptest.Ctx, "NoOpFilter", "newfilter", interfaces.UserFilter)
			Expect(err).To(BeNil())
			Expect(flt).ToNot(BeNil())
			err = globalUser.AddFilter(mptest.Ctx, flt)
			Expect(err).To(BeNil())
		})
		It("Should be able to Get Filters", func() {
			flts, err := globalUser.GetFilters(mptest.Ctx)
			Expect(err).To(BeNil())
			Expect(len(flts)).To(Equal(1))
		})
		It("Should be able to remove Filters", func() {
			flt, err := mptest.Mp.GetFilterService().Create(mptest.Ctx, "NoOpFilter", "delfilter", interfaces.UserFilter)
			Expect(err).To(BeNil())
			Expect(flt).ToNot(BeNil())
			err = globalUser.AddFilter(mptest.Ctx, flt)
			Expect(err).To(BeNil())
			err = globalUser.DelFilter(mptest.Ctx, flt)
			Expect(err).To(BeNil())
		})
		It("Should add Transport Recipients", func() {
			trp, err := mptest.Mp.GetTransportService().GetTransportProvider(mptest.Ctx, "noop")
			Expect(err).To(BeNil())
			Expect(trp).ToNot(BeNil())
			Expect(trp.GetName()).To(Equal("noop"))
			cfg, err := trp.LoadConfigFromJSON(mptest.Ctx, "{}")
			Expect(err).To(BeNil())
			Expect(cfg).To(BeAssignableToTypeOf(&noop.NoOpConfig{}))
			tpi, err := mptest.Mp.GetTransportService().CreateTransportInstance(mptest.Ctx, trp, "test", cfg)
			Expect(err).To(BeNil())
			Expect(tpi).ToNot(BeNil())
			Expect(tpi.GetName()).To(Equal("test"))
			tprcfg, err := tpi.LoadTransportReciepientConfig(mptest.Ctx, "{}")
			Expect(err).To(BeNil())
			Expect(tprcfg).To(BeAssignableToTypeOf(&noop.NoOpRecipientConfig{}))
			tpr, err := mptest.Mp.GetTransportService().Create(mptest.Ctx, tpi, "test", tprcfg)
			Expect(err).To(BeNil())
			Expect(tpr).ToNot(BeNil())
			err = globalUser.AddTransportRecipient(mptest.Ctx, tpr)
			Expect(err).To(BeNil())
		})
		It("SHould be able to get Transport Recipients", func() {
			trps := globalUser.GetTransportRecipients(mptest.Ctx)
			Expect(len(trps)).To(Equal(1))
		})
		It("SHould be able to Process a Message with Filters and Transport Recipients", func() {
			app, err := mptest.Mp.GetAppService().Get(mptest.Ctx, "MouthPiece")
			Expect(err).To(BeNil())
			Expect(app).ToNot(BeNil())
			newmsg := msg.NewMessage(mptest.Ctx, "Hello World", app)
			Expect(newmsg).ToNot(BeNil())
			err = globalUser.ProcessMessage(mptest.Ctx, newmsg)
			Expect(err).To(BeNil())
			_, ok := noopflt.Messages[newmsg.GetID()]
			Expect(ok).To(BeTrue())
			_, ok = noop.Messages[newmsg.GetID()]
			Expect(ok).To(BeTrue())
		})
		It("SHould be able to Delete a Transport Recipient", func() {
			trps := globalUser.GetTransportRecipients(mptest.Ctx)
			Expect(len(trps)).To(Equal(1))
			err := globalUser.DelTransportRecipient(mptest.Ctx, trps[0])
			Expect(err).To(BeNil())
			trps = globalUser.GetTransportRecipients(mptest.Ctx)
			Expect(len(trps)).To(Equal(0))
		})
	})
})
