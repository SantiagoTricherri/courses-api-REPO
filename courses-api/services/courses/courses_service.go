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
	SearchCourses(ctx context.Context, query string) ([]coursesDomain.Course, error)
	GetCoursesByUserID(ctx context.Context, userID int64) ([]coursesDomain.Course, error)
}

// Service estructura para el servicio de cursos
type Service struct {
	repository Repository
}

// NewService constructor para el servicio de cursos
func NewService(repository Repository) Service {
	return Service{
		repository: repository,
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

func (s Service) UpdateCourse(ctx context.Context, req coursesDomain.UpdateCourseRequest) (coursesDomain.CourseResponse, error) {
	course, err := s.repository.GetCourseByID(ctx, req.ID)
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

func (s Service) SearchCourses(ctx context.Context, query string) ([]coursesDomain.CourseResponse, error) {
	courses, err := s.repository.SearchCourses(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to search courses: %v", err)
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

func (s Service) GetCoursesByUserID(ctx context.Context, userID uint) ([]coursesDomain.CourseResponse, error) {
	courses, err := s.repository.GetCoursesByUserID(ctx, int64(userID))
	if err != nil {
		return nil, fmt.Errorf("failed to get courses by user ID: %v", err)
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
