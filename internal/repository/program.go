package repository

import (
	"gorm.io/gorm"
)

type Program struct {
	gorm.Model
	StartDate         string `json:"startDate" validate:"required"`
	Description       string `json:"description"`
	Duration          int    `json:"duration" validate:"required,min=0"`
	Split             string `json:"split" validate:"required"`
	WorkoutSessionIDs []int  `json:"workoutSessionIds,omitempty" gorm:"serializer:json"`
}
