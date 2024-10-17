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

	// Inicializar el contador de archivos
	filesRepositories.InitializeFileCounter(client, mongoConfig.Database, "files")

	// Configurar RabbitMQ
	rabbitConfig := rabbit.RabbitConfig{
		Username:  "tu_usuario",
		Password:  "tu_contraseña",
		Host:      "localhost",
		Port:      "5672",
		QueueName: "courses_queue",
	}
	rabbitQueue := rabbit.NewRabbit(rabbitConfig)

	// Crear instancias del repositorio
	courseRepo := coursesRepositories.NewMongo(mongoConfig)
	commentRepo := commentsRepositories.NewCommentsMongo(client, mongoConfig.Database, "comments")

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
	fileRepo := filesRepositories.NewMongo(client, mongoConfig.Database, "files")
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
