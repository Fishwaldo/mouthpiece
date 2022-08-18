package middleware

import (
	"context"
	"strconv"
	"net/http"

	"github.com/Fishwaldo/mouthpiece/internal/auth"
	. "github.com/Fishwaldo/mouthpiece/internal/log"
	"github.com/Fishwaldo/mouthpiece/internal/user"

	"github.com/danielgtaylor/huma"
	"github.com/go-pkgz/auth/token"
)

// UpdateAuthContext defines interface adding extras or modifying UserInfo in request context
type Middleware struct {
}



// Update user info in request context from go-pkgz/auth token.User to mouthpiece.User
func (a *Middleware) Update() func(http.Handler) http.Handler {
	f := func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// call update only if user info exists, otherwise do nothing
			if tknuser, err := token.GetUserInfo(r); err == nil {
				/* find out DB User */
				id, _ := strconv.Atoi(tknuser.ID)
				ctx := huma.ContextFromRequest(w, r)
				if dbUser, err := user.GetUserByID(ctx, uint(id)); err != nil {
					Log.Info("DBUser Not Found", "token", tknuser, "error", err)
					
					/* do Something */
					ctx.WriteError(http.StatusUnauthorized, "User not found", err)
					return
				} else {
					ok, res, err := auth.AuthService.AuthEnforcer.EnforceEx(dbUser.Email, r.URL.Path, r.Method)
					Log.V(1).Info("Access Control", "result", ok, "Policy", res, "Error", err)
					if !ok {
						huma.ContextFromRequest(w, r).WriteError(http.StatusForbidden, "Access Denied", err)
						return
					}
					r = r.WithContext(context.WithValue(r.Context(), user.CtxUserValue{}, tknuser))
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

