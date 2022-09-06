package users

import (
	"context"

	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/log"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm/clause"
)

type UsersService struct {
	ctx *interfaces.MPContext
}

func NewUsersService() *UsersService {
	us := &UsersService{
	}
	return us
}

func (us UsersService) Start(ctx *interfaces.MPContext) error {
	us.ctx = ctx
	db.Db.AutoMigrate(&User{})
	return nil
}

func (us UsersService) CreateUser(ctx context.Context, user interfaces.UserDetails) (interfaces.UserI, error) {
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		log.Log.Info("User Validation Error", "Error", err)
		return nil, err
	}
	var newuser User
	newuser.UserDetails = user

	tx := db.Db.WithContext(ctx).Omit("Password").Create(&newuser)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if dbuser, err := us.GetUser(ctx, newuser.ID); err == nil {
		/* Set the Users Initial Password */
		if err := dbuser.SetPassword(ctx, user.Password); err != nil {
			if tx := db.Db.WithContext(ctx).Delete(&dbuser); tx.Error != nil {
				log.Log.Info("Error Deleting User after failed Password", "Error", tx.Error)
			}
			return nil, err
		}
		/* New Users all Start with User Role */
		if !dbuser.AddRoleToUser(ctx, "user") {
			log.Log.Info("Error Adding User Role", "Error", err)
		}
		log.Log.Info("User Created", "User", dbuser)
		return dbuser, nil
	} else {
		return nil, err
	}
}

func (u UsersService) GetUsers(ctx context.Context) map[uint]interfaces.UserI {
	var users []User
	db.Db.WithContext(ctx).Find(&users)
	userMap := make(map[uint]interfaces.UserI)
	for _, user := range users {
		userMap[user.ID] = &user
	}
	return userMap
}

func (u UsersService) GetUserByEmail(ctx context.Context, email string) (user interfaces.UserI, err error) {
	var dbuser User
	tx := db.Db.WithContext(ctx).Preload(clause.Associations).First(&dbuser, "email = ?", email)
	return &dbuser, tx.Error
}
func (u UsersService) GetUser(ctx context.Context, id uint) (user interfaces.UserI, err error) {
	var dbuser User
	tx := db.Db.WithContext(ctx).Preload(clause.Associations).First(&dbuser, "id = ?", id)
	return &dbuser, tx.Error
}
