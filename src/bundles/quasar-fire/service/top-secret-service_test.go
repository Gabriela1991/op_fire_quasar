package service

import (
	"app/src/bundles/quasar-fire/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	topSecretSrv TopSecretService = NewTopSecretService()
)

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
	satellites.Satellites = satList
	return satellites
}

func TestDecryptQR_Ok(t *testing.T) {
	var object models.Satelites = getSatellites_Ok()

	response, _ := topSecretSrv.TopSecret(object)
	assert.NotNil(t, response)
}
