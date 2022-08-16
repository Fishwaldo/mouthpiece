package stdout

import (
	"fmt"
	. "github.com/Fishwaldo/mouthpiece/internal/log"
	"github.com/Fishwaldo/mouthpiece/internal/message"
	"github.com/Fishwaldo/mouthpiece/internal/transport"
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

func (t StdoutTransport) SendMessage(config transport.TransportConfig, msg msg.Message) (err error) {
	fmt.Println("=========================================================")
	fmt.Printf("Message: %s\n", msg.Body.Message)
	fmt.Println("=========================================================")
	transport.UpdateTransportStatus(t, msg, "sent")
	return nil
}

func (t StdoutTransport) Start() {
	Log.Info("Transport Started", "name", t.GetName())
}

func (t StdoutTransport) NewTransportConfig() {
	//	user.TransportConfigs = append(user.TransportConfigs, mouthpiece.TransportConfig{
	//		Transport: t.GetName(),
	//		Config: user.Username,
	//	})
	return
}
