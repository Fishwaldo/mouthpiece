package avatar

import (
	"bytes"
	"image/png"
	"github.com/aofei/cameron"
	"net/http"

	"github.com/Fishwaldo/mouthpiece/internal/server"
//	"github.com/Fishwaldo/mouthpiece/internal/restapi/auth"
	"github.com/danielgtaylor/huma/responses"
	"github.com/danielgtaylor/huma"
)

type avatarGetRequest struct {
	ID int `path:"id"`
}


func Setup(res *huma.Resource) error {
	avtapi := res.SubResource("/avatar")
	//	avtapi.Middleware(auth.RequireAuth)
	avtapi.Tags("Avatar")

	appavtapi := avtapi.SubResource("/app/{id}")
	
	appavtapi.Get("get Icon for App", "Get App Icon",
		responses.OK().ContentType("images/png"),
		responses.NotFound().ContentType("application/json"),
		).Run(appAvt)

	useravtapi := avtapi.SubResource("/user/{id}")
	useravtapi.Get("get Icon for User", "Get User Icon",
		responses.OK().ContentType("images/png"),
		responses.NotFound().ContentType("application/json"),
		).Run(userAvt)

	return nil
}

func appAvt(ctx huma.Context, input avatarGetRequest) {
	buf := bytes.Buffer{}
	app, err := server.Get().GetAppService().GetByID(ctx, input.ID)
	if err != nil {
		ctx.WriteError(http.StatusNotFound, "Internal Error", err)
		return
	}
	png.Encode(&buf, cameron.Identicon([]byte(app.GetName()), 540, 60))
	ctx.Header().Set("Content-Type", "image/png")
	ctx.Write(buf.Bytes())
}

func userAvt(ctx huma.Context, input avatarGetRequest) {
	buf := bytes.Buffer{}
	app, err := server.Get().GetUserService().GetByID(ctx, input.ID)
	if err != nil {
		ctx.WriteError(http.StatusNotFound, "Internal Error", err)
		return
	}
	png.Encode(&buf, cameron.Identicon([]byte(app.GetName()), 540, 60))
	ctx.Header().Set("Content-Type", "image/png")
	ctx.Write(buf.Bytes())
}