package repository

import (
	"gorm.io/gorm"
)

type Exercise struct {
	gorm.Model
	Name         string   `json:"name" validate:"required"`
	MuscleGroups []string `json:"muscleGroups" validate:"required" gorm:"serializer:json"`
	Equipment    string   `json:"equipment" validate:"required"`
	ImageURL     string   `json:"imageUrl" validate:"required"`
	Difficulty   int      `json:"difficulty" validate:"required,min=1"`
}
