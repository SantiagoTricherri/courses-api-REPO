package courses

import (
	"context"
	coursesDomain "courses-api/domain/courses"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Configuración para MongoDB
type MongoConfig struct {
	Host       string
	Port       string
	Username   string
	Password   string
	Database   string
	Collection string
}

// Estructura del repositorio Mongo
type Mongo struct {
	client     *mongo.Client
	database   string
	collection string
}

// Constante para la conexión
const (
	connectionURI = "mongodb://%s:%s"
)

// Nueva instancia de Mongo
func NewMongo(config MongoConfig) Mongo {
	credentials := options.Credential{
		Username: config.Username,
		Password: config.Password,
	}

	ctx := context.Background()
	uri := fmt.Sprintf(connectionURI, config.Host, config.Port)
	cfg := options.Client().ApplyURI(uri).SetAuth(credentials)

	client, err := mongo.Connect(ctx, cfg)
	if err != nil {
		log.Panicf("error connecting to mongo DB: %v", err)
	}

	return Mongo{
		client:     client,
		database:   config.Database,
		collection: config.Collection,
	}
}

// Métodos para operar con cursos
func (m Mongo) CreateCourse(ctx context.Context, course coursesDomain.Course) (coursesDomain.Course, error) {
	collection := m.client.Database(m.database).Collection(m.collection)
	_, err := collection.InsertOne(ctx, course)
	if err != nil {
		return coursesDomain.Course{}, fmt.Errorf("failed to insert course: %v", err)
	}
	return course, nil
}

func (m Mongo) GetCourses(ctx context.Context) ([]coursesDomain.Course, error) {
	var courses []coursesDomain.Course
	collection := m.client.Database(m.database).Collection(m.collection)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find courses: %v", err)
	}
	if err := cursor.All(ctx, &courses); err != nil {
		return nil, fmt.Errorf("failed to decode courses: %v", err)
	}
	return courses, nil
}

func (m Mongo) GetCourseByID(ctx context.Context, id int64) (coursesDomain.Course, error) {
	var course coursesDomain.Course
	collection := m.client.Database(m.database).Collection(m.collection)
	err := collection.FindOne(ctx, bson.M{"id": id}).Decode(&course)
	if err != nil {
		return coursesDomain.Course{}, fmt.Errorf("failed to find course: %v", err)
	}
	return course, nil
}

func (m Mongo) UpdateCourse(ctx context.Context, course coursesDomain.Course) (coursesDomain.Course, error) {
	collection := m.client.Database(m.database).Collection(m.collection)
	filter := bson.M{"id": course.ID}
	update := bson.M{"$set": course}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return coursesDomain.Course{}, fmt.Errorf("failed to update course: %v", err)
	}
	return course, nil
}

func (m Mongo) DeleteCourse(ctx context.Context, id int64) error {
	collection := m.client.Database(m.database).Collection(m.collection)
	_, err := collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return fmt.Errorf("failed to delete course: %v", err)
	}
	return nil
}

func (m Mongo) SearchCourses(ctx context.Context, query string) ([]coursesDomain.Course, error) {
	var courses []coursesDomain.Course
	collection := m.client.Database(m.database).Collection(m.collection)
	cursor, err := collection.Find(ctx, bson.M{
		"$or": []bson.M{
			{"name": bson.M{"$regex": query, "$options": "i"}},
			{"category": bson.M{"$regex": query, "$options": "i"}},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to search courses: %v", err)
	}
	if err := cursor.All(ctx, &courses); err != nil {
		return nil, fmt.Errorf("failed to decode courses: %v", err)
	}
	return courses, nil
}

func (m Mongo) GetCoursesByUserID(ctx context.Context, userID int64) ([]coursesDomain.Course, error) {
	// Implementar la lógica para obtener cursos por ID de usuario
	return []coursesDomain.Course{}, nil
}
