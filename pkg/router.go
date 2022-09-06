package mouthpiece

import (
	"context"


	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"
	"github.com/Fishwaldo/mouthpiece/pkg/msg"
)

func (mp MouthPiece) RouteMessage(ctx context.Context, msg *msg.Message) {
	if app, err := mp.appService.GetAppByName(ctx, msg.AppName); err == nil {
		if err = app.ProcessMessage(ctx, mp.userService, msg); err != nil {
			log.Log.Info("App Processing Message Failed", "App", app.GetName(), "Message", msg, "Error", err)
			return
		} else {
			log.Log.Info("App Processing Message Success", "App", app.GetName(), "Message", msg)
		}
	} else {
		log.Log.Error(mperror.ErrAppNotFound, "App Not Found", "App", msg.AppName)
	}
}

