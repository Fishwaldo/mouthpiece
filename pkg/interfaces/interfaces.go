package interfaces

import (
	"context"
	"github.com/Fishwaldo/mouthpiece/pkg/msg"
	"github.com/go-logr/logr"
)

type AppI interface {
	SetSvcCtx(*MPContext)
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
	AddFilter(ctx context.Context, filter FilterI) error
	DelFilter(ctx context.Context, filter FilterI) error
}

type AppServiceI interface {
	Start(*MPContext) error
	GetApps(context.Context) ([]AppDetails, error)
	GetApp(context.Context, uint) (AppI, error)
	GetAppByName(context.Context, string) (AppI, error)
	GetAppObj(context.Context, AppDetails) (AppI, error)
	CreateApp(context.Context, AppDetails) (AppI, error)
}

type UserI interface {
	GetID() uint
	ProcessMessage(context.Context, msg.Message) error
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
	Start(*MPContext) error
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
	Start(*MPContext) error
	CreateGroup(ctx context.Context, name string) (GroupI, error)
	DeleteGroup(ctx context.Context, group GroupI) error
	GetGroupByID(context.Context, uint) (GroupI, error)
	GetGroup(context.Context, string) (GroupI, error)
	SendMessageToUsers(ctx context.Context, msg *msg.Message, app AppI) error
	SendMessageToTransports(ctx context.Context, msg *msg.Message, app AppI) error
}

type FilterI interface {
	GetID() uint
	GetName() string
	ProcessMessage(context.Context, *msg.Message) (bool, error)
	GetConfig(item string) interface{}
}

type FilterServiceI interface {
	Start(*MPContext) error
	Get(context.Context, string, FilterType) FilterI
	GetByID(context.Context, uint, FilterType) FilterI
}

type TransportProvider interface {
	Start(*MPContext, logr.Logger) error
	GetName() string
	CreateInstance(context.Context, any) (TransportInstance, error)
	DeleteInstance(context.Context, TransportInstance) error
	GetInstance(context.Context, uint) (TransportInstance, error)
	GetInstanceByName(context.Context, string) (TransportInstance, error)
	GetInstances(context.Context) ([]uint, error)
}

type TransportInstance interface {
	Start(*MPContext) error
	Stop (context.Context) error
	GetName() string
	GetID() uint
	CreateGroupTransportRecipient(context.Context, uint, GroupI, any) (TransportRecipient, error)
	CreateUserTransportRecipient(context.Context, uint, UserI, any) (TransportRecipient, error)
	DeleteTransportRecipient(context.Context, TransportRecipient) error
	GetTransportRecipientByTransportID(context.Context, uint) (TransportRecipient, error)
	GetTransportReciepients(context.Context) ([]uint, error)
}

type TransportRecipient interface {
	ProcessGroupMessage(context.Context, msg.Message) error
	ProcessMessage(context.Context, msg.Message) error
	GetID() uint
	GetGroupID() uint
	GetUserID() uint
}

type TransportServiceI interface {
	Start(*MPContext) error
	CreateTransportInstance(context.Context, string, any) (TransportInstance, error)
	GetTransportInstances(context.Context) ([]uint, error)
	GetTransportInstance(context.Context, uint) (TransportInstance, error)
	GetTransportInstanceByName(context.Context, string) (TransportInstance, error)
	DeleteTransportInstance(context.Context, TransportInstance) error
	CreateGroupTransportRecipient(context.Context, TransportInstance, GroupI, any) (TransportRecipient, error)
	CreateUserTransportRecipient(context.Context, TransportInstance, UserI, any) (TransportRecipient, error)
	DeleteTransportRecipient(context.Context, TransportRecipient) error
	GetTransportReciepientsForGroup(context.Context, GroupI) ([]uint, error)
	GetTransportReciepientsForUser(context.Context, UserI) ([]uint, error)
	GetTransportReciepients(context.Context) ([]uint, error)
	GetTransportReciepient(context.Context, uint) (TransportRecipient, error)
}
