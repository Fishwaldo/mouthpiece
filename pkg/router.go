package mouthpiece

import (
	"context"

	"github.com/Fishwaldo/mouthpiece/pkg/errors"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/message"
)

func (mp MouthPiece) RouteMessage(ctx context.Context, msg *msg.Message) {
	if app, err := mp.appService.GetAppByName(ctx, msg.AppName); err == nil {
		if err = app.ProcessMessage(ctx, mp.userService, msg); err != nil {
			log.Log.Info("App Processing Message Failed", "App", app.GetName(), "Message", msg, "Error", err)
			return
		} else {
			if err := mp.groupService.SendMessageToUsers(ctx, *msg, app.GetID(), mp.userSender); err != nil {
				log.Log.Error(err, "Failed to send message to users", "App", app.GetName(), "Message", msg)
			}
			if err := mp.groupService.SendMessageToTransports(ctx, *msg, app.GetID(), mp.transportSender); err != nil {
				log.Log.Error(err, "Failed to send message to Transports", "App", app.GetName(), "Message", msg)
			}
		}
	} else {
		log.Log.Error(mperror.ErrAppNotFound, "App Not Found", "App", msg.AppName)
	}
}

func (mp MouthPiece) transportSender(ctx context.Context, msg msg.Message, tid uint) error {
	log.Log.Info("Sending Message to Transport", "Message", msg, "Transport", tid)
	return nil
}

func (mp MouthPiece) userSender(ctx context.Context, msg msg.Message, uid uint) error {
	if user, err := mp.userService.GetUser(ctx, uid); err != nil {
		log.Log.Error(err, "Failed to get userID for userSender", "User", uid)
		return err
	} else {
		if err := user.ProcessMessage(ctx, &msg); err != nil {
			log.Log.Error(err, "Failed to process message for userSender", "User", user.GetID(), "Message", msg)
		} else {
			log.Log.Info("Sending Message to User", "Message", msg, "User", user.GetID())
		}
	}
	return nil
}
