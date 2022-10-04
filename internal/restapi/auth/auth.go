package auth

import (
	"net/http"
	"strconv"
	"time"

	//	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/Fishwaldo/mouthpiece/internal/cache"
	"github.com/Fishwaldo/mouthpiece/internal/server"
	"github.com/Fishwaldo/mouthpiece/pkg/log"

	"github.com/go-logr/logr"

	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/responses"
	"github.com/spf13/viper"
)

var (
	llog logr.Logger
)

type passwordLoginResult struct {
	Status       string
	SessionToken string
	RefreshToken string
}

type passwordLogin struct {
	Body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"body"`
}

type tokenRefresh struct {
	Body struct {
		RefreshToken string `json:"refresh_token"`
	} `json:"body"`
}

type patRequest struct {
	Body struct {
		Duration time.Duration `json:"duration"`
	} `json:"body"`
}

type patResponse struct {
	Status string
	Token  string
}

type meResponse struct {
	Status string
	Email string
	Name string
}

type logoutResponse struct {
	Status string
}

type AuthData struct {
	Password             string
	PersonalAccessTokens map[string]int64
}

func init() {
	viper.SetDefault("jwt.session", 5)
	viper.SetDefault("jwt.refresh", 30)
	viper.SetDefault("user.pat.max", 5)
}

func Setup(res *huma.Resource) error {
	llog = log.Log.WithName("auth")

	/* setup a PAT/APP Cache */
	patCache = cache.GetNewCache[patCacheEntry]()

	authapi := res.SubResource("/auth")
	authapi.Tags("Auth")

	loginapi := authapi.SubResource("/password")

	loginapi.Post("login", "Login to Application with Username/Password",
		responses.OK().ContentType("application/json"),
		responses.OK().Model(passwordLoginResult{}),
		responses.NotAcceptable().ContentType("application/json"),
		responses.Unauthorized().ContentType("application/json"),
	).Run(login)

	meapi := authapi.SubResource("/me")
	meapi.Tags("Auth")
	meapi.Middleware(RequireAuth)

	meapi.Get("me", "Get information about the current user",
		responses.OK().ContentType("application/json"),
		responses.OK().Model(meResponse{}),
	).Run(getMe)


	tokenapi := authapi.SubResource("/tokenrefresh")
	tokenapi.Tags("Auth")
//	tokenapi.Middleware(RequireAuth)

	tokenapi.Post("refresh", "Refresh a token",
		responses.OK().ContentType("application/json"),
		responses.OK().Model(passwordLoginResult{}),
		responses.NotAcceptable().ContentType("application/json"),
		responses.Unauthorized().ContentType("application/json"),
	).Run(refresh)

	pattoken := authapi.SubResource("/pat")
	pattoken.Tags("Auth")
	pattoken.Middleware(RequireAuth)

	pattoken.Put("create", "Create a Personal Access Token",
		responses.OK().ContentType("application/json"),
		responses.OK().Model(patResponse{}),
		responses.NotAcceptable().ContentType("application/json"),
		responses.Unauthorized().ContentType("application/json"),
	).Run(createPAT)

	logoutapi := authapi.SubResource("/logout")
	logoutapi.Tags("Auth")
	logoutapi.Middleware(RequireAuth)

	logoutapi.Delete("logout", "Logout of Application",
		responses.OK().ContentType("application/json"),
		responses.OK().Model(logoutResponse{}),
	).Run(logout)




	return nil
}

func logout(ctx huma.Context) {
	ctx.WriteModel(http.StatusOK, logoutResponse{Status: "OK"})
}

func getMe(ctx huma.Context) {
	claims := GetClaimsFromContext(ctx)
	if claims == nil {
		ctx.WriteError(http.StatusUnauthorized, "Invalid User")
		return
	}
	id, err := strconv.Atoi(claims.Subject)
	if err != nil {
		llog.Error(err, "Can't convert ID to int")
		ctx.WriteError(http.StatusUnauthorized, "Invalid Token")
		return
	}
	user, err := server.Get().GetUserService().GetByID(ctx, id)
	if err != nil {
		llog.Error(err, "Error getting user")
		ctx.WriteError(http.StatusUnauthorized, "Internal Error")
		return
	}

	res := meResponse{
		Status: "OK",
		Email: user.GetEmail(),
		Name: user.GetName(),
	}
	ctx.WriteModel(http.StatusOK, res)
}

func login(ctx huma.Context, input passwordLogin) {
	llog.Info("Login Attempt", "username", input.Body.Username)

	user, err := server.Get().GetUserService().Get(ctx, input.Body.Username)
	if err != nil {
		llog.Error(err, "Error getting user")
		ctx.WriteError(http.StatusUnauthorized, "Internal Error")
		return
	}

	var ad AuthData
	err = user.GetAppData(ctx, "auth", &ad)
	if err != nil {
		llog.V(1).Error(err, "Initialize AuthData")
		ad.PersonalAccessTokens = make(map[string]int64)
	}

	llog.V(1).Info("Logging in On Tenant", "tenant", server.Get().GetTenant(ctx).Name, "User", user.GetEmail())
	llog.V(1).Info("Claims", "claims", GetClaimsFromContext(ctx))


		/* XXX EXTREMLY DANGEROUS */
	if !viper.GetBool("debug") {
		err = bcrypt.CompareHashAndPassword([]byte(ad.Password), []byte(input.Body.Password))
		if err != nil {
			llog.Error(err, "Error comparing password")
			ctx.WriteError(http.StatusUnauthorized, "Invalid Username/Password")
			return
		}
	}

	sesstok, err := newLoginToken(user)
	if err != nil {
		llog.Error(err, "Error creating token")
		ctx.WriteError(http.StatusUnauthorized, "Internal Error", err)
		return
	}
	refreshtok, err := newRefreshToken(user)
	if err != nil {
		llog.Error(err, "Error creating token")
		ctx.WriteError(http.StatusUnauthorized, "Internal Error", err)
		return
	}

	res := passwordLoginResult{
		Status:       "OK",
		SessionToken: sesstok,
		RefreshToken: refreshtok,
	}
	//ctx.Header().Set("Authorization", "Bearer "+tok)
	ctx.WriteModel(http.StatusOK, res)
	llog.Info("Login Success", "username", input.Body.Username)
}

func refresh(ctx huma.Context, input tokenRefresh) {
	llog.Info("Refresh Attempt")

	claims, err := parse(input.Body.RefreshToken)
	if err != nil {
		llog.Error(err, "Error parsing token")
		ctx.WriteError(http.StatusUnauthorized, "Invalid Token")
		return
	}

	id, err := strconv.Atoi(claims.Subject)
	if err != nil {
		llog.Error(err, "Can't convert ID to int")
		ctx.WriteError(http.StatusUnauthorized, "Invalid Token")
		return
	}

	if err := claims.Valid(); err != nil {
		llog.Error(err, "Token not valid")
		ctx.WriteError(http.StatusUnauthorized, "Invalid Token")
		return
	}

	if !claims.VerifyAudience("refresh", true) {
		llog.Error(nil, "Token is not a Refresh Token")
		ctx.WriteError(http.StatusUnauthorized, "Invalid Token")
		return
	}

	user, err := server.Get().GetUserService().GetByID(ctx, id)
	if err != nil {
		llog.Error(err, "Error getting user")
		ctx.WriteError(http.StatusUnauthorized, "Internal Error")
		return
	}

	sesstok, err := newLoginToken(user)
	if err != nil {
		llog.Error(err, "Error creating token")
		ctx.WriteError(http.StatusUnauthorized, "Internal Error", err)
		return
	}
	refreshtok, err := newRefreshToken(user)
	if err != nil {
		llog.Error(err, "Error creating token")
		ctx.WriteError(http.StatusUnauthorized, "Internal Error", err)
		return
	}

	res := passwordLoginResult{
		Status:       "OK",
		SessionToken: sesstok,
		RefreshToken: refreshtok,
	}
	//ctx.Header().Set("Authorization", "Bearer "+tok)
	ctx.WriteModel(http.StatusOK, res)
	llog.Info("Refresh Success")
}

func createPAT(ctx huma.Context, input patRequest) {
	llog.Info("Create PAT Attempt")

	claims := GetClaimsFromContext(ctx)

	id, err := strconv.Atoi(claims.Subject)
	if err != nil {
		llog.Error(err, "Can't convert ID to int")
		ctx.WriteError(http.StatusUnauthorized, "Invalid Token")
		return
	}

	user, err := server.Get().GetUserService().GetByID(ctx, id)
	if err != nil {
		llog.Error(err, "Error getting user")
		ctx.WriteError(http.StatusUnauthorized, "Internal Error")
		return
	}

	token, err := newPatToken(ctx, user, input.Body.Duration)
	if err != nil {
		llog.Error(err, "Error creating token")
		ctx.WriteError(http.StatusUnauthorized, "Internal Error", err)
		return
	}

	res := patResponse{
		Status: "OK",
		Token:  token,
	}
	ctx.WriteModel(http.StatusOK, res)
	llog.Info("Create PAT Success")
}
