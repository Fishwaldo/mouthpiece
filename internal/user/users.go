package user

import (
	"errors"
	
	. "github.com/Fishwaldo/mouthpiece/internal/log"
	"github.com/Fishwaldo/mouthpiece/internal/message"
	"github.com/Fishwaldo/mouthpiece/internal/transport"
	"github.com/Fishwaldo/mouthpiece/internal/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)


type User struct {
	gorm.Model	`json:"-"`
	Username string
	FirstName string
	LastName string
	Email string
	TransportConfigs []transport.TransportConfig `json:"transports,omitempty" gorm:"many2many:user_transports;"`
}

func InitializeUsers() {
	db.Db.AutoMigrate(&User{})
	var count int64
	db.Db.Model(&User{}).Count(&count)
	Log.V(1).Info("Initializing Users", "count", count)
	if (count == 0) {
		admin := &User{Username: "admin", FirstName: "Admin", LastName: "User", Email: "admin@example.com"}
		user := &User{Username: "user", FirstName: "User", LastName: "User", Email: "user@example.com"}
		t, _ := transport.GetTransport("stdout")
		t.NewTransportConfig()
		t.NewTransportConfig()
		Log.Info("Admin User", "user", admin)
		db.Db.Create(admin)
		db.Db.Create(user)
	}
} 

func GetUsers() []User {
	var users []User
	db.Db.Find(&users)
	return users
}

func GetUser(username string) (user *User, err error) {
	tx := db.Db.Preload(clause.Associations).First(&user, "username = ?", username)
	if tx.Error == gorm.ErrRecordNotFound {
		return nil, errors.New("User Not Found")
	}
	return
}

func (u User) ProcessMessage(msg msg.Message) (err error) {
	/* add User Fields to Message */
	msg.Body.Fields["username"] = u.Username
	msg.Body.Fields["first_name"] = u.FirstName
	msg.Body.Fields["last_name"] = u.LastName
	msg.Body.Fields["email"] = u.Email
	Log.V(1).Info("User Processing Message", "User", u.Username, "MessageID", msg.ID)
	for _, tc := range u.TransportConfigs {
		t, err := transport.GetTransport(tc.Transport); 
		if err != nil {
			Log.Info("Cant find Transport", "Transport", tc.Transport)
		}
		go t.SendMessage(tc, msg)
	}
	return
}
