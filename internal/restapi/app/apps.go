package app

import (
	"net/http"

	"github.com/Fishwaldo/mouthpiece/internal/server"
	"github.com/Fishwaldo/mouthpiece/internal/restapi/auth"

	"github.com/danielgtaylor/huma/responses"
	"github.com/danielgtaylor/huma"
)

type appListResponse struct {
	ID	int
	Name string
	Description string
}



func Setup(res *huma.Resource) error {
	// //AuthService.AddResourceURL("/v1/apps/", "apigroup:apps")
	appapi := res.SubResource("/apps")
	appapi.Middleware(auth.RequireAuth)
	appapi.Tags("Apps")

	appapi.Get("get-apps", "Get A List of Applications",
		responses.OK().ContentType("application/json"),
		responses.OK().Model([]appListResponse{}),
		).Run(appList)

	// appapi.Put("create-app", "Create a Application",
	// 	responses.OK().ContentType("application/json"),
	// 	responses.OK().Headers("Set-Cookie"),
	// 	responses.OK().Model(interfaces.AppDetails{}),
	// 	responses.NotAcceptable().ContentType("application/json"),
	// 	responses.NotAcceptable().Headers("Set-Cookie"),
	// ).Run(func(ctx huma.Context, input struct {
	// 	Body interfaces.AppDetails
	// }) {
	// 	if app, err := mps.GetAppService().CreateApp(ctx, input.Body); err != nil {
	// 		checkErrors(ctx, err)
	// 		ctx.WriteError(http.StatusNotAcceptable, "Create Failed")
	// 	} else {
	// 		ctx.WriteModel(http.StatusOK, app.GetDetails())
	// 	}
	// })

	// app := appapi.SubResource("/{appid}")
	// app.Get("get-app", "Get Application Details",
	// 	responses.OK().ContentType("application/json"),
	// 	responses.OK().Headers("Set-Cookie"),
	// 	responses.OK().Model(interfaces.AppDetails{}),
	// 	responses.NotFound(),
	// 	responses.NotAcceptable().Headers("Set-Cookie"),
	// ).Run(func(ctx huma.Context, input struct {
	// 	Appid uint `path:"appid"`
	// }) {
	// 	if app, err := mps.GetAppService().GetApp(ctx, input.Appid); err != nil {
	// 		checkErrors(ctx, err)
	// 		ctx.WriteHeader(http.StatusNotFound)
	// 	} else {
	// 		ctx.WriteModel(http.StatusOK, app.GetDetails())
	// 	}
	// })

	return nil
}

func appList(ctx huma.Context) {
	if apps, err := server.Get().GetAppService().GetAll(ctx); err != nil {
		ctx.WriteError(http.StatusInternalServerError, "Internal Error", err)
	} else {
		var appList []appListResponse
		for _, app := range apps {
			appList = append(appList, appListResponse{
				ID: app.GetID(),
				Name: app.GetName(),
				Description: app.GetDescription(),
			})
		}
		ctx.WriteModel(http.StatusOK, appList)
	}
}