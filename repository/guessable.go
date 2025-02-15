package repository

import "github.com/LockBlock-dev/javascriptdle/entity"

type GuessableRepository interface {
	GetAll() ([]entity.GuessableItem, error)
	FindByMethod(string) (*entity.GuessableItem, error)
}
