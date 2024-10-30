package courses

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"search-api/domain/courses"
)

// HTTPConfig contiene la configuración necesaria para conectarse a la API de Cursos
type HTTPConfig struct {
	Host string
	Port string
}

// HTTP representa el cliente HTTP para interactuar con la API de Cursos
type HTTP struct {
	baseURL func(courseID string) string
}

// NewHTTP crea un nuevo cliente HTTP para la API de Cursos
func NewHTTP(config HTTPConfig) HTTP {
	return HTTP{
		baseURL: func(courseID string) string {
			return fmt.Sprintf("http://%s:%s/courses/%s", config.Host, config.Port, courseID)
		},
	}
}

// GetCourseByID obtiene los detalles de un curso usando su ID
func (repository HTTP) GetCourseByID(ctx context.Context, id string) (courses.Course, error) {
	// Realiza la solicitud GET a la API de Cursos
	resp, err := http.Get(repository.baseURL(id))
	if err != nil {
		return courses.Course{}, fmt.Errorf("error al obtener el curso (%s): %w", id, err)
	}
	defer resp.Body.Close()

	// Verifica si la respuesta es exitosa
	if resp.StatusCode != http.StatusOK {
		return courses.Course{}, fmt.Errorf("error al obtener el curso (%s): código de estado %d", id, resp.StatusCode)
	}

	// Lee el cuerpo de la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return courses.Course{}, fmt.Errorf("error al leer la respuesta para el curso (%s): %w", id, err)
	}

	// Deserializa los datos del curso en la estructura Course
	var course courses.Course
	if err := json.Unmarshal(body, &course); err != nil {
		return courses.Course{}, fmt.Errorf("error al deserializar los datos del curso (%s): %w", id, err)
	}

	return course, nil
}
