package service

import (
	err_template "github.com/blattaria7/go-template/internal/errors"

	"go.uber.org/zap"

	"github.com/blattaria7/go-template/internal/repository"
)

type Service struct {
	repo   repository.Repositorier
	logger *zap.Logger
}

func NewService(repo repository.Repositorier, logger *zap.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}

type Servicer interface {
	Get(id string) (string, error)
}

func (s *Service) Get(id string) (string, error) {
	result, err := s.repo.Get(id)

	switch err {
	case nil:
		if result == "" {
			return "", err_template.ErrValueNotFound
		}

		return result, nil

	default:
		return "", err_template.ErrInternalServerError
	}

}
