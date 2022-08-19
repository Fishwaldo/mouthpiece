package stdout

import (
	"context"
	"fmt"
	"github.com/Fishwaldo/mouthpiece/pkg/log"
	"github.com/Fishwaldo/mouthpiece/pkg/message"
	"github.com/Fishwaldo/mouthpiece/pkg/transport"
)

type StdoutTransport struct {
}

func init() {
	tp := NewStdoutTransport()
	transport.RegisterTransport(tp)
}

func NewStdoutTransport() transport.ITransport {
	return &StdoutTransport{}
}

func (t StdoutTransport) GetName() string {
	return "stdout"
}

func (t StdoutTransport) SendMessage(ctx context.Context, config transport.TransportConfig, msg msg.Message) (err error) {
	fmt.Println("=========================================================")
	fmt.Printf("Message: %s\n", msg.Body.Message)
	fmt.Println("=========================================================")
	transport.UpdateTransportStatus(ctx, t, msg, "sent")
	return nil
}

func (t StdoutTransport) Start() {
	log.Log.Info("Transport Started", "name", t.GetName())
}

func (t StdoutTransport) NewTransportConfig(ctx context.Context) {
	//	user.TransportConfigs = append(user.TransportConfigs, mouthpiece.TransportConfig{
	//		Transport: t.GetName(),
	//		Config: user.Username,
	//	})
	return
}
