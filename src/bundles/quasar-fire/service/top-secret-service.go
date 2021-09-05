package service

import (
	"app/src/bundles/quasar-fire/models"
	"errors"
)

var (
	locationSrv LocationService = NewLocationService()
	messageSrv  MessageService  = NewMessageService()
)

type topSecretService struct{}

type TopSecretService interface {
	TopSecret(satellites models.Satelites) (models.TopSecretResponse, error)
	TopSecretSplit(satellites models.TopSecretSplitReq, satName string) (models.TopSecretResponse, error)
}

func NewTopSecretService() TopSecretService {
	return &topSecretService{}
}

func (*topSecretService) TopSecret(satellites models.Satelites) (models.TopSecretResponse, error) {
	var topSecretResp models.TopSecretResponse
	if dist := len(satellites.Satellites); dist > 3 || dist < 3 {
		return topSecretResp, errors.New("Distances array out of bounds, expected 3 distances")
	}

	x, y := locationSrv.GetLocation(satellites.Satellites[0].Distance, satellites.Satellites[1].Distance, satellites.Satellites[2].Distance)
	msg := messageSrv.GetMessage(satellites.Satellites[0].Message, satellites.Satellites[1].Message, satellites.Satellites[2].Message)

	topSecretResp.Message = msg
	topSecretResp.Position = models.Positions{
		AxisX: x,
		AxisY: y,
	}
	return topSecretResp, nil
}

func (*topSecretService) TopSecretSplit(satellite models.TopSecretSplitReq, satName string) (models.TopSecretResponse, error) {
	var topSecretResp models.TopSecretResponse

	x, y, err := locationSrv.CalculateCircle(satName, satellite.Distance)
	msg := messageSrv.GetMessage(satellite.Message)

	if err != nil {
		return topSecretResp, err
	}

	if msg == "" {
		return topSecretResp, errors.New("There is not enough information to get the message")
	}

	topSecretResp.Message = msg
	topSecretResp.Position = models.Positions{
		AxisX: x,
		AxisY: y,
	}

	return topSecretResp, nil
}
