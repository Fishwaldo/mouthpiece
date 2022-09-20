package db

import (
	"context"
	"fmt"

	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/ent"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/tenant"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/migrate"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/rules"
	_ "github.com/Fishwaldo/mouthpiece/pkg/ent/runtime"

	"entgo.io/ent/dialect"
)

var DbClient *ent.Client

func Initialize(driver, dbconnectstring string) (*ent.Client, error) {
	var err error
	switch driver {
	case
		dialect.MySQL,
		dialect.Postgres,
		dialect.SQLite:
		if DbClient, err = ent.Open(driver, dbconnectstring, ent.Log(dbLog)); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported driver %q", driver)
	}

	DbClient.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true))

	/* create our default Tenant if it doesn't exist */

	tntrole := &rules.UserViewer{
		Role: rules.GlobalAdmin,
	}
	ctx := rules.NewContext(context.Background(), tntrole)
	if DbClient.Tenant.Query().CountX(ctx) == 0 {
		DbClient.Tenant.Create().SetName("default").Save(ctx)
	}

	return DbClient.Debug(), nil
}

func dbLog(args ...interface{}) {
	logline := fmt.Sprint(args...)
	log.Log.WithName("DB").Info(logline)
}

func GetDefaultTenant() *ent.Tenant {
	tntrole := &rules.UserViewer{
		Role: rules.GlobalAdmin,
	}
	ctx := rules.NewContext(context.Background(), tntrole)
	tnt, err := DbClient.Tenant.Query().Where(tenant.Name("default")).Only(ctx)
	if err != nil {
		log.Log.WithName("DB").Error(err, "Unable to find default tenant")
		return nil
	}
	return tnt
}