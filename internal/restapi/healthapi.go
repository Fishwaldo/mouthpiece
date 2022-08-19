package restapi

import (
	"github.com/Fishwaldo/mouthpiece/internal/health"
	"github.com/alexliesenfeld/health"
	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/responses"
	"net/http"
)

func setupHealth(Server *huma.Router) {
	healthChecker.StartHealth()
	// Declare the root resource and a GET operation on it.
	healthapi := Server.Resource("/health")
	healthapi.Tags("Health")
	healthapi.Get("get-health", "Get Health of the Service",
		responses.OK().ContentType("application/json"),
		responses.OK().Headers("Content-Type"),
		responses.OK().Model(health.CheckerResult{}),
		responses.InternalServerError().ContentType("application/json"),
		responses.InternalServerError().Headers("Content-Type"),
		responses.InternalServerError().Model(health.CheckerResult{}),
	).Run(func(ctx huma.Context) {
		test := healthChecker.HealthChecker.Check(ctx)
		status := http.StatusOK
		if test.Status != health.StatusUp {
			status = http.StatusInternalServerError
		}
		ctx.WriteModel(status, test)
	})
}
