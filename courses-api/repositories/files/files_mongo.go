package files

/*
import (
	"context"
	"fmt"

	"courses-api/domain/files" // Importar el domain de files

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repositorio MongoDB para archivos
type Mongo struct {
	client     *mongo.Client
	database   string
	collection string
}

// Constructor del repositorio Mongo
func NewMongo(client *mongo.Client, db, collection string) Mongo {
	return Mongo{
		client:     client,
		database:   db,
		collection: collection,
	}
}

// Crear archivo
func (m Mongo) CreateFile(ctx context.Context, file files.File) (files.File, error) {
	_, err := m.client.Database(m.database).Collection(m.collection).InsertOne(ctx, file)
	if err != nil {
		return files.File{}, fmt.Errorf("failed to insert file: %v", err)
	}
	return file, nil
}

// Obtener archivos por ID de curso
func (m Mongo) GetFilesByCourseID(ctx context.Context, courseID int64) ([]files.File, error) {
	var filesData []files.File
	cursor, err := m.client.Database(m.database).Collection(m.collection).Find(ctx, bson.M{"course_id": courseID})
	if err != nil {
		return nil, fmt.Errorf("failed to get files: %v", err)
	}
	if err := cursor.All(ctx, &filesData); err != nil {
		return nil, fmt.Errorf("failed to decode files: %v", err)
	}
	return filesData, nil
}
*/
