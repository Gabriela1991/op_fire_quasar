package service

import (
	"app/src/bundles/health-check/models"
)

type service struct{}

type Health models.HealthModel

type HealthService interface {
	HealthCheck() Health
}

func NewHealthService() HealthService {
	return &service{}
}

func (*service) HealthCheck() Health {
	response := Health{
		Status: "UP",
	}
	return response
}
