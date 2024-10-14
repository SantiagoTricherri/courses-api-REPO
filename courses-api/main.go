package main

import (
	"context"
	"log"
	"os"
	"time"

	commentsController "courses-api/controllers/comments"
	coursesController "courses-api/controllers/courses"
	commentsRepositories "courses-api/repositories"
	coursesRepositories "courses-api/repositories"
	coursesRouter "courses-api/router/courses"
	commentsServices "courses-api/services/comments"
	coursesServices "courses-api/services/courses"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Configuración del cliente MongoDB
	mongoConfig := coursesRepositories.MongoConfig{
		Host:       "localhost",
		Port:       "27017",
		Username:   "root",
		Password:   "root",
		Database:   "courses-api",
		Collection: "courses",
	}

	// Crear cliente de MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(
		options.Credential{
			Username: mongoConfig.Username,
			Password: mongoConfig.Password,
		},
	)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Error al crear el cliente de MongoDB: %v", err)
	}

	// Conectar con MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		log.Fatalf("Error al conectar con MongoDB: %v", err)
	}

	// Inicializar el contador de cursos
	coursesRepositories.InitializeCounter(client, mongoConfig.Database, mongoConfig.Collection)

	// Inicializar el contador de comentarios
	commentsRepositories.InitializeCommentCounter(client, mongoConfig.Database, "comments")

	// Crear instancias del repositorio, servicio y controlador
	courseRepo := coursesRepositories.NewMongo(mongoConfig)
	commentRepo := commentsRepositories.NewCommentsMongo(client, mongoConfig.Database, "comments")

	courseService := coursesServices.NewService(courseRepo, commentRepo)
	courseController := coursesController.NewController(courseService)

	commentsService := commentsServices.NewService(commentRepo, courseRepo)
	commentsController := commentsController.NewController(commentsService)

	// Configurar las rutas
	router := coursesRouter.SetupRouter(courseController, commentsController)

	// Leer el puerto desde las variables de entorno
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("Usando puerto por defecto:", port)
	}

	// Iniciar el servidor
	if err := router.Run(":" + port); err != nil {
		log.Fatal("No se pudo iniciar el servidor:", err)
	}
}
