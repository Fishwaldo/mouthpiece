package mock_interfaces



import (
	"context"
	"testing"
	"os"
	"time"
	"database/sql"

	"github.com/Fishwaldo/mouthpiece/pkg"
	"github.com/Fishwaldo/mouthpiece/pkg/ent"
	_ "github.com/Fishwaldo/mouthpiece/pkg/ent/runtime"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"

	"github.com/go-logr/logr"
	"github.com/go-logr/logr/testr"
	_ "github.com/mattn/go-sqlite3"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var GClient *ent.Client
var Ctx context.Context
var Tst *testing.T
var Mp *mouthpiece.MouthPiece
var App interfaces.AppI
var GlobalLogger logr.Logger

var DBBeforeSuite = BeforeSuite(func() {
	interfaces.Config.ExpireFilters = 2 * time.Second

	if len(os.Getenv("MP_DEBUG_LOG")) > 0 {
		GlobalLogger = testr.NewWithOptions(Tst, testr.Options{Verbosity: 10})
	} else {
		GlobalLogger = logr.Discard()
	}
	db, err := sql.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	Expect(err).To(BeNil())
	Expect(db).ToNot(BeNil())
	Mp = mouthpiece.NewMouthPiece(context.Background(), "sqlite3", db, GlobalLogger)
	Expect(Mp).ToNot(BeNil())
	
	Ctx = Mp.SetAdminTenant(context.Background())
	Expect(Ctx).ToNot(BeNil())

	Expect(Mp.Start(Ctx)).To(BeNil())


	App, err = Mp.GetAppService().Get(Ctx, "MouthPiece")
	Expect(err).To(BeNil())
	Expect(App).ToNot(BeNil())
})

var DBAfterSuite = AfterSuite(func() {
	Mp.Close()
})

func TransportSend(tpi interfaces.TransportInstanceImpl) {
	ctrl := gomock.NewController(GinkgoT())
	mockMessage := NewMockMessageI(ctrl)

	mockMessage.EXPECT().String().Return("Test Message").AnyTimes()
	mockMessage.EXPECT().GetFields(gomock.Any()).Return(map[string]string{"test": "value"}, nil).AnyTimes()
	mockMessage.EXPECT().GetMetadataFields(gomock.Any()).Return(map[string]interface{}{"test": "value"}, nil).AnyTimes()
	mockMessage.EXPECT().GetID().Return(uuid.New()).AnyTimes()

	user := NewMockUserI(ctrl)
	user.EXPECT().GetEmail().Return("test@test.com").AnyTimes()

	tpr := NewMockTransportRecipient(ctrl)
	tpr.EXPECT().GetRecipientType(gomock.Any()).Return(interfaces.TransportRecipientTypeUser).AnyTimes()
	tpr.EXPECT().GetUser(gomock.Any()).Return(user, nil).AnyTimes()
	tpr.EXPECT().GetName().Return("test").AnyTimes()

	err := tpi.Send(Ctx, tpr, mockMessage)
	Expect(err).To(BeNil())

	grp := NewMockGroupI(ctrl)
	grp.EXPECT().GetName().Return("test").AnyTimes()

	tpr2 := NewMockTransportRecipient(ctrl)
	tpr2.EXPECT().GetGroup(gomock.Any()).Return(grp, nil).AnyTimes()
	tpr2.EXPECT().GetName().Return("test").AnyTimes()
	tpr2.EXPECT().GetRecipientType(gomock.Any()).Return(interfaces.TransportRecipientTypeGroup).AnyTimes()

	err = tpi.Send(Ctx, tpr2, mockMessage)
	Expect(err).To(BeNil())
}
