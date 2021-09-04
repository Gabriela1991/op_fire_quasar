package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	healthsrv HealthService = NewHealthService()
)

const (
	STATUS string = "UP"
)

func TestHealthcheck(t *testing.T) {

	responseHealth := healthsrv.HealthCheck()

	assert.NotNil(t, responseHealth.Status)
	assert.Equal(t, STATUS, responseHealth.Status)
}
