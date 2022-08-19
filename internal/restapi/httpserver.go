package restapi

import (
	"context"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/Fishwaldo/mouthpiece/frontend"
	mpserver "github.com/Fishwaldo/mouthpiece/internal"
	mouthpiece "github.com/Fishwaldo/mouthpiece/pkg"

	"github.com/Fishwaldo/mouthpiece/pkg/log"

	"github.com/danielgtaylor/huma"
	//"github.com/danielgtaylor/huma/middleware"
	hmw "github.com/danielgtaylor/huma/middleware"

	//	"github.com/danielgtaylor/huma/responses"
	"github.com/go-chi/chi"
	//"github.com/go-chi/chi/middleware"

	"github.com/go-pkgz/rest"

	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("frontend.path", "frontend/dist")
	viper.SetDefault("frontend.external", false)
}

type RestAPI struct {
	huma *huma.Router
	mux  *chi.Mux
}

func NewRestAPI(mps *mouthpiece.MouthPiece) *RestAPI {
	bi := mouthpiece.GetVersionInfo()

	restapi := &RestAPI{}
	restapi.huma = huma.New(bi.Name, bi.GitVersion)
	huma.AddAllowedHeaders("App-Name", "Author", "App-Version", "X-Request-Id", "Set-Cookie")
	restapi.huma.DisableSchemaProperty()
	restapi.mux = chi.NewRouter()

	hmw.NewLogger = mpserver.GetHumaLogger

	restapi.mux.Use(hmw.DefaultChain)
	restapi.mux.Use(rest.AppInfo(bi.Name, "Fishwaldo", bi.GitVersion))
	restapi.mux.Use(rest.Ping)
	restapi.mux.Use(rest.RealIP)
	restapi.mux.Use(rest.Trace)
	restapi.mux.Use(CleanPath)

	if viper.GetBool("debug") {
		log.Log.Info("Enabling Debug Endpoints")
		bench := rest.NewBenchmarks()
		bench.WithTimeRange(time.Hour)
		restapi.mux.Use(bench.Handler)
		restapi.mux.Use(rest.Metrics("10.0.0.0/8"))
		restapi.mux.Mount("/debug", rest.Profiler("10.0.0.0/8"))
		restapi.mux.Get("/bench", func(w http.ResponseWriter, r *http.Request) {
			resp := struct {
				OneMin     rest.BenchmarkStats `json:"1min"`
				FiveMin    rest.BenchmarkStats `json:"5min"`
				FifteenMin rest.BenchmarkStats `json:"15min"`
				Hourly	   rest.BenchmarkStats `json:"hourly"`
			}{
				bench.Stats(time.Minute),
				bench.Stats(time.Minute * 5),
				bench.Stats(time.Minute * 15),
				bench.Stats(time.Hour),
			}
			rest.RenderJSON(w, resp)
		})
	}

	restapi.fileServer("/static", getFrontendFiles())
	restapi.mux.Handle("/", http.RedirectHandler("/static/", http.StatusMovedPermanently))

	restapi.mux.Handle("/*", restapi.huma)
	setupHealth(restapi.huma)
	v1api := restapi.huma.Resource("/v1")
	//mw := NewMiddleware()
	//v1api.Middleware(mw.Update(mps.GetUserService().GetUser))

	setupApps(v1api, mps)
	return restapi
}

func (restapi *RestAPI) Start() {
	// The HTTP Server
	server := &http.Server{Addr: "0.0.0.0:8080", Handler: restapi.mux}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		log.Log.Info("Server Shutting Down")

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Log.Error(shutdownCtx.Err(), "graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Log.Error(err, "Fatal Shutdown")
		}
		//cancel()
		serverStopCtx()
	}()

	// Run the server
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Log.Error(err, "Server Shutdown Error")
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}

// FileServer conveniently sets up a http.FileServer handler to serve static files from a http.FileSystem.
// Borrowed from https://github.com/go-chi/chi/blob/master/_examples/fileserver/main.go
func (restapi *RestAPI) fileServer(path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	//log.Log.Printf("[INFO] serving static files from %v", root)
	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		restapi.mux.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	restapi.mux.Get(path, func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
}

func getFrontendFiles() http.FileSystem {
	if viper.GetBool("frontend.external") {
		log.Log.Info("Serving frontend from external location", "path", viper.GetString("frontend.path"))
		return http.Dir(viper.GetString("frontend.path"))
	} else {
		log.Log.Info("Serving frontend from Bundled Files")
		subdir, err := fs.Sub(frontend.FrontEndFiles, "dist")
		if err != nil {
			log.Log.Error(err, "Failed to get subdir")
		}
		return http.FS(subdir)
	}

}
