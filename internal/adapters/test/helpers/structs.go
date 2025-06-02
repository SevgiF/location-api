package helpers

import "time"

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}
type RequestAndResponses struct {
	Request      any
	ResponseBody any
	StatusCode   int
}

type ValidationResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Location struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	MarkerColor string    `json:"marker_color"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
