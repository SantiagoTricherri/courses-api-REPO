package search

import (
    "context"
    "fmt"
    "log"
    "search-api/domain/courses"
)

// Repository define las operaciones necesarias en el índice de SolR
type Repository interface {
    Index(ctx context.Context, course courses.CourseUpdate) (string, error)
    Update(ctx context.Context, course courses.CourseUpdate) error
    Delete(ctx context.Context, id string) error
    Search(ctx context.Context, query string, limit int, offset int) ([]courses.CourseUpdate, error)
}

// Service representa el servicio de búsqueda
type Service struct {
    repository Repository
}

// NewService crea una nueva instancia del servicio de búsqueda
func NewService(repository Repository) Service {
    return Service{
        repository: repository,
    }
}

// HandleCourseUpdate procesa las actualizaciones de cursos recibidas desde RabbitMQ
func (service Service) HandleCourseUpdate(courseUpdate courses.CourseUpdate) {
    ctx := context.Background()

    switch courseUpdate.Operation {
    case "CREATE":
        // Indexar el nuevo curso en SolR
        if _, err := service.repository.Index(ctx, courseUpdate); err != nil {
            log.Printf("Error al indexar el curso (%s): %v", courseUpdate.CourseID, err)
        } else {
            log.Printf("Curso indexado exitosamente: %s", courseUpdate.CourseID)
        }

    case "UPDATE":
        // Actualizar el curso existente en SolR
        if err := service.repository.Update(ctx, courseUpdate); err != nil {
            log.Printf("Error al actualizar el curso (%s): %v", courseUpdate.CourseID, err)
        } else {
            log.Printf("Curso actualizado exitosamente: %s", courseUpdate.CourseID)
        }

    case "DELETE":
        // Eliminar el curso del índice de SolR
        if err := service.repository.Delete(ctx, courseUpdate.CourseID); err != nil {
            log.Printf("Error al eliminar el curso (%s): %v", courseUpdate.CourseID, err)
        } else {
            log.Printf("Curso eliminado exitosamente: %s", courseUpdate.CourseID)
        }

    default:
        log.Printf("Operación desconocida: %s", courseUpdate.Operation)
    }
}
