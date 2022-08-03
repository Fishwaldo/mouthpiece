package transport

import (
	"errors"

	. "github.com/Fishwaldo/mouthpiece/internal/log"
	"github.com/Fishwaldo/mouthpiece/internal/message"
	"github.com/Fishwaldo/mouthpiece/internal/db"
	"gorm.io/gorm"
)

type TransportConfig struct {
	gorm.Model	`json:"-"`
	UserID uint	`json:"-"`
	Transport string
	Config string
}

type ITransport interface {
	GetName() string
	Start()
	SendMessage(config TransportConfig, message msg.Message) (err error)
	NewTransportConfig()
}

var transports map[string]ITransport

func RegisterTransport(transport ITransport) {
	if transports == nil {
		transports = make(map[string]ITransport)
	}
	transports[transport.GetName()] = transport
}

func InitializeTransports() {
	db.Db.AutoMigrate(&TransportConfig{})
}

func StartTransports() {
	for k, t := range transports {
		Log.Info("Starting Transport", "transport", k)
		t.Start()
	}
}

func GetTransport(name string) (ITransport, error) {
	if t, ok := transports[name]; ok {
		return t, nil
	}
	return nil, errors.New("Transport Not Found")
}

func GetTransports() []string {
	var a []string
	for k := range transports {
		a = append(a, k)
	}
	return a
}

func UpdateTransportStatus(t ITransport, m msg.Message, status string) {
	Log.Info("Transport Status", "status", status, "MessageID", m.ID, "Transport", t.GetName())
}
