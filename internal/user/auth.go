package user

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/Fishwaldo/mouthpiece/internal/errors"
	. "github.com/Fishwaldo/mouthpiece/internal/log"
	"github.com/go-pkgz/auth/token"
)

type CtxUserValue struct{}


func dbAuthProvider(user, pass string) (ok bool, err error) {
	user = strings.TrimSpace(user)
	Log.Info("Direct Login", "user", user)
	dbUser, err := GetUser(context.Background(), user)
	Log.Info("User", "user", dbUser, "error", err)

	if err == mperror.ErrUserNotFound {
		Log.Info("User not found", "user", user)
		return false, nil
	}
	if !dbUser.CheckPassword(context.Background(), pass) {
		Log.Info("Password Invalid", "user", user)
		return false, nil
	}
	return true, nil
}

// Called when the Tokens are created/refreshed.
func MapClaimsToUser(claims token.Claims) token.Claims {
	//Log.Info("Map Claims To User", "claims", claims)
	//	if claims.User != nil {
	//		if user, err := GetUser(claims.User.Name); err != nil {
	//			Log.Info("User not found", "user", claims.User.Name)
	//			claims.User.SetBoolAttr("valid", false)
	//		} else {
	//			claims.User.SetStrAttr("backenduser", user.Username)
	//			claims.User.SetBoolAttr("valid", true)
	//		}
	//	}
	return claims
}

// called on every access to the API
func UserValidator(token string, claims token.Claims) bool {
	//Log.Info("User Validator", "user", claims.User.Name)
	if claims.User != nil {
		if user, _ := GetUser(context.Background(), claims.User.Name); user != nil {
			claims.User.ID = fmt.Sprintf("%d", user.ID)
			return true
		}
	}	
	return false
}

func GetUserFromContext(ctx context.Context) (bool, *User) {
	v := ctx.Value(CtxUserValue{}).(token.User)
	if id, _ := strconv.Atoi(v.ID); id > 0 {
		if user, _ := GetUserByID(ctx, uint(id)); user != nil {
			return true, user
		}
	}
	return false, nil
}