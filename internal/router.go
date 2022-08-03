package mouthpiece

import (
	. "github.com/Fishwaldo/mouthpiece/internal/log"
	"github.com/Fishwaldo/mouthpiece/internal/errors"
	"github.com/Fishwaldo/mouthpiece/internal/app"
	"github.com/Fishwaldo/mouthpiece/internal/message"
)

func RouteMessage(msg *msg.Message) {
	if app, err := app.FindApp(msg.AppName); err == nil {
		app.ProcessMessage(msg)
	} else {
		Log.Error(mperror.ErrAppNotFound, "App Not Found", "App", msg.AppName)
	}
}