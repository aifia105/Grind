package repository

import (
	"time"

	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Name            string  `json:"name" validate:"required"`
	CaloriesPer100g int     `json:"caloriesPer100g"`
	ProteinPer100g  float64 `json:"proteinPer100g"`
	CarbsPer100g    float64 `json:"carbsPer100g"`
	FatPer100g      float64 `json:"fatPer100g"`
}

type Meal struct {
	gorm.Model
	Date          time.Time `json:"date" validate:"required"`
	MealType      string    `json:"mealType" validate:"required"`
	Foods         []Food    `json:"foods" validate:"required" gorm:"serializer:json"`
	TotalCalories int       `json:"totalCalories" validate:"required,min=0"`
}

type MealFood struct {
	gorm.Model
	MealID   uint    `json:"mealId" validate:"required"`
	FoodID   uint    `json:"foodId" validate:"required"`
	Quantity float64 `json:"quantity" validate:"required,min=0"`
}
