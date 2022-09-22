package app

import (
	"context"
	"fmt"
	"sync"

	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"

	"github.com/Fishwaldo/mouthpiece/pkg/ent"
	"github.com/go-logr/logr"
)

type App struct {
	dbApp *ent.DbApp
	lock  sync.RWMutex
	log   logr.Logger
}

func newApp(ctx context.Context, log logr.Logger, appname string, desc string) (*App, error) {
	newlogger := log.WithName("App").WithValues("App", appname)
	dbApp, err := db.DbClient.DbApp.Create().
		SetName(appname).
		SetDescription(desc).
		SetStatus(interfaces.Enabled).
		Save(ctx)
	if err != nil {
		newlogger.Error(err, "Error creating app")
		return nil, mperror.FilterErrors(err)
	}
	app := &App{
		dbApp: dbApp,
		log:   newlogger,
	}
	if err := app.init(); err != nil {
		app.log.Error(err, "Error initializing app")
		return nil, mperror.FilterErrors(err)
	}
	app.log.Info("Created New App")
	return app, nil
}

func (a *App) init() error {
	return mperror.FilterErrors(nil)
}

func (app *App) Load(ctx context.Context, logger logr.Logger, a any) error {
	app.lock.Lock()
	defer app.lock.Unlock()

	var ok bool
	app.dbApp, ok = a.(*ent.DbApp)
	if !ok {
		logger.Error(nil, "Can't Load App From incompatible dbType", "type", fmt.Sprintf("%T", a))
		return mperror.ErrInvalidType
	}
	app.log = logger.WithName("App").WithValues("App", app.dbApp.Name)
	app.log.V(1).Info("Loaded App")
	return app.init()
}

func (app *App) Save(ctx context.Context) (err error) {
	app.lock.Lock()
	defer app.lock.Unlock()
	mydb, err := app.dbApp.Update().Save(ctx)
	if err == nil {
		app.dbApp = mydb
	}
	return mperror.FilterErrors(err)
}

func (app *App) GetID() int {
	app.lock.RLock()
	defer app.lock.RUnlock()

	return app.dbApp.ID
}

func (app *App) GetName() string {
	app.lock.RLock()
	defer app.lock.RUnlock()

	return app.dbApp.Name
}

func (app *App) SetName(ctx context.Context, name string) (err error) {
	app.lock.Lock()
	defer app.lock.Unlock()

	dbtmp, err := app.dbApp.Update().SetName(name).Save(ctx)
	if err == nil {
		app.dbApp = dbtmp
	}
	return mperror.FilterErrors(err)
}

func (app *App) GetDescription() string {
	app.lock.RLock()
	defer app.lock.RUnlock()

	return app.dbApp.Description
}

func (app *App) SetDescription(ctx context.Context, description string) (err error) {
	app.lock.Lock()
	defer app.lock.Unlock()

	dbtmp, err := app.dbApp.Update().SetDescription(description).Save(ctx)
	if err == nil {
		app.dbApp = dbtmp
	}
	return mperror.FilterErrors(err)
}

func (app *App) GetIcon() string {
	app.lock.RLock()
	defer app.lock.RUnlock()

	return app.dbApp.Icon
}

func (app *App) SetIcon(ctx context.Context, icon string) (err error) {
	app.lock.Lock()
	defer app.lock.Unlock()

	dbtmp, err := app.dbApp.Update().SetIcon(icon).Save(ctx)
	if err == nil {
		app.dbApp = dbtmp
	}
	return mperror.FilterErrors(err)
}

func (app *App) GetURL() string {
	app.lock.RLock()
	defer app.lock.RUnlock()

	return app.dbApp.URL
}

func (app *App) SetURL(ctx context.Context, url string) (err error) {
	app.lock.Lock()
	defer app.lock.Unlock()

	dbtmp, err := app.dbApp.Update().SetURL(url).Save(ctx)
	if err == nil {
		app.dbApp = dbtmp
	}
	return mperror.FilterErrors(err)
}

func (app *App) GetStatus() interfaces.AppStatus {
	app.lock.RLock()
	defer app.lock.RUnlock()

	return app.dbApp.Status
}

func (app *App) SetStatus(ctx context.Context, status interfaces.AppStatus) (err error) {
	app.lock.Lock()
	defer app.lock.Unlock()

	dbtmp, err := app.dbApp.Update().SetStatus(status).Save(ctx)
	if err == nil {
		app.dbApp = dbtmp
	}
	return mperror.FilterErrors(err)
}

func (app *App) ProcessMessage(ctx context.Context, msg interfaces.MessageI) (err error) {
	app.lock.Lock()
	defer app.lock.Unlock()

	app.log.V(1).Info("App Processing Message", "MessageID", msg.GetID())

	/* populate Message Fields with App Data */
	if len(app.dbApp.Description) > 0 {
		msg.SetMetadata(ctx, interfaces.MD_AppDescription, app.dbApp.Description)
	}
	if len(app.dbApp.Icon) > 0 {
		msg.SetMetadata(ctx, interfaces.MD_AppIcon, app.dbApp.Icon)
	}
	if len(app.dbApp.URL) > 0 {
		msg.SetMetadata(ctx, interfaces.MD_AppURL, app.dbApp.URL)
	}
	msg.SetMetadata(ctx, interfaces.MD_AppName, app.dbApp.Name)

	// make sure our Filters are loaded
	var flts []*ent.DbFilter
	if flts, err = app.dbApp.Edges.FiltersOrErr(); err != nil {
		if flts, err = app.dbApp.QueryFilters().All(ctx); err != nil {
			app.log.Error(err, "Error loading Filters for App")
			return mperror.FilterErrors(err)
		}
	}

	app.log.V(1).Info("Loaded App Filters", "Filters", flts)
	for _, appfilter := range flts {
		app.log.V(1).Info("App Processing Message with Filter", "Filter", appfilter.Name)
		flt, err := interfaces.GetFilterService(ctx).Load(ctx, appfilter)
		if err != nil {
			app.log.Error(err, "Error loading Filter", "Filter", appfilter.Name)
			continue
		}
		if ok, err := flt.ProcessMessage(ctx, msg); err != nil {
			app.log.Error(err, "Error processing message with filter", "Filter", flt.GetName())
			continue
		} else if ok == interfaces.FilterMatch {
			app.log.Info("App Filter Matched Message", "Filter", flt.GetName(), "Message", msg.GetID())
			break
		} else if ok == interfaces.FilterNoMatch {
			app.log.Info("App Filter Did Not Match Message", "Filter", flt.GetName(), "Message", msg.GetID())
			return nil
		} else if ok == interfaces.FilterPass {
			app.log.Info("App Filter Passed Message", "Filter", flt.GetName(), "Message", msg.GetID())
			continue
		}

	}
	var grps []*ent.DbGroup
	if grps, err = app.dbApp.Edges.GroupsOrErr(); err != nil {
		if grps, err = app.dbApp.QueryGroups().All(ctx); err != nil {
			app.log.Error(err, "Error loading Groups for App")
			return mperror.FilterErrors(err)
		}
	}
	for _, grp := range grps {
		app.log.V(1).Info("App Processing Message with Group", "Group", grp.Name)
		grpobj, err := interfaces.GetGroupService(ctx).Load(ctx, grp)
		if err != nil {
			app.log.Error(err, "Error loading Group", "Group", grp.Name)
			continue
		}
		if err := grpobj.ProcessMessage(ctx, msg); err != nil {
			app.log.Error(err, "Error processing message with group", "Group", grpobj.GetName())
			continue
		}
	}
	return nil
}

func (app *App) AddFilter(ctx context.Context, filter interfaces.FilterI) error {
	app.lock.Lock()
	defer app.lock.Unlock()

	var err error
	dbtmp, err := app.dbApp.Update().AddFilterIDs(filter.GetID()).Save(ctx)
	if err != nil {
		app.log.Error(err, "Error adding filter to app", "Filter", filter.GetName())
		return mperror.FilterErrors(err)
	} 
	app.dbApp = dbtmp
	return nil
}

func (app *App) DelFilter(ctx context.Context, filter interfaces.FilterI) error {
	app.lock.Lock()
	defer app.lock.Unlock()

	var err error
	dbtmp, err := app.dbApp.Update().RemoveFilterIDs(filter.GetID()).Save(ctx)
	if err != nil {
		app.log.Error(err, "Error removing filter from app", "Filter", filter.GetName())
		return mperror.FilterErrors(err)
	}
	app.dbApp = dbtmp
	return nil
}

func (app *App) GetFilters(ctx context.Context) (flts []interfaces.FilterI, err error) {
	app.lock.Lock()
	defer app.lock.Unlock()

	// make sure our Filters are loaded
	var dbflts []*ent.DbFilter
	if dbflts, err = app.dbApp.Edges.FiltersOrErr(); err != nil {
		if dbflts, err = app.dbApp.QueryFilters().All(ctx); err != nil {
			app.log.Error(err, "Error loading Filters for App")
			return nil, mperror.FilterErrors(err)
		}
	}

	for _, f := range dbflts {
		flt, err := interfaces.GetFilterService(ctx).Load(ctx, f)
		if err != nil {
			app.log.Error(err, "Error loading Filter", "Filter", f.Name)
			continue
		}
		flts = append(flts, flt)
	}
	return flts, nil
}

var _ interfaces.AppI = (*App)(nil)
