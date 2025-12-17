package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/kwame-Owusu/lista/internal/models"
)

func SaveTodos(todos []models.Todo, filename string) error {
	if !strings.HasSuffix(filename, ".json") {
		return fmt.Errorf("invalid filename: must end with .json")
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("creating file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(todos); err != nil {
		return fmt.Errorf("encoding todos: %w", err)
	}

	return nil
}

func LoadTodos(filename string) ([]models.Todo, error) {
	if !strings.HasSuffix(filename, ".json") {
		return nil, fmt.Errorf("invalid filename: must end with .json")
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	var todos []models.Todo
	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, fmt.Errorf("unmarshaling todos: %w", err)
	}

	return todos, nil
}
