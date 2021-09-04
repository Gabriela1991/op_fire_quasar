package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"app/src/bundles/health-check/models"
	"app/src/bundles/health-check/service"

	"github.com/stretchr/testify/assert"
)

var (
	healthSrv   service.HealthService = service.NewHealthService()
	healthCheck HealthcheckController = NewHealthcheckController(healthSrv)
)

const (
	STATUS string = "UP"
)

func TestGetHealthcheck(t *testing.T) {

	req, _ := http.NewRequest("GET", "health", nil)

	handler := http.HandlerFunc(healthCheck.GetHealthcheck)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Handler return wrong status code: got %v want %v", status, http.StatusOK)
	}

	var healthResponse models.HealthModel
	json.NewDecoder(io.Reader(response.Body)).Decode(&healthResponse)

	assert.NotNil(t, healthResponse.Status)
	assert.Equal(t, STATUS, healthResponse.Status)

}
