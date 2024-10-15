package courses

type Course struct {
	ID           int64   `bson:"id"`
	Name         string  `bson:"name"`
	Description  string  `bson:"description"`
	Category     string  `bson:"category"`
	Duration     string  `bson:"duration"`
	InstructorID int64   `bson:"instructor_id"`
	ImageID      string  `bson:"image_id"`
	Capacity     int     `bson:"capacity"`
	Rating       float64 `bson:"rating"`
}

type CreateCourseRequest struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Category     string `json:"category" binding:"required"`
	Duration     string `json:"duration" binding:"required"`
	InstructorID int64  `json:"instructor_id" binding:"required"`
	ImageID      string `json:"image_id" binding:"required"`
	Capacity     int    `json:"capacity" binding:"required"`
}

type UpdateCourseRequest struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Category     string  `json:"category"`
	Duration     string  `json:"duration"`
	InstructorID int64   `json:"instructor_id"`
	ImageID      string  `json:"image_id"`
	Capacity     int     `json:"capacity"`
	Rating       float64 `json:"rating"`
}

type CourseResponse struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Category     string  `json:"category"`
	Duration     string  `json:"duration"`
	InstructorID int64   `json:"instructor_id"`
	ImageID      string  `json:"image_id"`
	Capacity     int     `json:"capacity"`
	Rating       float64 `json:"rating"`
}
