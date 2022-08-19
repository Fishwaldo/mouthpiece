package interfaces

import (
	"context"

	msg "github.com/Fishwaldo/mouthpiece/pkg/message"
)

type AppI interface {
	ProcessMessage(context.Context, UserServicierI, *msg.Message) error
	GetName() string
	SetName(string) error
	GetID() uint
	GetDetails() AppDetails
	SetDetails(AppDetails) error
	GetDescription() string
	SetDescription(string) error
	GetIcon() string
	SetIcon(string) error
	GetURL() string
	SetURL(string) error
}

type AppServiceI interface {
	GetApps(context.Context) map[uint]AppI
	GetApp(context.Context, uint) (AppI, error)
	GetAppByName(context.Context, string) (AppI, error)
	CreateApp(context.Context, AppDetails) (AppI, error)
}

type UserI interface {
	GetID() uint
	ProcessMessage(context.Context, *msg.Message) error
	CheckPassword(context.Context, string) bool
	SetPassword(context.Context, string) error
	AddRoleToUser(context.Context, string) bool
	GetDetails() UserDetails
	SetDetails(UserDetails) error
	GetEmail() string
	SetEmail(string) error
	GetFirstName() string
	SetFirstName(string) error
	GetLastName() string
	SetLastName(string) error
}

type UserServicierI interface {
	GetUser(context.Context, uint) (UserI, error)
	GetUserByEmail(context.Context, string) (UserI, error)
	CreateUser(context.Context, UserDetails) (UserI, error)
}

type UserSender func(context.Context, msg.Message, uint) error
type TransportSender func(context.Context, msg.Message, uint) error

type GroupI interface {
	GetID() uint
	GetName() string
	SetName(string) error
	AddUserToGroup(ctx context.Context, user uint) bool
	DelUserFromGroup(ctx context.Context, user uint) bool
	GetUsers() []uint
	AddAppToGroup(context.Context, uint) bool
	DelAppFromGroup(context.Context, uint) bool
	//GetApps() []
	AddTransportToGroup(ctx context.Context, tid uint) bool
	DelTransportFromGroup(ctx context.Context, tid uint) bool
	GetTransports() []uint
}

type GroupServiceI interface {
	GetGroup(context.Context, uint) (GroupI, error)
	SendMessageToUsers(ctx context.Context, msg msg.Message, appid uint, sender UserSender) error
	SendMessageToTransports(ctx context.Context, msg msg.Message, appid uint, sender UserSender) error
}
