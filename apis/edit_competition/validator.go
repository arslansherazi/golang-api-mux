package edit_competition_api

import "mime/multipart"

type Validator struct {
	Title         string `validate:"min=10,max=100"`
	Description   string `validate:"min=100,max=5000"`
	Latitude      float64
	Longitude     float64
	Address       string `validate:"max=1000"`
	StartingDate  string `validate:"min=10,max=10"`
	StartingTime  string `validate:"min=5,max=5"`
	EndingTime    string
	DeletedImages []string
	AddedImages   []multipart.File
}
