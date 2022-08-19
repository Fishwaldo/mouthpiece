package users

import (
	"context"

	"github.com/Fishwaldo/mouthpiece/pkg/db"
	mperror "github.com/Fishwaldo/mouthpiece/pkg/errors"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/log"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UsersService struct {
}

func NewUsersService() *UsersService {
	us := &UsersService{}

	return us
}

func (us UsersService) Init() {
	db.Db.Debug().AutoMigrate(&User{})
	var count int64
	db.Db.Debug().Model(&User{}).Count(&count)
	log.Log.V(1).Info("Initializing Users", "count", count)
	if count == 0 {
		log.Log.Info("Creating Default Users")
		admin := interfaces.UserDetails{FirstName: "Admin", LastName: "User", Email: "admin@example.com", Password: "password"}
		if admin, err := us.CreateUser(context.Background(), admin); err == nil {
			admin.AddRoleToUser(context.Background(), "admin")
			log.Log.Info("Created Default Admin admin@example.com")
		}
		user := interfaces.UserDetails{FirstName: "User", LastName: "User", Email: "user@example.com", Password: "password"}
		if _, err := us.CreateUser(context.Background(), user); err == nil {
			log.Log.Info("Created Default User user@example.com")
		}
	}
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
	tx := db.Db.WithContext(ctx).Preload(clause.Associations).First(&user, "email = ?", email)
	if tx.Error == gorm.ErrRecordNotFound {
		return nil, mperror.ErrUserNotFound
	}
	return user, nil
}
func (u UsersService) GetUser(ctx context.Context, id uint) (user interfaces.UserI, err error) {
	tx := db.Db.WithContext(ctx).Preload(clause.Associations).First(&user, "ID = ?", id)
	if tx.Error == gorm.ErrRecordNotFound {
		return nil, mperror.ErrUserNotFound
	}
	return user, nil
}
