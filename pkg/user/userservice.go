package user

import (
	"context"
	"fmt"

	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/ent"
	"github.com/Fishwaldo/mouthpiece/pkg/ent/dbuser"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"

	"github.com/go-logr/logr"
)

type UserService struct {
	log logr.Logger
}

func NewUsersService(_ context.Context, logger logr.Logger) *UserService {
	us := &UserService{
		log: logger.WithName("UserService"),
	}
	us.log.V(1).Info("New User Service")
	return us
}


func (us *UserService) Start(ctx context.Context) error {

	return nil
}

func (us *UserService) Create(ctx context.Context, email string, name string) (interfaces.UserI, error) {
	if ok, err := us.Exists(ctx, name); err != nil {
		return nil, mperror.FilterErrors(err)
	} else if ok {
		return nil, mperror.ErrUserExists
	}
	app, err := newUser(ctx, us.log, email, name)
	if err != nil {
		us.log.Error(err, "Error creating app", "name", name)
		return nil, mperror.FilterErrors(err)
	} else {
		return app, nil
	}
}

func (us *UserService) Delete(ctx context.Context, user interfaces.UserI) error {
	if ok, err := us.Exists(ctx, user.GetName()); err != nil {
		return mperror.FilterErrors(err)
	} else if !ok {
		return mperror.ErrUserNotFound
	}
	if err := db.DbClient.DbUser.DeleteOneID(user.GetID()).Exec(ctx); err != nil {
		us.log.Error(err, "Error deleting User", "Email", user.GetEmail())
		return mperror.FilterErrors(err)
	}
	return nil
}

func (us *UserService) Get(ctx context.Context, email string) (interfaces.UserI, error) {
	db_app, err := db.DbClient.DbUser.Query().Where(dbuser.Email(email)).Only(ctx)
	if err != nil {
		us.log.Error(err, "Error getting User", "Email", email)
		return nil, mperror.FilterErrors(err)
	}
	user, err := us.Load(ctx, db_app)
	return user, mperror.FilterErrors(err)
}

func (us *UserService) GetByID(ctx context.Context, id int) (interfaces.UserI, error) {
	db_app, err := db.DbClient.DbUser.Query().Where(dbuser.ID(id)).Only(ctx)
	if err != nil {
		us.log.Error(err, "Error getting User", "ID", id)
		return nil, mperror.FilterErrors(err)
	}
	user, err := us.Load(ctx, db_app)
	return user, mperror.FilterErrors(err)
}


func (us *UserService) GetAll(ctx context.Context) (users []interfaces.UserI, err error) {
	var dbusers []*ent.DbUser
	if dbusers, err = db.DbClient.DbUser.Query().All(ctx); err != nil {
		us.log.Error(err, "Error getting all users")
		return nil, mperror.FilterErrors(err)
	}

	for _, dbuser := range dbusers {
		if user, err := us.Load(ctx, dbuser); err != nil {
			us.log.Error(err, "Error loading user", "email", dbuser.Email)
			return nil, mperror.FilterErrors(err)
		} else {
			users = append(users, user)
		}
	}
	return users, nil
}

func (us *UserService) Load(ctx context.Context, dbuser any) (interfaces.UserI, error) {
	entUser, ok := dbuser.(*ent.DbUser)
	if !ok {
		us.log.Error(mperror.ErrInvalidType, "Error loading user", "type", fmt.Sprintf("%T", dbuser))
		return nil, mperror.ErrInvalidType
	}
	user := &User{}
	if err := user.Load(ctx, us.log, entUser); err != nil {
		us.log.Error(err, "Error loading user", "email", entUser.Email)
		return nil, mperror.FilterErrors(err)
	}
	return user, nil
}


func (us *UserService) Exists(ctx context.Context, name string) (bool, error) {
	if ok, err := db.DbClient.DbUser.Query().Where(dbuser.Name(name)).Exist(ctx); err != nil {
		us.log.Error(err, "Error checking if user exists", "name", name)
		return false, mperror.FilterErrors(err)
	} else {
		return ok, nil
	}
}

func (us *UserService) ExistsByID(ctx context.Context, id int) (bool, error) {
	if ok, err :=  db.DbClient.DbUser.Query().Where(dbuser.ID(id)).Exist(ctx); err != nil {
		us.log.Error(err, "Error checking if user exists", "id", id)
		return false, mperror.FilterErrors(err)
	} else {
		return ok, nil
	}
}


var _ interfaces.UserServiceI = (*UserService)(nil)