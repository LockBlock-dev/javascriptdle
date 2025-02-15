package data

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/LockBlock-dev/javascriptdle/entity"
)

func LoadDataFromJson(path string) ([]entity.GuessableItem, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("JSON file is empty")
	}

	var items []entity.GuessableItem

	err = json.Unmarshal(data, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}
