package mock_interfaces

import (
	"context"
	"testing"
	"os"
	"time"

	"github.com/Fishwaldo/mouthpiece/pkg"
	"github.com/Fishwaldo/mouthpiece/pkg/ent"
	_ "github.com/Fishwaldo/mouthpiece/pkg/ent/runtime"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"

	"github.com/go-logr/logr"
	"github.com/go-logr/logr/testr"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var GClient *ent.Client
var Ctx context.Context
var Tst *testing.T
var Mp *mouthpiece.MouthPiece
var App interfaces.AppI

var DBBeforeSuite = BeforeSuite(func() {
	interfaces.Config.ExpireFilters = 2 * time.Second

	var logger logr.Logger
	if len(os.Getenv("MP_DEBUG_LOG")) > 0 {
		logger = testr.NewWithOptions(Tst, testr.Options{Verbosity: 10})
	} else {
		logger = logr.Discard()
	}
	Mp = mouthpiece.NewMouthPiece(context.Background(), "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", logger)
	Expect(Mp).ToNot(BeNil())
	
	Ctx = Mp.SetAdminTenant(context.Background())
	Expect(Ctx).ToNot(BeNil())

	Expect(Mp.Start(Ctx)).To(BeNil())

	var err error

	App, err = Mp.GetAppService().Get(Ctx, "MouthPiece")
	Expect(err).To(BeNil())
	Expect(App).ToNot(BeNil())
})

var DBAfterSuite = AfterSuite(func() {
	Mp.Close()
})
