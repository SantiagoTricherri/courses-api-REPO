package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"search-api/clients/queues"
	"search-api/controllers/search"
	"search-api/repositories/courses"
	"search-api/services/search"
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
		QueueName: "courses-news",
	})

	// Configuración del cliente HTTP para la API de Cursos
	coursesAPI := courses.NewHTTP(courses.HTTPConfig{
		Host: "courses-api",
		Port: "8081",
	})

	// Inicialización del servicio de búsqueda
	searchService := search.NewService(solrRepo, coursesAPI)

	// Inicialización del controlador de búsqueda
	searchController := search.NewController(searchService)

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
