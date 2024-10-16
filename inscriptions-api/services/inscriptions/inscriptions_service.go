package service

import (
	"context"
	domain "inscriptions-api/domain/inscriptions"
)

type Repository interface {
	CreateInscription(ctx context.Context, userID, courseID uint) (*domain.Inscription, error)
	GetInscriptions(ctx context.Context) ([]domain.Inscription, error)
	GetInscriptionsByUser(ctx context.Context, userID uint) ([]domain.Inscription, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) CreateInscription(ctx context.Context, userID, courseID uint) (*domain.Inscription, error) {
	return s.repository.CreateInscription(ctx, userID, courseID)
}

func (s *Service) GetInscriptions(ctx context.Context) ([]domain.Inscription, error) {
	return s.repository.GetInscriptions(ctx)
}

func (s *Service) GetInscriptionsByUser(ctx context.Context, userID uint) ([]domain.Inscription, error) {
	return s.repository.GetInscriptionsByUser(ctx, userID)
}
