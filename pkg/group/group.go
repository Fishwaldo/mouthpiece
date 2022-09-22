package group

import (
	"context"
	"fmt"
	"sync"

	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/ent"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"
	"github.com/go-logr/logr"
)

type Group struct {
	dbGroup *ent.DbGroup
	lock    sync.RWMutex
	log     logr.Logger
}

func newGroup(ctx context.Context, logger logr.Logger, name string, description string) (*Group, error) {
	newlogger := logger.WithName("Group").WithValues("name", name)
	dbgrp, err := db.DbClient.DbGroup.Create().
		SetName(name).
		SetDescription(description).
		Save(ctx)
	if err != nil {
		newlogger.Error(err, "Error creating group")
		return nil, mperror.FilterErrors(err)
	}
	group := &Group{
		dbGroup: dbgrp,
		log: newlogger,
	}
	if err := group.init(); err != nil {
		group.log.Error(err, "Error initializing group")
		return nil, mperror.FilterErrors(err)
	}
	group.log.V(1).Info("New Group Created")
	return group, nil
}

func (g *Group) init() error {
	return mperror.FilterErrors(nil)
}

func (g *Group) Load(ctx context.Context, logger logr.Logger, newgroup any) error {
	g.lock.Lock()
	defer g.lock.Unlock()
	var ok bool
	g.dbGroup, ok = newgroup.(*ent.DbGroup)
	if !ok {
		logger.Error(mperror.ErrInvalidType, "Invalid type passed to Load", "type", fmt.Sprintf("%T", newgroup))
		return mperror.ErrInvalidType
	}
	g.log = logger.WithName("Group").WithValues("name", g.dbGroup.Name)
	return g.init()
}

func (g *Group) Save(ctx context.Context) (err error) {
	g.lock.Lock()
	defer g.lock.Unlock()

	dbtmp, err := g.dbGroup.Update().Save(ctx)
	if err != nil {
		g.log.Error(err, "Error saving group")
		return mperror.FilterErrors(err)
	}
	g.dbGroup = dbtmp
	return nil
}

func (g *Group) GetID() int {
	g.lock.RLock()
	defer g.lock.RUnlock()

	return g.dbGroup.ID
}

func (g *Group) GetName() string {
	g.lock.RLock()
	defer g.lock.RUnlock()

	return g.dbGroup.Name
}

func (g *Group) SetName(ctx context.Context, name string) (err error) {
	g.lock.Lock()
	defer g.lock.Unlock()

	dbtmp, err := g.dbGroup.Update().SetName(name).Save(ctx)
	if err != nil {
		g.log.Error(err, "Error setting name")
		return mperror.FilterErrors(err)
	}
	g.dbGroup = dbtmp
	return nil
}

func (g *Group) GetDescription() string {
	g.lock.RLock()
	defer g.lock.RUnlock()

	return g.dbGroup.Description
}

func (g *Group) SetDescription(ctx context.Context, description string) (err error) {
	g.lock.Lock()
	defer g.lock.Unlock()

	dbtmp, err := g.dbGroup.Update().SetDescription(description).Save(ctx)
	if err != nil {
		g.log.Error(err, "Error setting description")
		return mperror.FilterErrors(err)
	}
	g.dbGroup = dbtmp
	return nil
}

func (g *Group) GetApps(ctx context.Context) []interfaces.AppI {
	g.lock.RLock()
	defer g.lock.RUnlock()

	if apps, err := g.dbGroup.QueryApps().All(ctx); err != nil {
		g.log.Error(err, "Error getting Apps for Group")
		return nil
	} else {
		appsvc := interfaces.GetAppService(ctx)
		var ret []interfaces.AppI
		for _, app := range apps {
			if appi, err := appsvc.Load(ctx, app); err != nil {
				g.log.Error(err, "Error Loading App for Group", "app", app.ID)
			} else {
				ret = append(ret, appi)
			}
		}
		return ret
	}
}

func (g *Group) AddApp(ctx context.Context, app interfaces.AppI) (err error) {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.log.Info("Adding App to Group", "app", app.GetName())
	dbtmp, err := g.dbGroup.Update().AddAppIDs(app.GetID()).Save(ctx)
	if err != nil {
		g.log.Error(err, "Error adding app to group")
		return mperror.FilterErrors(err)
	}
	g.dbGroup = dbtmp
	return nil
}

func (g *Group) DelApp(ctx context.Context, app interfaces.AppI) (err error) {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.log.Info("Deleting App from Group", "app", app)
	dbtmp, err := g.dbGroup.Update().RemoveAppIDs(app.GetID()).Save(ctx)
	if err != nil {
		g.log.Error(err, "Error deleting app from group")
		return mperror.FilterErrors(err)
	}
	g.dbGroup = dbtmp
	return nil
}

func (g *Group) GetUsers(ctx context.Context) []interfaces.UserI {
	g.lock.RLock()
	defer g.lock.RUnlock()

	if users, err := g.dbGroup.QueryUsers().All(ctx); err != nil {
		g.log.Error(err, "Error getting Users for Group")
		return nil
	} else {
		usersvc := interfaces.GetUserService(ctx)
		var ret []interfaces.UserI
		for _, user := range users {
			if useri, err := usersvc.Load(ctx, user); err != nil {
				g.log.Error(err, "Error Loading User for Group", "user", user.ID)
			} else {
				ret = append(ret, useri)
			}
		}
		return ret
	}
}

func (g *Group) AddUser(ctx context.Context, user interfaces.UserI) (err error) {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.log.Info("Adding User to Group")
	dbtmp, err := g.dbGroup.Update().AddUserIDs(user.GetID()).Save(ctx)
	if err != nil {
		g.log.Error(err, "Error adding user to group")
		return mperror.FilterErrors(err)
	}
	g.dbGroup = dbtmp
	return nil
}

func (g *Group) DelUser(ctx context.Context, user interfaces.UserI) (err error) {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.log.Info("Deleting User from Group", "user", user)
	dbtmp, err := g.dbGroup.Update().RemoveUserIDs(user.GetID()).Save(ctx)
	if err != nil {
		g.log.Error(err, "Error deleting user from group")
		return mperror.FilterErrors(err)
	}
	g.dbGroup = dbtmp
	return nil
}

func (g *Group) GetTransportRecipients(ctx context.Context) []interfaces.TransportRecipient {
	g.lock.RLock()
	defer g.lock.RUnlock()

	if tpr, err := g.dbGroup.QueryTransportRecipients().All(ctx); err != nil {
		g.log.Error(err, "Error getting TransportRecipients for Group")
		return nil
	} else {
		tprsvc := interfaces.GetTransportService(ctx)
		var ret []interfaces.TransportRecipient
		for _, t := range tpr {
			if tpr, err := tprsvc.Load(ctx, t); err != nil {
				g.log.Error(err, "Error Loading TransportRecipient for Group", "tpr", t.ID)
			} else {
				ret = append(ret, tpr)
			}
		}
		return ret
	}
}

func (g *Group) AddTransportRecipient(ctx context.Context, tid interfaces.TransportRecipient) (err error) {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.log.Info("Adding Transport to Group", "tid", tid)
	dbtmp, err := g.dbGroup.Update().AddTransportRecipientIDs(tid.GetID()).Save(ctx)
	if err != nil {
		g.log.Error(err, "Error adding transport to group")
		return mperror.FilterErrors(err)
	}
	g.dbGroup = dbtmp
	return nil
}

func (g *Group) DelTransportRecipient(ctx context.Context, tid interfaces.TransportRecipient) (err error) {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.log.Info("Deleting Transport from Group", "tid", tid)
	dbtmp, err := g.dbGroup.Update().RemoveTransportRecipientIDs(tid.GetID()).Save(ctx)
	if err != nil {
		g.log.Error(err, "Error deleting transport from group")
		return mperror.FilterErrors(err)
	}
	g.dbGroup = dbtmp
	return nil
}

func (g *Group) ProcessMessage(ctx context.Context, msg interfaces.MessageI) (err error) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.log.Info("Processing Message for Group", "msg", msg.GetID())
 
	var users []*ent.DbUser
	if users, err = g.dbGroup.Edges.UsersOrErr(); err != nil {
		if users, err = g.dbGroup.QueryUsers().All(ctx); err != nil {
			g.log.Error(err, "Error getting Users for Group")
			return mperror.ErrInternalError
		}
	}
	for _, user := range users {
		g.log.V(1).Info("Group Processing Message for User", "user", user.Email, "msg", msg.GetID())
		usrobj, err := interfaces.GetUserService(ctx).Load(ctx, user)
		if err != nil {
			g.log.Error(err, "Error Loading User for Group", "user", user.Email)
			continue
		}
		if err = usrobj.ProcessMessage(ctx, msg.Clone()); err != nil {
			g.log.Error(err, "Error Processing Message for User", "user", user.Email, "msg", msg.GetID())
		}
	}

	var tpr []*ent.DbTransportRecipients
	if tpr, err = g.dbGroup.Edges.TransportRecipientsOrErr(); err != nil {
		if tpr, err = g.dbGroup.QueryTransportRecipients().All(ctx); err != nil {
			g.log.Error(err, "Error loading TransportRecipients for Group")
			return err
		}
	}
	for _, tr := range tpr {
		g.log.V(1).Info("Group Processing Message with TransportRecipient", "TransportRecipient", tr.Name)
		trobj, err := interfaces.GetTransportService(ctx).Load(ctx, tr)
		if err != nil {
			g.log.Error(err, "Error loading TransportRecipient", "TransportRecipient", tr.Name)
			continue
		}
		if err := trobj.ProcessMessage(ctx, msg.Clone()); err != nil {
			g.log.Error(err, "Error processing message with transport recipient", "TransportRecipient", trobj.GetName())
			continue
		}
	}

	return nil
}

var _ interfaces.GroupI = (*Group)(nil)
