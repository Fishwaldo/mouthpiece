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
	"github.com/Fishwaldo/mouthpiece/internal"
	mouthpiece "github.com/Fishwaldo/mouthpiece/pkg"


	"github.com/Fishwaldo/mouthpiece/pkg/log"

	"github.com/danielgtaylor/huma"
	//"github.com/danielgtaylor/huma/middleware"
	hmw "github.com/danielgtaylor/huma/middleware"

	//	"github.com/danielgtaylor/huma/responses"
	"github.com/go-chi/chi"
	//	"github.com/go-chi/chi/middleware"

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
	restapi := &RestAPI{}
	restapi.huma = huma.New("test", "test")

	restapi.huma.DisableSchemaProperty()
	restapi.mux = chi.NewRouter()

	hmw.NewLogger = mpserver.GetHumaLogger

	restapi.mux.Use(hmw.DefaultChain)

	restapi.fileServer("/static", getFrontendFiles())

	restapi.mux.Handle("/*", restapi.huma)
	setupHealth(restapi.huma)
	v1api := restapi.huma.Resource("/v1")
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
