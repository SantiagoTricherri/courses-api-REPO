package courses

import (
	"context"
	coursesDomain "courses-api/domain/courses"
	"fmt"
)

// Repository interface para las operaciones de curso
type Repository interface {
	CreateCourse(ctx context.Context, course coursesDomain.Course) (coursesDomain.Course, error)
	GetCourses(ctx context.Context) ([]coursesDomain.Course, error)
	GetCourseByID(ctx context.Context, id int64) (coursesDomain.Course, error)
	UpdateCourse(ctx context.Context, course coursesDomain.Course) (coursesDomain.Course, error)
	DeleteCourse(ctx context.Context, id int64) error
}

// CommentsRepository interface para las operaciones de comentarios
type CommentsRepository interface {
	DeleteCommentsByCourseID(ctx context.Context, courseID int64) error
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
	// Primero, eliminamos los comentarios asociados al curso
	err := s.commentsRepository.DeleteCommentsByCourseID(ctx, id)
	if err != nil {
		return fmt.Errorf("error al eliminar los comentarios del curso: %v", err)
	}

	// Luego, eliminamos el curso
	err = s.repository.DeleteCourse(ctx, id)
	if err != nil {
		return fmt.Errorf("error al eliminar el curso: %v", err)
	}

	return nil
}
