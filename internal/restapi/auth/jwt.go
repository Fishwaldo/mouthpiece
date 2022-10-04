package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Fishwaldo/mouthpiece/internal/server"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/danielgtaylor/huma"
	"github.com/google/uuid"

	gc "github.com/eko/gocache/v3/cache"
	"github.com/eko/gocache/v3/store"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type mpCustomClaims struct {
	jwt.RegisteredClaims
}

type patCacheEntry struct {
	/* ExpiresAt is when the PAT/App Token Expires */
	ExpiresAt time.Time
	/* The ID of the PAT or App Token */
	PATid string
	/* the ID of the User or App */
	ID int
	/* True if this is a App Token */
	App bool
}

var (
	patCache *gc.MetricCache[patCacheEntry]
)

func newLoginToken(user interfaces.UserI) (string, error) {
	claims := &mpCustomClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(viper.GetInt("jwt.session")) * time.Minute)),
			Issuer:    "mouthpiece",
			Audience:  []string{"login"},
			Subject:   fmt.Sprintf("%d", user.GetID()),
			ID:        uuid.NewString(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(viper.GetString("secret")))
	return tokenString, err
}

func newRefreshToken(user interfaces.UserI) (string, error) {
	claims := &mpCustomClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(viper.GetInt("jwt.refresh")) * time.Minute)),
			Issuer:    "mouthpiece",
			Audience:  []string{"refresh"},
			Subject:   fmt.Sprintf("%d", user.GetID()),
			ID:        uuid.NewString(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(viper.GetString("secret")))
	return tokenString, err
}

func newPatToken(ctx huma.Context, user interfaces.UserI, expire time.Duration) (string, error) {
	claims := &mpCustomClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire * time.Hour)),
			Issuer:    "mouthpiece",
			Audience:  []string{"pat"},
			Subject:   fmt.Sprintf("%d", user.GetID()),
			ID:        uuid.NewString(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(viper.GetString("secret")))
	if err != nil {
		return "", err
	}
	var ad AuthData
	/* ignore the Error in case there is no AppData there already */
	_ = user.GetAppData(ctx, "auth", &ad)



	if ad.PersonalAccessTokens == nil {
		ad.PersonalAccessTokens = make(map[string]int64)
	}
	if len(ad.PersonalAccessTokens) >= viper.GetInt("user.pat.max") {
		return "", errors.New("Maximum number of personal access tokens reached")
	}
	ad.PersonalAccessTokens[claims.ID] = claims.ExpiresAt.Time.Unix()
	if err := user.SetAppData(ctx, "auth", ad); err != nil {
		return "", err
	}
	/* add it to the Cache */
	patCache.Set(ctx, claims.ID, patCacheEntry{
		ExpiresAt: claims.ExpiresAt.Time,
		PATid:     claims.ID,
		ID:        user.GetID(),
		App:       false,
	}, store.WithExpiration(time.Until(claims.ExpiresAt.Time)))

	return tokenString, err
}

func newAppToken(app interfaces.AppI, expire time.Duration) (string, error) {
	claims := &mpCustomClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire)),
			Issuer:    "mouthpiece",
			Audience:  []string{"app"},
			Subject:   fmt.Sprintf("%d", app.GetID()),
			ID:        uuid.NewString(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(viper.GetString("secret")))
	return tokenString, err
}

func parse(token string) (*mpCustomClaims, error) {
	claims := &mpCustomClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("secret")), nil
	})
	if err != nil {
		return claims, err
	}

	return claims, nil
}

func ParseJWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var jwttoken string
		ctx := r.Context()
		ctx = server.Get().SetUserTenant(ctx)

		if r.URL.Path == "/api/tokenrefresh" {
			llog.V(1).Info("Skipping JWT Auth for token refresh")
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		header := r.Header.Get("Authorization")
		if len(header) == 0 {
			cookie, err := r.Cookie("jwt")
			if errors.Is(err, http.ErrNoCookie) {
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			} else {
				jwttoken = cookie.Value
			}
		} else {
			if len(header) > 7 && strings.ToUpper(header[0:6]) == "BEARER" {
				jwttoken = header[7:]
			}
		}

		if len(jwttoken) == 0 {
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		claims, err := parse(jwttoken)
		if err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
			llog.Error(err, "Error parsing JWT")
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		if !claims.VerifyExpiresAt(time.Now(), true) && claims.VerifyAudience("login", true) {
			llog.Error(nil, "Token Expired", "claims", claims.ExpiresAt.Time)
			ctx := huma.ContextFromRequest(w, r)
			ctx.WriteHeader(http.StatusUnauthorized)
			ctx.Write([]byte("Token Expired"))
			return
		}

		if err := claims.Valid(); err != nil {
			llog.Error(err, "Token Didn't Validate")
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		if !claims.VerifyIssuer("mouthpiece", true) {
			llog.Error(err, "Invalid Issuer", "issuer", claims.Issuer)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		if claims.VerifyAudience("login", true) {
			llog.V(1).Info("JWT Login Token Successful", "user", claims.Subject)
			next.ServeHTTP(w, r.WithContext(newContextWithClaims(ctx, claims)))
		} else if claims.VerifyAudience("pat", true) {
			/* lets check if the PAT is still valid */
			userid, err := strconv.Atoi(claims.Subject)
			if err != nil {
				llog.Error(err, "Invalid Subject")
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
			if err := checkPAT(ctx, claims.ID, userid); err != nil {
				llog.Error(err, "Invalid PAT")
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			llog.V(1).Info("JWT PAT Token Successful", "user", claims.Subject)
			next.ServeHTTP(w, r.WithContext(newContextWithClaims(ctx, claims)))
		} else if claims.VerifyAudience("app", true) {
			llog.V(1).Info("JWT App Token Successful", "app", claims.Subject)
			next.ServeHTTP(w, r.WithContext(newContextWithClaims(ctx, claims)))
		} else {
			llog.Error(nil, "Invalid audience", "Audience", claims.Audience)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

func checkPAT(ctx context.Context, patid string, userid int) error {

	if pce, err := patCache.Get(ctx, patid); err == nil {
		if (pce.PATid == patid) && (pce.ID == userid) && pce.ExpiresAt.After(time.Now()) {
			llog.V(1).Info("PAT Cache Hit", "PCE", pce)
			return nil
		}
	}

	var ad AuthData
	user, err := server.Get().GetUserService().GetByID(ctx, userid)
	if err != nil {
		llog.V(1).Error(err, "Error getting user")
		return err
	}

	if err := user.GetAppData(ctx, "auth", &ad); err != nil {
		return err
	}
	if ad.PersonalAccessTokens == nil {
		return errors.New("No PATs found")
	}
	if _, ok := ad.PersonalAccessTokens[patid]; !ok {
		return errors.New("PAT not found")
	}
	/* create a cache entry */
	pce := patCacheEntry{
		PATid:     patid,
		ID:        userid,
		ExpiresAt: time.Unix(ad.PersonalAccessTokens[patid], 0),
	}
	if err := patCache.Set(ctx, patid, pce); err != nil {
		llog.Error(err, "Error setting PAT Cache")
	}

	return nil
}

type jwtContext struct {
}

func (j *jwtContext) String() string {
	return "jwtContext"
}

func newContextWithClaims(ctx context.Context, claims *mpCustomClaims) context.Context {
	return context.WithValue(ctx, jwtContext{}, claims)
}

func GetClaimsFromContext(ctx context.Context) *mpCustomClaims {
	val := ctx.Value(jwtContext{})
	if val == nil {
		return nil
	}
	if _, ok := val.(*mpCustomClaims); !ok {
		return nil
	}
	return ctx.Value(jwtContext{}).(*mpCustomClaims)
}

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := GetClaimsFromContext(r.Context())
		if claims == nil {
			ctx := huma.ContextFromRequest(w, r)
			ctx.WriteHeader(http.StatusForbidden)
			ctx.Write([]byte("Unauthorized"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
