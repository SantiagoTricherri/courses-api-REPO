package comments

import (
	"context"
	"net/http"
	"strconv"

	commentsDTOs "courses-api/DTOs/comments"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateComment(ctx context.Context, courseID int64, req commentsDTOs.CreateCommentRequestDTO) (commentsDTOs.CommentResponseDTO, error)
	GetCommentsByCourseID(ctx context.Context, courseID int64) ([]commentsDTOs.CommentResponseDTO, error)
}

type Controller struct {
	service Service
}

func NewController(service Service) Controller {
	return Controller{service: service}
}

func (ctrl Controller) AddCommentToCourse(ctx *gin.Context) {
	courseID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de curso inválido"})
		return
	}

	var req commentsDTOs.CreateCommentRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido: " + err.Error()})
		return
	}

	comment, err := ctrl.service.CreateComment(ctx.Request.Context(), courseID, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al agregar comentario: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, comment)
}

func (ctrl Controller) GetCommentsByCourseID(ctx *gin.Context) {
	courseID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	comments, err := ctrl.service.GetCommentsByCourseID(ctx.Request.Context(), courseID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener comentarios: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, comments)
}