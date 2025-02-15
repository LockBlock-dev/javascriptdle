package repository

import (
	"fmt"
	"log"
	"sync"

	"github.com/LockBlock-dev/javascriptdle/data"
	"github.com/LockBlock-dev/javascriptdle/entity"
)

type MemoryGuessableRepository struct {
	guessableItems map[string]entity.GuessableItem
	lock           *sync.RWMutex
}

func NewMemoryGuessableRepository() *MemoryGuessableRepository {
	items, err := data.LoadDataFromJson("./data.json")
	if err != nil {
		log.Panicf("Error loading data from JSON file: %v\n", err)
	}

	repo := MemoryGuessableRepository{
		guessableItems: map[string]entity.GuessableItem{},
		lock:           &sync.RWMutex{},
	}

	for _, item := range items {
		repo.guessableItems[fmt.Sprintf("%s.%s", item.Object, item.Method)] = item
	}

	return &repo
}

func (r *MemoryGuessableRepository) GetAll() ([]entity.GuessableItem, error) {
	values := []entity.GuessableItem{}

	for _, value := range r.guessableItems {
		values = append(values, value)
	}

	return values, nil
}

func (r *MemoryGuessableRepository) FindByMethod(userGuess string) (*entity.GuessableItem, error) {
	guess, ok := r.guessableItems[userGuess]

	if !ok {
		return nil, fmt.Errorf("not found!")
	}

	return &guess, nil
}
