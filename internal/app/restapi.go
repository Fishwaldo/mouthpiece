package app

import (
	"net/http"

	"github.com/Fishwaldo/mouthpiece/internal/auth"

	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/responses"
)

func InitializeAppRestAPI(res *huma.Resource) error {
	auth.AuthService.AddResourceURL("/v1/apps/", "apigroup:apps")
	appapi := res.SubResource("/apps/")

	appapi.Get("get-apps", "Get A List of Applications",
		responses.OK().ContentType("application/json"),
		responses.OK().Headers("Set-Cookie"),
		responses.OK().Model([]App{}),
	).Run(func(ctx huma.Context) {
		ctx.WriteModel(http.StatusOK, GetApps(ctx))
	})


	appapi.Put("create-app", "Create a Application",
		responses.OK().ContentType("application/json"),
		responses.OK().Headers("Set-Cookie"),
		responses.OK().Model(&App{}),
		responses.NotAcceptable().ContentType("application/json"),
		responses.NotAcceptable().Headers("Set-Cookie"),
	).Run(func(ctx huma.Context, input struct {
		Body AppDetails
	}) {
		if app, err := CreateApp(ctx, input.Body); err != nil {
			ctx.WriteError(http.StatusNotAcceptable, "Database Error", err)
		} else {
			ctx.WriteModel(http.StatusOK, app)
		}
	})

	return nil
}
