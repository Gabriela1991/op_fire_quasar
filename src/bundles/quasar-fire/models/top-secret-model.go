package models

type TopSecretResponse struct {
	Position Positions `json:"position"`
	// Message received
	Message string `json:"message"`
}

type Positions struct {
	// X axix coordinate
	AxisX float32 `json:"x"`
	// Y axix coordinate
	AxisY float32 `json:"y"`
}

type TopSecretSplitReq struct {
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}
