package service

import (
	"app/src/bundles/quasar-fire/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	topSecretSrv TopSecretService = NewTopSecretService()
)

func getSatellites_Err() models.Satelites {
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
	satellites.Satellites = satList
	return satellites
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
		Distance: 115.5,
		Message:  []string{" ", "es", "un", " ", "secreto"},
	}
	satList = append(satList, sat)
	satellites.Satellites = satList
	return satellites
}

func TestTopSecret_Ok(t *testing.T) {
	var object models.Satelites = getSatellites_Ok()

	response, _ := topSecretSrv.TopSecret(object)
	assert.NotNil(t, response)
}

func TestTopSecret_Err(t *testing.T) {
	var object models.Satelites = getSatellites_Err()

	response, _ := topSecretSrv.TopSecret(object)
	assert.NotNil(t, response)
}

func TestTopSecretSplit_Ok(t *testing.T) {
	object := models.TopSecretSplitReq{
		Distance: 100.0,
		Message:  []string{"este", "es", "un", "mensaje", "prueba"},
	}

	response, _ := topSecretSrv.TopSecretSplit(object, "sato")
	assert.NotNil(t, response)
}

func TestTopSecretSplit_MsgErr(t *testing.T) {
	object := models.TopSecretSplitReq{
		Distance: 100.0,
		Message:  []string{"este", " ", " ", "mensaje", " "},
	}

	response, _ := topSecretSrv.TopSecretSplit(object, "sato")
	assert.NotNil(t, response)
}

func TestTopSecretSplit_SatelliteErr(t *testing.T) {
	object := models.TopSecretSplitReq{
		Distance: 100.0,
		Message:  []string{"este", " ", " ", "mensaje", " "},
	}

	response, _ := topSecretSrv.TopSecretSplit(object, "prueba")
	assert.NotNil(t, response)
}
