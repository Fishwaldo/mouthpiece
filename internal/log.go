package mpserver

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var zapLog *zap.Logger

func GetHumaLogger() (*zap.Logger, error) {
	return zapLog.Named("huma"), nil
}
func GetZapLogger() (*zap.Logger, error) {
	return zapLog, nil
}

func InitLogger() logr.Logger {
	var cfg zap.Config
	var lvl zapcore.Level
	var err error
	if lvl, err = zapcore.ParseLevel(viper.GetString("log.level")); err != nil {
		panic(err)
	}
	if viper.GetBool("debug") {
		fmt.Printf("Debug Enabled at %s level\n", viper.GetString("log.level"))
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cfg.OutputPaths = []string{"stdout"}
		cfg.EncoderConfig.EncodeTime = iso8601UTCTimeEncoder
		cfg.Level = zap.NewAtomicLevelAt(lvl)
		cfg.DisableStacktrace = false
		var err error
		if zapLog, err = cfg.Build(); err != nil {
			panic(err)
		}
	} else {
		fmt.Printf("Debug Disabled. Logging to file %s at %s level\n", filepath.Join(viper.GetString("log.dir"), "mouthpiece.log"), viper.GetString("log.level"))
		lumberJackLogger := &lumberjack.Logger{
			Filename:   filepath.Join(viper.GetString("log.dir"), "mouthpiece.log"),
			MaxSize:    viper.GetInt("log.maxsize"), // megabytes
			MaxBackups: viper.GetInt("log.maxbackups"),
			MaxAge:     viper.GetInt("log.maxage"), //days
			Compress:   viper.GetBool("log.compress"),
		}
		ws := zapcore.AddSync(lumberJackLogger)
		enccfg := zap.NewProductionEncoderConfig()
		enccfg.EncodeTime = iso8601UTCTimeEncoder
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(enccfg),
			ws,
			zap.NewAtomicLevelAt(lvl),
		)

		zapLog = zap.New(core)
		zap.ReplaceGlobals(zapLog)
	}
	zap.RedirectStdLog(zapLog)
	return zapr.NewLogger(zapLog)

}

// A UTC variation of ZapCore.ISO8601TimeEncoder with millisecond precision
func iso8601UTCTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000Z"))
}
