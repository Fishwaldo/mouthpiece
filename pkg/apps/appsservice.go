package apps

import (
	"context"
	"errors"

	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/errors"
	"github.com/Fishwaldo/mouthpiece/pkg/filter"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/validate"

	"github.com/go-playground/validator/v10"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)


type AppService struct {
}

func NewAppsService() *AppService {
	db.Db.AutoMigrate(&App{})
	db.Db.AutoMigrate(&ApplicationFilters{})
	return &AppService{}
}

func (a AppService) GetApps(ctx context.Context) map[uint]interfaces.AppI {
	var apps []App
	
	appMap := make(map[uint]interfaces.AppI)
	if tx := db.Db.WithContext(ctx).Find(&apps); tx.Error != nil {
		log.Log.Error(tx.Error, "Error Finding Apps")
	}
	for _, app := range apps {
		appMap[app.ID] = app
	}
	return appMap
}

func (a AppService) GetAppByName(ctx context.Context, app_name string) (app interfaces.AppI, err error) {
	var dbapp App
	tx := db.Db.WithContext(ctx).Preload("Filters").First(&dbapp, "app_name = ?", app_name)
	log.Log.V(1).Info("Finding App", "App", app_name, "Result", tx, "app", dbapp)
	return &dbapp, tx.Error
}

func (a AppService) GetApp(ctx context.Context, appid uint) (app interfaces.AppI, err error) {
	var dbApp App
	tx := db.Db.WithContext(ctx).Preload("Filters").First(&dbApp, "id = ?", appid)
	log.Log.V(1).Info("Finding App", "App", appid, "Result", tx, "app", dbApp)
	return &dbApp, tx.Error
}

func (a AppService) CreateApp(ctx context.Context, app interfaces.AppDetails) (newapp interfaces.AppI, err error) {
	newapp, err = a.GetAppByName(ctx, app.AppName)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		log.Log.Info("Creating New App", "App", app)
		var dbApp App
		copier.Copy(&dbApp, &app)
		if filter.FindFilter("CopyShortMessage") != nil {
			dbApp.Filters = append(dbApp.Filters, ApplicationFilters{Name: "CopyShortMessage"})
		}
		if filter.FindFilter("FindSeverity") != nil {
			dbApp.Filters = append(dbApp.Filters, ApplicationFilters{Name: "FindSeverity"})
		}
		if err := validate.Get().Struct(dbApp); err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				log.Log.Info("CreateApp: Validation Error", "Error", e)
			}
			return nil, err.(validator.ValidationErrors)
		}
		result := db.Db.WithContext(ctx).Create(&dbApp)
		if result.Error != nil {
			return newapp, result.Error
		}
		return a.GetApp(ctx, dbApp.ID)
	}
	log.Log.Error(err, "App Already Exists", "App", newapp)
	return nil, mperror.ErrAppExists
}
