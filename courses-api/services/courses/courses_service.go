package courses

import (
	"context"
	commentsDomain "courses-api/domain/comments"
	coursesDomain "courses-api/domain/courses"
	"fmt"
	"time"
)

// Repository interface para las operaciones de curso
type Repository interface {
	CreateCourse(ctx context.Context, course coursesDomain.Course) (coursesDomain.Course, error)
	GetCourses(ctx context.Context) ([]coursesDomain.Course, error)
	GetCourseByID(ctx context.Context, id int64) (coursesDomain.Course, error)
	UpdateCourse(ctx context.Context, course coursesDomain.Course) (coursesDomain.Course, error)
	DeleteCourse(ctx context.Context, id int64) error
}

type CommentsRepository interface {
	CreateComment(ctx context.Context, comment commentsDomain.Comment) (commentsDomain.Comment, error)
	GetCommentsByCourseID(ctx context.Context, courseID int64) ([]commentsDomain.Comment, error)
}

// Service estructura para el servicio de cursos
type Service struct {
	repository         Repository
	commentsRepository CommentsRepository
}

// NewService constructor para el servicio de cursos
func NewService(repository Repository, commentsRepository CommentsRepository) Service {
	return Service{
		repository:         repository,
		commentsRepository: commentsRepository,
	}
}

func (s Service) CreateCourse(ctx context.Context, req coursesDomain.CreateCourseRequest) (coursesDomain.CourseResponse, error) {
	course := coursesDomain.Course{
		Name:         req.Name,
		Description:  req.Description,
		Category:     req.Category,
		Duration:     req.Duration,
		InstructorID: req.InstructorID,
	}

	createdCourse, err := s.repository.CreateCourse(ctx, course)
	if err != nil {
		return coursesDomain.CourseResponse{}, fmt.Errorf("failed to create course: %v", err)
	}

	return coursesDomain.CourseResponse{
		ID:           createdCourse.ID,
		Name:         createdCourse.Name,
		Description:  createdCourse.Description,
		Category:     createdCourse.Category,
		Duration:     createdCourse.Duration,
		InstructorID: createdCourse.InstructorID,
	}, nil
}

func (s Service) GetCourses(ctx context.Context) ([]coursesDomain.CourseResponse, error) {
	courses, err := s.repository.GetCourses(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get courses: %v", err)
	}

	var coursesDTO []coursesDomain.CourseResponse
	for _, course := range courses {
		coursesDTO = append(coursesDTO, coursesDomain.CourseResponse{
			ID:           course.ID,
			Name:         course.Name,
			Description:  course.Description,
			Category:     course.Category,
			Duration:     course.Duration,
			InstructorID: course.InstructorID,
		})
	}

	return coursesDTO, nil
}

func (s Service) GetCourseByID(ctx context.Context, id int64) (coursesDomain.CourseResponse, error) {
	course, err := s.repository.GetCourseByID(ctx, id)
	if err != nil {
		return coursesDomain.CourseResponse{}, fmt.Errorf("failed to get course: %v", err)
	}

	return coursesDomain.CourseResponse{
		ID:           course.ID,
		Name:         course.Name,
		Description:  course.Description,
		Category:     course.Category,
		Duration:     course.Duration,
		InstructorID: course.InstructorID,
	}, nil
}

func (s Service) UpdateCourse(ctx context.Context, id int64, req coursesDomain.UpdateCourseRequest) (coursesDomain.CourseResponse, error) {
	course, err := s.repository.GetCourseByID(ctx, id)
	if err != nil {
		return coursesDomain.CourseResponse{}, fmt.Errorf("course not found: %v", err)
	}

	if req.Name != "" {
		course.Name = req.Name
	}
	if req.Description != "" {
		course.Description = req.Description
	}
	if req.Category != "" {
		course.Category = req.Category
	}
	if req.Duration != "" {
		course.Duration = req.Duration
	}
	if req.InstructorID != 0 {
		course.InstructorID = req.InstructorID
	}

	updatedCourse, err := s.repository.UpdateCourse(ctx, course)
	if err != nil {
		return coursesDomain.CourseResponse{}, fmt.Errorf("failed to update course: %v", err)
	}

	return coursesDomain.CourseResponse{
		ID:           updatedCourse.ID,
		Name:         updatedCourse.Name,
		Description:  updatedCourse.Description,
		Category:     updatedCourse.Category,
		Duration:     updatedCourse.Duration,
		InstructorID: updatedCourse.InstructorID,
	}, nil
}

func (s Service) DeleteCourse(ctx context.Context, id int64) error {
	return s.repository.DeleteCourse(ctx, id)
}

func (s Service) CreateComment(ctx context.Context, courseID int64, req commentsDomain.CreateCommentRequest) (commentsDomain.CommentResponse, error) {
	// Primero, verificamos si el curso existe
	_, err := s.repository.GetCourseByID(ctx, courseID)
	if err != nil {
		return commentsDomain.CommentResponse{}, fmt.Errorf("el curso con ID %d no existe: %v", courseID, err)
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
		return commentsDomain.CommentResponse{}, fmt.Errorf("error al crear el comentario: %v", err)
	}

	return commentsDomain.CommentResponse{
		ID:        createdComment.ID,
		CourseID:  createdComment.CourseID,
		UserID:    createdComment.UserID,
		Content:   createdComment.Content,
		Rating:    createdComment.Rating,
		CreatedAt: createdComment.CreatedAt,
	}, nil
}

func (s Service) GetCommentsByCourseID(ctx context.Context, courseID int64) ([]commentsDomain.CommentResponse, error) {
	commentsDB, err := s.commentsRepository.GetCommentsByCourseID(ctx, courseID)
	if err != nil {
		return nil, fmt.Errorf("failed to get comments: %v", err)
	}

	var commentsResponse []commentsDomain.CommentResponse
	for _, comment := range commentsDB {
		commentsResponse = append(commentsResponse, commentsDomain.CommentResponse{
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
