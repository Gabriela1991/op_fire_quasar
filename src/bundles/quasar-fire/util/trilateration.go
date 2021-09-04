package util

import (
	"app/src/bundles/quasar-fire/models"
	"math"
)

type Trilateration struct{}

func (*Trilateration) CalculateTrilateration(circles []models.Circle) (float32, float32) {

	p2p1Distance := math.Pow(math.Pow(circles[1].AxisX-circles[0].AxisX, 2)+math.Pow(circles[1].AxisY-circles[0].AxisY, 2), 0.5)

	var ex models.Circle
	ex.AxisX = (circles[1].AxisX - circles[0].AxisX) / p2p1Distance
	ex.AxisY = (circles[1].AxisY - circles[0].AxisY) / p2p1Distance

	var aux models.Circle
	aux.AxisX = circles[2].AxisX - circles[0].AxisX
	aux.AxisY = circles[2].AxisY - circles[0].AxisY

	//componente X
	i := ex.AxisX*aux.AxisX + ex.AxisY*aux.AxisY

	//the unit vector in the y direction.
	var aux2 models.Circle
	aux2.AxisX = circles[2].AxisX - circles[0].AxisX - i*ex.AxisX
	aux2.AxisY = circles[2].AxisY - circles[0].AxisY - i*ex.AxisY

	var ey models.Circle
	ey.AxisX = aux2.AxisX / norm(aux2)
	ey.AxisY = aux2.AxisY / norm(aux2)

	//componente Y
	j := ey.AxisX*aux.AxisX + ey.AxisY*aux.AxisY

	//coordenadas
	x := (math.Pow(circles[0].Radio, 2) - math.Pow(circles[1].Radio, 2) + math.Pow(p2p1Distance, 2)) / (2 * p2p1Distance)
	y := (math.Pow(circles[0].Radio, 2)-math.Pow(circles[2].Radio, 2)+math.Pow(i, 2)+math.Pow(j, 2))/(2*j) - i*x/j

	//resultado
	finalX := circles[0].AxisX + x*ex.AxisX + y*ey.AxisX
	finalY := circles[0].AxisY + x*ex.AxisY + y*ey.AxisY
	return float32(finalX), float32(finalY)
}

func norm(p models.Circle) float64 {
	return math.Pow(math.Pow(p.AxisX, 2)+math.Pow(p.AxisY, 2), .5)
}
