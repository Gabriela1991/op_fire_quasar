package service

import (
	"app/src/bundles/quasar-fire/models"
	"app/src/bundles/quasar-fire/util"
	"errors"
	"strings"
)

var (
	trilateration util.Trilateration
)

type locationService struct{}

type LocationService interface {
	GetLocation(distances ...float32) (x, y float32)
	CalculateCircle(sat_name string, distance float32) (float32, float32, error)
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

func (*locationService) CalculateCircle(sat_name string, distance float32) (float32, float32, error) {
	isSat, circle := validateCircle(sat_name, distance)

	if !isSat {
		return 0, 0, errors.New("Satellite invalid")
	}

	x := (-circle.AxisX) / 2
	y := (-circle.AxisY) / 2
	return float32(x), float32(y), nil
}

func validateCircle(sat_name string, distance float32) (bool, models.Circle) {
	distances := []float32{distance, distance, distance}
	circles := getCircles(distances)

	switch strings.ToLower(sat_name) {
	case "kenoby":
		return true, circles[0]
	case "skywalker":
		return true, circles[1]
	case "sato":
		return true, circles[2]
	default:
		return false, circles[0]
	}
}
