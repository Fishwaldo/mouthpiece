package interfaces

import (
	"context"
	"time"
)

//CtxUserValue Context Key to get token.User value from Context
type CtxUserValue struct{}

//go:generate go run github.com/dmarkham/enumer -type=AppStatus -json -text -sql
type AppStatus int

const (
	Enabled AppStatus = iota
	Disabled
)

func (AppStatus) Values() []string {
	return AppStatusStrings()
}

type CacheAble struct {
	lastUsed time.Time
}

func (c *CacheAble) GetLastUsed() time.Time {
	return c.lastUsed
}
func (c *CacheAble) SetLastUsed() {
	c.lastUsed = time.Now()
}

type AppDetails struct {
	ID          uint   `doc:"App ID" gorm:"primary_key"`
	AppName     string `doc:"Application Name" pattern:"^[a-z0-9]+$" gorm:"unique;uniqueIndex" validate:"required,max=255,alphanum"`
	Status      string `doc:"Status of Application" enum:"Enabled,Disabled" default:"Enabled" validate:"required,oneof=Enabled Disabled"`
	Description string `doc:"Description of Application" validate:"required,max=255"`
	Icon        string `doc:"Icon of Application" validate:"url"`
	URL         string `doc:"URL of Application" validate:"url"`
}

type UserDetails struct {
	ID        uint   `doc:"User ID" gorm:"primary_key"`
	Email     string `doc:"Email" validate:"required,email"`
	FirstName string `doc:"First Name" validate:"required"`
	LastName  string `doc:"Last Name" validate:"required"`
	Password  string `doc:"Password" json:"-" writeOnly:"true" validate:"required"`
}

//go:generate go run github.com/dmarkham/enumer -type=FilterType -json -text -sql
type FilterType int

const (
	InvalidFilter FilterType = iota
	AppFilter
	UserFilter
	TransportFilter
)

func (FilterType) Values() []string {
	return FilterTypeStrings()
}

type ctxKey struct {
	key string
}

var (
	MpctxKey = ctxKey{key: "mp"}
)

type MpService interface {
	GetUserService() UserServiceI
	GetAppService() AppServiceI
	GetGroupService() GroupServiceI
	GetFilterService() FilterServiceI
	GetTransportService() TransportServiceI
}

func GetUserService(ctx context.Context) UserServiceI {
	return ctx.Value(MpctxKey).(MpService).GetUserService()
}

func GetAppService(ctx context.Context) AppServiceI {
	return ctx.Value(MpctxKey).(MpService).GetAppService()
}

func GetGroupService(ctx context.Context) GroupServiceI {
	return ctx.Value(MpctxKey).(MpService).GetGroupService()
}

func GetFilterService(ctx context.Context) FilterServiceI {
	return ctx.Value((MpctxKey)).(MpService).GetFilterService()
}

func GetTransportService(ctx context.Context) TransportServiceI {
	return ctx.Value(MpctxKey).(MpService).GetTransportService()
}

const (
	MD_AppName        = "X-AppName"
	MD_AppDescription = "X-AppDescription"
	MD_AppURL         = "X-AppURL"
	MD_AppIcon        = "X-AppIcon"
	MD_UserEmail      = "X-UserEmail"
	MD_UserName       = "X-UserName"
)


