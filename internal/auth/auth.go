package auth

import (
	"time"
	//	"strings"
	"context"
	"crypto/sha1"
	"fmt"
	"net/http"
	"os"
//	"strings"

	dbauth "github.com/Fishwaldo/mouthpiece/internal/auth/db"
	telegramauth "github.com/Fishwaldo/mouthpiece/internal/auth/telegram"
	"github.com/Fishwaldo/mouthpiece/internal/db"
	. "github.com/Fishwaldo/mouthpiece/internal/log"

	"github.com/spf13/viper"

	pkauth "github.com/go-pkgz/auth"
	"github.com/go-pkgz/auth/avatar"
	"github.com/go-pkgz/auth/provider"
	"github.com/go-pkgz/auth/token"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"

	//"github.com/casbin/casbin/v2/log"
	defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"
	gormadapter "github.com/casbin/gorm-adapter/v3"

	"golang.org/x/oauth2/github"
)

type Auth struct {
	Service      *pkauth.Service
	AuthEnforcer *casbin.Enforcer
}

var AuthService *Auth

type AuthLogger struct {
}

var AL *AuthLogger

func (AL AuthLogger) Logf(format string, args ...interface{}) {
	Log.WithName("Auth").Info("Authentication", "message", fmt.Sprintf(format, args...))
}

type AuthConfig struct {
	CredChecker func(username string, password string) (ok bool, err error)
	MapClaimsToUser token.ClaimsUpdFunc
	Validator token.ValidatorFunc
	Host string
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
	viper.SetDefault("auth.microsoft.signups", false)
	viper.SetDefault("auth.secret", "secret")
	viper.SetDefault("auth.debug", false)
	//viper.SetDefault("auth.avatar.cachedir", os.MkdirTemp("", "mouthpiece_avatar"))
}

func customGitHubProvider() (cred pkauth.Client, ch provider.CustomHandlerOpt) {
	ch = provider.CustomHandlerOpt{
		Endpoint: github.Endpoint,
		InfoURL:  "https://api.github.com/user",
		MapUserFn: func(data provider.UserData, _ []byte) token.User {
			userInfo := token.User{
				ID:      "github_" + token.HashID(sha1.New(), data.Value(viper.GetString("auth.github.fields.UserName"))),
				Name:    data.Value(viper.GetString("auth.github.fields.UserName")),
				Email:   data.Value(viper.GetString("auth.github.fields.Email")),
				Picture: data.Value(viper.GetString("auth.github.fields.Avatar")),
			}
			userInfo.SetStrAttr("fullname", data.Value(viper.GetString("auth.github.fields.FullName")))
			userInfo.SetStrAttr("source", "github")
			return userInfo
		},
		Scopes: []string{},
	}
	cred = pkauth.Client{
		Cid:     viper.GetString("auth.github.client_id"),
		Csecret: viper.GetString("auth.github.client_secret"),
	}

	return cred, ch
}


func InitAuth(Config AuthConfig) {
	AL = &AuthLogger{}
	AuthService = &Auth{}

	var avatarcachedir string
	if viper.IsSet("auth.avatar.cachedir") {
		avatarcachedir = viper.GetString("auth.avatar.cachedir")
	} else { 
		avatarcachedir, _ = os.MkdirTemp("", "mouthpiece_avatar")
	}
	options := pkauth.Opts{
		SecretReader: token.SecretFunc(func(_ string) (string, error) { // secret key for JWT, ignores aud
			return viper.GetString("auth.secret"), nil
		}),
		TokenDuration:  time.Minute * 5, // short token, refreshed automatically
		CookieDuration: time.Hour * 24,  // cookie fine to keep for long time
		DisableXSRF:    true,            // don't disable XSRF in real-life applications!
		Issuer:         "mouthpiece",    // part of token, just informational
		URL:            Config.Host,            // base url of the protected service
		//AdminPasswd:       "password",												  // admin password
		AvatarStore:       avatar.NewLocalFS(avatarcachedir), // stores avatars locally
		AvatarResizeLimit: 200,                                         // resizes avatars to 200x200
		ClaimsUpd: token.ClaimsUpdFunc(Config.MapClaimsToUser),
		Validator: Config.Validator,
		Logger:      AL,   // optional logger for auth library
		UseGravatar: true, // for verified provider use gravatar service
	}

	// create auth service
	AuthService.Service = pkauth.NewService(options)
	if viper.GetBool("auth.dev.enabled") {
		Log.Info("Auth Dev Mode Enabled!")
		AuthService.Service.AddProvider("dev", "", "")
		// run dev/test oauth2 server on :8084
		go func() {
			devAuthServer, err := AuthService.Service.DevAuth() // peak dev oauth2 server
			devAuthServer.GetEmailFn = func(username string) string {
				return "admin@example.com"
			}
			if err != nil {
				Log.Error(err, "[PANIC] failed to start dev oauth2 server")
			}
			devAuthServer.Run(context.Background())

		}()
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
	if viper.GetBool("auth.microsoft.enabled") {
		Log.Info("Auth Microsoft Enabled!")
		AuthService.Service.AddProvider("microsoft", os.Getenv("AEXMPL_MS_APIKEY"), os.Getenv("AEXMPL_MS_APISEC"))
	}
	/* direct loging (username/password) is always handled */
	dbprovider := dbauth.DirectHandler{
		L:            AL,
		ProviderName: "direct",
		Issuer:       options.Issuer,
		TokenService: AuthService.Service.TokenService(),
		AvatarSaver:  AuthService.Service.AvatarProxy(),
		CredChecker:  Config.CredChecker,
	}
	AuthService.Service.AddCustomHandler(dbprovider)

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
	InitCasbin()
	Log.Info("Auth service started")
}
func InitCasbin() {
	cdb, err := gormadapter.NewAdapterByDB(db.Db)
	if err != nil {
		Log.Error(err, "Failed to Setup Casbin Auth Adapter")
	}

	AuthService.AuthEnforcer, err = casbin.NewEnforcer("config/auth_model.conf", cdb)
	if err != nil {
		Log.Error(err, "Failed to setup Casbin")
	}
	AuthService.AuthEnforcer.EnableLog(viper.GetBool("auth.debug"))
	AuthService.AuthEnforcer.EnableAutoSave(true)
	AuthService.AuthEnforcer.SetRoleManager(defaultrolemanager.NewRoleManager(10))
	if err := AuthService.AuthEnforcer.LoadModel(); err != nil {
		Log.Error(err, "Failed to load Casbin model")
	}

	if err := AuthService.AuthEnforcer.LoadPolicy(); err != nil {
		Log.Error(err, "Failed to Load Casbin Policy")
	}
	if !AuthService.AuthEnforcer.AddNamedMatchingFunc("g2", "KeyMatch3", util.KeyMatch3) {
		Log.Error(nil, "Failed to add g2 matching function")
	}
	AuthService.AuthEnforcer.AddPolicy("role:admin", "apigroup:apps", "PUT")
	AuthService.AuthEnforcer.AddPolicy("role:user", "apigroup:apps", "GET")
	AuthService.AuthEnforcer.AddPolicy("role:admin", "apigroup:message", "PUT")
	AuthService.AuthEnforcer.AddPolicy("role:user", "apigroup:message", "GET")
	AuthService.AuthEnforcer.AddPolicy("role:admin", "apigroup:users", "PUT")
	AuthService.AuthEnforcer.AddPolicy("role:user", "apigroup:users", "GET")
	AuthService.AuthEnforcer.AddPolicy("role:admin", "apigroup:transports", "PUT")
	AuthService.AuthEnforcer.AddPolicy("role:user", "apigroup:transports", "GET")

	//	AuthService.AuthEnforcer.AddPolicy("role:user", "apigroup:apps", "GET")
	AuthService.AuthEnforcer.AddRoleForUser("role:admin", "role:user")

	//	AuthService.AuthEnforcer.AddRoleForUser("admin", "role:admin")
	//	AuthService.AuthEnforcer.AddRoleForUser("dev_user", "role:admin")
	p, _ := AuthService.AuthEnforcer.GetImplicitPermissionsForUser("admin@example.com")
	fmt.Printf("Admin Permissions: %+v\n", p)

	AuthService.AuthEnforcer.SavePolicy()
	rm := AuthService.AuthEnforcer.GetPolicy()
	Log.Info("Casbin Policy", "policy", rm)
	Log.Info("Casbin User Roles", "Roles", AuthService.AuthEnforcer.GetGroupingPolicy())
	Log.Info("Casbin API Groups", "API Groups", AuthService.AuthEnforcer.GetNamedGroupingPolicy("g2"))

}

func (a *Auth) AddResourceURL(url string, group string) bool {
	ok, err := a.AuthEnforcer.AddNamedGroupingPolicy("g2", url, group)
	if err != nil {
		Log.Error(err, "Failed to add g2 policy", "url", url, "group", group)
	}
	return ok
}
