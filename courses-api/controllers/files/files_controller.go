package files

/*
import (
	"context"
	"encoding/base64"
	"net/http"
	"strconv"

	"courses-api/domain/files" // Importar el domain de files

	"github.com/gin-gonic/gin"
)

// Interface del servicio de archivos
type Service interface {
	CreateFile(ctx context.Context, req files.CreateFileRequestDTO) (files.FileResponseDTO, error)
	GetFilesByCourseID(ctx context.Context, courseID int64) ([]files.FileResponseDTO, error)
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
	var req files.CreateFileRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido: " + err.Error()})
		return
	}

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
	courseID, err := strconv.ParseInt(ctx.Param("courseID"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	files, err := ctrl.service.GetFilesByCourseID(ctx.Request.Context(), courseID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener archivos: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, files)
}
*/
