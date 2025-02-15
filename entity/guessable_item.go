package entity

type GuessableItem struct {
	Object          string `json:"object"`
	Method          string `json:"method"`
	ECMAversion     string `json:"ECMAversion"`
	ReturnType      string `json:"returnType"`
	MandatoryParams int    `json:"mandatoryParams"`
	OptionalParams  int    `json:"optionalParams"`
	TotalParams     int    `json:"totalParams"`
}

func NewGuessableItem() *GuessableItem {
	return &GuessableItem{}
}
