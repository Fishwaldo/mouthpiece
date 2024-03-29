package app

import (
	"errors"
	"context"

	"github.com/Fishwaldo/mouthpiece/internal/db"
	"github.com/Fishwaldo/mouthpiece/internal/errors"
	"github.com/Fishwaldo/mouthpiece/internal/filter"
	. "github.com/Fishwaldo/mouthpiece/internal/log"
	"github.com/Fishwaldo/mouthpiece/internal/message"
	"github.com/Fishwaldo/mouthpiece/internal/user"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type AppDetails struct {
	AppName     string `doc:"Application Name" pattern:"^[a-z0-9]+$" gorm:"unique;uniqueIndex"`
	Status      string `doc:"Status of Application" enum:"Enabled,Disabled" default:"Enabled"`
	Description string `doc:"Description of Application"`
	Icon        string `doc:"Icon of Application"`
	URL         string `doc:"URL of Application"`
}

type ApplicationFilters struct {
	gorm.Model `json:"-"`
	AppID      uint `json:"-"`
	Name       string
}

type App struct {
	gorm.Model `json:"-"`
	AppDetails
	AssociatedUsers []*user.User `gorm:"many2many:app_user;"`
	Filters         []ApplicationFilters
}

func InitializeApps() {
	db.Db.AutoMigrate(&App{})
	db.Db.AutoMigrate(&ApplicationFilters{})
}

func GetApps(ctx context.Context) []App {
	var apps []App
	db.Db.WithContext(ctx).Preload("AssociatedUsers").Preload("AssociatedUsers.TransportConfigs").Preload("Filters").Find(&apps)
	return apps
}
func FindApp(ctx context.Context, app_name string) (app *App, err error) {
	tx := db.Db.WithContext(ctx).Preload("AssociatedUsers").Preload("AssociatedUsers.TransportConfigs").Preload("Filters").First(&app, "app_name = ?", app_name)
	Log.V(1).Info("Finding App", "App", app_name, "Result", tx, "app", app)
	return app, tx.Error
}

func AppExists(ctx context.Context, app_name string) bool {
	var app App
	tx := db.Db.WithContext(ctx).First(&app, "app_name = ?", app_name)
	return tx.Error == nil
}

func CreateApp(ctx context.Context, app AppDetails) (newapp *App, err error) {
	newapp, err = FindApp(ctx, app.AppName)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		Log.Info("Creating New App", "App", app)
		var dbApp App
		copier.Copy(&dbApp, &app)
		adminuser, _ := user.GetUser(ctx, "admin@example.com")
		dbApp.AssociatedUsers = append(dbApp.AssociatedUsers, adminuser)
		normaluser, _ := user.GetUser(ctx, "user@example.com")
		dbApp.AssociatedUsers = append(dbApp.AssociatedUsers, normaluser)
		if filter.FindFilter("CopyShortMessage") != nil {
			dbApp.Filters = append(dbApp.Filters, ApplicationFilters{Name: "CopyShortMessage"})
		}
		if filter.FindFilter("FindSeverity") != nil {
			dbApp.Filters = append(dbApp.Filters, ApplicationFilters{Name: "FindSeverity"})
		}
		result := db.Db.WithContext(ctx).Create(&dbApp)
		if result.Error != nil {
			return newapp, result.Error
		}
		return FindApp(ctx, app.AppName)
	}
	Log.Error(err, "App Already Exists", "App", newapp)
	return newapp, mperror.ErrAppExists
}

func (app App) ProcessMessage(ctx context.Context, msg *msg.Message) error {
	Log.V(1).Info("App Processing Message", "App", app.AppName, "MessageID", msg.ID)
	/* populate Message Fields with App Data */
	msg.Body.Fields["app_description"] = app.Description
	msg.Body.Fields["app_icon"] = app.Icon
	msg.Body.Fields["app_url"] = app.URL
	for _, appfilter := range app.Filters {
		flt := filter.FindFilter(appfilter.Name)
		if flt != nil {
			Log.V(1).Info("App Processing Message with Filter", "Filter", appfilter)
			ok, _ := flt.ProcessMessage(ctx, msg)
			if !ok {
				Log.Info("App Filter Blocked Message", "App", app.AppName, "Filter", appfilter, "Message", msg)
				return nil
			}
		} else {
			Log.Info("App Filter Not Found", "App", app.AppName, "Filter", appfilter)
		}
	}
	for _, user := range app.AssociatedUsers {
		user.ProcessMessage(ctx, *msg)
	}
	return nil
}
