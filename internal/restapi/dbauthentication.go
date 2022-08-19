package restapi

import (
	"context"
	"crypto/rand"
	"crypto/sha1" //nolint
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
	"time"

	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/go-pkgz/rest"
	"github.com/golang-jwt/jwt"

	"github.com/go-pkgz/auth/logger"
	"github.com/go-pkgz/auth/provider"
	"github.com/go-pkgz/auth/token"
)

const (
	// MaxHTTPBodySize defines max http body size
	MaxHTTPBodySize = 1024 * 1024
)

type ICredChecker func(user string, password string) (ok bool, err error)

// DirectHandler implements non-oauth2 provider authorizing user in traditional way with storage
// with users and hashes
type DirectHandler struct {
	logger.L
	ProviderName string
	TokenService provider.TokenService
	Issuer       string
	AvatarSaver  provider.AvatarSaver
	CredChecker  ICredChecker
}

// credentials holds user credentials
type credentials struct {
	User     string `json:"user"`
	Password string `json:"passwd"`
	Audience string `json:"aud"`
}

// Name of the handler
func (p DirectHandler) Name() string { return p.ProviderName }

// LoginHandler checks "user" and "passwd" against data store and makes jwt if all passed.
//
// GET /something?user=name&passwd=xyz&aud=bar&sess=[0|1]
//
// POST /something?sess[0|1]
// Accepts application/x-www-form-urlencoded or application/json encoded requests.
//
// application/x-www-form-urlencoded body example:
// user=name&passwd=xyz&aud=bar
//
// application/json body example:
// {
//   "user": "name",
//   "passwd": "xyz",
//   "aud": "bar",
// }
func (p DirectHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	creds, err := p.getCredentials(w, r)
	if err != nil {
		rest.SendErrorJSON(w, r, p.L, http.StatusBadRequest, err, "failed to parse credentials")
		return
	}
	sessOnly := r.URL.Query().Get("sess") == "1"

	ok, err := p.CredChecker(creds.User, creds.Password)
	if err != nil {
		rest.SendErrorJSON(w, r, p.L, http.StatusInternalServerError, err, "failed to check user credentials")
		return
	}
	if !ok {
		rest.SendErrorJSON(w, r, p.L, http.StatusForbidden, nil, "incorrect user or password")
		return
	}

	userID := p.ProviderName + "_" + token.HashID(sha1.New(), creds.User)

	u := token.User{
		Name:  creds.User,
		ID:    userID,
		Email: creds.User,
	}
	u, err = setAvatar(p.AvatarSaver, u, &http.Client{Timeout: 5 * time.Second})
	if err != nil {
		rest.SendErrorJSON(w, r, p.L, http.StatusInternalServerError, err, "failed to save avatar to proxy")
		return
	}

	cid, err := randToken()
	if err != nil {
		rest.SendErrorJSON(w, r, p.L, http.StatusInternalServerError, err, "can't make token id")
		return
	}

	claims := token.Claims{
		User: &u,
		StandardClaims: jwt.StandardClaims{
			Id:       cid,
			Issuer:   p.Issuer,
			Audience: creds.Audience,
		},
		SessionOnly: sessOnly,
	}

	if _, err = p.TokenService.Set(w, claims); err != nil {
		rest.SendErrorJSON(w, r, p.L, http.StatusInternalServerError, err, "failed to set token")
		return
	}
	rest.RenderJSON(w, claims.User)
}

// getCredentials extracts user and password from request
func (p DirectHandler) getCredentials(w http.ResponseWriter, r *http.Request) (credentials, error) {

	// GET /something?user=name&passwd=xyz&aud=bar
	if r.Method == "GET" {
		return credentials{
			User:     r.URL.Query().Get("user"),
			Password: r.URL.Query().Get("passwd"),
			Audience: r.URL.Query().Get("aud"),
		}, nil
	}

	if r.Method != "POST" {
		return credentials{}, fmt.Errorf("method %s not supported", r.Method)
	}

	if r.Body != nil {
		r.Body = http.MaxBytesReader(w, r.Body, MaxHTTPBodySize)
	}
	contentType := r.Header.Get("Content-Type")
	if contentType != "" {
		mt, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
		if err != nil {
			return credentials{}, err
		}
		contentType = mt
	}

	// POST with json body
	if contentType == "application/json" {
		var creds credentials
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			return credentials{}, fmt.Errorf("failed to parse request body: %w", err)
		}
		return creds, nil
	}

	// POST with form
	if err := r.ParseForm(); err != nil {
		return credentials{}, fmt.Errorf("failed to parse request: %w", err)
	}

	return credentials{
		User:     r.Form.Get("user"),
		Password: r.Form.Get("passwd"),
		Audience: r.Form.Get("aud"),
	}, nil
}

// AuthHandler doesn't do anything for direct login as it has no callbacks
func (p DirectHandler) AuthHandler(w http.ResponseWriter, r *http.Request) {}

// LogoutHandler - GET /logout
func (p DirectHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	p.TokenService.Reset(w)
}

// setAvatar saves avatar and puts proxied URL to u.Picture
func setAvatar(ava provider.AvatarSaver, u token.User, client *http.Client) (token.User, error) {
	if ava != nil {
		avatarURL, e := ava.Put(u, client)
		if e != nil {
			return u, fmt.Errorf("failed to save avatar for: %w", e)
		}
		u.Picture = avatarURL
		return u, nil
	}
	return u, nil // empty AvatarSaver ok, just skipped
}

func randToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("can't get random: %w", err)
	}
	s := sha1.New()
	if _, err := s.Write(b); err != nil {
		return "", fmt.Errorf("can't write randoms to sha1: %w", err)
	}
	return fmt.Sprintf("%x", s.Sum(nil)), nil
}

func dbAuthProvider(user, pass string) (ok bool, err error) {
	// 	user = strings.TrimSpace(user)
	// 	log.Log.Info("Direct Login", "user", user)
	// 	dbUser, err := u.GetUserByEmail(context.Background(), user)
	// 	log.Log.Info("User", "user", dbUser, "error", err)

	// 	if err == mperror.ErrUserNotFound {
	// 		log.Log.Info("User not found", "user", user)
	// 		return false, nil
	// 	}
	// 	if !dbUser.CheckPassword(context.Background(), pass) {
	// 		log.Log.Info("Password Invalid", "user", user)
	// 		return false, nil
	// 	}
	return true, nil
}

// // Called when the Tokens are created/refreshed.
func MapClaimsToUser(claims token.Claims) token.Claims {
	// 	//Log.Info("Map Claims To User", "claims", claims)
	// 	//	if claims.User != nil {
	// 	//		if user, err := GetUser(claims.User.Name); err != nil {
	// 	//			Log.Info("User not found", "user", claims.User.Name)
	// 	//			claims.User.SetBoolAttr("valid", false)
	// 	//		} else {
	// 	//			claims.User.SetStrAttr("backenduser", user.Username)
	// 	//			claims.User.SetBoolAttr("valid", true)
	// 	//		}
	// 	//	}
	return claims
}

// // called on every access to the API
func UserValidator(token string, claims token.Claims) bool {
	// 	//Log.Info("User Validator", "user", claims.User.Name)
	// 	if claims.User != nil {
	// 		if user, _ := u.GetUserByEmail(context.Background(), claims.User.Name); user != nil {
	// 			claims.User.ID = fmt.Sprintf("%d", user.GetID())
	// 			return true
	// 		}
	// 	}
	// 	return false
	return true
}

func GetUserFromContext(ctx context.Context) (bool, interfaces.UserI) {
	// 	v := ctx.Value(interfaces.CtxUserValue{}).(token.User)
	// 	if id, _ := strconv.Atoi(v.ID); id > 0 {
	// 		if user, _ := u.GetUser(ctx, uint(id)); user != nil {
	// 			return true, user
	// 		}
	// 	}
	return false, nil
}
