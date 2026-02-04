package todo

import (
	"fmt"
	"time"
)

type Manager struct {
	repo  *Repository
	items []Item
}

func NewManager(filename string) (*Manager, error) {
	repo := NewRepository(filename)
	items, err := repo.Load()
	if err != nil {
		return nil, err
	}
	return &Manager{
		repo:  repo,
		items: items,
	}, nil
}

func (m *Manager) Add(task string) error {
	newItem := Item {
		ID: len(m.items) + 1,
		Task: task,
		Done: false,
		CreatedAt: time.Now(),
	}
	m.items = append(m.items, newItem)
	return m.repo.Save(m.items)
}

func (m *Manager) GetAll() []Item {
	return m.items
}

func (m *Manager) MarkAsDone(id int) error {
	found := false
	for i, item := range m.items {
		if item.ID == id {
			m.items[i].Done = true
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("item with ID %d not found", id)
	}

	return m.repo.Save(m.items)
}
