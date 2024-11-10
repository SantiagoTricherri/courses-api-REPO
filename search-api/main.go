package main

import (
	"log"
	"search-api/clients/queues"
	searchController "search-api/controllers/search"
	"search-api/repositories/courses"
	searchService "search-api/services/search"

	"github.com/gin-gonic/gin"
)

func main() {
	// Configuración de SolR
	solrRepo := courses.NewSolr(courses.SolrConfig{
		Host:       "solr",    // SolR host
		Port:       "8983",    // SolR port
		Collection: "courses", // Nombre de la colección en SolR
	})

	// Configuración de RabbitMQ
	eventsQueue := queues.NewRabbit(queues.RabbitConfig{
		Host:      "rabbitmq",
		Port:      "5672",
		Username:  "root",
		Password:  "root",
		QueueName: "courses_queue",
	})

	// Configuración del cliente HTTP para la API de Cursos
	coursesAPI := courses.NewHTTP(courses.HTTPConfig{
		Host: "courses-api",
		Port: "8081",
	})

	// Inicialización del servicio de búsqueda
	searchService := searchService.NewService(solrRepo, coursesAPI)

	// Inicialización del controlador de búsqueda
	searchController := searchController.NewController(searchService)

	// Lanzar el consumidor de RabbitMQ
	if err := eventsQueue.StartConsumer(searchService.HandleCourseUpdate); err != nil {
		log.Fatalf("Error al ejecutar el consumidor: %v", err)
	}

	// Configuración del router con Gin
	router := gin.Default()
	router.GET("/search", searchController.Search)

	// Ejecutar la API en el puerto 8082
	if err := router.Run(":8082"); err != nil {
		log.Fatalf("Error al ejecutar la aplicación: %v", err)
	}
}
