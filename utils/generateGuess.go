package utils

import (
	"log"

	"github.com/LockBlock-dev/javascriptdle/entity"
	"github.com/LockBlock-dev/javascriptdle/repository"

	"math/rand"
)

var TodaysGuess *entity.GuessableItem

func GenerateTodaysGuess(guessableRepo repository.GuessableRepository) {
	items, err := guessableRepo.GetAll()
	if err != nil {
		TodaysGuess = nil
		return
	}

	TodaysGuess = &items[rand.Intn(len(items))]

	log.Printf("Today's guess is: %+v\n", *TodaysGuess)
}
