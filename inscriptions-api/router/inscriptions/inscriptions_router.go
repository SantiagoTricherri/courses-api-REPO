package router

import (
	controller "inscriptions-api/controllers/inscriptions"

	"github.com/gin-gonic/gin"
)

// MapRoutes mapea las rutas del controlador de inscripciones.
func MapRoutes(r *gin.Engine, ctrl *controller.Controller) {
	r.POST("/inscriptions", ctrl.CreateInscription)
	r.GET("/inscriptions", ctrl.GetInscriptions)
	r.GET("/users/:userID/inscriptions", ctrl.GetInscriptionsByUser)
}
