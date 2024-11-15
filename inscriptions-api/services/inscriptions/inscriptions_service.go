package service

import (
	"context"
	"errors"
	"fmt"
	"inscriptions-api/clients"
	domain "inscriptions-api/domain/inscriptions"
)

type Repository interface {
	CreateInscription(ctx context.Context, userID, courseID uint) (*domain.Inscription, error)
	GetInscriptions(ctx context.Context) ([]domain.Inscription, error)
	GetInscriptionsByUser(ctx context.Context, userID uint) ([]domain.Inscription, error)
	GetInscriptionsByCourse(ctx context.Context, courseID uint) ([]domain.Inscription, error)
}

type Service struct {
	repository Repository
	httpClient *clients.HTTPClient
}

func NewService(repository Repository, httpClient *clients.HTTPClient) *Service {
	return &Service{repository: repository, httpClient: httpClient}
}

func (s *Service) CreateInscription(ctx context.Context, userID, courseID uint) (*domain.Inscription, error) {
	// Verificar si el usuario existe (usando la implementación temporal)
	if err := s.httpClient.CheckUserExists(userID); err != nil {
		return nil, fmt.Errorf("failed to verify user: %v", err)
	}

	// Verificar si el curso existe y obtener su capacidad
	course, err := s.httpClient.GetCourseDetails(courseID)
	if err != nil {
		return nil, fmt.Errorf("failed to verify course: %v", err)
	}

	// Obtener el número actual de inscripciones para el curso
	currentInscriptions, err := s.repository.GetInscriptionsByCourse(ctx, courseID)
	if err != nil {
		return nil, fmt.Errorf("failed to get current inscriptions: %v", err)
	}

	// Verificar si hay cupos disponibles
	if len(currentInscriptions) >= course.Capacity {
		return nil, errors.New("course is at full capacity")
	}

	// Crear la inscripción
	inscription, err := s.repository.CreateInscription(ctx, userID, courseID)
	if err != nil {
		return nil, fmt.Errorf("failed to create inscription: %v", err)
	}

	return inscription, nil
}

func (s *Service) GetInscriptions(ctx context.Context) ([]domain.Inscription, error) {
	return s.repository.GetInscriptions(ctx)
}

func (s *Service) GetInscriptionsByUser(ctx context.Context, userID uint) ([]domain.Inscription, error) {
	return s.repository.GetInscriptionsByUser(ctx, userID)
}

func (s *Service) GetInscriptionsByCourse(ctx context.Context, courseID uint) ([]domain.Inscription, error) {
	// Verificar si el curso existe
	if err := s.httpClient.CheckCourseExists(courseID); err != nil {
		return nil, fmt.Errorf("failed to verify course: %v", err)
	}

	return s.repository.GetInscriptionsByCourse(ctx, courseID)
}
