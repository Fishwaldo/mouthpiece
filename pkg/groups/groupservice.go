package groups

import (
	"context"

	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/msg"

	"github.com/jinzhu/copier"
)

type GroupsService struct {
	ctx *interfaces.MPContext
}

func NewGroupsService() *GroupsService {
	return &GroupsService{
	}
}

func (gs *GroupsService) Start(ctx *interfaces.MPContext) error {
	gs.ctx = ctx
	db.Db.AutoMigrate(&Group{}, &GroupAppMember{}, &GroupUserMember{}, &GroupTransportMember{})
	return nil
}

func (gs GroupsService) CreateGroup(ctx context.Context, name string) (interfaces.GroupI, error) {
	log.Log.Info("Group: Creating Group", "name", name)
	var group = Group{
		Name: name,
	}
	tx := db.Db.Create(&group)
	if tx.Error != nil {
		log.Log.Error(tx.Error, "Error Creating Group", "name", name)
		return nil, tx.Error
	}
	return group, tx.Error
}

func (gs GroupsService) DeleteGroup(ctx context.Context, group interfaces.GroupI) error {
	log.Log.Info("Group: Deleting Group", "name", group.GetName())
	tx := db.Db.WithContext(ctx).First(&group, group.GetID())
	if tx.Error != nil {
		return tx.Error
	}
	tx = db.Db.WithContext(ctx).Delete(&group)
	return tx.Error
}

func (gs GroupsService) GetGroupByID(ctx context.Context, id uint) (interfaces.GroupI, error) {
	var group Group
	tx := db.Db.WithContext(ctx).Preload("Apps").Preload("Users").Preload("Transports").Find(&group, id)
	return group, tx.Error
}

func (gs GroupsService) GetGroup(ctx context.Context, name string) (interfaces.GroupI, error) {
	var group Group
	tx := db.Db.WithContext(ctx).Preload("Apps").Preload("Users").Preload("Transports").First(&group, "name = ?", name)
	return group, tx.Error
}

func (g GroupsService) GetGroupsForApp(ctx context.Context, app interfaces.AppI) ([]interfaces.GroupI, error) {
	var groups []Group
	tx := db.Db.WithContext(ctx).Preload("Apps", "id = ?", app.GetID()).Find(&groups).Error
	var groupList []interfaces.GroupI
	for _, group := range groups {
		gp, _ := g.GetGroupByID(ctx, group.GetID())
		groupList = append(groupList, gp)
	}
	return groupList, tx
}

func (g GroupsService) GetGroupsForUser(ctx context.Context, user interfaces.UserI) ([]interfaces.GroupI, error) {
	var groups []Group
	tx := db.Db.WithContext(ctx).Preload("Users", "id = ?", user.GetID()).Find(&groups).Error
	var groupList []interfaces.GroupI
	for _, group := range groups {
		gp, _ := g.GetGroupByID(ctx, group.GetID())
		groupList = append(groupList, gp)
	}
	return groupList, tx
}

func (g GroupsService) GetGroupsForTransport(ctx context.Context, tid uint) ([]interfaces.GroupI, error) {
	var groups []Group
	tx := db.Db.WithContext(ctx).Preload("Transports", "id = ?", tid).Find(&groups).Error
	var groupList []interfaces.GroupI
	for _, group := range groups {
		gp, _ := g.GetGroupByID(ctx, group.GetID())
		groupList = append(groupList, gp)
	}
	return groupList, tx
}

func (g GroupsService) SendMessageToUsers(ctx context.Context, sendmsg *msg.Message, app interfaces.AppI) error {
	log.Log.V(1).Info("Group: Sending Message to Users", "app", app.GetName())
	var sendto []uint
	if groups, err := g.GetGroupsForApp(ctx, app); err != nil {
		log.Log.Error(err, "Error Getting Groups for App", "app", app.GetName())
		return err
	} else {
		for _, group := range groups {
			sendto = append(sendto, group.GetUsers()...)
		}
	}
	sendto = removeDuplicate(sendto)
	log.Log.V(1).Info("Group: Sending Message to Users", "app", app.GetName(), "userid", sendto)
	for _, userid := range sendto {
		if user, err := g.ctx.GetUserService().GetUser(ctx, userid); err != nil {
			log.Log.Error(err, "Error Getting User", "userid", userid)
		} else {
			var usrmsg msg.Message
			usrmsg.Body.Fields = make(map[string]interface{})
			copier.Copy(usrmsg, sendmsg)
			if err := user.ProcessMessage(ctx, usrmsg); err != nil {
				log.Log.Error(err, "Error Processing Message for User", "userid", userid)
			}
		}
	}
	return nil
}

func (g GroupsService) SendMessageToTransports(ctx context.Context, sendmsg *msg.Message, app interfaces.AppI) error {
	log.Log.V(1).Info("Group: Sending Message to Transports", "app", app.GetName())
	var sendto []uint
	if groups, err := g.GetGroupsForApp(ctx, app); err != nil {
		log.Log.Error(err, "Error Getting Groups for App", "app", app.GetName())
		return err
	} else {
		for _, group := range groups {
			sendto = append(sendto, group.GetTransports()...)
		}
	}
	sendto = removeDuplicate(sendto)
	log.Log.V(1).Info("Group: Sending Message to Transports", "app", app.GetName(), "transportid", sendto)
	for _, tid := range sendto {
		if transport, err := g.ctx.GetTransportService().GetTransportReciepient(ctx, tid); err != nil {
			log.Log.Error(err, "Error Getting Transport", "transportid", tid)
		} else {
			var usrmsg msg.Message
			usrmsg.Body.Fields = make(map[string]interface{})
			copier.Copy(usrmsg, sendmsg)
			if err := transport.ProcessGroupMessage(ctx, usrmsg); err != nil {
				log.Log.Error(err, "Error Processing Message for Group Transport", "transportid", tid)
			}
		}
	}
	return nil
}

func removeDuplicate[T uint | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
