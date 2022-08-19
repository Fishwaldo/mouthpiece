package groups

import (
	"context"

	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/message"
)

type GroupsService struct {
}

func NewGroupsService() *GroupsService {
	db.Db.Debug().AutoMigrate(&Group{})
	return &GroupsService{}
}

func (gs GroupsService) CreateGroup(ctx context.Context, name string) (interfaces.GroupI, error) {
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

func (gs GroupsService) DeleteGroup(ctx context.Context, group Group) error {
	tx := db.Db.WithContext(ctx).First(&group, group.GetID())
	if tx.Error != nil {
		return tx.Error
	}
	tx = db.Db.WithContext(ctx).Delete(&group)
	return tx.Error
}

func (gs GroupsService) GetGroup(ctx context.Context, id uint) (interfaces.GroupI, error) {
	var group Group
	tx := db.Db.WithContext(ctx).First(&group, id)
	return group, tx.Error
}

func (g GroupsService) GetGroupsForApp(ctx context.Context, app uint) ([]interfaces.GroupI, error) {
	var groups []Group
	tx := db.Db.WithContext(ctx).Where("apps = ?", app).Find(&groups)
	var groupList []interfaces.GroupI
	for _, group := range groups {
		groupList = append(groupList, group)
	}
	return groupList, tx.Error
}

func (g GroupsService) GetGroupsForUser(ctx context.Context, id uint) ([]interfaces.GroupI, error) {
	var groups []Group
	tx := db.Db.WithContext(ctx).Where("users = ?", id).Find(&groups)
	var groupList []interfaces.GroupI
	for _, group := range groups {
		groupList = append(groupList, group)
	}
	return groupList, tx.Error
}

func (g GroupsService) SendMessageToUsers(ctx context.Context, msg msg.Message, appid uint, sender interfaces.UserSender) error {
	var sendto []uint
	if groups, err := g.GetGroupsForApp(ctx, appid); err != nil {
		log.Log.Error(err, "Error Getting Groups for App", "appid", appid)
		return err
	} else {
		for _, group := range groups {
			sendto = append(sendto, group.GetUsers()...)
		}
	}
	sendto = removeDuplicate(sendto)
	for _, userid := range sendto {
		if err := sender(ctx, msg, userid); err != nil {
			log.Log.Error(err, "Error Sending Message to User", "userid", userid)
		}
	}
	return nil
}

func (g GroupsService) SendMessageToTransports(ctx context.Context, msg msg.Message, appid uint, sender interfaces.UserSender) error {
	var sendto []uint
	if groups, err := g.GetGroupsForApp(ctx, appid); err != nil {
		log.Log.Error(err, "Error Getting Groups for App", "appid", appid)
		return err
	} else {
		for _, group := range groups {
			sendto = append(sendto, group.GetTransports()...)
		}
	}
	sendto = removeDuplicate(sendto)
	for _, userid := range sendto {
		if err := sender(ctx, msg, userid); err != nil {
			log.Log.Error(err, "Error Sending Message to Transports", "userid", userid)
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
