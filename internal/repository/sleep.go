package repository

import (
	"errors"

	"gorm.io/gorm"
)

type Sleep struct {
	gorm.Model
	Date       string  `json:"date" validate:"required"`
	HoursSlept float64 `json:"hoursSlept" validate:"required,min=0,max=24"`
	Quality    int     `json:"quality" validate:"required,min=1,max=5"`
}

func (s *Sleep) ValidateQuality() error {
	if s.Quality < 1 || s.Quality > 5 {
		return errors.New("sleep quality must be between 1 and 5")
	}
	return nil
}

func (s *Sleep) Validate() error {
	if err := s.ValidateQuality(); err != nil {
		return err
	}
	if s.HoursSlept < 0 || s.HoursSlept > 24 {
		return errors.New("hours slept must be between 0 and 24")
	}
	if s.Date == "" {
		return errors.New("date is required")
	}
	return nil
}
