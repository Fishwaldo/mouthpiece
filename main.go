/*
MIT License

Copyright (c) 2021 Justin Hammond

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	//"fmt"
	//	"context"
	"fmt"
	"io/fs"
	"net/http"

	//	"reflect"
	"strings"
	//	"unsafe"
	"embed"
	"encoding/json"
	"os"
	"runtime/debug"

	"github.com/Fishwaldo/mouthpiece/frontend"
	_ "github.com/Fishwaldo/mouthpiece/frontend"
	mouthpiece "github.com/Fishwaldo/mouthpiece/internal"
	"github.com/Fishwaldo/mouthpiece/internal/app"
	"github.com/Fishwaldo/mouthpiece/internal/auth"
	"github.com/Fishwaldo/mouthpiece/internal/db"
	"github.com/Fishwaldo/mouthpiece/internal/filter"
	. "github.com/Fishwaldo/mouthpiece/internal/log"
	msg "github.com/Fishwaldo/mouthpiece/internal/message"
	"github.com/Fishwaldo/mouthpiece/internal/middleware"
	"github.com/Fishwaldo/mouthpiece/internal/transport"
	"github.com/Fishwaldo/mouthpiece/internal/user"

	healthChecker "github.com/Fishwaldo/mouthpiece/internal/health"
	_ "github.com/Fishwaldo/mouthpiece/internal/transport/stdout"
	_ "github.com/Fishwaldo/mouthpiece/internal/transport/telegram"

	"github.com/alexliesenfeld/health"

	"github.com/go-chi/chi"

	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/cli"
	"github.com/danielgtaylor/huma/responses"
	"github.com/spf13/viper"
)

//go:embed config
var ConfigFiles embed.FS

func init() {
	viper.SetDefault("frontend.path", "frontend/dist")
	viper.SetDefault("frontend.external", false)
}

// FileServer conveniently sets up a http.FileServer handler to serve static files from a http.FileSystem.
// Borrowed from https://github.com/go-chi/chi/blob/master/_examples/fileserver/main.go
func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	//log.Printf("[INFO] serving static files from %v", root)
	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
}

func printBuildInfo() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("Getting build info failed (not in module mode?)!")
		return
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(bi); err != nil {
		panic(err)
	}
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/mouthpiece")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("$HOME/.mouthpiece")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			//viper.SetConfigType("yaml")
			fmt.Printf("No Config File Found. Writing Defaults\n")
			if err := viper.WriteConfigAs("config.yaml"); err != nil {
				panic(fmt.Errorf("fatal writting config file: %w \n", err))
			}
		} else {
			panic(fmt.Errorf("fatal error config file: %w \n", err))
		}
	}

	bi := mouthpiece.GetVersionInfo()
	bi.Name = "MouthPiece"
	bi.Description = "Messaging Server"
	if bi.CheckFontName("starwars") {
		bi.FontName = "starwars"
	}

	fmt.Println(bi.String())

	// Create a new router & CLI with default middleware.
	InitLogger()
	db.InitializeDB()
	humucli := cli.NewRouter(bi.Name, bi.GitVersion)
	humucli.DisableSchemaProperty()
	humucli.PreStart(transport.InitializeTransports)
	humucli.PreStart(msg.InitializeMessage)
	humucli.PreStart(user.InitializeUsers)
	humucli.PreStart(app.InitializeApps)
	humucli.PreStart(transport.StartTransports)
	humucli.PreStart(filter.InitFilter)
	//	app.PreStart()
	humucli.PreStart(healthChecker.StartHealth)
	humucli.GatewayClientCredentials("mouthpiece", "/oauth2/token", nil)
	humucli.GatewayAuthCode("mouthpiece2", "/oauth2/token", "/oauth2/token", nil)
	humucli.GatewayBasicAuth("basic")

	user.AuthConfig.Host = fmt.Sprintf("http://arm64-1.dmz.dynam.ac:%v", viper.Get("Port"))
	user.AuthConfig.ConfigDir = ConfigFiles
	auth.InitAuth(user.AuthConfig)
	m := auth.AuthService.Service.Middleware()
	p := middleware.Middleware{}

	authRoutes, avaRoutes := auth.AuthService.Service.Handlers()
	mux := humucli.Resource("/").GetMux()
	mux.Mount("/auth", authRoutes)
	mux.Mount("/avatar", avaRoutes)

	var httpfiles http.FileSystem
	if viper.GetBool("frontend.external") {
		Log.Info("Serving frontend from external location", "path", viper.GetString("frontend.path"))
		httpfiles = http.Dir(viper.GetString("frontend.path"))
	} else {
		Log.Info("Serving frontend from Bundled Files")
		subdir, err := fs.Sub(frontend.FrontEndFiles, "dist")
		if err != nil {
			Log.Error(err, "Failed to get subdir")
		}
		httpfiles = http.FS(subdir)
	}
	fileServer(mux, "/static", httpfiles)

	// Declare the root resource and a GET operation on it.
	humucli.Resource("/health").Get("get-health", "Get Health of the Service",
		responses.OK().ContentType("application/json"),
		responses.OK().Headers("Content-Type"),
		responses.OK().Model(health.CheckerResult{}),
		responses.InternalServerError().ContentType("application/json"),
		responses.InternalServerError().Headers("Content-Type"),
		responses.InternalServerError().Model(health.CheckerResult{}),
	).Run(func(ctx huma.Context) {
		test := healthChecker.HealthChecker.Check(ctx)
		status := http.StatusOK
		if test.Status != health.StatusUp {
			status = http.StatusInternalServerError
		}
		ctx.WriteModel(status, test)
	})

	humucli.Resource("/config/frontend").Get("get-config", "Get Config of the Service",
		responses.OK().ContentType("application/json"),
		responses.OK().Headers("Content-Type"),
		responses.OK().Model(&mouthpiece.FEConfig{}),
	).Run(func(ctx huma.Context) {
		ctx.WriteModel(http.StatusOK, mouthpiece.GetFEConfig())
	})

	v1api := humucli.Resource("/v1")
	v1api.Middleware(m.Trace)
	v1api.Middleware(p.Update())

	auth.AuthService.AddResourceURL("/v1/message/{application}", "apigroup:message")
	v1api.SubResource("/message/{application}").Post("post-message", "Post Message to the Service",
		responses.OK().ContentType("application/json"),
		responses.OK().Model(&msg.MessageResult{}),
		responses.NotFound().ContentType("application/json"),
	).Run(func(ctx huma.Context, input msg.Message) {
		Log.Info("Recieved Message", "message", input)
		if app.AppExists(input.AppName) {
			if err := input.ProcessMessage(); err == nil {
				mouthpiece.RouteMessage(&input)
				ctx.WriteModel(http.StatusOK, input.Result)
			} else {
				ctx.WriteError(http.StatusInternalServerError, err.Error())
			}
		} else {
			ctx.WriteError(http.StatusNotFound, "Application Not Found")
		}
	})

	auth.AuthService.AddResourceURL("/v1/apps/", "apigroup:apps")
	appapi := v1api.SubResource("/apps/")
	appapi.Get("get-apps", "Get A List of Applications",
		responses.OK().ContentType("application/json"),
		responses.OK().Headers("Set-Cookie"),
		responses.OK().Model([]app.App{}),
	).Run(func(ctx huma.Context) {
		ctx.WriteModel(http.StatusOK, app.GetApps())
	})
	appapi.Put("create-app", "Create a Application",
		responses.OK().ContentType("application/json"),
		responses.OK().Headers("Set-Cookie"),
		responses.OK().Model(&app.App{}),
		responses.NotAcceptable().ContentType("application/json"),
		responses.NotAcceptable().Headers("Set-Cookie"),
	).Run(func(ctx huma.Context, input struct {
		Body app.AppDetails
	}) {
		if app, err := app.CreateApp(input.Body); err != nil {
			ctx.WriteError(http.StatusNotAcceptable, "Database Error", err)
		} else {
			ctx.WriteModel(http.StatusOK, app)
		}
	})

	auth.AuthService.AddResourceURL("/v1/users/", "apigroup:users")
	userapi := v1api.SubResource("/users/")
	userapi.Get("get-users", "Get A List of Users",
		responses.OK().ContentType("application/json"),
		responses.OK().Headers("Set-Cookie"),
		responses.OK().Model([]user.User{}),
	).Run(func(ctx huma.Context) {
		ctx.WriteModel(http.StatusOK, user.GetUsers())
	})

	auth.AuthService.AddResourceURL("/v1/users/{userid}/transports/", "apigroup:users")
	usertransports := v1api.SubResource("/users/{userid}/transports/")
	usertransports.Get("get-user-transports", "Get A List of Transports for a User",
		responses.OK().ContentType("application/json"),
		responses.OK().Headers("Set-Cookie"),
		responses.OK().Model([]string{}),
		responses.NotFound().ContentType("application/json"),
	).Run(func(ctx huma.Context, input struct {
		User uint `path:"userid"`
	}) {
		if user, err := user.GetUserByID(input.User); err != nil {
			ctx.WriteError(http.StatusNotFound, "User Not Found", err)
		} else {
			var transport []string
			for _, t := range user.TransportConfigs {
				transport = append(transport, t.Transport)
			}
			ctx.WriteModel(http.StatusOK, transport)
		}
	})
	auth.AuthService.AddResourceURL("/v1/users/{userid}/transports/{transportid}/", "apigroup:users")
	usertransportdetails := v1api.SubResource("/users/{userid}/transports/{transportid}/")
	usertransportdetails.Get("get-user-transport-details", "Get Details for a Transport for a User",
		responses.OK().ContentType("application/json"),
		responses.OK().Headers("Set-Cookie"),
		responses.OK().Model(transport.TransportConfig{}),
		responses.NotFound().ContentType("application/json"),
	).Run(func(ctx huma.Context, input struct {
		User      uint   `path:"userid"`
		Transport string `path:"transportid"`
	}) {
		if user, err := user.GetUserByID(input.User); err != nil {
			ctx.WriteError(http.StatusNotFound, "User Not Found", err)
		} else {
			ok := false
			for _, t := range user.TransportConfigs {
				if t.Transport == input.Transport {
					ctx.WriteModel(http.StatusOK, t)
					ok = true
				}
			}
			if !ok {
				ctx.WriteError(http.StatusNotFound, "Transport Not Found")
			}
		}
	})
	auth.AuthService.AddResourceURL("/v1/transports/", "apigroup:transports")
	transportapi := v1api.SubResource("/transports/")
	transportapi.Get("get-transports", "Get A List of Transports",
		responses.OK().ContentType("application/json"),
		responses.OK().Headers("Set-Cookie"),
		responses.OK().Model([]string{}),
	).Run(func(ctx huma.Context) {
		ctx.WriteModel(http.StatusOK, transport.GetTransports())
	})

	// Run the CLI. When passed no arguments, it starts the server.
	humucli.Run()
}
