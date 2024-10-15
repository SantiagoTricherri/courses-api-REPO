package courses

// CreateCourseRequest representa la solicitud para crear un curso
type CreateCourseRequestDTO struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Category     string `json:"category" binding:"required"`
	Duration     string `json:"duration" binding:"required"`
	InstructorID uint   `json:"instructor_id" binding:"required"`
	ImageID      string `json:"image_id" binding:"required"`
	Capacity     int    `json:"capacity" binding:"required"`
}

// UpdateCourseRequest representa la solicitud para actualizar un curso
type UpdateCourseRequestDTO struct {
	ID           uint    `json:"id" binding:"required"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Category     string  `json:"category"`
	Duration     string  `json:"duration"`
	InstructorID uint    `json:"instructor_id"`
	ImageID      string  `json:"image_id"`
	Capacity     int     `json:"capacity"`
	Rating       float64 `json:"rating"`
}

// CourseResponse representa la respuesta con los detalles de un curso
type CourseResponseDTO struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Category     string  `json:"category"`
	Duration     string  `json:"duration"`
	InstructorID uint    `json:"instructor_id"`
	ImageID      string  `json:"image_id"`
	Capacity     int     `json:"capacity"`
	Rating       float64 `json:"rating"`
}
