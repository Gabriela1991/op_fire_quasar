package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"app/src/bundles/quasar-fire/models"
	"app/src/bundles/quasar-fire/service"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var (
	topSecrectSrvTest    service.TopSecretService = service.NewTopSecretService()
	fireQuasarController FireQuasarController     = NewFireQuasarController(topSecrectSrvTest)
)

func TestTopSecret_BadRequest(t *testing.T) {

	req, _ := http.NewRequest("POST", "topsecret", nil)

	handler := http.HandlerFunc(fireQuasarController.TopSecret)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	assert.Equal(t, http.StatusBadRequest, response.Code)

}

func TestTopSecret_NotFound(t *testing.T) {

	req, _ := http.NewRequest("POST", "topsecre", nil)

	handler := http.HandlerFunc(fireQuasarController.TopSecret)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	assert.Equal(t, http.StatusBadRequest, response.Code)

}

func TestTopSecret_Ok(t *testing.T) {
	satellites := getSatellites_Ok()
	jsonR, _ := json.Marshal(&satellites)
	bodyresp := strings.NewReader(string(jsonR))

	req, _ := http.NewRequest("POST", "topsecret", bodyresp)

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(fireQuasarController.TopSecret)
	handler.ServeHTTP(response, req)

	var tsResponse models.TopSecretResponse
	json.NewDecoder(io.Reader(response.Body)).Decode(&tsResponse)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.NotNil(t, tsResponse.Message)

}

func TestTopSecret_errorJson(t *testing.T) {

	var satellites = "prueba"
	jsonR, _ := json.Marshal(&satellites)
	bodyresp := strings.NewReader(string(jsonR))

	req, _ := http.NewRequest("POST", "topsecret", bodyresp)

	handler := http.HandlerFunc(fireQuasarController.TopSecret)

	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func getSatellites_Ok() models.Satelites {
	var satellites models.Satelites
	var satList []models.Satelite

	// Kenoby
	var sat = models.Satelite{
		Name:     "Kenoby",
		Distance: 100.0,
		Message:  []string{"este", " ", " ", "mensaje", " "},
	}
	satList = append(satList, sat)

	// Skywalker
	sat = models.Satelite{
		Name:     "Skywalker",
		Distance: 115.5,
		Message:  []string{" ", "es", " ", " ", "secreto"},
	}
	satList = append(satList, sat)

	// Sato
	sat = models.Satelite{
		Name:     "Sato",
		Distance: 142.7,
		Message:  []string{"este", " ", "un", " ", " "},
	}
	satList = append(satList, sat)
	satellites.Satellites = satList
	return satellites
}

func TestTopSecretSplit_BadRequest(t *testing.T) {

	req, _ := http.NewRequest("GET", "topsecret_split", nil)

	handler := http.HandlerFunc(fireQuasarController.TopSecret)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	assert.Equal(t, http.StatusBadRequest, response.Code)

}

func TestTopSecretSplit_NotFound(t *testing.T) {

	req, _ := http.NewRequest("GET", "topsecret_s", nil)

	handler := http.HandlerFunc(fireQuasarController.TopSecret)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	assert.Equal(t, http.StatusBadRequest, response.Code)

}

func TestTopSecretSplit_Ok(t *testing.T) {
	topSecretSplitreq := models.TopSecretSplitReq{
		Distance: 100,
		Message:  []string{"Esta", "es", "una", "prueba"},
	}
	jsonR, _ := json.Marshal(&topSecretSplitreq)
	bodyresp := strings.NewReader(string(jsonR))

	vars := map[string]string{
		"satellite_name": "skywalker",
	}

	req, _ := http.NewRequest("GET", "topsecret_split/skywalker", bodyresp)
	req = mux.SetURLVars(req, vars)

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(fireQuasarController.TopSecretSplit)
	handler.ServeHTTP(response, req)

	var tsResponse models.TopSecretResponse
	json.NewDecoder(io.Reader(response.Body)).Decode(&tsResponse)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.NotNil(t, tsResponse.Message)
}

func TestTopSecretSplit_SatErr(t *testing.T) {
	topSecretSplitreq := models.TopSecretSplitReq{
		Distance: 100,
		Message:  []string{"Esta", "es", "una", "prueba"},
	}
	jsonR, _ := json.Marshal(&topSecretSplitreq)
	bodyresp := strings.NewReader(string(jsonR))

	vars := map[string]string{
		"satellite_name": "prueba",
	}

	req, _ := http.NewRequest("GET", "topsecret_split/prueba", bodyresp)
	req = mux.SetURLVars(req, vars)

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(fireQuasarController.TopSecretSplit)
	handler.ServeHTTP(response, req)

	var tsResponse models.TopSecretResponse
	json.NewDecoder(io.Reader(response.Body)).Decode(&tsResponse)
	assert.Equal(t, http.StatusNotFound, response.Code)
	assert.NotNil(t, tsResponse.Message)
}
