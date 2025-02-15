package entity

type GuessResult struct {
	Label     string `json:"label"`
	IsCorrect bool   `json:"isCorrect"`
}

func NewGuessResult() *GuessResult {
	return &GuessResult{}
}
