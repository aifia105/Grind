package repository

import (
	"time"

	"gorm.io/gorm"
)

type BodyMeasurement struct {
	gorm.Model
	Date    time.Time `json:"date" validate:"required"`
	Weight  float64   `json:"weight" validate:"required,min=0"`
	BodyFat float64   `json:"bodyFat" validate:"required,min=0"`
	Waist   float64   `json:"waist" validate:"required,min=0"`
	Arms    float64   `json:"arms" validate:"required,min=0"`
	Thighs  float64   `json:"thighs" validate:"required,min=0"`
	Photos  []string  `json:"photos" gorm:"serializer:json"`
}
