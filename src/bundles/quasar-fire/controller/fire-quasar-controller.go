package controller

import (
	"app/src/bundles/quasar-fire/models"
	"app/src/bundles/quasar-fire/service"
	"encoding/json"
	"net/http"

	// "strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	topSecretSrv service.TopSecretService
)

type controllers struct{}

type FireQuasarController interface {
	TopSecret(response http.ResponseWriter, request *http.Request)
	TopSecretSplit(response http.ResponseWriter, request *http.Request)
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
		strEr := `{"message": "Empty request"}`
		log.Error(strEr)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(strEr))
		return
	}

	err := json.NewDecoder(request.Body).Decode(&satellites)
	if err != nil {
		strEr := `{"message": "Incorrect JSON"}`
		log.Error(strEr)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(strEr))
		return
	}

	// *** Se llama al servicio *** //
	response.Header().Set("Content-type", "application-json")
	resp, error := topSecretSrv.TopSecret(satellites)

	// *** Se valida si ocurrio algun error en el service *** //
	if error != nil {
		strEr := `{"message": "` + error.Error() + `"}`
		log.Error(strEr)
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(strEr))
		return
	}

	log.Info("-- Call Endpoint: TopSecret Complete Successfull--")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resp)
}

func (*controllers) TopSecretSplit(response http.ResponseWriter, request *http.Request) {
	log.Info("-- Call Endpoint TopSecretSplit --")
	response.Header().Set("Content-Type", "application/json")
	var topSecretReq models.TopSecretSplitReq

	// *** Se valida el request *** //
	if request.Body == nil {
		strEr := `{"message": "Empty request"}`
		log.Error(strEr)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(strEr))
		return
	}

	err := json.NewDecoder(request.Body).Decode(&topSecretReq)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(request)
	sat_name, er := vars["satellite_name"]

	if !er || sat_name == "" {
		strEr := `{"message": "Missing var path satellite_name"}`
		log.Error(strEr)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(strEr))
		return
	}

	// *** Se llama al servicio *** //
	response.Header().Set("Content-type", "application-json")
	resp, error := topSecretSrv.TopSecretSplit(topSecretReq, sat_name)

	// *** Se valida si ocurrio algun error en el service *** //
	if error != nil {
		strEr := `{"message": "` + error.Error() + `"}`
		log.Error(strEr)
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(strEr))
		return
	}

	log.Info("-- Call Endpoint: TopSecret Complete Successfull--")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resp)
}
