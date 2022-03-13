package services

import (
	"go.uber.org/zap"

	err_template "github.com/blattaria7/go-template/internal/errors"
)

type Service struct {
	repo   Repository
	logger *zap.Logger
}

func NewService(repo Repository, logger *zap.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}

type Repository interface {
	Get(id string) (string, error)
}

func (s *Service) Get(id string) (string, error) {
	result, err := s.repo.Get(id)
	if err != nil {
		return "", err_template.ErrInternalServerError
	}

	if result == "" {
		return "", err_template.ErrValueNotFound
	}

	return result, nil
}
