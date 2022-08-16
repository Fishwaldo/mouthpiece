package healthChecker

import (
	"context"
	_ "fmt"
	"time"

	"github.com/Fishwaldo/mouthpiece/internal/db"
	. "github.com/Fishwaldo/mouthpiece/internal/log"

	"github.com/alexliesenfeld/health"

	httpCheck "github.com/hellofresh/health-go/v4/checks/http"
)

var HealthChecker health.Checker

func StartHealth() {

	HealthChecker = health.NewChecker(
		health.WithTimeout(10*time.Second),
		//health.WithInterceptors(interceptors.BasicLogger()),
		health.WithCheck(health.Check{
			Name: "Google",
			Check: httpCheck.New(httpCheck.Config{
				URL: "https://www.google.com/",
			}),
		}),
		health.WithCheck(health.Check{
			Name: "Database",
			Check: func(ctx context.Context) error {
				sqlDB, err := db.Db.DB()
				if err != nil {
					return err
				}

				return sqlDB.PingContext(ctx)
			},
		}),
		health.WithInterceptors(BasicLogger()),
	)
	HealthChecker.Start()
}

// BasicLogger is a basic logger that is mostly used to showcase this library.
func BasicLogger() health.Interceptor {
	return func(next health.InterceptorFunc) health.InterceptorFunc {
		return func(ctx context.Context, name string, state health.CheckState) health.CheckState {
			now := time.Now()
			result := next(ctx, name, state)
			Log.V(1).Info("processed health check request",
				"check", name, "seconds", time.Now().Sub(now).Seconds(), "result", result.Status)
			return result
		}
	}
}
