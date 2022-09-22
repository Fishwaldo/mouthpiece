package mouthpiece

import (
	"context"

	"github.com/Fishwaldo/mouthpiece/pkg/app"
	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/ent"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/rules"
	"github.com/Fishwaldo/mouthpiece/pkg/filter"
	"github.com/Fishwaldo/mouthpiece/pkg/group"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/transport"
	"github.com/Fishwaldo/mouthpiece/pkg/user"

	//	_ "github.com/Fishwaldo/mouthpiece/pkg/filter/evalfilter"
	"github.com/Fishwaldo/mouthpiece/pkg/filter/severity"
	"github.com/Fishwaldo/mouthpiece/pkg/transport/console"

	"github.com/go-logr/logr"
)

type MouthPiece struct {
	appService       interfaces.AppServiceI
	userService      interfaces.UserServiceI
	groupService     interfaces.GroupServiceI
	filterService    interfaces.FilterServiceI
	transportService interfaces.TransportServiceI
	log              logr.Logger
	globalAdminRole  *rules.UserViewer
	tenantAdminRole  *rules.UserViewer
	tenantUserRole   *rules.UserViewer
}

func NewMouthPiece(ctx context.Context, dbtype string, dbconnstring string, logger logr.Logger) *MouthPiece {
	mplog := log.InitLogger(logger)
	if _, err := db.Initialize(dbtype, dbconnstring); err != nil {
		mplog.Error(err, "Error Initializing DB")
		return nil
	} else {
		mp := &MouthPiece{
			log:              mplog,
			userService:      user.NewUsersService(ctx, mplog),
			appService:       app.NewAppService(ctx, mplog),
			groupService:     group.NewGroupService(ctx, mplog),
			filterService:    filter.NewFilterService(ctx, mplog),
			transportService: transport.NewTransportService(ctx, mplog),
			globalAdminRole: &rules.UserViewer{
				Role: rules.GlobalAdmin,
				T:    db.GetDefaultTenant(),
			},
			tenantAdminRole: &rules.UserViewer{
				Role: rules.Admin,
				T:    db.GetDefaultTenant(),
			},
			tenantUserRole: &rules.UserViewer{
				Role: rules.View,
				T:    db.GetDefaultTenant(),
			},
		}
		return mp
	}
}

func (mp MouthPiece) Start(ctx context.Context) error {
	mp.log.Info("Starting MouthPeice Services")
	mp.userService.Start(ctx)
	mp.appService.Start(ctx)
	mp.groupService.Start(ctx)
	mp.filterService.Start(ctx)
	mp.transportService.Start(ctx)
	mp.SeedMouthPieceApp(ctx)
	mp.log.Info("Moutpiece Services Started")
	return nil
}

func (mp MouthPiece) withContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, interfaces.MpctxKey, mp)
}

func (mp MouthPiece) GetAppService() interfaces.AppServiceI {
	return mp.appService
}

func (mp MouthPiece) GetUserService() interfaces.UserServiceI {
	return mp.userService
}

func (mp MouthPiece) GetGroupService() interfaces.GroupServiceI {
	return mp.groupService
}

func (mp MouthPiece) GetFilterService() interfaces.FilterServiceI {
	return mp.filterService
}

func (mp MouthPiece) GetTransportService() interfaces.TransportServiceI {
	return mp.transportService
}

func (mp MouthPiece) SeedMouthPieceApp(ctx context.Context) error {
	if !interfaces.Config.SeedDB {
		mp.log.V(1).Info("Skipping DB Seed")
		return nil
	}
	if _, err := mp.appService.Get(ctx, "MouthPiece"); err != nil && ent.IsNotFound(err) {
		log.Log.Info("Seeding Database")
		var admin interfaces.UserI
		var user interfaces.UserI
		if admin, err = mp.userService.Create(ctx, "admin@example.com", "Admin User"); err == nil {
			log.Log.Info("Created Default Admin admin@example.com")
		} else {
			log.Log.Error(err, "Error Creating Default Admin")
			panic(err)
		}
		if user, err = mp.userService.Create(ctx, "user@example.com", "Normal User"); err == nil {
			log.Log.Info("Created Default User user@example.com")
		} else {
			log.Log.Error(err, "Error Creating Default User")
			panic(err)
		}

		/* create Console Provider */
		consoleprovider, err := mp.transportService.GetTransportProvider(ctx, "console")
		if err != nil {
			log.Log.Error(err, "Error getting Console Transport Provider")
		}

		/* create console output */
		consoleconfig := &console.ConsoleConfig{}
		consoleinstance, err := mp.transportService.CreateTransportInstance(ctx, consoleprovider, "consoleinstance", consoleconfig)
		if err != nil {
			log.Log.Error(err, "Error Creating Console Instance")
		}

		adminrecptcfg := &console.ConsoleRecipientConfig{}

		/* create Console Recipient for Admin/User */
		adminconsolerecipient, err := mp.transportService.Create(ctx, consoleinstance, "adminreciptient", adminrecptcfg)
		if err != nil {
			log.Log.Error(err, "Error Creating Admin Console Recipient")
		}

		err = admin.AddTransportRecipient(ctx, adminconsolerecipient)
		if err != nil {
			log.Log.Error(err, "Error Adding Admin Console Recipient")
		}

		userrecptcfg := &console.ConsoleRecipientConfig{}

		userconsolerecipient, err := mp.transportService.Create(ctx, consoleinstance, "userrecipient", userrecptcfg)
		if err != nil {
			log.Log.Error(err, "Error Creating User Console Recipient")
		}

		err = user.AddTransportRecipient(ctx, userconsolerecipient)
		if err != nil {
			log.Log.Error(err, "Error Adding User Console Recipient")
		}

		app, err := mp.appService.Create(ctx, "MouthPiece", "MouthPiece App")
		if err != nil {
			log.Log.Error(err, "Error Creating MouthPiece App")
		}

		flt, err := mp.filterService.Create(ctx, "SeverityFilter", "SevFilter", interfaces.UserFilter)
		if err != nil {
			log.Log.Error(err, "Error Creating SevFilter")
		}
		svfltconfig := &severity.SeverityFilterConfig {
			Op: severity.SeverityFilterOpGT,
			Severity: 3,
		}
		if err := flt.SetConfig(ctx, svfltconfig); err != nil {
			log.Log.Error(err, "Error Setting SevFilter Config")
		}
		if err := admin.AddFilter(ctx, flt); err != nil {
			log.Log.Error(err, "Error Adding SevFilter to Admin")
		}

		if grp, err := mp.groupService.Create(ctx, "MouthPiece", "Default MouthPiece App Group"); err != nil {
			log.Log.Error(err, "Error creating MouthPiece group")
		} else {
			grp.AddApp(ctx, app)
			grp.AddUser(ctx, admin)
			grp.AddUser(ctx, user)
		}
	} else {
		mp.log.Error(err, "Error Seeding MouthPiece App")
		panic(err)
	}
	mp.log.Info("Seeding Database Complete")
	return nil
}

func (mp MouthPiece) Close() {

}

func (mp MouthPiece) SetAdminTenant(ctx context.Context) context.Context {
	return mp.withContext(rules.NewContext(ctx, mp.tenantAdminRole))
}

func (mp MouthPiece) SetUserTenant(ctx context.Context) context.Context {
	return mp.withContext(rules.NewContext(ctx, mp.tenantUserRole))
}

func (mp MouthPiece) RouteMessage(ctx context.Context, msg interfaces.MessageI) error {
	ctx = mp.SetAdminTenant(ctx)
	log.Log.Info("Routing Message", "msg", msg)
	if app, err := msg.GetApp(ctx); err != nil {
		log.Log.Error(err, "Error getting app from message")
		return err
	} else {
		if err := app.ProcessMessage(ctx, msg); err != nil {
			log.Log.Error(err, "Error routing message")
			return err
		}
	}
	return nil
}

var _ interfaces.MpService = &MouthPiece{}
