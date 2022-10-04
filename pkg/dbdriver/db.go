package dbdriver

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Fishwaldo/mouthpiece/pkg/ent"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/migrate"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/rules"
	_ "github.com/Fishwaldo/mouthpiece/pkg/ent/runtime"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/tenant"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
)

var DbClient *ent.Client
var SqlDrv *sql.DB

func Initialize(drivername string, driver *sql.DB) (*ent.Client, error) {

	if !(strings.EqualFold(drivername, dialect.MySQL) || strings.EqualFold(drivername, dialect.Postgres)  || strings.EqualFold(drivername, dialect.SQLite)) {
			return nil, mperror.ErrUnsupportedDBType
	}

	client := entsql.OpenDB(drivername, driver)

	DbClient = ent.NewClient(ent.Driver(client), ent.Log(dbLog))

	if interfaces.Config.DebugSQL {
		DbClient = DbClient.Debug()
	}

	if err := DbClient.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
		log.Log.WithName("DB").Error(err, "Unable to create schema")
		return nil, err
	}

	/* create our default Tenant if it doesn't exist */

	tntrole := &rules.UserViewer{
		Role: rules.GlobalAdmin,
	}
	ctx := rules.NewContext(context.Background(), tntrole)
	if DbClient.Tenant.Query().CountX(ctx) == 0 {
		DbClient.Tenant.Create().SetName("default").Save(ctx)
	}
	SqlDrv = driver
	return DbClient, nil
}

func dbLog(args ...interface{}) {
	//fmt.Print(args...)
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