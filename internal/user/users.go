package user

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"

	"github.com/Fishwaldo/mouthpiece/internal/auth"
	"github.com/Fishwaldo/mouthpiece/internal/db"
	"github.com/Fishwaldo/mouthpiece/internal/errors"
	. "github.com/Fishwaldo/mouthpiece/internal/log"
	"github.com/Fishwaldo/mouthpiece/internal/message"
	"github.com/Fishwaldo/mouthpiece/internal/transport"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	gorm.Model       `json:"-"`
	ID               uint                        `gorm:"primarykey"`
	Email            string                      `validate:"required,email"`
	FirstName        string                      `validate:"required"`
	LastName         string                      `validate:"required"`
	Password         string                      `json:"-" writeOnly:"true" validate:"required"`
	TransportConfigs []transport.TransportConfig `json:"transports,omitempty" gorm:"many2many:user_transports;" validate:"-"`
}

var AuthConfig auth.AuthConfig

func init() {
	AuthConfig = auth.AuthConfig{
		CredChecker:     dbAuthProvider,
		MapClaimsToUser: MapClaimsToUser,
		Validator:       UserValidator,
	}
}

func CreateUser(user *User) error {
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		Log.Info("User Validation Error", "Error", err)
		return err
	}
	tx := db.Db.Omit("Password").Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	if dbuser, err := GetUser(user.Email); err == nil {
		/* Set the Users Initial Password */
		if err := dbuser.SetPassword(user.Password); err != nil {
			if tx := db.Db.Delete(&dbuser); tx.Error != nil {
				Log.Info("Error Deleting User after failed Password", "Error", tx.Error)
				return err
			}
			return err
		}
		/* New Users all Start with User Role */
		if !dbuser.addUserRole("user") {
			Log.Info("Error Adding User Role", "Error", err)
		}
		return nil
	} else {
		return err
	}
}

func (u *User) addUserRole(role string) bool {
	_, err := auth.AuthService.AuthEnforcer.AddRoleForUser(u.Email, fmt.Sprintf("role:%s", role))
	if err != nil {
		Log.Info("Failed to add role for user", "email", u.Email, "role", role, "error", err)
		return false
	}
	return true
}

func (u *User) CheckPassword(password string) bool {
	Log.Info("Checking Password", "email", u.Email)
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		Log.Info("Password Check Failed", "Error", err)
		return false
	}
	return true
}

func (u *User) SetPassword(password string) error {
	Log.Info("Setting Password", "Email", u.Email)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		Log.Info("Error Generating SetPassword Hash", "Error", err)
		return err
	}
	if tx := db.Db.Model(&u).Update("password", string(hashedPassword)); tx.Error != nil {
		Log.Info("Error Setting Password", "Error", tx.Error)
		return tx.Error
	}
	return nil
}

func InitializeUsers() {
	db.Db.AutoMigrate(&User{})
	var count int64
	db.Db.Model(&User{}).Count(&count)
	Log.V(1).Info("Initializing Users", "count", count)
	if count == 0 {
		Log.Info("Creating Default Users")
		admin := &User{FirstName: "Admin", LastName: "User", Email: "admin@example.com", Password: "password"}
		if err := CreateUser(admin); err == nil {
			admin.addUserRole("admin")
			Log.Info("Created Default Admin admin@example.com")
		}
		if err := CreateUser(&User{FirstName: "User", LastName: "User", Email: "user@example.com", Password: "password"}); err == nil {
			Log.Info("Created Default User user@example.com")
		}
	}
}

func GetUsers() []User {
	var users []User
	db.Db.Find(&users)
	return users
}

func GetUser(email string) (user *User, err error) {
	tx := db.Db.Preload(clause.Associations).First(&user, "email = ?", email)
	if tx.Error == gorm.ErrRecordNotFound {
		return nil, mperror.ErrUserNotFound
	}
	return
}
func GetUserByID(id uint) (user *User, err error) {
	tx := db.Db.Preload(clause.Associations).First(&user, "ID = ?", id)
	if tx.Error == gorm.ErrRecordNotFound {
		return nil, mperror.ErrUserNotFound
	}
	return
}

func (u User) ProcessMessage(msg msg.Message) (err error) {
	/* add User Fields to Message */
	msg.Body.Fields["first_name"] = u.FirstName
	msg.Body.Fields["last_name"] = u.LastName
	msg.Body.Fields["email"] = u.Email
	Log.V(1).Info("User Processing Message", "Email", u.Email, "MessageID", msg.ID)
	for _, tc := range u.TransportConfigs {
		t, err := transport.GetTransport(tc.Transport)
		if err != nil {
			Log.Info("Cant find Transport", "Transport", tc.Transport)
		}
		go t.SendMessage(tc, msg)
	}
	return
}
