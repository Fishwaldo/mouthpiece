package restapi

import (
	"net/http"

	mouthpiece "github.com/Fishwaldo/mouthpiece/pkg"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"

	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/responses"
)

func setupApps(res *huma.Resource, mps *mouthpiece.MouthPiece) error {
	//AuthService.AddResourceURL("/v1/apps/", "apigroup:apps")
	appapi := res.SubResource("/apps")
	appapi.Tags("Apps")

	appapi.Get("get-apps", "Get A List of Applications",
		responses.OK().ContentType("application/json"),
		responses.OK().Headers("Set-Cookie"),
		responses.OK().Model([]interfaces.AppDetails{}),
	).Run(func(ctx huma.Context) {
		if apps, err := mps.GetAppService().GetApps(ctx); err != nil {
			ctx.WriteError(http.StatusInternalServerError, "Internal Error", err)
		} else {
			ctx.WriteModel(http.StatusOK, apps)
		}
	})

	appapi.Put("create-app", "Create a Application",
		responses.OK().ContentType("application/json"),
		responses.OK().Headers("Set-Cookie"),
		responses.OK().Model(interfaces.AppDetails{}),
		responses.NotAcceptable().ContentType("application/json"),
		responses.NotAcceptable().Headers("Set-Cookie"),
	).Run(func(ctx huma.Context, input struct {
		Body interfaces.AppDetails
	}) {
		if app, err := mps.GetAppService().CreateApp(ctx, input.Body); err != nil {
			checkErrors(ctx, err)
			ctx.WriteError(http.StatusNotAcceptable, "Create Failed")
		} else {
			ctx.WriteModel(http.StatusOK, app.GetDetails())
		}
	})

	app := appapi.SubResource("/{appid}")
	app.Get("get-app", "Get Application Details",
		responses.OK().ContentType("application/json"),
		responses.OK().Headers("Set-Cookie"),
		responses.OK().Model(interfaces.AppDetails{}),
		responses.NotFound(),
		responses.NotAcceptable().Headers("Set-Cookie"),
	).Run(func(ctx huma.Context, input struct {
		Appid uint `path:"appid"`
	}) {
		if app, err := mps.GetAppService().GetApp(ctx, input.Appid); err != nil {
			checkErrors(ctx, err)
			ctx.WriteHeader(http.StatusNotFound)
		} else {
			ctx.WriteModel(http.StatusOK, app.GetDetails())
		}
	})

	return nil
}
