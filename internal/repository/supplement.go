package repository

import (
	"gorm.io/gorm"
)

type Supplement struct {
	gorm.Model
	Name      string  `json:"name" validate:"required"`
	Dosage    float64 `json:"dosage" validate:"required,min=0"`
	TimeTaken string  `json:"timeTaken" validate:"required"`
}
