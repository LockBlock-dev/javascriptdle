package controller

import (
	"net/http"

	"github.com/LockBlock-dev/javascriptdle/components"
	"github.com/LockBlock-dev/javascriptdle/entity"
	"github.com/LockBlock-dev/javascriptdle/repository"
	"github.com/LockBlock-dev/javascriptdle/utils"
	"github.com/LockBlock-dev/javascriptdle/views"
)

type GuessController struct {
	guessableRepo repository.GuessableRepository
}

func NewGuessController(guessableRepo repository.GuessableRepository) *GuessController {
	return &GuessController{
		guessableRepo: guessableRepo,
	}
}

func (c *GuessController) Home(w http.ResponseWriter, r *http.Request) {
	items, err := c.guessableRepo.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	views.Home(items).Render(r.Context(), w)
}

func (c *GuessController) Guess(w http.ResponseWriter, r *http.Request) {
	userGuess := r.FormValue("guess")
	guess, err := c.guessableRepo.FindByMethod(userGuess)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := []entity.GuessResult{
		{
			Label:     guess.Object,
			IsCorrect: utils.TodaysGuess.Object == guess.Object,
		},
		{
			Label:     guess.Method,
			IsCorrect: utils.TodaysGuess.Method == guess.Method,
		},
		{
			Label:     guess.ECMAversion,
			IsCorrect: utils.TodaysGuess.ECMAversion == guess.ECMAversion,
		},
		{
			Label:     guess.ReturnType,
			IsCorrect: utils.TodaysGuess.ReturnType == guess.ReturnType,
		},
	}

	components.GuessResult(result).Render(r.Context(), w)
}
