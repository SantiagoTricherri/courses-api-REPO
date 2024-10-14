package comments

import (
	"context"
	"fmt"
	"time"

	commentsDTOs "courses-api/DTOs/comments"
	commentsDomain "courses-api/domain/comments"
	coursesDomain "courses-api/domain/courses"
)

type CommentsRepository interface {
	CreateComment(ctx context.Context, comment commentsDomain.Comment) (commentsDomain.Comment, error)
	GetCommentsByCourseID(ctx context.Context, courseID int64) ([]commentsDomain.Comment, error)
	DeleteCommentsByCourseID(ctx context.Context, courseID int64) error
}

type CoursesRepository interface {
	GetCourseByID(ctx context.Context, id int64) (coursesDomain.Course, error)
}

type Service struct {
	commentsRepository CommentsRepository
	coursesRepository  CoursesRepository
}

func NewService(commentsRepo CommentsRepository, coursesRepo CoursesRepository) Service {
	return Service{
		commentsRepository: commentsRepo,
		coursesRepository:  coursesRepo,
	}
}

func (s Service) CreateComment(ctx context.Context, courseID int64, req commentsDTOs.CreateCommentRequestDTO) (commentsDTOs.CommentResponseDTO, error) {
	// Verificar si el curso existe
	_, err := s.coursesRepository.GetCourseByID(ctx, courseID)
	if err != nil {
		return commentsDTOs.CommentResponseDTO{}, fmt.Errorf("el curso con ID %d no existe: %v", courseID, err)
	}

	comment := commentsDomain.Comment{
		CourseID:  courseID,
		UserID:    req.UserID,
		Content:   req.Content,
		Rating:    req.Rating,
		CreatedAt: time.Now().Unix(),
	}

	createdComment, err := s.commentsRepository.CreateComment(ctx, comment)
	if err != nil {
		return commentsDTOs.CommentResponseDTO{}, fmt.Errorf("error al crear el comentario: %v", err)
	}

	return commentsDTOs.CommentResponseDTO{
		ID:        createdComment.ID,
		CourseID:  createdComment.CourseID,
		UserID:    createdComment.UserID,
		Content:   createdComment.Content,
		Rating:    createdComment.Rating,
		CreatedAt: createdComment.CreatedAt,
	}, nil
}

func (s Service) GetCommentsByCourseID(ctx context.Context, courseID int64) ([]commentsDTOs.CommentResponseDTO, error) {
	// Verificar si el curso existe
	_, err := s.coursesRepository.GetCourseByID(ctx, courseID)
	if err != nil {
		return nil, fmt.Errorf("el curso con ID %d no existe: %v", courseID, err)
	}

	commentsDB, err := s.commentsRepository.GetCommentsByCourseID(ctx, courseID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener los comentarios: %v", err)
	}

	var commentsResponse []commentsDTOs.CommentResponseDTO
	for _, comment := range commentsDB {
		commentsResponse = append(commentsResponse, commentsDTOs.CommentResponseDTO{
			ID:        comment.ID,
			CourseID:  comment.CourseID,
			UserID:    comment.UserID,
			Content:   comment.Content,
			Rating:    comment.Rating,
			CreatedAt: comment.CreatedAt,
		})
	}

	return commentsResponse, nil
}
