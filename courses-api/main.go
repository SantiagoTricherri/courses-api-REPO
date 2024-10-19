package main

import (
	"context"
	"log"
	"os"
	"time"

	"courses-api/clients/rabbit"
	commentsController "courses-api/controllers/comments"
	coursesController "courses-api/controllers/courses"
	filesController "courses-api/controllers/files"
	commentsRepositories "courses-api/repositories/comments"
	coursesRepositories "courses-api/repositories/courses"
	filesRepositories "courses-api/repositories/files"
	coursesRouter "courses-api/router/courses"
	commentsServices "courses-api/services/comments"
	coursesServices "courses-api/services/courses"
	filesServices "courses-api/services/files"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Configuración del cliente MongoDB
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	// Crear cliente de MongoDB
	clientOptions := options.Client().ApplyURI(mongoURI)
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
	coursesRepositories.InitializeCounter(client, "courses-api", "courses")

	// Inicializar el contador de comentarios
	commentsRepositories.InitializeCommentCounter(client, "courses-api", "comments")

	// Inicializar el contador de archivos
	filesRepositories.InitializeFileCounter(client, "courses-api", "files")

	// Configurar RabbitMQ
	rabbitURI := os.Getenv("RABBITMQ_URI")
	if rabbitURI == "" {
		rabbitURI = "amqp://guest:guest@localhost:5672/"
	}
	rabbitConfig := rabbit.RabbitConfig{
		URI:       rabbitURI,
		QueueName: "courses_queue",
	}
	rabbitQueue := rabbit.NewRabbit(rabbitConfig)

	// Crear instancias del repositorio
	courseRepo := coursesRepositories.NewMongo(coursesRepositories.MongoConfig{
		Host:       "mongodb", // Nombre del servicio en docker-compose
		Port:       "27017",
		Username:   "root",
		Password:   "root",
		Database:   "courses-api",
		Collection: "courses",
	})
	commentRepo := commentsRepositories.NewCommentsMongo(client, "courses-api", "comments")

	// Crear el servicio de cursos
	courseService := coursesServices.NewService(
		courseRepo,
		commentRepo,
		rabbitQueue,
	)

	// Crear el controlador de cursos
	courseController := coursesController.NewController(courseService)

	// Crear instancias para comentarios
	commentService := commentsServices.NewService(commentRepo, courseRepo)
	commentController := commentsController.NewController(commentService)

	// Crear instancias para archivos
	fileRepo := filesRepositories.NewMongo(client, "courses-api", "files")
	fileService := filesServices.NewService(fileRepo)
	fileController := filesController.NewController(fileService)

	// Configurar las rutas
	router := coursesRouter.SetupRouter(courseController, commentController, fileController)

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
