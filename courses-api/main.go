package main

import (
	coursesController "courses-api/controllers/courses" // Ajusta el path si es necesario
	coursesRepositories "courses-api/repositories"      // Ajusta el path si es necesario
	coursesRouter "courses-api/router/courses"          // Ajusta el path si es necesario
	coursesServices "courses-api/services/courses"      // Ajusta el path si es necesario
	"log"
	"os"
)

// MongoConfig define la configuración para la conexión a MongoDB
type MongoConfig = coursesRepositories.MongoConfig

func main() {
	// Configuración de la conexión a MongoDB
	mongoConfig := MongoConfig{
		Host:       "localhost",   // Cambia esto si es necesario
		Port:       "27017",       // Cambia esto si es necesario
		Username:   "root",        // Cambia esto por tu usuario
		Password:   "root",        // Cambia esto por tu contraseña
		Database:   "courses-api", // Nombre de tu base de datos
		Collection: "courses",     // Nombre de tu colección
	}

	// Crear una nueva conexión Mongo
	courseRepository := coursesRepositories.NewMongo(mongoConfig)

	// Crear instancia del servicio de cursos
	courseService := coursesServices.NewService(courseRepository)

	// Crear instancia del controlador con el servicio inyectado
	courseController := coursesController.NewController(courseService)

	// Configurar el router con el controlador
	r := coursesRouter.SetupRouter(courseController)

	// Leer el puerto desde la variable de entorno o usar el puerto por defecto 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("Usando puerto por defecto:", port)
	}

	// Iniciar el servidor
	if err := r.Run(":" + port); err != nil {
		log.Fatal("No se pudo iniciar el servidor:", err)
	}
}
