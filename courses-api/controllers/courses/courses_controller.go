package courses

import (
	"context"
	coursesDomain "courses-api/domain/courses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Interface que define los métodos del servicio
type Service interface {
	CreateCourse(ctx context.Context, req coursesDomain.CreateCourseRequest) (coursesDomain.CourseResponse, error)
	GetCourses(ctx context.Context) ([]coursesDomain.CourseResponse, error)
	GetCourseByID(ctx context.Context, id int64) (coursesDomain.CourseResponse, error)
	UpdateCourse(ctx context.Context, req coursesDomain.UpdateCourseRequest) (coursesDomain.CourseResponse, error)
	DeleteCourse(ctx context.Context, id int64) error
}

// Controller estructura del controlador
type Controller struct {
	service Service
}

// NewController constructor del controlador
func NewController(service Service) Controller {
	return Controller{service: service}
}

// Crear curso
func (ctrl Controller) CreateCourse(ctx *gin.Context) {
	var req coursesDomain.CreateCourseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido: " + err.Error()})
		return
	}
	course, err := ctrl.service.CreateCourse(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear curso: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, course)
}

// Obtener todos los cursos
func (ctrl Controller) GetCourses(ctx *gin.Context) {
	courses, err := ctrl.service.GetCourses(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al listar cursos: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, courses)
}

// Obtener curso por ID
func (ctrl Controller) GetCourseByID(ctx *gin.Context) {
	courseID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	course, err := ctrl.service.GetCourseByID(ctx.Request.Context(), courseID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener curso: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, course)
}

// Actualizar curso
func (ctrl Controller) UpdateCourse(ctx *gin.Context) {
	var req coursesDomain.UpdateCourseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido: " + err.Error()})
		return
	}
	course, err := ctrl.service.UpdateCourse(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar curso: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, course)
}

// Eliminar curso
func (ctrl Controller) DeleteCourse(ctx *gin.Context) {
	courseID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if err := ctrl.service.DeleteCourse(ctx.Request.Context(), courseID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar curso: " + err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}
