package storage

import (
	"fmt"
)

type Storage struct {
	items map[string]string
}

func NewStorage(items map[string]string) *Storage {
	return &Storage{
		items: items,
	}
}

type Storager interface {
	Get(id string) (string, error)
}

func (s *Storage) Get(id string) (string, error) {
	value, ok := s.items[id]
	if !ok {
		return "", fmt.Errorf("value %s not found", id)
	}

	return value, nil
}
