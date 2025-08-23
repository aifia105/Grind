package repository

import (
	"time"

	"gorm.io/gorm"
)

type Nutrition struct {
	gorm.Model
	Date               time.Time `json:"date" validate:"required"`
	GoalCaloriesPerDay int       `json:"goalCaloriesPerDay" validate:"required,min=0"`
	CaloriesPerDay     int       `json:"caloriesPerDay"`
	GoalProteinPerDay  int       `json:"goalProteinPerDay" validate:"required,min=0"`
	ProteinPerDay      int       `json:"proteinPerDay"`
	GoalCarbsPerDay    int       `json:"goalCarbsPerDay" validate:"required,min=0"`
	CarbsPerDay        int       `json:"carbsPerDay"`
	GoalFatPerDay      int       `json:"goalFatPerDay" validate:"required,min=0"`
	FatPerDay          int       `json:"fatPerDay"`
	TotalCalories      int       `json:"totalCalories" validate:"required,min=0"`
	TotalProtein       int       `json:"totalProtein" validate:"required,min=0"`
	TotalCarbs         int       `json:"totalCarbs" validate:"required,min=0"`
	TotalFat           int       `json:"totalFat" validate:"required,min=0"`
	WaterIntake        float64   `json:"waterIntake"`
	GoalWaterIntake    float64   `json:"goalWaterIntake" validate:"required,min=0"`
}
