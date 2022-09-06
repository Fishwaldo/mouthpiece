package interfaces

import (
	"context"

	"github.com/Fishwaldo/mouthpiece/pkg/log"
)

//CtxUserValue Context Key to get token.User value from Context
type CtxUserValue struct{}

type AppDetails struct {
	ID          uint   `doc:"App ID" gorm:"primary_key"`
	AppName     string `doc:"Application Name" pattern:"^[a-z0-9]+$" gorm:"unique;uniqueIndex" validate:"required,max=255,alphanum"`
	Status      string `doc:"Status of Application" enum:"Enabled,Disabled" default:"Enabled" validate:"required,oneof=Enabled Disabled"`
	Description string `doc:"Description of Application" validate:"required,max=255"`
	Icon        string `doc:"Icon of Application" validate:"url"`
	URL         string `doc:"URL of Application" validate:"url"`
}

// TableName overrides the table name used by UserDetails to `users`
func (AppDetails) TableName() string {
	return "apps"
}

type UserDetails struct {
	ID        uint   `doc:"User ID" gorm:"primary_key"`
	Email     string `doc:"Email" validate:"required,email"`
	FirstName string `doc:"First Name" validate:"required"`
	LastName  string `doc:"Last Name" validate:"required"`
	Password  string `doc:"Password" json:"-" writeOnly:"true" validate:"required"`
}

// TableName overrides the table name used by UserDetails to `users`
func (UserDetails) TableName() string {
	return "users"
}

const (
	InvalidFilter = iota
	AppFilter
	UserFilter
	TransportFilter
)

type FilterType int

func (ft FilterType) String() string {
	return [...]string{"InvalidFilter", "AppFilter", "UserFilter", "TransportFilter"}[ft]
}

type MPContext struct {
	context.Context
}

type ctxKey struct {
	key string
}

var (
	userctxKey      = ctxKey{key: "user"}
	appctxKey       = ctxKey{key: "app"}
	groupctxKey     = ctxKey{key: "group"}
	filterctxKey    = ctxKey{key: "filter"}
	transportctxKey = ctxKey{key: "transport"}
)

func NewContext(ctx context.Context) *MPContext {
	return &MPContext{ctx}
}

func (c *MPContext) SetUserService(usersvc UserServicierI) {
	c.Context = context.WithValue(c.Context, userctxKey, usersvc)
}

func (c *MPContext) GetUserService() UserServicierI {
	return c.Context.Value(userctxKey).(UserServicierI)
}

func (c *MPContext) SetAppService(appsvc AppServiceI) {
	c.Context = context.WithValue(c.Context, appctxKey, appsvc)
}

func (c *MPContext) GetAppService() AppServiceI {
	return c.Context.Value(appctxKey).(AppServiceI)
}

func (c *MPContext) SetGroupService(groupsvc GroupServiceI) {
	c.Context = context.WithValue(c.Context, groupctxKey, groupsvc)
}

func (c *MPContext) GetGroupService() GroupServiceI {
	return c.Context.Value(groupctxKey).(GroupServiceI)
}

func (c *MPContext) SetFilterService(filtersvc FilterServiceI) {
	c.Context = context.WithValue(c.Context, filterctxKey, filtersvc)
}

func (c *MPContext) GetFilterService() FilterServiceI {
	if c.Context.Value((filterctxKey)).(FilterServiceI) != nil {
		return c.Context.Value((filterctxKey)).(FilterServiceI)
	} else {
		log.Log.Error(nil, "Filter service is nil")
		return nil
	}
}

func (c *MPContext) SetTransportService(tps TransportInstance) {
	c.Context = context.WithValue(c.Context, transportctxKey, tps)
}

func (c *MPContext) GetTransportService() TransportServiceI {
	return c.Context.Value(transportctxKey).(TransportServiceI)
}
