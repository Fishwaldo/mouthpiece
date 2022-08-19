package users

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"

	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/message"
	"github.com/Fishwaldo/mouthpiece/pkg/transport"
	"github.com/go-playground/validator/v10"
)

type User struct {
	interfaces.UserDetails
	TransportConfigs []transport.TransportConfig `json:"transports,omitempty" gorm:"many2many:user_transports;" validate:"-"`
}

func (u User) AddRoleToUser(ctx context.Context, role string) bool {
	//TODO: Move to Internal
	// _, err := auth.AuthService.AuthEnforcer.AddRoleForUser(u.Email, fmt.Sprintf("role:%s", role))
	// if err != nil {
	// 	log.Log.Info("Failed to add role for user", "email", u.Email, "role", role, "error", err)
	// 	return false
	// }
	return true
}

func (u User) CheckPassword(ctx context.Context, password string) bool {
	log.Log.Info("Checking Password", "email", u.Email)
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		log.Log.Info("Password Check Failed", "Error", err)
		return false
	}
	return true
}

func (u User) SetPassword(ctx context.Context, password string) error {
	log.Log.Info("Setting Password", "Email", u.Email)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Log.Info("Error Generating SetPassword Hash", "Error", err)
		return err
	}
	if tx := db.Db.WithContext(ctx).Model(&u).Update("password", string(hashedPassword)); tx.Error != nil {
		log.Log.Info("Error Setting Password", "Error", tx.Error)
		return tx.Error
	}
	return nil
}

func (u User) GetEmail() string {
	return u.Email
}

func (u User) SetEmail(email string) error {
	var details interfaces.UserDetails = u.UserDetails
	details.Email = email
	return u.SetDetails(details)
}

func (u User) GetID() uint {
	return u.ID
}

func (u User) GetDetails() interfaces.UserDetails {
	return u.UserDetails
}

func (u *User) SetDetails(details interfaces.UserDetails) error {
	validate := validator.New()
	if err := validate.Struct(details); err != nil {
		log.Log.Info("SetDetails Validation Error", "Error", err)
		return err
	}
	if tx := db.Db.Model(&u).Updates(details); tx.Error != nil {
		log.Log.Error(tx.Error, "Error updating User details", "email", u.Email)
		return tx.Error
	} else {
		u.UserDetails = details
	}
	return nil
}

func (u User) GetFirstName() string {
	return u.FirstName
}

func (u User) SetFirstName(fname string) error {
	var details interfaces.UserDetails = u.UserDetails
	details.FirstName = fname
	return u.SetDetails(details)
}

func (u User) GetLastName() string {
	return u.LastName
}

func (u User) SetLastName(lname string) error {
	var details interfaces.UserDetails = u.UserDetails
	details.LastName = lname
	return u.SetDetails(details)
}

func (u User) ProcessMessage(ctx context.Context, msg *msg.Message) (err error) {
	/* add User Fields to Message */
	msg.Body.Fields["first_name"] = u.FirstName
	msg.Body.Fields["last_name"] = u.LastName
	msg.Body.Fields["email"] = u.Email
	log.Log.V(1).Info("User Processing Message", "Email", u.Email, "MessageID", msg.ID)
	for _, tc := range u.TransportConfigs {
		t, err := transport.GetTransport(ctx, tc.Transport)
		if err != nil {
			log.Log.Info("Cant find Transport", "Transport", tc.Transport)
		}
		go t.SendMessage(ctx, tc, *msg)
	}
	return
}
