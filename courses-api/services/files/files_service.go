package files

/*
import (
	"context"
	"fmt"

	"courses-api/domain/files" // Importar el domain de files
)

// Interface del repositorio
type Repository interface {
	CreateFile(ctx context.Context, file files.File) (files.File, error)
	GetFilesByCourseID(ctx context.Context, courseID int64) ([]files.File, error)
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
func (s Service) CreateFile(ctx context.Context, req files.CreateFileRequestDTO) (files.FileResponseDTO, error) {
	file := files.File{
		Name:     req.Name,
		Content:  []byte(req.Content),
		UserID:   req.UserID,
		CourseID: req.CourseID,
	}

	createdFile, err := s.repository.CreateFile(ctx, file)
	if err != nil {
		return files.FileResponseDTO{}, fmt.Errorf("failed to create file: %v", err)
	}

	return files.FileResponseDTO{
		ID:       createdFile.ID,
		Name:     createdFile.Name,
		Content:  createdFile.Content,
		UserID:   createdFile.UserID,
		CourseID: createdFile.CourseID,
	}, nil
}

// Obtener archivos por ID de curso
func (s Service) GetFilesByCourseID(ctx context.Context, courseID int64) ([]files.FileResponseDTO, error) {
	filesData, err := s.repository.GetFilesByCourseID(ctx, courseID)
	if err != nil {
		return nil, fmt.Errorf("failed to get files: %v", err)
	}

	var response []files.FileResponseDTO
	for _, f := range filesData {
		response = append(response, files.FileResponseDTO{
			ID:       f.ID,
			Name:     f.Name,
			Content:  f.Content,
			UserID:   f.UserID,
			CourseID: f.CourseID,
		})
	}
	return response, nil
}
*/
