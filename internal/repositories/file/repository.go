package file

import (
	"fmt"

	"go.uber.org/zap"
)

type Repository struct {
	items  map[string]string
	logger *zap.Logger
}

func NewRepository(items map[string]string, logger *zap.Logger) *Repository {
	return &Repository{
		items:  items,
		logger: logger,
	}
}

func (r *Repository) Get(id string) (string, error) {
	value, ok := r.items[id]
	if !ok {
		return "", fmt.Errorf("value %s not found", id)
	}

	return value, nil
}
