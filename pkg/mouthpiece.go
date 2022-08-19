package mouthpiece

import (
	"github.com/Fishwaldo/mouthpiece/pkg/apps"
	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/groups"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/users"

	"github.com/go-logr/logr"
	"gorm.io/gorm"
)

type MouthPiece struct {
	appService   interfaces.AppServiceI
	userService  interfaces.UserServicierI
	groupService interfaces.GroupServiceI
	db           *gorm.DB
	log          logr.Logger
}

func NewMouthPiece(dbconn gorm.Dialector, logger logr.Logger) *MouthPiece {

	mp := MouthPiece{
		log:          log.InitLogger(logger),
		db:           db.Initialize(dbconn),
		userService:  users.NewUsersService(),
		appService:   apps.NewAppsService(),
		groupService: groups.NewGroupsService(),
	}
	return &mp
}

func (mp MouthPiece) GetAppService() interfaces.AppServiceI {
	return mp.appService
}
func (mp MouthPiece) GetUserService() interfaces.UserServicierI {
	return mp.userService
}
