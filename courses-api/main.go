package main

import (
	"context"
	"log"
	"os"
	"time"

	coursesController "courses-api/controllers/courses"
	coursesRepositories "courses-api/repositories"
	coursesRouter "courses-api/router/courses"
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

	// Inicializar el contador basado en el último ID en la colección
	coursesRepositories.InitializeCounter(client, mongoConfig.Database, mongoConfig.Collection)

	// Crear instancias del repositorio, servicio y controlador
	courseRepo := coursesRepositories.NewMongo(mongoConfig)
	courseService := coursesServices.NewService(courseRepo)
	courseController := coursesController.NewController(courseService)

	// Configurar las rutas
	router := coursesRouter.SetupRouter(courseController)

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
