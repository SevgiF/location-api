package domain

import (
	"time"
)

type Location struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Name        string    `json:"name" gorm:"type:varchar(128);not null" validate:"required,min=3,max=128"`
	Latitude    float64   `json:"latitude" gorm:"type:decimal(10,8);not null" validate:"required"`
	Longitude   float64   `json:"longitude" gorm:"type:decimal(11,8);not null" validate:"required"`
	MarkerColor string    `json:"marker_color" gorm:"type:varchar(6);not null" validate:"required,hexadecimal"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type TargetLocation struct {
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}
type LocationRouting struct {
	Name     string  `json:"name"`
	Distance float64 `json:"distance"`
}

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}
