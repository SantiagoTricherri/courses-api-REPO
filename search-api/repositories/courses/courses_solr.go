package courses

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "log"
    "search-api/domain/courses"
    "github.com/stevenferrer/solr-go"
)

// SolrConfig contiene la configuración necesaria para conectarse a SolR
type SolrConfig struct {
    Host       string // Dirección del host de SolR
    Port       string // Puerto de SolR
    Collection string // Nombre de la colección en SolR
}

// Solr representa el cliente de SolR
type Solr struct {
    Client     *solr.JSONClient
    Collection string
}

// NewSolr inicializa un nuevo cliente de SolR
func NewSolr(config SolrConfig) Solr {
    baseURL := fmt.Sprintf("http://%s:%s", config.Host, config.Port)
    client := solr.NewJSONClient(baseURL)

    return Solr{
        Client:     client,
        Collection: config.Collection,
    }
}


// Index añade un nuevo curso al índice de SolR
func (s Solr) Index(ctx context.Context, course courses.CourseUpdate) (string, error) {
    doc := map[string]interface{}{
        "id":          course.CourseID,
        "name":        course.Name,
        "category":    course.Category,
        "description": course.Description,
    }

    indexRequest := map[string]interface{}{
        "add": []interface{}{doc},
    }

    body, err := json.Marshal(indexRequest)
    if err != nil {
        return "", fmt.Errorf("error al serializar el documento del curso: %w", err)
    }

    resp, err := s.Client.Update(ctx, s.Collection, solr.JSON, bytes.NewReader(body))
    if err != nil {
        return "", fmt.Errorf("error al indexar el curso: %w", err)
    }
    if resp.Error != nil {
        return "", fmt.Errorf("error al indexar el curso en SolR: %v", resp.Error)
    }

    if err := s.Client.Commit(ctx, s.Collection); err != nil {
        return "", fmt.Errorf("error al confirmar los cambios en SolR: %w", err)
    }

    return course.CourseID, nil
}


// Update modifica un curso existente en el índice de SolR
func (s Solr) Update(ctx context.Context, course courses.CourseUpdate) error {
    doc := map[string]interface{}{
        "id":          course.CourseID,
        "name":        course.Name,
        "category":    course.Category,
        "description": course.Description,
    }

    updateRequest := map[string]interface{}{
        "add": []interface{}{doc},
    }

    body, err := json.Marshal(updateRequest)
    if err != nil {
        return fmt.Errorf("error al serializar el documento del curso: %w", err)
    }

    resp, err := s.Client.Update(ctx, s.Collection, solr.JSON, bytes.NewReader(body))
    if err != nil {
        return fmt.Errorf("error al actualizar el curso: %w", err)
    }
    if resp.Error != nil {
        return fmt.Errorf("error al actualizar el curso en SolR: %v", resp.Error)
    }

    if err := s.Client.Commit(ctx, s.Collection); err != nil {
        return fmt.Errorf("error al confirmar los cambios en SolR: %w", err)
    }

    return nil
}


// Delete elimina un curso del índice de SolR
func (s Solr) Delete(ctx context.Context, id string) error {
    deleteRequest := map[string]interface{}{
        "delete": map[string]interface{}{
            "id": id,
        },
    }

    body, err := json.Marshal(deleteRequest)
    if err != nil {
        return fmt.Errorf("error al serializar el documento del curso para eliminar: %w", err)
    }

    resp, err := s.Client.Update(ctx, s.Collection, solr.JSON, bytes.NewReader(body))
    if err != nil {
        return fmt.Errorf("error al eliminar el curso: %w", err)
    }
    if resp.Error != nil {
        return fmt.Errorf("error al eliminar el curso en SolR: %v", resp.Error)
    }

    if err := s.Client.Commit(ctx, s.Collection); err != nil {
        return fmt.Errorf("error al confirmar la eliminación en SolR: %w", err)
    }

    return nil
}
