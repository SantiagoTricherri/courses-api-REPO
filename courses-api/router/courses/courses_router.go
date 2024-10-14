package router

// SIN MIDDLWARES, VER!!

import (
	"courses-api/controllers/comments"
	"courses-api/controllers/courses"
	"courses-api/controllers/files"

	"github.com/gin-gonic/gin"
)

// Función para configurar las rutas
func SetupRouter(courseController courses.Controller, commentController comments.Controller, fileController files.Controller) *gin.Engine {
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
		coursesGroup.POST("/:id/files", fileController.CreateFile)
		coursesGroup.GET("/:id/files", fileController.GetFilesByCourseID)
	}

	return r
}
