package messages

import (
	"net/http"
	"time"
	"fmt"

	"github.com/Fishwaldo/mouthpiece/internal/restapi/auth"
	"github.com/Fishwaldo/mouthpiece/internal/server"

	"github.com/Fishwaldo/mouthpiece/pkg/msg"

	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/responses"

	"github.com/google/uuid"
)

type msgListGet struct {
	Page int `json:"page" query:"page"`
	Size int `json:"size" query:"size"`
	OrderBy string `json:"order_by" query:"orderBy"`
	OrderDir string `json:"order_dir" query:"orderDir"`
}

func (m *msgListGet) Resolve(ctx huma.Context, r *http.Request) {
	fmt.Printf("%s\n", r.URL.Query())
}


type msgListData struct {
	ID        uuid.UUID 		  
	Message   string
	ShortMsg  *string
	Topic     *string
	Severity  int
	TimeStamp time.Time
	Fields    map[string]string
	AppID     int
}

type msgListResponse struct {
	Data []msgListData `json:"data"`
	Count int 		   `json:"count"`
}

func Setup(res *huma.Resource) error {
	// //AuthService.AddResourceURL("/v1/apps/", "apigroup:apps")
	appapi := res.SubResource("/messages")
	appapi.Middleware(auth.RequireAuth)
	appapi.Tags("Messages")

	appapi.Get("get-messages", "Get A List of Messages",
		responses.OK().ContentType("application/json"),
		responses.OK().Model(msgListResponse{}),
	).Run(msgList)

	return nil
}

func msgList(ctx huma.Context, input msgListGet) {
	//fmt.Printf("Page: %d, Size: %d Sort: %s %s\n", input.Page, input.Size, input.OrderBy, input.OrderDir)

	var options []func(any)any
	if (input.Size > 0) {
		options = append(options, msg.WithPaginate(input.Page, input.Size))
	}
	if (len(input.OrderBy) > 0) {
		options = append(options, msg.WithSort(input.OrderBy, input.OrderDir))
	}

	msgs, err := server.Get().GetMsgService().GetMessages(ctx, options...)
	if err != nil {
		ctx.WriteError(http.StatusInternalServerError, "Error getting messages")
		return
	}
	count, err := server.Get().GetMsgService().GetMessageCount(ctx)
	if err != nil {
		ctx.WriteError(http.StatusInternalServerError, "Error getting message count")
		return
	}
	var msgmodel msgListResponse

	for _, msg := range msgs {
		flds, _ := msg.GetFields(ctx)
		app, _ := msg.GetApp(ctx)
		nemsg := msgListData{
			ID:        msg.GetID(),
			Message:   msg.GetMessage(),
			ShortMsg:  msg.GetShortMsg(),
			Topic:     msg.GetTopic(),
			Severity:  msg.GetSeverity(),
			TimeStamp: msg.GetTimestamp(),
			Fields:    flds,
			AppID:     app.GetID(),
		}
		msgmodel.Data = append(msgmodel.Data, nemsg)
	}
	msgmodel.Count = count
	ctx.WriteModel(http.StatusOK, msgmodel)
}
