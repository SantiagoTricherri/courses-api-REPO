package files

import (
	"context"
	filesDTOs "courses-api/DTOs/files"
	filesDomain "courses-api/domain/files"
	"fmt"
)

// Interface del repositorio
type Repository interface {
	CreateFile(ctx context.Context, file filesDomain.File) (filesDomain.File, error)
	GetFilesByCourseID(ctx context.Context, courseID int64) ([]filesDomain.File, error)
}

// Servicio de archivos
type Service struct {
	repository Repository
}

// Constructor del servicio
func NewService(repo Repository) Service {
	return Service{repository: repo}
}

// Crear archivo
func (s Service) CreateFile(ctx context.Context, req filesDTOs.CreateFileRequestDTO) (filesDTOs.FileResponseDTO, error) {
	file := filesDomain.File{
		Name:     req.Name,
		Content:  []byte(req.Content),
		UserID:   req.UserID,
		CourseID: req.CourseID,
	}

	createdFile, err := s.repository.CreateFile(ctx, file)
	if err != nil {
		return filesDTOs.FileResponseDTO{}, fmt.Errorf("failed to create file: %v", err)
	}

	return filesDTOs.FileResponseDTO{
		ID:       createdFile.ID,
		Name:     createdFile.Name,
		Content:  createdFile.Content,
		UserID:   createdFile.UserID,
		CourseID: createdFile.CourseID,
	}, nil
}

// Obtener archivos por ID de curso
func (s Service) GetFilesByCourseID(ctx context.Context, courseID int64) ([]filesDTOs.FileResponseDTO, error) {
	filesData, err := s.repository.GetFilesByCourseID(ctx, courseID)
	if err != nil {
		return nil, fmt.Errorf("failed to get files: %v", err)
	}

	var response []filesDTOs.FileResponseDTO
	for _, f := range filesData {
		response = append(response, filesDTOs.FileResponseDTO{
			ID:       f.ID,
			Name:     f.Name,
			Content:  f.Content,
			UserID:   f.UserID,
			CourseID: f.CourseID,
		})
	}
	return response, nil
}
