package mouthpiece

import (
	"context"

	"github.com/Fishwaldo/mouthpiece/internal/app"
	"github.com/Fishwaldo/mouthpiece/internal/errors"
	. "github.com/Fishwaldo/mouthpiece/internal/log"
	"github.com/Fishwaldo/mouthpiece/internal/message"
)

func RouteMessage(ctx context.Context, msg *msg.Message) {
	if app, err := app.FindApp(ctx, msg.AppName); err == nil {
		app.ProcessMessage(ctx, msg)
	} else {
		Log.Error(mperror.ErrAppNotFound, "App Not Found", "App", msg.AppName)
	}
}
