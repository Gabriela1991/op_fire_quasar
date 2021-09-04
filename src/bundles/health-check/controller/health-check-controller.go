package controller

import (
	"app/src/bundles/health-check/service"
	"encoding/json"
	"net/http"
)

var (
	healthService service.HealthService
)

type controller struct{}

type HealthcheckController interface {
	GetHealthcheck(response http.ResponseWriter, request *http.Request)
}

func NewHealthcheckController(service service.HealthService) HealthcheckController {
	healthService = service
	return &controller{}
}

func (*controller) GetHealthcheck(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application-json")
	var healthObj = healthService.HealthCheck()
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(healthObj)
}
