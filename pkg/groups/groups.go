package groups

import (
	"context"

	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
)

type Group struct {
	ID         uint   `gorm:"primary_key"`
	Name       string `gorm:"not null"`
	Apps       []uint `gorm:"many2many:group_apps"`
	Users      []uint `gorm:"many2many:group_users"`
	Transports []uint `gorm:"many2many:group_apps"`
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
	return g.Apps
}
func (g Group) GetUsers() []uint {
	return g.Users
}
func (g Group) GetTransports() []uint {
	return g.Transports
}
func (g Group) AddAppToGroup(ctx context.Context, app uint) bool {
	tx := db.Db.Model(&g).Association("Apps").Append(app)
	if tx != nil {
		log.Log.Error(tx, "Error Adding App to Group", "app", app, "group", g.GetID())
		return false
	}
	return true
}

func (g Group) DelAppFromGroup(ctx context.Context, app uint) bool {
	tx := db.Db.Model(&g).Association("Apps").Delete(app)
	if tx != nil {
		log.Log.Error(tx, "Error Deleting App from Group", "app", app, "group", g.GetID())
		return false
	}
	return true
}

func (g Group) AddUserToGroup(ctx context.Context, user uint) bool {
	tx := db.Db.Model(&g).Association("Users").Append(user)
	if tx != nil {
		log.Log.Error(tx, "Error Adding User to Group", "user", user, "group", g.GetID())
		return false
	}
	return true
}

func (g Group) DelUserFromGroup(ctx context.Context, user uint) bool {
	tx := db.Db.Model(&g).Association("Users").Delete(user)
	if tx != nil {
		log.Log.Error(tx, "Error Deleting User from Group", "user", user, "group", g.GetID())
		return false
	}
	return true
}

func (g Group) AddTransportToGroup(ctx context.Context, tid uint) bool {
	tx := db.Db.Model(&g).Association("Transports").Append(tid)
	if tx != nil {
		log.Log.Error(tx, "Error Adding Transport to Group", "tid", tid, "group", g.GetID())
		return false
	}
	return true
}

func (g Group) DelTransportFromGroup(ctx context.Context, tid uint) bool {
	tx := db.Db.Model(&g).Association("Transports").Delete(tid)
	if tx != nil {
		log.Log.Error(tx, "Error Deleting Transport from Group", "tid", tid, "group", g.GetID())
		return false
	}
	return true
}
