package apps

import (
	//	"errors"
	"context"
//	"fmt"

	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/filter"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/message"
	"github.com/Fishwaldo/mouthpiece/pkg/validate"

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

func (app App) validateSaveField(val interfaces.AppDetails, fields...string) error {
	if len(fields) > 0 { 
		if err := validate.Get().StructPartial(val, fields...); err != nil {
			log.Log.Info("Validation Error", "Error", err)
			return err;
		}
	} else {
		if err := validate.Get().Struct(val); err != nil {
			log.Log.Info("Validation Error", "Error", err)
			return err;
		}
	}
	app.AppDetails = val
	if tx := db.Db.Model(&app).Updates(app); tx.Error != nil {
		log.Log.Error(tx.Error, "Error updating application details", "AppName", app.AppName)
		return tx.Error
	}
	return nil
}

func (app App) GetName() string {
	return app.AppName
}

func (app App) SetName(name string) error {
	var details interfaces.AppDetails = app.AppDetails
	details.AppName = name
	return app.validateSaveField(details, "AppName")
}

func (app App) GetID() uint {
	return app.ID
}

func (app App) GetDetails() interfaces.AppDetails {
	return app.AppDetails
}
func (app App) SetDetails(details interfaces.AppDetails) error {
	return app.validateSaveField(details)
}

func (app App) GetDescription() string {
	return app.Description
}

func (app App) SetDescription(description string) error {
	var details interfaces.AppDetails = app.GetDetails()
	details.Description = description
	return app.validateSaveField(details, "Description")	
}

func (app App) GetIcon() string {
	return app.Icon
}

func (app App) SetIcon(icon string) error {
	var details interfaces.AppDetails = app.AppDetails
	details.Icon = icon
	return app.validateSaveField(details, "Icon")
}

func (app App) GetURL() string {
	return app.URL
}

func (app App) SetURL(icon string) error {
	var details interfaces.AppDetails = app.AppDetails
	details.Icon = icon
	return app.validateSaveField(details, "URL")
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
