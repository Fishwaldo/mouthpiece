package msg

import (
	"github.com/Fishwaldo/mouthpiece/pkg/db"
	"github.com/Fishwaldo/mouthpiece/pkg/errors"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"gorm.io/gorm"
	"time"
)

func InitializeMessage() {
	db.Db.AutoMigrate(&Message{})
	db.Db.AutoMigrate(&MessageResult{})
}

type Message struct {
	gorm.Model
	AppName string `path:"application" doc:"Application Name" Example:"MyApp"`
	Body    struct {
		Message   string                 `json:"message" doc:"Message to be Sent"`
		ShortMsg  string                 `json:"shortmessage,omitempty" doc:"Short Message to be Sent"`
		Topic     string                 `json:"topic,omitempty" doc:"Topic of Message"`
		Severity  string                 `json:"severity,omitempty" doc:"Severity of Message" default:"INFO"`
		Timestamp time.Time              `json:"timestamp,omitempty" doc:"Timestamp of Message"`
		Fields    map[string]interface{} `json:"fields,omitempty" doc:"Additional Fields" gorm:"-"`
	} `json:"body" doc:"Message Body" gorm:"embedded"`
	Result *MessageResult `json:"result,omitempty" doc:"Result of Message"`
}

type MessageResult struct {
	MessageID uint   `json:"message_id" doc:"Message ID"`
	Status    string `json:"status" doc:"Status of Message"`
}

func (msg *Message) ProcessMessage() (err error) {
	if len(msg.AppName) == 0 {
		return mperror.ErrAppNotFound
	} else {
		log.Log.V(1).Info("Processing Message", "Message", msg)
		if tx := db.Db.Create(&msg); tx.Error != nil {
			panic("Can't Save Message to Database: %s" + tx.Error.Error())
		}
		msg.Result = &MessageResult{MessageID: msg.ID, Status: "Queued"}
		return nil
	}
}
