package groups

import (
	"context"

	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
)

type Group struct {
	ID         uint                   `gorm:"primary_key"`
	Name       string                 `gorm:"not null"`
	Apps       []GroupAppMember       `gorm:"many2many:group_apps"`
	Users      []GroupUserMember      `gorm:"many2many:group_users"`
	Transports []GroupTransportMember `gorm:"many2many:group_transports"`
}

type GroupAppMember struct {
	ID      uint `gorm:"primary_key"`
	GroupID uint
	AppID   uint
}

type GroupUserMember struct {
	ID      uint `gorm:"primary_key"`
	GroupID uint
	UserID  uint
}

type GroupTransportMember struct {
	ID          uint `gorm:"primary_key"`
	GroupID     uint
	TransportID uint
}

func (g Group) GetID() uint {
	return g.ID
}
func (g Group) GetName() string {
	return g.Name
}
func (g Group) SetName(name string) error {
	return db.Db.Model(&g).Update("Name", name).Error
}
func (g Group) GetApps() []uint {
	apps := []uint{}
	for _, app := range g.Apps {
		apps = append(apps, app.AppID)
	}
	return apps
}
func (g Group) GetUsers() []uint {
	users := []uint{}
	for _, user := range g.Users {
		users = append(users, user.UserID)
	}
	return users
}
func (g Group) GetTransports() []uint {
	tps := []uint{}
	for _, tp := range g.Transports {
		tps = append(tps, tp.TransportID)
	}
	return tps
}
func (g Group) AddAppToGroup(ctx context.Context, app uint) bool {
	log.Log.Info("Adding App to Group", "app", app, "group", g.GetID())
	member := &GroupAppMember{
		GroupID: g.GetID(),
		AppID:   app,
	}
	tx := db.Db.Model(&g).Association("Apps").Append(member)
	if tx != nil {
		log.Log.Error(tx, "Error Adding App to Group", "app", app, "group", g.GetID())
		return false
	}
	g.Apps = append(g.Apps, *member)
	return true
}

func (g Group) DelAppFromGroup(ctx context.Context, app uint) bool {
	log.Log.Info("Deleting App from Group", "app", app, "group", g.GetID())
	member := &GroupAppMember{
		GroupID: g.GetID(),
		AppID:   app,
	}
	tx := db.Db.Model(&g).Association("Apps").Delete(member)
	if tx != nil {
		log.Log.Error(tx, "Error Deleting App from Group", "app", app, "group", g.GetID())
		return false
	}
	return true
}

func (g Group) AddUserToGroup(ctx context.Context, user uint) bool {
	log.Log.Info("Adding User to Group", "user", user, "group", g.GetID())
	member := &GroupUserMember{
		GroupID: g.GetID(),
		UserID:  user,
	}
	tx := db.Db.Model(&g).Association("Users").Append(member)
	if tx != nil {
		log.Log.Error(tx, "Error Adding User to Group", "user", user, "group", g.GetID())
		return false
	}
	g.Users = append(g.Users, *member)
	return true
}

func (g Group) DelUserFromGroup(ctx context.Context, user uint) bool {
	log.Log.Info("Deleting User from Group", "user", user, "group", g.GetID())
	member := &GroupUserMember{
		GroupID: g.GetID(),
		UserID:  user,
	}
	tx := db.Db.Model(&g).Association("Users").Delete(member)
	if tx != nil {
		log.Log.Error(tx, "Error Deleting User from Group", "user", user, "group", g.GetID())
		return false
	}
	return true
}

func (g Group) AddTransportToGroup(ctx context.Context, tid uint) bool {
	log.Log.Info("Adding Transport to Group", "tid", tid, "group", g.GetID())
	member := &GroupTransportMember{
		GroupID:     g.GetID(),
		TransportID: tid,
	}
	tx := db.Db.Model(&g).Association("Transports").Append(member)
	if tx != nil {
		log.Log.Error(tx, "Error Adding Transport to Group", "tid", tid, "group", g.GetID())
		return false
	}
	g.Transports = append(g.Transports, *member)
	return true
}

func (g Group) DelTransportFromGroup(ctx context.Context, tid uint) bool {
	log.Log.Info("Deleting Transport from Group", "tid", tid, "group", g.GetID())
	member := &GroupTransportMember{
		GroupID:     g.GetID(),
		TransportID: tid,
	}
	tx := db.Db.Model(&g).Association("Transports").Delete(member)
	if tx != nil {
		log.Log.Error(tx, "Error Deleting Transport from Group", "tid", tid, "group", g.GetID())
		return false
	}
	return true
}
