package apps

import (
	//	"errors"
	"context"

	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/filter"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/message"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ApplicationFilters struct {
	gorm.Model `json:"-"`
	AppID      uint `json:"-"`
	Name       string
}

type App struct {
	interfaces.AppDetails
	Filters []ApplicationFilters
}

func (app App) GetName() string {
	return app.AppName
}

func (app App) SetName(name string) error {
	var details interfaces.AppDetails = app.AppDetails
	details.AppName = name
	return app.SetDetails(details)
}

func (app App) GetID() uint {
	return app.ID
}

func (app App) GetDetails() interfaces.AppDetails {
	return app.AppDetails
}
func (app *App) SetDetails(details interfaces.AppDetails) error {
	validate := validator.New()
	if err := validate.Struct(details); err != nil {
		log.Log.Info("SetDetails Validation Error", "Error", err)
		return err
	}
	if tx := db.Db.Model(&app).Updates(details); tx.Error != nil {
		log.Log.Error(tx.Error, "Error updating application details", "AppName", app.AppName)
		return tx.Error
	} else {
		app.AppDetails = details
	}
	return nil
}

func (app App) GetDescription() string {
	return app.Description
}

func (app App) SetDescription(description string) error {
	var details interfaces.AppDetails = app.AppDetails
	details.Description = description
	return app.SetDetails(details)
}

func (app App) GetIcon() string {
	return app.Icon
}

func (app App) SetIcon(icon string) error {
	var details interfaces.AppDetails = app.AppDetails
	details.Icon = icon
	return app.SetDetails(details)
}

func (app App) GetURL() string {
	return app.URL
}

func (app App) SetURL(icon string) error {
	var details interfaces.AppDetails = app.AppDetails
	details.Icon = icon
	return app.SetDetails(details)
}

func (app App) ProcessMessage(ctx context.Context, us interfaces.UserServicierI, msg *msg.Message) error {
	log.Log.V(1).Info("App Processing Message", "App", app.AppName, "MessageID", msg.ID)
	/* populate Message Fields with App Data */
	msg.Body.Fields["app_description"] = app.Description
	msg.Body.Fields["app_icon"] = app.Icon
	msg.Body.Fields["app_url"] = app.URL
	for _, appfilter := range app.Filters {
		flt := filter.FindFilter(appfilter.Name)
		if flt != nil {
			log.Log.V(1).Info("App Processing Message with Filter", "Filter", appfilter)
			ok, _ := flt.ProcessMessage(ctx, msg)
			if !ok {
				log.Log.Info("App Filter Blocked Message", "App", app.AppName, "Filter", appfilter, "Message", msg)
				return nil
			}
		} else {
			log.Log.Info("App Filter Not Found", "App", app.AppName, "Filter", appfilter)
		}
	}
	return nil
}
