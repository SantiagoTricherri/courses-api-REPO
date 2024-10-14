package courses

type Course struct {
	ID           int64  `bson:"id"`
	Name         string `bson:"name"`
	Description  string `bson:"description"`
	Category     string `bson:"category"`
	Duration     string `bson:"duration"`
	InstructorID int64  `bson:"instructor_id"`
}

type CreateCourseRequest struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Category     string `json:"category" binding:"required"`
	Duration     string `json:"duration" binding:"required"`
	InstructorID int64  `json:"instructor_id" binding:"required"`
}

type UpdateCourseRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Category     string `json:"category"`
	Duration     string `json:"duration"`
	InstructorID int64  `json:"instructor_id"`
}

type CourseResponse struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Category     string `json:"category"`
	Duration     string `json:"duration"`
	InstructorID int64  `json:"instructor_id"`
}
