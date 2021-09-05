package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	locationTestSrv LocationService = NewLocationService()
)

func TestGetMultipleLocation_Ok(t *testing.T) {
	x, y := locationTestSrv.GetLocation(100, -500, 100)
	assert.NotNil(t, x)
	assert.NotNil(t, y)
}

func TestGetSingleLocation_OkSato(t *testing.T) {
	x, y, _ := locationTestSrv.CalculateCircle("sato", 100)
	assert.NotNil(t, x)
	assert.NotNil(t, y)
}

func TestGetSingleLocation_OkKenoby(t *testing.T) {
	x, y, _ := locationTestSrv.CalculateCircle("kenoby", 100)
	assert.NotNil(t, x)
	assert.NotNil(t, y)
}

func TestGetSingleLocation_OkSkywalker(t *testing.T) {
	x, y, _ := locationTestSrv.CalculateCircle("skywalker", 100)
	assert.NotNil(t, x)
	assert.NotNil(t, y)
}

func TestGetSingleLocation_SatelliteErr(t *testing.T) {
	x, _, err := locationTestSrv.CalculateCircle("prueba", 100)
	assert.Equal(t, x, float32(0))
	assert.NotNil(t, err)
}
