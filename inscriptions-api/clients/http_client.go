package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type HTTPClient struct {
	usersAPIURL   string
	coursesAPIURL string
	client        *http.Client
}

func NewHTTPClient(usersAPIURL, coursesAPIURL string) *HTTPClient {
	return &HTTPClient{
		usersAPIURL:   usersAPIURL,
		coursesAPIURL: coursesAPIURL,
		client:        &http.Client{},
	}
}

func (c *HTTPClient) CheckUserExists(userID uint) error {
	// Implementación temporal: aceptar usuarios con ID par
	if userID%2 == 0 {
		return nil
	}
	return errors.New("user does not exist")
}

func (c *HTTPClient) CheckCourseExists(courseID uint) error {
	resp, err := c.client.Get(fmt.Sprintf("%s/courses/%d", c.coursesAPIURL, courseID))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("course with ID %d not found", courseID)
	}

	return nil
}

type CourseDetails struct {
	ID       uint `json:"id"`
	Capacity int  `json:"capacity"`
	// Add other fields if needed
}

func (c *HTTPClient) GetCourseDetails(courseID uint) (*CourseDetails, error) {
	resp, err := c.client.Get(fmt.Sprintf("%s/courses/%d", c.coursesAPIURL, courseID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("course with ID %d not found", courseID)
	}

	var course CourseDetails
	if err := json.NewDecoder(resp.Body).Decode(&course); err != nil {
		return nil, fmt.Errorf("error decoding course details: %v", err)
	}

	return &course, nil
}
