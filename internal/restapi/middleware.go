package restapi

import (
	"context"
	"net/http"
	"path"
	"strconv"

	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/log"

	"github.com/danielgtaylor/huma"
	"github.com/go-chi/chi"
	"github.com/go-pkgz/auth/token"
)

// UpdateAuthContext defines interface adding extras or modifying UserInfo in request context
type Middleware struct {
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

type getUserFn func(context.Context, uint) (interfaces.UserI, error)

// Update user info in request context from go-pkgz/auth token.User to mouthpiece.User
func (a *Middleware) Update(getUser getUserFn) func(http.Handler) http.Handler {
	f := func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// call update only if user info exists, otherwise do nothing
			if tknuser, err := token.GetUserInfo(r); err == nil {
				/* find out DB User */
				id, _ := strconv.Atoi(tknuser.ID)
				ctx := huma.ContextFromRequest(w, r)
				if dbUser, err := getUser(ctx, uint(id)); err != nil {
					log.Log.Info("DBUser Not Found", "token", tknuser, "error", err)

					/* do Something */
					ctx.WriteError(http.StatusUnauthorized, "User not found", err)
					return
				} else {
					ok, res, err := AuthService.AuthEnforcer.EnforceEx(dbUser.GetEmail(), r.URL.Path, r.Method)
					log.Log.V(1).Info("Access Control", "result", ok, "Policy", res, "Error", err)
					if !ok {
						huma.ContextFromRequest(w, r).WriteError(http.StatusForbidden, "Access Denied", err)
						return
					}
					r = r.WithContext(context.WithValue(r.Context(), interfaces.CtxUserValue{}, tknuser))
				}
				h.ServeHTTP(w, r)
				return
			} else {
				ctx := huma.ContextFromRequest(w, r)
				ctx.WriteError(http.StatusUnauthorized, "Access Denied")
			}
		}
		return http.HandlerFunc(fn)
	}
	return f
}

// CleanPath middleware will clean out double slash mistakes from a user's request path.
// For example, if a user requests /users//1 or //users////1 will both be treated as: /users/1
func CleanPath(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())

		routePath := rctx.RoutePath
		if routePath == "" {
			if r.URL.RawPath != "" {
				routePath = r.URL.RawPath
			} else {
				routePath = r.URL.Path
			}
			rctx.RoutePath = path.Clean(routePath)
		}

		next.ServeHTTP(w, r)
	})
}
