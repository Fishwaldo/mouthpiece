package app

import (
	"fmt"
	"net/http"
	"context"

	"github.com/Fishwaldo/mouthpiece/internal/restapi/auth"
	"github.com/Fishwaldo/mouthpiece/internal/server"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"

	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/responses"

)

type appListResponse struct {
	ID          int
	Name        string
	Description string
}

type appGetRequest struct {
	ID int `path:"id"`
}

type appFiltersResponse struct {
	ID          int    `doc:"ID of the Application" readOnly:"true"`
	Name        string `doc:"Name of the Filter" readOnly:"true"`
	Description string `doc:"Description of the Filter" readOnly:"true"`
	Type        string `doc:"Type of the Filter" readOnly:"true"`
}

type appGroupResponse struct {
	ID          int    `doc:"ID of the Application" readOnly:"true"`
	Name        string `doc:"Name of the Filter" readOnly:"true"`
	Description string `doc:"Description of the Filter" readOnly:"true"`
}

type appGetResponse struct {
	ID          int                  `doc:"ID of the Application" readOnly:"true" json:"id"`
	Name        string               `doc:"Name of the Application" pattern:"^[a-zA-Z0-9_]+$" minLength:"3" maxLength:"32" json:"name"`
	Description string               `doc:"Name of the Application" pattern:"^[a-zA-Z0-9_]+$" minLength:"0" maxLength:"255" json:"description,omitempty"`
	URL         string               `doc:"URL of the Application" json:"URL,omitempty"`
	Icon        string               `doc:"Icon of the Application" json:"Icon,omitempty"`
	Status      string  			 `enum:"Enabled,Disabled" doc:"Status of the Application" json:"status"`
	Filters     []appFiltersResponse `doc:"Filters of the Application" json:"filters,omitempty" readOnly:"true"`
	Groups      []appGroupResponse   `doc:"Groups of the Application" json:"groups,omitempty" readOnly:"true"`
}

type appPutRequest struct {
	ID   int `doc:"ID of the Application" readOnly:"true" path:"id"`
	Body struct {
		Name        string               `doc:"Name of the Application" pattern:"^[a-zA-Z0-9_]+$" minLength:"3" maxLength:"32" json:"name"`
		Description string               `doc:"Name of the Application" pattern:"^[a-zA-Z0-9_]+$" minLength:"0" maxLength:"255" json:"description,omitempty"`
		URL         string               `doc:"URL of the Application" json:"URL,omitempty"`
		Icon        string               `doc:"Icon of the Application" json:"Icon,omitempty"`
		Status      string  			 `enum:"Enabled,Disabled" doc:"Status of the Application" json:"status"`
	}
}

type appPostRequest struct {
	Body struct {
		Name        string               `doc:"Name of the Application" pattern:"^[a-zA-Z0-9_]+$" minLength:"3" maxLength:"32" json:"name"`
		Description string               `doc:"Name of the Application" pattern:"^[a-zA-Z0-9_]+$" minLength:"0" maxLength:"255" json:"description,omitempty"`
		URL         string               `doc:"URL of the Application" json:"URL,omitempty"`
		Icon        string               `doc:"Icon of the Application" json:"Icon,omitempty"`
		Status      string  			 `enum:"Enabled,Disabled" doc:"Status of the Application" json:"status"`
	}
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

	getapi := appapi.SubResource("/{id}")
	getapi.Get("Get-App", "Get App Details",
		responses.OK().ContentType("application/json"),
		responses.OK().Model(appGetResponse{}),
		responses.NotFound().ContentType("application/json"),
	).Run(appGet)

	getapi.Put("Put-App", "Update a App",
		responses.OK().ContentType("application/json"),
		responses.OK().Model(appGetResponse{}),
		responses.NotAcceptable(),
		responses.NotFound().ContentType("application/json"),
	).Run(appSet)

	newapi := appapi.SubResource("/new")
	newapi.Post("Post-App", "Create a App",
		responses.OK().ContentType("application/json"),
		responses.OK().Model(appGetResponse{}),
		responses.NotAcceptable(),
		responses.Conflict(),
	).Run(appPost)

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
				ID:          app.GetID(),
				Name:        app.GetName(),
				Description: app.GetDescription(),
			})
		}
		ctx.WriteModel(http.StatusOK, appList)
	}
}

func appGet(ctx huma.Context, input appGetRequest) {
	if app, err := server.Get().GetAppService().GetByID(ctx, input.ID); err != nil {
		ctx.WriteError(http.StatusNotFound, "Not Found", err)
	} else {
		response := apptoGetResponse(ctx, app)
		ctx.WriteModel(http.StatusOK, response)
	}

}

func appSet(ctx huma.Context, input appPutRequest) {
	if app, err := server.Get().GetAppService().GetByID(ctx, input.ID); err != nil {
		ctx.WriteError(http.StatusNotFound, "Not Found", err)
	} else {
		if input.Body.Name != app.GetName() {
			if err := app.SetName(ctx, input.Body.Name); err != nil {
				ctx.WriteError(http.StatusNotAcceptable, "Name Not Acceptable", err)
				return
			}
		}
		if input.Body.Description != app.GetDescription() {
			if err := app.SetDescription(ctx, input.Body.Description); err != nil {
				ctx.WriteError(http.StatusNotAcceptable, "Description Not Acceptable", err)
				return
			}
		}
		if input.Body.URL != app.GetURL() {
			if err := app.SetURL(ctx, input.Body.URL); err != nil {
				ctx.WriteError(http.StatusNotAcceptable, "URL Not Acceptable", err)
				return
			}
		}
		if input.Body.Icon != app.GetIcon() {
			if err := app.SetIcon(ctx, input.Body.Icon); err != nil {
				ctx.WriteError(http.StatusNotAcceptable, "Icon Not Acceptable", err)
				return
			}
		}
		if input.Body.Status != app.GetStatus().String() {
			var status interfaces.AppStatus
			status.Scan(input.Body.Status)
			if err := app.SetStatus(ctx, status); err != nil {
				ctx.WriteError(http.StatusNotAcceptable, "Status Not Acceptable", err)
				return
			}
		}
		res := apptoGetResponse(ctx, app)
		ctx.WriteModel(http.StatusOK, res)
	}
}

func appPost(ctx huma.Context, input appPostRequest) {
		if ok, _ := server.Get().GetAppService().Exists(ctx, input.Body.Name); ok {
			ctx.WriteError(http.StatusConflict, "App Already Exists", fmt.Errorf("%s Exists", input.Body.Name))
			return;
		} else {
			if app, err := server.Get().GetAppService().Create(ctx, input.Body.Name, input.Body.Description); err != nil {
				ctx.WriteError(http.StatusNotAcceptable, "Create Failed", err)
				return;
			} else {
				if input.Body.URL != "" {
					if err := app.SetURL(ctx, input.Body.URL); err != nil {
						ctx.WriteError(http.StatusNotAcceptable, "Set URL Failed", err)
						return;
					}
				}
				if input.Body.Icon != "" {
					if err := app.SetIcon(ctx, input.Body.Icon); err != nil {
						ctx.WriteError(http.StatusNotAcceptable, "Set Icon Failed", err)
						return;
					}
				}
				var status interfaces.AppStatus
				status.Scan(input.Body.Status)
				if err := app.SetStatus(ctx, status); err != nil {
					ctx.WriteError(http.StatusNotAcceptable, "Set Status Failed", err)
					return;
				}

				if res, err := server.Get().GetAppService().GetByID(ctx, app.GetID()); err != nil {
					ctx.WriteError(http.StatusInternalServerError, "Not Found", err)
					return;
				} else {
					response := apptoGetResponse(ctx, res)
					ctx.WriteModel(http.StatusOK, response)
				}
			}
		}
}

func apptoGetResponse(ctx context.Context, input interfaces.AppI) appGetResponse {
	response := appGetResponse{
		ID:          input.GetID(),
		Name:        input.GetName(),
		Description: input.GetDescription(),
		URL:         input.GetURL(),
		Icon:        input.GetIcon(),
		Status:      input.GetStatus().String(),
	}
	flts, _ := input.GetFilters(ctx)
	for _, flt := range flts {
		appfilter := appFiltersResponse{
			ID:          flt.GetID(),
			Name:        flt.GetName(),
			Description: flt.GetDescription(),
			Type:        flt.GetType(),
		}
		response.Filters = append(response.Filters, appfilter)
	}
	grps, _ := input.GetGroups(ctx)
	for _, grp := range grps {
		appgroup := appGroupResponse{
			ID:          grp.GetID(),
			Name:        grp.GetName(),
			Description: grp.GetDescription(),
		}
		response.Groups = append(response.Groups, appgroup)
	}
	return response
}