package user

import (
	"strings"

	"github.com/Fishwaldo/mouthpiece/internal/errors"
	. "github.com/Fishwaldo/mouthpiece/internal/log"
	"github.com/go-pkgz/auth/token"

)

func dbAuthProvider(user, pass string) (ok bool, err error) {
	user = strings.TrimSpace(user)
	Log.Info("Direct Login", "user", user, "pass", pass)
	dbUser, err := GetUser(user)
	Log.Info("User", "user", dbUser, "error", err)

	if err == mperror.ErrUserNotFound {
		Log.Info("User not found", "user", user)
		return false, nil
	}
	if !dbUser.CheckPassword(pass) {
		Log.Info("Password Invalid", "user", user)
		return false, nil
	}
	return true, nil
}

// Called when the Tokens are created/refreshed. 
func MapClaimsToUser(claims token.Claims) token.Claims {
	Log.Info("Map Claims To User", "claims", claims)
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
func UserValidator(token string, claims token.Claims) (bool) {
	Log.Info("User Validator", "token", token, "claims", claims)
	if claims.User != nil {
		return true
	}
	return false
}