package mouthpiece

import (
	"context"

	"github.com/Fishwaldo/mouthpiece/pkg/apps"
	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/filter"
	"github.com/Fishwaldo/mouthpiece/pkg/groups"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/msg"
	"github.com/Fishwaldo/mouthpiece/pkg/transport"
	"github.com/Fishwaldo/mouthpiece/pkg/users"

	_ "github.com/Fishwaldo/mouthpiece/pkg/filter/evalfilter"
	"github.com/Fishwaldo/mouthpiece/pkg/transport/console"

	"github.com/go-logr/logr"
	"gorm.io/gorm"
)

type MouthPiece struct {
	serviceContext   *interfaces.MPContext
	appService       interfaces.AppServiceI
	userService      interfaces.UserServicierI
	groupService     interfaces.GroupServiceI
	filterService    interfaces.FilterServiceI
	transportService interfaces.TransportServiceI
	db               *gorm.DB
	log              logr.Logger
}

func NewMouthPiece(svcctx *interfaces.MPContext, dbconn gorm.Dialector, logger logr.Logger) *MouthPiece {

	mp := MouthPiece{
		serviceContext:   svcctx,
		log:              log.InitLogger(logger),
		db:               db.Initialize(dbconn),
		userService:      users.NewUsersService(),
		appService:       apps.NewAppsService(),
		groupService:     groups.NewGroupsService(),
		filterService:    filter.NewFilterService(),
		transportService: transport.NewTransportService(),
	}
	mp.serviceContext.SetAppService(mp.appService)
	mp.serviceContext.SetUserService(mp.userService)
	mp.serviceContext.SetGroupService(mp.groupService)
	mp.serviceContext.SetFilterService(mp.filterService)

	msg.InitializeMessage()
	return &mp
}

func (mp MouthPiece) Start() error {
	mp.userService.Start(mp.serviceContext)
	mp.appService.Start(mp.serviceContext)
	mp.groupService.Start(mp.serviceContext)
	mp.filterService.Start(mp.serviceContext)
	mp.transportService.Start(mp.serviceContext)
	return nil
}

func (mp MouthPiece) GetAppService() interfaces.AppServiceI {
	return mp.appService
}

func (mp MouthPiece) GetUserService() interfaces.UserServicierI {
	return mp.userService
}

func (mp MouthPiece) SeedMouthPieceApp(ctx context.Context) error {
	if _, err := mp.appService.GetAppByName(ctx, "MouthPiece"); err != nil {
		log.Log.Info("Creating Default Users")
		admind := interfaces.UserDetails{FirstName: "Admin", LastName: "User", Email: "admin@example.com", Password: "password"}
		var admin interfaces.UserI
		var user interfaces.UserI
		if admin, err = mp.userService.CreateUser(context.Background(), admind); err == nil {
			admin.AddRoleToUser(context.Background(), "admin")
			log.Log.Info("Created Default Admin admin@example.com")
		}
		userd := interfaces.UserDetails{FirstName: "User", LastName: "User", Email: "user@example.com", Password: "password"}
		if user, err = mp.userService.CreateUser(context.Background(), userd); err == nil {
			log.Log.Info("Created Default User user@example.com")
		}

		/* create console output */
		consoleconfig := console.ConsoleInstanceConfig{}
		consoleconfig.Name = "StdOut"
		consoleconfig.Description = "Standard Output"
		consoleinstance, err := mp.transportService.CreateTransportInstance(context.Background(), "console", consoleconfig)
		if err != nil {
			log.Log.Error(err, "Error Creating StdOut Console Instance")
		}
		/* create console reciptient for admin */
		app, _ := mp.appService.CreateApp(mp.serviceContext, interfaces.AppDetails{
			AppName:     "MouthPiece",
			Status:      "Enabled",
			Description: "Internal MouthPiece App",
			Icon:        "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
			URL:         "https://github.com/Fishwaldo/mouthpiece",
		})
		flt := mp.filterService.Get(ctx, "CopyShortMessage", interfaces.AppFilter)
		app.AddFilter(mp.serviceContext, flt)
		if grp, err := mp.groupService.CreateGroup(mp.serviceContext, "MouthPiece"); err != nil {
			log.Log.Error(err, "Error creating MouthPiece group")
		} else {
			grp.AddAppToGroup(mp.serviceContext, app.GetID())
			grp.AddUserToGroup(mp.serviceContext, admin.GetID())
			grp.AddUserToGroup(mp.serviceContext, user.GetID())
			reciptientconfig := &console.ConsoleRecipientConfig{}
			reciptientconfig.UserID = admin.GetID()
			reciptientconfig.Name = "Group Stdout"
			reciptientconfig.Description = "Group Stdout Console Reciptient"
			atr, err := mp.transportService.CreateGroupTransportRecipient(context.Background(),consoleinstance, grp, reciptientconfig)		
			if err != nil {
				log.Log.Error(err, "Error Creating StdOut Console Reciptient")
			}
			grp.AddTransportToGroup(mp.serviceContext, atr.GetID())
		}

	}
	return nil
}
