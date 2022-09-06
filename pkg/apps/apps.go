package apps

import (
	"context"

	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"
	"github.com/Fishwaldo/mouthpiece/pkg/msg"
	"github.com/Fishwaldo/mouthpiece/pkg/validate"
)

type ApplicationFilters struct {
	ID    uint `doc:"Filter ID" gorm:"primary_key"`
	AppID uint `json:"-"`
	Name  string
}

type App struct {
	mpctx *interfaces.MPContext
	interfaces.AppDetails
	Filters []AppFilter `json:"-" gorm:"many2many:app2flt;"`
}

type AppFilter struct {
	ID       uint `doc:"Filter ID" gorm:"primary_key"`
	AppID    uint `json:"-"`
	FilterID uint `json:"-"`
}

// TableName Override the table name for the App table
// to ensure its consistant with the table name defined in the Interfaces Package
func (app *App) TableName() string {
	return "apps"
}

func (app *App) validateSaveField(val interfaces.AppDetails, fields ...string) error {
	if len(fields) > 0 {
		if err := validate.Get().StructPartial(val, fields...); err != nil {
			log.Log.Info("Validation Error", "Error", err)
			return err
		}
	} else {
		if err := validate.Get().Struct(val); err != nil {
			log.Log.Info("Validation Error", "Error", err)
			return err
		}
	}
	app.AppDetails = val
	if tx := db.Db.Model(&app).Updates(app); tx.Error != nil {
		log.Log.Error(tx.Error, "Error updating application details", "AppName", app.AppName)
		return tx.Error
	}
	return nil
}

func (app *App) SetSvcCtx(ctx *interfaces.MPContext) {
	app.mpctx = ctx
}

func (app *App) GetName() string {
	return app.AppName
}

func (app *App) SetName(name string) error {
	var details interfaces.AppDetails = app.AppDetails
	details.AppName = name
	return app.validateSaveField(details, "AppName")
}

func (app *App) GetID() uint {
	return app.ID
}

func (app *App) GetDetails() interfaces.AppDetails {
	return app.AppDetails
}
func (app *App) SetDetails(details interfaces.AppDetails) error {
	return app.validateSaveField(details)
}

func (app *App) GetDescription() string {
	return app.Description
}

func (app *App) SetDescription(description string) error {
	var details interfaces.AppDetails = app.GetDetails()
	details.Description = description
	return app.validateSaveField(details, "Description")
}

func (app *App) GetIcon() string {
	return app.Icon
}

func (app *App) SetIcon(icon string) error {
	var details interfaces.AppDetails = app.AppDetails
	details.Icon = icon
	return app.validateSaveField(details, "Icon")
}

func (app *App) GetURL() string {
	return app.URL
}

func (app *App) SetURL(icon string) error {
	var details interfaces.AppDetails = app.AppDetails
	details.Icon = icon
	return app.validateSaveField(details, "URL")
}

func (app *App) ProcessMessage(ctx context.Context, us interfaces.UserServicierI, msg *msg.Message) error {
	log.Log.V(1).Info("App Processing Message", "App", app.AppName, "MessageID", msg.ID)
	/* populate Message Fields with App Data */
	msg.Body.Fields["app_description"] = app.Description
	msg.Body.Fields["app_icon"] = app.Icon
	msg.Body.Fields["app_url"] = app.URL
	// make sure our Filters are loaded
	if len(app.Filters) == 0 {
		log.Log.Info("Loading Filters?", "App", app.AppName)
		if tx := db.Db.Debug().Model(&app).Association("Filters").Find(&app.Filters); tx != nil {
			log.Log.Error(tx, "Error loading application filters", "AppName", app.AppName)
			return tx
		}
	}
	log.Log.V(1).Info("App Filters", "App", app.AppName, "Filters", app.Filters)
	for _, appfilter := range app.Filters {
		flt := app.mpctx.GetFilterService().GetByID(ctx, appfilter.FilterID, interfaces.AppFilter)
		log.Log.V(1).Info("App Processing Message with Filter", "Filter", flt.GetName())
		if ok, err := flt.ProcessMessage(ctx, msg); err != nil {
			log.Log.Error(err, "Error processing message with filter", "Filter", flt.GetName())
			continue
		} else if !ok {
			log.Log.Info("App Filter Blocked Message", "App", app.AppName, "Filter", flt.GetName(), "Message", msg)
			return nil
		}
	}
	if err := app.mpctx.GetGroupService().SendMessageToUsers(ctx, msg, app); err != nil {
		log.Log.Error(err, "Error sending message to users", "App", app.AppName)
	}
	if err := app.mpctx.GetGroupService().SendMessageToTransports(ctx, msg, app); err != nil {
		log.Log.Error(err, "Error sending message to transports", "App", app.AppName)
	}
	return nil
}

func (app *App) AddFilter(ctx context.Context, filter interfaces.FilterI) error {
	if app.ID == 0 {
		log.Log.Error(nil, "Error adding filter to application, application not saved", "AppName", app.AppName)
		return mperror.ErrAppNotFound
	}
	if filter.GetID() == 0 {
		log.Log.Error(nil, "Error adding filter to application, filter not saved", "AppName", app.AppName, "Filter", filter.GetName())
		return mperror.ErrFilterNotFound
	}
	appflt := AppFilter{AppID: app.ID, FilterID: filter.GetID()}

	if tx := db.Db.Model(&app).Association("Filters").Append(&appflt); tx != nil {
		log.Log.Error(tx, "Error adding filter to application", "AppName", app.AppName, "Filter", filter.GetName())
		return tx
	}
	app.Filters = append(app.Filters, appflt)
	return nil
}

func removeCopy(slice []AppFilter, i int) []AppFilter {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func (app *App) DelFilter(ctx context.Context, filter interfaces.FilterI) error {
	if app.ID == 0 {
		log.Log.Error(nil, "Error adding filter to application, application not saved", "AppName", app.AppName)
		return mperror.ErrAppNotFound
	}
	if filter.GetID() == 0 {
		log.Log.Error(nil, "Error adding filter to application, filter not saved", "AppName", app.AppName, "Filter", filter.GetName())
		return mperror.ErrFilterNotFound
	}
	appflt := AppFilter{AppID: app.ID, FilterID: filter.GetID()}
	if tx := db.Db.Model(&app).Association("Filters").Delete(&appflt); tx != nil {
		log.Log.Error(tx, "Error deleting filter from application", "AppName", app.AppName, "Filter", filter.GetName())
		return tx
	}
	for i, appfilter := range app.Filters {
		if appfilter.ID == appflt.ID {
			app.Filters = removeCopy(app.Filters, i)
			break
		}
	}
	return nil
}
