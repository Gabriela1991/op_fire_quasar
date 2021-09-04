package service

import (
	"app/src/bundles/quasar-fire/models"
	"app/src/bundles/quasar-fire/util"
)

var (
	trilateration util.Trilateration
)

type locationService struct{}

type LocationService interface {
	GetLocation(distances ...float32) (x, y float32)
}

func NewLocationService() LocationService {
	return &locationService{}
}

func (*locationService) GetLocation(distances ...float32) (float32, float32) {
	x, y := trilateration.CalculateTrilateration(getCircles(distances))
	return x, y
}

func getCircles(distances []float32) []models.Circle {
	var circles []models.Circle

	// Kenoby
	var circle = models.Circle{
		AxisX: -500.0,
		AxisY: -200.0,
		Radio: float64(distances[0]),
	}
	circles = append(circles, circle)

	// Skywalker
	circle = models.Circle{
		AxisX: 100.0,
		AxisY: -100.0,
		Radio: float64(distances[1]),
	}
	circles = append(circles, circle)

	// Sato
	circle = models.Circle{
		AxisX: 500.0,
		AxisY: 100.0,
		Radio: float64(distances[2]),
	}
	circles = append(circles, circle)
	return circles
}
