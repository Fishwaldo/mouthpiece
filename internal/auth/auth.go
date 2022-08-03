package auth

import (
	"time"
//	"strings"
	"os"
	"fmt"
	"net/http"
	"context"
	"crypto/sha1"

	. "github.com/Fishwaldo/mouthpiece/internal/log"
//	"github.com/Fishwaldo/mouthpiece/internal"
	"github.com/Fishwaldo/mouthpiece/internal/user"
	"github.com/Fishwaldo/mouthpiece/internal/auth/telegram"

	"github.com/spf13/viper"

	pkauth "github.com/go-pkgz/auth"
	"github.com/go-pkgz/auth/avatar"
	//"github.com/go-pkgz/auth/middleware"
	"github.com/go-pkgz/auth/provider"
	"github.com/go-pkgz/auth/token"
	//"github.com/go-pkgz/auth/logger"

	"golang.org/x/oauth2/github"
)

type Auth struct {
	Service *pkauth.Service
}

var AuthService *Auth

type AuthLogger struct {
}

var AL *AuthLogger

func (AL AuthLogger) Logf(format string, args ...interface{}) {
	Log.WithName("Auth").Info("Authentication", "message", fmt.Sprintf(format, args...))
}

func init() {
	viper.SetDefault("auth.github.enabled", false)
	viper.SetDefault("auth.github.signups", true)
	viper.SetDefault("auth.github.fields.UserName", "login")
	viper.SetDefault("auth.github.fields.Email", "email")
	viper.SetDefault("auth.github.fields.Avatar", "avatar_url")
	viper.SetDefault("auth.github.fields.FullName", "name")
	viper.SetDefault("auth.dev.enabled", false)
	viper.SetDefault("auth.email.enabled", false)
}

func customGitHubProvider() (cred pkauth.Client, ch provider.CustomHandlerOpt) {
	ch = provider.CustomHandlerOpt {
		Endpoint: github.Endpoint,
		InfoURL: "https://api.github.com/user",
		MapUserFn: func(data provider.UserData, _ []byte) token.User {
			userInfo := token.User{
				ID: "github_" + token.HashID(sha1.New(), data.Value(viper.GetString("auth.github.fields.UserName"))),
				Name: data.Value(viper.GetString("auth.github.fields.UserName")),
				Email: data.Value(viper.GetString("auth.github.fields.Email")),
				Picture: data.Value(viper.GetString("auth.github.fields.Avatar")),
			}
			userInfo.SetStrAttr("fullname", data.Value(viper.GetString("auth.github.fields.FullName")))
			userInfo.SetStrAttr("source", "github")
			return userInfo
		},
		Scopes: []string{},
	}
	cred = pkauth.Client{
		Cid: viper.GetString("auth.github.client_id"),
		Csecret: viper.GetString("auth.github.client_secret"),
	}

	return cred, ch
}

func mapClaimsToUser(claims token.Claims) token.Claims {
	if claims.User != nil { 
		if user, err := user.GetUser(claims.User.Name); err != nil {
			Log.Info("User not found", "user", claims.User.Name)
			claims.User.SetBoolAttr("valid", false)
		} else {
			claims.User.SetStrAttr("backenduser", user.Username)
			claims.User.SetBoolAttr("valid", true)
		}
	}
	fmt.Printf("Update Claims %+v\n", claims)
	return claims
}



func InitAuth(host string) {
	AL = &AuthLogger{}
	AuthService = &Auth{}

	options := pkauth.Opts{
		SecretReader: token.SecretFunc(func(_ string) (string, error) { // secret key for JWT, ignores aud
			return "secret", nil
		}),
		TokenDuration:     time.Minute,                                 // short token, refreshed automatically
		CookieDuration:    time.Hour * 24,                              // cookie fine to keep for long time
		DisableXSRF:       true,                                        // don't disable XSRF in real-life applications!
		Issuer:            "my-demo-service",                           // part of token, just informational
		URL:               host,                     					// base url of the protected service
		AdminPasswd:       "password",												  // admin password
		AvatarStore:       avatar.NewLocalFS("/tmp/demo-auth-service"), // stores avatars locally
		AvatarResizeLimit: 200,                                         // resizes avatars to 200x200
		ClaimsUpd: token.ClaimsUpdFunc(mapClaimsToUser),
//		Validator: token.ValidatorFunc(func(_ string, claims token.Claims) bool { // rejects some tokens
//			if claims.User != nil {
//				if strings.HasPrefix(claims.User.ID, "github_") { // allow all users with github auth
//					return true
//				}
//				if strings.HasPrefix(claims.User.ID, "microsoft_") { // allow all users with ms auth
//					return true
//				}
//				if strings.HasPrefix(claims.User.ID, "patreon_") { // allow all users with ms auth
//					return true
//				}
//				if strings.HasPrefix(claims.User.Name, "dev_") { // non-guthub allow only dev_* names
//					return true
//				}
//				return strings.HasPrefix(claims.User.Name, "custom123_")
//			}
//			return false
//		}),
		Logger:      AL, 			// optional logger for auth library
		UseGravatar: true,          // for verified provider use gravatar service
	}

	// create auth service
	AuthService.Service = pkauth.NewService(options)
	if viper.GetBool("auth.dev.enabled") {
		Log.Info("Auth Dev Mode Enabled!")
		AuthService.Service.AddProvider("dev", "", "")
	}
	if viper.GetBool("auth.github.enabled") {
		if !viper.IsSet("auth.github.client_id") {
			Log.Error(nil, "Github auth is enabled but client_id is not set")
		} else {
			if !viper.IsSet("auth.github.client_secret") {
				Log.Error(nil, "Github auth is enabled but client_secret is not set")
			} else { 					
				Log.Info("Auth Github Enabled!")
				gcred, gch := customGitHubProvider()
				AuthService.Service.AddCustomProvider("github", gcred, gch)
			}
		}
	}
	AuthService.Service.AddProvider("microsoft", os.Getenv("AEXMPL_MS_APIKEY"), os.Getenv("AEXMPL_MS_APISEC"))


	if viper.GetBool("auth.email.enabled") {
		Log.Info("Auth Email Enabled!")
		AuthService.Service.AddVerifProvider("email",
			"To confirm use {{.Token}}\nor follow http://arm64-1.dmz.dynam.ac:8888/auth/email/login?token={{.Token}}",
			provider.SenderFunc(func(address string, text string) error { // sender just prints token
				fmt.Printf("CONFIRMATION for %s\n%s\n", address, text)
				return nil
			}),
		)
	}

	if viper.GetBool("auth.telegram.enabled") {
		if viper.IsSet("auth.telegram.token") {
			Log.Info("Auth Telegram Enabled!")
			// add telegram provider
			telegram := telegramauth.TelegramHandler{
				ProviderName: "telegram",
				ErrorMsg:     "❌ Invalid auth request. Please try clicking link again.",
				SuccessMsg:   "✅ You have successfully authenticated!",
				Telegram:     telegramauth.NewTelegramAPI(viper.GetString("auth.telegram.token"), http.DefaultClient),
				L:            AL,
				TokenService: AuthService.Service.TokenService(),
				AvatarSaver:  AuthService.Service.AvatarProxy(),
			}

			go func() {
				err := telegram.Run(context.Background())
				if err != nil {
					Log.Error(err, "[PANIC] failed to start telegram")
				}
			}()

			AuthService.Service.AddCustomHandler(&telegram)
		} else {
			Log.Error(nil, "Telegram auth is enabled but token is not set")
		}
	}

	// run dev/test oauth2 server on :8084
	go func() {
		devAuthServer, err := AuthService.Service.DevAuth() // peak dev oauth2 server
		devAuthServer.GetEmailFn = func(username string) string {
			return username + "@dynam.com"
		}
		if err != nil {
			Log.Error(err, "[PANIC] failed to start dev oauth2 server")
		}
		devAuthServer.Run(context.Background())
	}()
	Log.Info("Auth service started")
}


// UpdateAuthContext defines interface adding extras or modifying UserInfo in request context
type UpdateAuthContext struct {
}


type CtxUserValue struct{}

// Update user info in request context from go-pkgz/auth token.User to mouthpiece.User
func (a *UpdateAuthContext) Update() func(http.Handler) http.Handler {
	f := func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// call update only if user info exists, otherwise do nothing
			if user, err := token.GetUserInfo(r); err == nil {
				/* find out DB User */

				r = r.WithContext(context.WithValue(r.Context(), CtxUserValue{}, user))
			}
			h.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
	return f
}