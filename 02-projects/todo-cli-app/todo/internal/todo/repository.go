package todo

import (
	"encoding/json"
	"os"
)

type Repository struct {
	Filename string
}

func NewRepository(filename string) *Repository {
	return &Repository{Filename: filename}
}

func (r *Repository) Load() ([]Item, error) {
	if _, err := os.Stat(r.Filename); os.IsNotExist(err) {
		return []Item{}, nil
	}

	data, err := os.ReadFile(r.Filename)
	if err != nil {
		return nil, err
	}

	var items []Item
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *Repository) Save(items []Item) error {
	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.Filename, data, 0644)
}
