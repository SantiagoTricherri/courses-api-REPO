package files

import (
	"context"
	"encoding/base64"
	"net/http"
	"strconv"

	filesDTOs "courses-api/DTOs/files" // Importamos los DTOs de archivos

	"github.com/gin-gonic/gin"
)

// Interface del servicio de archivos
type Service interface {
	CreateFile(ctx context.Context, req filesDTOs.CreateFileRequestDTO) (filesDTOs.FileResponseDTO, error)
	GetFilesByCourseID(ctx context.Context, courseID int64) ([]filesDTOs.FileResponseDTO, error)
}

// Controller de archivos
type Controller struct {
	service Service
}

// Constructor del controlador
func NewController(service Service) Controller {
	return Controller{service: service}
}

// Crear un archivo
func (ctrl Controller) CreateFile(ctx *gin.Context) {
	courseID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de curso inválido en la URL"})
		return
	}

	var req filesDTOs.CreateFileRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido: " + err.Error()})
		return
	}

	// Asignar el courseID de la URL al DTO
	req.CourseID = courseID

	decodedContent, err := base64.StdEncoding.DecodeString(req.Content)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar: " + err.Error()})
		return
	}
	req.Content = string(decodedContent)

	file, err := ctrl.service.CreateFile(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al subir archivo: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, file)
}

// Obtener archivos por ID de curso
func (ctrl Controller) GetFilesByCourseID(ctx *gin.Context) {
	courseID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de curso inválido en la URL"})
		return
	}

	files, err := ctrl.service.GetFilesByCourseID(ctx.Request.Context(), courseID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener archivos: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, files)
}
