package server

import (
	"app/src/bundles/health-check/controller"
	"app/src/bundles/health-check/models"
	"app/src/bundles/health-check/service"
	router "app/src/http"
)

var (
	healthSrv        service.HealthService            = service.NewHealthService()
	healthController controller.HealthcheckController = controller.NewHealthcheckController(healthSrv)
)

// Health response payload
// swagger:response healthRes
type swaggHealthRes struct {
	// in:body
	Body models.HealthModel
}

func HealthCheckRoute(httpRouter router.Router) {
	// swagger:operation GET healthCheck health getHealthCheck
	// ---
	// summary: Return an HealthCheck of application.
	// description: If the application is run, health will be returned else Error Not Found (404) will be returned.
	// parameters:
	// responses:
	//   "200":
	//     "$ref": "#/responses/healthRes"
	//   "404":
	//     "$ref": "#/responses/notFoundReq"
	httpRouter.GET("/healthCheck", healthController.GetHealthcheck)
}
