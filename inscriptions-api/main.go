package main

import (
	dao "inscriptions-api/DAOs/inscriptions"
	controller "inscriptions-api/controllers/inscriptions"
	repositories "inscriptions-api/repositories/inscriptions"
	router "inscriptions-api/router/inscriptions"
	service "inscriptions-api/services/inscriptions"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Conexión a la base de datos.
	db, err := repositories.Connect()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Inicialización de DAO, servicio y controlador.
	inscriptionDAO := dao.NewInscriptionDAO(db)
	inscriptionService := service.NewService(inscriptionDAO)
	inscriptionController := controller.NewController(inscriptionService)

	// Configuración del router.
	r := gin.Default()
	router.MapRoutes(r, inscriptionController)

	// Iniciar el servidor.
	log.Println("Server running on port 8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
