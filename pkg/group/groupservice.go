package group

import (
	"context"
	"fmt"

	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/ent"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbgroup"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"

	"github.com/go-logr/logr"
)

type GroupService struct {
	log logr.Logger
}

func NewGroupService(_ context.Context, logger logr.Logger) *GroupService {
	newlog := logger.WithName("GroupService")
	newlog.V(1).Info("New Group Service")
	return &GroupService{
		log: newlog,
	}
}

func (gs *GroupService) Start(ctx context.Context) error {
	return nil
}

func (gs *GroupService) Create(ctx context.Context, name string, description string) (grp interfaces.GroupI, err error) {
	if ok, err := gs.Exists(ctx, name); err != nil {
		gs.log.V(2).Error(err, "Error checking if group exists", "name", name)
		return nil, mperror.ErrInternalError
	} else if ok {
		return nil, mperror.ErrGroupExists
	}


	gs.log.Info("Group: Creating Group", "name", name)
	grp, err = newGroup(ctx, gs.log, name, description)
	if err != nil {
		gs.log.Error(err, "Error creating group", "name", name)
		return nil, err
	} else {
		return grp, nil
	}
}

func (gs *GroupService) Delete(ctx context.Context, group interfaces.GroupI) error {
	if ok, err := gs.Exists(ctx, group.GetName()); err != nil {
		gs.log.V(1).Error(err, "Error checking if group exists", "name", group.GetName())
		return mperror.ErrInternalError
	} else if !ok {
		return mperror.ErrGroupNotFound
	}

	gs.log.Info("Group: Deleting Group", "name", group.GetName())
	err := db.DbClient.DbGroup.DeleteOneID(group.GetID()).Exec(ctx)
	if err != nil {
		gs.log.Error(err, "Error deleting group", "name", group.GetName())
	}
	return err
}

func (gs *GroupService) GetByID(ctx context.Context, id int) (interfaces.GroupI, error) {

	grp, err := db.DbClient.DbGroup.Query().Where(dbgroup.IDEQ(id)).Only(ctx)
	if err != nil {
		gs.log.V(1).Error(err, "Error getting group by ID", "id", id)
		return nil, mperror.ErrGroupNotFound
	}
	return gs.Load(ctx, grp)
}

func (gs *GroupService) Get(ctx context.Context, name string) (interfaces.GroupI, error) {

	grp, err := db.DbClient.DbGroup.Query().Where(dbgroup.NameEQ(name)).Only(ctx)
	if err != nil {
		gs.log.V(1).Error(err, "Error getting group", "name", name)
		return nil, mperror.ErrGroupNotFound
	}
	return gs.Load(ctx, grp)
}

func (gs *GroupService) GetAll(ctx context.Context) ([]interfaces.GroupI, error) {

	groups, err := db.DbClient.DbGroup.Query().All(ctx)
	if err != nil {
		gs.log.Error(err, "Error getting all groups")
		return nil, err
	}
	var g []interfaces.GroupI
	for _, grp := range groups {
		newg, err := gs.Load(ctx, grp)
		if err != nil {
			gs.log.Error(err, "Error loading group", "name", grp.Name)
		} else {
			g = append(g, newg)
		}
	}
	return g, nil
}


func (gs *GroupService) Load(ctx context.Context, dbGroup any) (interfaces.GroupI, error) {

	var ok bool
	var grp *ent.DbGroup
	grp, ok = dbGroup.(*ent.DbGroup)
	if !ok {
		gs.log.Error(mperror.ErrInvalidType, "Error loading group", "type", fmt.Sprintf("%T", dbGroup))
		return nil, mperror.ErrInvalidType
	}
	g := &Group{}
	if err := g.Load(ctx, gs.log, grp); err != nil {
		gs.log.Error(err, "Error loading group", "name", grp.Name)
		return nil, err
	}
	return g, nil
}

func (gs *GroupService) Exists(ctx context.Context, name string) (bool, error) {
	if ok, err := db.DbClient.DbGroup.Query().Where(dbgroup.NameEQ(name)).Exist(ctx); err != nil {
		gs.log.V(1).Error(err, "Error checking if group exists", "name", name)
		return false, mperror.ErrInternalError
	} else {
		return ok, nil
	}
}

func (gs *GroupService) ExistsByID(ctx context.Context, id int) (bool, error) {
	if ok, err :=  db.DbClient.DbGroup.Query().Where(dbgroup.IDEQ(id)).Exist(ctx); err != nil {
		gs.log.V(1).Error(err, "Error checking if group exists", "id", id)
		return false, mperror.ErrInternalError
	} else {
		return ok, nil
	}
}

// func (g GroupsService) GetGroupsForApp(ctx context.Context, app interfaces.AppI) ([]interfaces.GroupI, error) {
// 	var groups []Group
// 	tx := db.Db.WithContext(ctx).Preload("Apps", "id = ?", app.GetID()).Find(&groups).Error
// 	var groupList []interfaces.GroupI
// 	for _, group := range groups {
// 		gp, _ := g.GetGroupByID(ctx, group.GetID())
// 		groupList = append(groupList, gp)
// 	}
// 	return groupList, tx
// }

// func (g GroupsService) GetGroupsForUser(ctx context.Context, user interfaces.UserI) ([]interfaces.GroupI, error) {
// 	var groups []Group
// 	tx := db.Db.WithContext(ctx).Preload("Users", "id = ?", user.GetID()).Find(&groups).Error
// 	var groupList []interfaces.GroupI
// 	for _, group := range groups {
// 		gp, _ := g.GetGroupByID(ctx, group.GetID())
// 		groupList = append(groupList, gp)
// 	}
// 	return groupList, tx
// }

// func (g GroupsService) GetGroupsForTransport(ctx context.Context, tid uint) ([]interfaces.GroupI, error) {
// 	var groups []Group
// 	tx := db.Db.WithContext(ctx).Preload("Transports", "id = ?", tid).Find(&groups).Error
// 	var groupList []interfaces.GroupI
// 	for _, group := range groups {
// 		gp, _ := g.GetGroupByID(ctx, group.GetID())
// 		groupList = append(groupList, gp)
// 	}
// 	return groupList, tx
// }

// func (g GroupsService) SendMessageToUsers(ctx context.Context, sendmsg *msg.Message, app interfaces.AppI) error {
// 	log.Log.V(1).Info("Group: Sending Message to Users", "app", app.GetName())
// 	var sendto []uint
// 	if groups, err := g.GetGroupsForApp(ctx, app); err != nil {
// 		log.Log.Error(err, "Error Getting Groups for App", "app", app.GetName())
// 		return err
// 	} else {
// 		for _, group := range groups {
// 			sendto = append(sendto, group.GetUsers()...)
// 		}
// 	}
// 	sendto = removeDuplicate(sendto)
// 	log.Log.V(1).Info("Group: Sending Message to Users", "app", app.GetName(), "userid", sendto)
// 	for _, userid := range sendto {
// 		if user, err := g.ctx.GetUserService().GetUser(ctx, userid); err != nil {
// 			log.Log.Error(err, "Error Getting User", "userid", userid)
// 		} else {
// 			var usrmsg msg.Message
// 			usrmsg.Body.Fields = make(map[string]interface{})
// 			copier.Copy(usrmsg, sendmsg)
// 			if err := user.ProcessMessage(ctx, usrmsg); err != nil {
// 				log.Log.Error(err, "Error Processing Message for User", "userid", userid)
// 			}
// 		}
// 	}
// 	return nil
// }

// func (g GroupsService) SendMessageToTransports(ctx context.Context, sendmsg *msg.Message, app interfaces.AppI) error {
// 	log.Log.V(1).Info("Group: Sending Message to Transports", "app", app.GetName())
// 	var sendto []uint
// 	if groups, err := g.GetGroupsForApp(ctx, app); err != nil {
// 		log.Log.Error(err, "Error Getting Groups for App", "app", app.GetName())
// 		return err
// 	} else {
// 		for _, group := range groups {
// 			sendto = append(sendto, group.GetTransports()...)
// 		}
// 	}
// 	sendto = removeDuplicate(sendto)
// 	log.Log.V(1).Info("Group: Sending Message to Transports", "app", app.GetName(), "transportid", sendto)
// 	for _, tid := range sendto {
// 		if transport, err := g.ctx.GetTransportService().GetTransportReciepient(ctx, tid); err != nil {
// 			log.Log.Error(err, "Error Getting Transport", "transportid", tid)
// 		} else {
// 			var usrmsg msg.Message
// 			usrmsg.Body.Fields = make(map[string]interface{})
// 			copier.Copy(usrmsg, sendmsg)
// 			if err := transport.ProcessGroupMessage(ctx, usrmsg); err != nil {
// 				log.Log.Error(err, "Error Processing Message for Group Transport", "transportid", tid)
// 			}
// 		}
// 	}
// 	return nil
// }

// func removeDuplicate[T uint | int](sliceList []T) []T {
// 	allKeys := make(map[T]bool)
// 	list := []T{}
// 	for _, item := range sliceList {
// 		if _, value := allKeys[item]; !value {
// 			allKeys[item] = true
// 			list = append(list, item)
// 		}
// 	}
// 	return list
// }

var _ interfaces.GroupServiceI = (*GroupService)(nil)
