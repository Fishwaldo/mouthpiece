package app

import (
	"context"

	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/ent"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbapp"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"

	"github.com/go-logr/logr"
)

type AppService struct {
	log logr.Logger
}

func NewAppService(ctx context.Context, log logr.Logger) *AppService {
	newlog := log.WithName("AppService")
	newlog.V(1).Info("New App Service")
	return &AppService{
		log: newlog,
	}
}

func (a *AppService) Start(ctx context.Context) error {
	return nil
}

func (a *AppService) Create(ctx context.Context, name string, description string) (newapp interfaces.AppI, err error) {
	if ok, err := a.Exists(ctx, name); err != nil {
		return nil, mperror.ErrInternalError
	} else if ok {
		return nil, mperror.ErrAppExists
	}
	app, err := newApp(ctx, a.log, name, description)
	if err != nil {
		a.log.Error(err, "Error creating app", "name", name)
		return nil, mperror.ErrInternalError
	} else {
		return app, nil
	}
}

func (a *AppService) Delete(ctx context.Context, app interfaces.AppI) error {
	if ok, err := a.Exists(ctx, app.GetName()); err != nil {
		return mperror.ErrInternalError
	} else if !ok {
		return mperror.ErrAppNotFound
	}
	if err := db.DbClient.DbApp.DeleteOneID(app.GetID()).Exec(ctx); err != nil {
		a.log.Error(err, "Error deleting app", "name", app.GetName())
		return mperror.ErrInternalError
	}
	return nil
}

func (a *AppService) GetByID(ctx context.Context, id int) (app interfaces.AppI, err error) {
	db_app, err := db.DbClient.DbApp.Query().Where(dbapp.ID(id)).Only(ctx)
	if err != nil {
		a.log.V(1).Error(err, "Error getting app", "id", id)
		return nil, mperror.FilterErrors(err)
	}
	return a.Load(ctx, db_app)
}

func (a *AppService) Get(ctx context.Context, name string) (app interfaces.AppI, err error) {
	db_app, err := db.DbClient.DbApp.Query().Where(dbapp.Name(name)).Only(ctx)
	if err != nil {
		a.log.V(1).Error(err, "Error getting app", "name", name)
		return nil, mperror.FilterErrors(err)
	}
	return a.Load(ctx, db_app)
}

func (a *AppService) GetAll(ctx context.Context) ([]interfaces.AppI, error) {
	apps, err := db.DbClient.DbApp.Query().All(ctx)
	if err != nil {
		return nil, mperror.ErrInternalError
	}
	var appDetails []interfaces.AppI
	for _, app := range apps {
		myapp, err := a.Load(ctx, app)
		if err != nil {
			a.log.Error(err, "Error loading app", "name", app.Name)
			continue
		}
		appDetails = append(appDetails, myapp)
	}
	return appDetails, nil
}

func (a *AppService) Exists(ctx context.Context, name string) (bool, error) {
	if ok, err := db.DbClient.DbApp.Query().Where(dbapp.Name(name)).Exist(ctx); err != nil {
		a.log.Error(err, "Error checking if app exists", "name", name)
		return false, mperror.ErrInternalError
	} else {
		return ok, nil
	}
}

func (a *AppService) ExistsByID(ctx context.Context, id int) (bool, error) {
	if ok, err :=  db.DbClient.DbApp.Query().Where(dbapp.ID(id)).Exist(ctx); err != nil {
		a.log.Error(err, "Error checking if app exists", "id", id)
		return false, mperror.ErrInternalError
	} else {
		return ok, nil
	}
}

func (a *AppService) Load(ctx context.Context, db_app any) (interfaces.AppI, error) {
	entApp, ok := db_app.(*ent.DbApp)
	if !ok {
		a.log.Error(nil, "Error loading app")
		return nil, mperror.ErrInvalidType
	}
	app := &App{}
	if err := app.Load(ctx, a.log, entApp); err != nil {
		a.log.Error(err, "Error loading app", "name", entApp.Name)
		return nil, mperror.ErrInvalidType
	}
	return app, nil
}

var _ interfaces.AppServiceI = (*AppService)(nil)
