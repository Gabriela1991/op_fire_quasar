package controller

import (
	"app/src/bundles/quasar-fire/models"
	"app/src/bundles/quasar-fire/service"
	"encoding/json"
	"net/http"

	// "strconv"

	log "github.com/sirupsen/logrus"
)

var (
	topSecretSrv service.TopSecretService
)

type controllers struct{}

type FireQuasarController interface {
	TopSecret(response http.ResponseWriter, request *http.Request)
}

func NewFireQuasarController(service service.TopSecretService) FireQuasarController {
	topSecretSrv = service
	return &controllers{}
}

func (*controllers) TopSecret(response http.ResponseWriter, request *http.Request) {
	log.Info("-- Call Endpoint TopSecret --")
	response.Header().Set("Content-Type", "application/json")
	var satellites models.Satelites

	// *** Se valida el request *** //
	if request.Body == nil {
		http.Error(response, "Empty request", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(request.Body).Decode(&satellites)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	// *** Se llama al servicio *** //
	response.Header().Set("Content-type", "application-json")
	resp, error := topSecretSrv.TopSecret(satellites)

	// *** Se valida si ocurrio algun error en el service *** //
	if error != nil {
		response.WriteHeader(http.StatusNotFound)
		log.Error(`{"message": "` + error.Error() + `"}`)
		return
	}

	log.Info("-- Call Endpoint: TopSecret Complete Successfull--")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resp)
}
