package interfaces

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	"github.com/google/uuid"
)

type CacheAbleI interface {
	GetLastUsed() time.Time
	SetLastUsed()
}

type MessageI interface {
	Load(context.Context, any) error
	Save(context.Context, AppI) error
	GetID() uuid.UUID
	GetMessage() string
	SetMessage(ctx context.Context, newmessage string) (err error)
	GetShortMsg() *string
	SetShortMsg(ctx context.Context, shortmsg string) (err error)
	GetTopic() *string
	SetTopic(ctx context.Context, topic string) (err error)
	GetSeverity() int
	SetSeverity(ctx context.Context, sev int) (err error)
	GetTimestamp() time.Time
	SetTimestamp(ctx context.Context, ts time.Time) (err error)
	SetFields(ctx context.Context, fields map[string]string) error
	GetFields(ctx context.Context) (flds map[string]string, err error)
	GetField(ctx context.Context, key string) (value string, err error)
	SetField(ctx context.Context, key string, value string) (err error)
	SetMetadata(ctx context.Context, key string, value any) (err error)
	GetMetadata(ctx context.Context, key string) (value any, err error)
	GetMetadataFields(ctx context.Context) (flds map[string]any, err error)
	GetApp(ctx context.Context) (AppI, error)
	ProcessMessage(ctx context.Context) (err error)
	String() string
	Clone() MessageI
}


type MessageServiceI interface {
	// Start starts the MessageService
	Start(context.Context) error
	Get(ctx context.Context, id uuid.UUID) (MessageI, error)
	GetMessages(ctx context.Context, option ...func(query any) any) ([]MessageI, error)
	GetMessageCount(ctx context.Context) (int, error)
	Load(ctx context.Context, db_msg any) (msg MessageI, err error)
}

// AppI represents an application
type AppI interface {
	// Load loads the App from a DB record
	Load(context.Context, logr.Logger, any) error
	// Save saves the App to the DB
	Save(context.Context) error
	// GetID returns the ID of the App
	GetID() int
	// GetName returns the name of the App
	GetName() string
	// SetName sets the name of the App
	SetName(context.Context, string) error
	// GetDescription returns the description of the App
	GetDescription() string
	// SetDescription sets the description of the App
	SetDescription(context.Context, string) error
	// GetIcon returns the icon of the App
	GetIcon() string
	// SetIcon sets the icon of the App
	SetIcon(context.Context, string) error
	// GetURL returns the URL of the App
	GetURL() string
	// SetURL sets the URL of the App
	SetURL(context.Context, string) error
	// GetStatus sets the Status of the App
	GetStatus() AppStatus
	// SetStatus sets the Status of the App
	SetStatus(context.Context, AppStatus) error
	// ProcessMessage processes a message
	ProcessMessage(context.Context, MessageI) error
	// AddFilter adds a filter to the App
	AddFilter(context.Context, FilterI) error
	// DelFilter removes a filter from the App
	DelFilter(context.Context, FilterI) error
	// GetFilters returns the filters of the App
	GetFilters(context.Context) ([]FilterI, error)
	GetGroups(ctx context.Context) ([]GroupI, error)
	GetAppData(ctx context.Context, name string, data any) (err error)
	SetAppData(ctx context.Context, name string, data any) (err error)
}

// AppServiceI represents the AppService
type AppServiceI interface {
	// Start starts the AppService
	Start(context.Context) error
	// Create creates a new App
	Create(context.Context, string, string) (AppI, error)
	// Delete deletes an App
	Delete(context.Context, AppI) error
	// GetByID returns an App by ID
	GetByID(context.Context, int) (AppI, error)
	// Get returns an App by name
	Get(context.Context, string) (AppI, error)
	// GetAll returns all Apps
	GetAll(context.Context) ([]AppI, error)
	// Load loads an App from a DB record
	Load(context.Context, any) (AppI, error)
	// Exists checks if an App exists
	Exists(context.Context, string) (bool, error)
	// ExistsByID checks if an App exists by ID
	ExistsByID(context.Context, int) (bool, error)
}

type UserI interface {
	Load(context.Context, logr.Logger, any) error
	Save(context.Context) error
	GetID() int
	//	CheckPassword(context.Context, string) bool
	//	SetPassword(context.Context, string) error
	//	AddRoleToUser(context.Context, string) bool
	GetEmail() string
	SetEmail(context.Context, string) error
	GetName() string
	SetName(context.Context, string) error
	GetDescription() string
	SetDescription(context.Context, string) error
	// AddFilter adds a filter to the App
	AddFilter(context.Context, FilterI) error
	// DelFilter removes a filter from the App
	DelFilter(context.Context, FilterI) error
	// GetFilters returns the filters of the App
	GetFilters(context.Context) ([]FilterI, error)
	//GetTransports Get the Transports in the Group
	GetTransportRecipients(context.Context) []TransportRecipient
	//AddTransport Add a Transport to the Group
	AddTransportRecipient(context.Context, TransportRecipient) error
	//DelTransport Remove a Transport from the Group
	DelTransportRecipient(context.Context, TransportRecipient) error

	SetFields(ctx context.Context, fields map[string]string) (err error)
	GetFields(ctx context.Context) (flds map[string]string, err error)
	GetField(ctx context.Context, key string) (value string, err error)
	SetField(ctx context.Context, key string, value string) (err error)

	ProcessMessage(context.Context, MessageI) error
	GetAppData(ctx context.Context, name string, data any) (err error)
	SetAppData(ctx context.Context, name string, data any) (err error)
}

type UserServiceI interface {
	Start(context.Context) error
	Create(context.Context, string, string) (UserI, error)
	Delete(context.Context, UserI) error
	GetByID(context.Context, int) (UserI, error)
	Get(context.Context, string) (UserI, error)
	GetAll(context.Context) ([]UserI, error)
	Load(context.Context, any) (UserI, error)
	Exists(context.Context, string) (bool, error)
	ExistsByID(context.Context, int) (bool, error)
}

type UserSender func(context.Context, MessageI, uint) error
type TransportSender func(context.Context, MessageI, uint) error

// GroupI is the interface for a group.
// Groups Contain Both Users and TransportRecipients (for Channel/Broadcast Transports)
// Apps can be added to groups, and messages from Apps should be broadcast to all members of the Group.
type GroupI interface {
	//Load Load a group from a database Pointer
	Load(context.Context, logr.Logger, any) error
	//Save the Group to the database
	Save(context.Context) error
	//GetID Get the ID of the Group
	GetID() int
	//GetName Get the Name of the Group
	GetName() string
	//SetName Set the Name of the Group
	SetName(context.Context, string) error
	//GetDescription Get the Description of the Group
	GetDescription() string
	//SetDescription Set the Description of the Group
	SetDescription(context.Context, string) error

	//GetApps Get the Apps in the Group
	GetApps(context.Context) []AppI
	//AddApp Add a App to the Group
	AddApp(context.Context, AppI) error
	//DelApp Remove a App from the Group
	DelApp(context.Context, AppI) error

	//GetUsers Get the Users in the Group
	GetUsers(context.Context) []UserI
	//AddUser Add a User to the Group
	AddUser(ctx context.Context, user UserI) error
	//DelUser Remove a User from the Group
	DelUser(ctx context.Context, user UserI) error

	//GetTransports Get the Transports in the Group
	GetTransportRecipients(context.Context) []TransportRecipient
	//AddTransport Add a Transport to the Group
	AddTransportRecipient(context.Context, TransportRecipient) error
	//DelTransport Remove a Transport from the Group
	DelTransportRecipient(context.Context, TransportRecipient) error

	//ProcessMessage Process a Message for the Group
	ProcessMessage(context.Context, MessageI) error
	GetAppData(ctx context.Context, name string, data any) (err error)
	SetAppData(ctx context.Context, name string, data any) (err error)
}

//GroupServiceI is the Factory/Manager for all Groups
type GroupServiceI interface {
	Start(context.Context) error
	Create(context.Context, string, string) (GroupI, error)
	Delete(context.Context, GroupI) error
	GetByID(context.Context, int) (GroupI, error)
	Get(context.Context, string) (GroupI, error)
	GetAll(context.Context) ([]GroupI, error)
	Load(context.Context, any) (GroupI, error)
	Exists(context.Context, string) (bool, error)
	ExistsByID(context.Context, int) (bool, error)
}

type FilterAction int

const (
	FilterPass FilterAction = iota
	FilterMatch
	FilterNoMatch
)

type FilterI interface {
	CacheAbleI
	Load(context.Context, logr.Logger, any) error
	Save(context.Context) error
	GetID() int
	GetName() string
	SetName(context.Context, string) error
	GetDescription() string
	SetDescription(context.Context, string) error
	GetType() string
	GetFilterImplementation() FilterImplI
	ProcessMessage(context.Context, MessageI) (FilterAction, error)
	GetConfig(context.Context) (MarshableConfigI, error)
	SetConfig(context.Context, MarshableConfigI) error
	GetAppData(ctx context.Context, name string, data any) (err error)
	SetAppData(ctx context.Context, name string, data any) (err error)
}

type FilterServiceI interface {
	Load(context.Context, any) (FilterI, error)
	Get(context.Context, string, FilterType) (FilterI, error)
	GetByID(context.Context, int, FilterType) (FilterI, error)
	Create(context.Context, string, string, FilterType) (FilterI, error)
	Start(context.Context) error
}

type FilterImplI interface {
	FilterName() string
	Init(context.Context) error
	SetConfig(context.Context, MarshableConfigI) error
	GetConfig(context.Context) (MarshableConfigI, error)
	Process(context.Context, MessageI) (FilterAction, error)
}

type MarshableConfigI interface {
	AsJSON() (string, error)
	FromJSON(string) error
}

type TransportProvider interface {
	GetName() string
	CreateInstance(context.Context, logr.Logger, string, MarshableConfigI) (TransportInstanceImpl, error)
	LoadConfigFromJSON(context.Context, string) (MarshableConfigI, error)
}

type TransportInstance interface {
	Load(context.Context, logr.Logger, any) error
	Start(context.Context) error
	Stop(context.Context) error
	GetID() int
	GetName() string
	SetName(ctx context.Context, name string) error
	GetDescription() string
	SetDescription(ctx context.Context, description string) error
	SetConfig(context.Context, MarshableConfigI) error
	GetConfig(context.Context) (MarshableConfigI, error)
	Send(context.Context, TransportRecipient, MessageI) error
	ValidateTransportRecipientConfig(context.Context, MarshableConfigI) error
	LoadTransportReciepientConfig(context.Context, string) (MarshableConfigI, error)
	GetAppData(ctx context.Context, name string, data any) (err error)
	SetAppData(ctx context.Context, name string, data any) (err error)
}

type TransportInstanceImpl interface {
	Init(context.Context) error
	Start(context.Context) error
	Stop(context.Context) error
	Send(context.Context, TransportRecipient, MessageI) error
	GetConfig(context.Context) MarshableConfigI
	SetConfig(context.Context, MarshableConfigI) error
	ValidateTransportRecipientConfig(context.Context, MarshableConfigI) error
	LoadTransportReciepientConfig(context.Context, string) (MarshableConfigI, error)
}

type TransportRecipientType int

const (
	TransportRecipientTypeNotSet TransportRecipientType = iota
	TransportRecipientTypeUser
	TransportRecipientTypeGroup
)

type TransportRecipient interface {
	Load(context.Context, logr.Logger, any) error
	GetID() int
	GetName() string
	SetName(context.Context, string) error
	GetDescription() string
	SetDescription(context.Context, string) error
	GetConfig() (MarshableConfigI, error)
	SetConfig(context.Context, MarshableConfigI) error
	SetUser(context.Context, UserI) error
	GetUser(context.Context) (UserI, error)
	SetGroup(context.Context, GroupI) error
	GetGroup(context.Context) (GroupI, error)
	GetRecipientType(context.Context) TransportRecipientType
	ProcessMessage(context.Context, MessageI) error
	GetAppData(ctx context.Context, name string, data any) (err error)
	SetAppData(ctx context.Context, name string, data any) (err error)
}

type TransportServiceI interface {
	Start(context.Context) error
	Stop(context.Context) error
	Create(context.Context, TransportInstance, string, MarshableConfigI) (TransportRecipient, error)
	Delete(context.Context, TransportRecipient) error
	GetByID(context.Context, int) (TransportRecipient, error)
	Get(context.Context, string) (TransportRecipient, error)
	GetAll(context.Context) ([]TransportRecipient, error)
	Load(context.Context, any) (TransportRecipient, error)
	Exists(context.Context, string) bool
	ExistsByID(context.Context, int) bool

	CreateTransportInstance(context.Context, TransportProvider, string, MarshableConfigI) (TransportInstance, error)
	GetTransportInstance(context.Context, string) (TransportInstance, error)
	GetTransportInstanceByID(context.Context, int) (TransportInstance, error)
	GetAllTransportInstances(context.Context) ([]TransportInstance, error)
	DeleteTransportInstance(context.Context, TransportInstance) error
	ExistsTransportInstance(context.Context, string) bool
	ExistsByIDTransportInstance(context.Context, int) bool
	LoadTransportInstance(context.Context, any) (TransportInstance, error)

	GetTransportProvider(context.Context, string) (TransportProvider, error)
	GetAllTransportProviders(context.Context) ([]TransportProvider, error)
}
