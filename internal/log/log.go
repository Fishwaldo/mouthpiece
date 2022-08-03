package log

import (
	"fmt"
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"github.com/danielgtaylor/huma/middleware"
)

var Log logr.Logger
var zapLog *zap.Logger

func InitLogger() {
	zapLog, err := middleware.NewDefaultLogger()
    if err != nil {
        panic(fmt.Sprintf("Initilize Logging Failed (%v)?", err))
    }
    Log = zapr.NewLogger(zapLog)

    Log.Info("Logging Started")
}