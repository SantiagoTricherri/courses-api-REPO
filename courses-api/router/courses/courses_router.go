package router

// SIN MIDDLWARES, VER!!

import (
	"courses-api/controllers/comments" // Ajusta el path si es necesario
	"courses-api/controllers/courses"  // Ajusta el path si es necesario

	"github.com/gin-gonic/gin"
)

// Función para configurar las rutas
func SetupRouter(courseController courses.Controller, commentController comments.Controller) *gin.Engine {
	r := gin.Default() // Sin middleware adicional

	// Rutas para cursos
	coursesGroup := r.Group("/courses")
	{
		coursesGroup.POST("", courseController.CreateCourse)       // Crear curso
		coursesGroup.GET("", courseController.GetCourses)          // Obtener todos los cursos
		coursesGroup.GET("/:id", courseController.GetCourseByID)   // Obtener curso por ID
		coursesGroup.PUT("/:id", courseController.UpdateCourse)    // Actualizar curso
		coursesGroup.DELETE("/:id", courseController.DeleteCourse) // Eliminar curso
		coursesGroup.POST("/:id/comments", commentController.AddCommentToCourse)
		coursesGroup.GET("/:id/comments", commentController.GetCommentsByCourseID)

	}

	return r
}
