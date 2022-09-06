package apps

import (
	"context"
	"errors"

	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"
	"github.com/Fishwaldo/mouthpiece/pkg/validate"

	"github.com/go-playground/validator/v10"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type AppService struct {
	ctx *interfaces.MPContext
}

func NewAppsService() *AppService {
	return &AppService{
	}
}

func (a *AppService) Start(ctx *interfaces.MPContext) error {
	db.Db.AutoMigrate(&App{})
	a.ctx = ctx
	return nil
}

func (a *AppService) GetApps(ctx context.Context) ([]interfaces.AppDetails, error) {
	var apps []App

	var appDetails []interfaces.AppDetails
	if tx := db.Db.WithContext(ctx).Find(&apps); tx.Error != nil {
		log.Log.Error(tx.Error, "Error Finding Apps")
		return appDetails, tx.Error
	}
	for _, app := range apps {
		appDetails = append(appDetails, app.AppDetails)
	}
	return appDetails, nil
}

func (a *AppService) GetAppByName(ctx context.Context, app_name string) (app interfaces.AppI, err error) {
	var dbapp App
	tx := db.Db.WithContext(ctx).Preload("Filters").First(&dbapp, "app_name = ?", app_name)
	if tx.Error != nil {
		log.Log.V(1).Error(tx.Error, "GetAppByName", "App", app_name)
	}
	dbapp.SetSvcCtx(a.ctx)
	return &dbapp, tx.Error
}

func (a *AppService) GetApp(ctx context.Context, appid uint) (app interfaces.AppI, err error) {
	var dbApp App
	tx := db.Db.WithContext(ctx).Preload("Filters").First(&dbApp, "id = ?", appid)
	if tx.Error != nil {
		log.Log.V(1).Error(tx.Error, "GetApp", "App", appid)
	}
	dbApp.SetSvcCtx(a.ctx)
	return &dbApp, tx.Error
}

func (a *AppService) GetAppObj(ctx context.Context, appDet interfaces.AppDetails) (app interfaces.AppI, err error) {
	return a.GetApp(ctx, appDet.ID)
}

func (a *AppService) CreateApp(ctx context.Context, app interfaces.AppDetails) (newapp interfaces.AppI, err error) {
	newapp, err = a.GetAppByName(ctx, app.AppName)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		log.Log.Info("Creating New App", "App", app)
		var dbApp App
		copier.Copy(&dbApp, &app)
		if err := validate.Get().Struct(dbApp); err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				log.Log.Info("CreateApp: Validation Error", "Error", e)
			}
			return nil, err.(validator.ValidationErrors)
		}
		result := db.Db.WithContext(ctx).Create(&dbApp)
		if result.Error != nil {
			return nil, result.Error
		}
		return a.GetApp(ctx, dbApp.ID)
	}
	log.Log.Error(err, "App Already Exists", "App", newapp)
	return nil, mperror.ErrAppExists
}
