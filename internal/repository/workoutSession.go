package repository

import (
	"gorm.io/gorm"
)

type WorkoutSession struct {
	gorm.Model
	ProgramID uint          `json:"programId" validate:"required"`
	Duration  int           `json:"duration"`
	Exercises []ExerciseSet `json:"exercises" validate:"required" gorm:"serializer:json"`
	Notes     string        `json:"notes"`
}

type ExerciseSet struct {
	gorm.Model
	WorkoutSessionID uint    `json:"workoutSessionId" validate:"required"`
	ExerciseID       uint    `json:"exerciseId" validate:"required"`
	SetNumber        int     `json:"setNumber" validate:"required,min=1"`
	Reps             int     `json:"reps" validate:"required,min=1"`
	Weight           float64 `json:"weight" validate:"required,min=0"`
	RestTime         int     `json:"restTime" validate:"required,min=0"`
}
