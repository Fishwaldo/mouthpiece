package messages

import (
	"net/http"
	"time"

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


type msgListResponse struct {
	Data []msgResponse `json:"data"`
	Count int 		   `json:"count"`
}

type msgGet struct {
	ID	string `path:"msgid"`
}

type msgResponse struct {
	ID        uuid.UUID 		  
	Message   string
	ShortMsg  *string
	Topic     *string
	Severity  int
	TimeStamp time.Time
	Fields    map[string]string
	AppID     int
}



func Setup(res *huma.Resource) error {
	// //AuthService.AddResourceURL("/v1/apps/", "apigroup:apps")
	msglistapi := res.SubResource("/messages")
	msglistapi.Middleware(auth.RequireAuth)
	msglistapi.Tags("Messages")

	msglistapi.Get("get-messages", "Get A List of Messages",
		responses.OK().ContentType("application/json"),
		responses.OK().Model(msgListResponse{}),
	).Run(msgList)

	msggetapi := msglistapi.SubResource("/{msgid}")
	msggetapi.Get("get-message", "Get a Single Message",
		responses.OK().ContentType("application/json"),
		responses.OK().Model(msgResponse{}),
		responses.NotFound().ContentType("application/json"),
	).Run(getMsg)
	

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
		nemsg := msgResponse{
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

func getMsg(ctx huma.Context, input msgGet) {
	//fmt.Printf("Page: %d, Size: %d Sort: %s %s\n", input.Page, input.Size, input.OrderBy, input.OrderDir)

	uuid, err := uuid.Parse(input.ID)
	if err != nil {
		ctx.WriteError(http.StatusBadRequest, "Invalid UUID")
		return
	}
	msg, err := server.Get().GetMsgService().Get(ctx, uuid)
	if err != nil {
		ctx.WriteError(http.StatusNotFound, "Error getting messages")
		return
	}
	flds, _ := msg.GetFields(ctx)
	app, _ := msg.GetApp(ctx)
	nemsg := msgResponse{
		ID:        msg.GetID(),
		Message:   msg.GetMessage(),
		ShortMsg:  msg.GetShortMsg(),
		Topic:     msg.GetTopic(),
		Severity:  msg.GetSeverity(),
		TimeStamp: msg.GetTimestamp(),
		Fields:    flds,
		AppID:     app.GetID(),
	}
	ctx.WriteModel(http.StatusOK, nemsg)
}