package log

import (
	"github.com/go-logr/logr"

	"github.com/spf13/viper"
)

var Log logr.Logger

//var zapLog *zap.Logger

func init() {
	viper.SetDefault("debug", false)
	viper.SetDefault("log.dir", "logs")
	viper.SetDefault("log.maxsize", 1)
	viper.SetDefault("log.maxbackups", 3)
	viper.SetDefault("log.maxage", 7)
	viper.SetDefault("log.compress", true)
	viper.SetDefault("log.level", "info")

}

func InitLogger(passedlog logr.Logger) logr.Logger {
	Log = passedlog.WithName("MP")
	Log.Info("Logging Started", "level", viper.GetString("log.level"))
	return Log
}
