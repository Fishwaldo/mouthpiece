package transport

import (
	"errors"
	"context"

	"github.com/Fishwaldo/mouthpiece/internal/db"
	. "github.com/Fishwaldo/mouthpiece/internal/log"
	"github.com/Fishwaldo/mouthpiece/internal/message"
	"gorm.io/gorm"
)

type TransportConfig struct {
	gorm.Model `json:"-"`
	UserID     uint `json:"-"`
	Transport  string
	Config     string
}

type ITransport interface {
	GetName() string
	Start()
	SendMessage(ctx context.Context, config TransportConfig, message msg.Message) (err error)
	NewTransportConfig(ctx context.Context)
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

func GetTransport(ctx context.Context, name string) (ITransport, error) {
	if t, ok := transports[name]; ok {
		return t, nil
	}
	return nil, errors.New("Transport Not Found")
}

func GetTransports(ctx context.Context) []string {
	var a []string
	for k := range transports {
		a = append(a, k)
	}
	return a
}

func UpdateTransportStatus(ctx context.Context, t ITransport, m msg.Message, status string) {
	Log.Info("Transport Status", "status", status, "MessageID", m.ID, "Transport", t.GetName())
}
