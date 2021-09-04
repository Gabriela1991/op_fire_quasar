package models

type Satelite struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

type Satelites struct {
	Satellites []Satelite `json:"satellites"`
}
