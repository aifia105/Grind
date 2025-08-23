package service

import (
	"github.com/aifia105/grind/internal/repository"
	"gorm.io/gorm"
)

type ExerciseService struct {
	DB *gorm.DB
}

func NewExerciseService(db *gorm.DB) *ExerciseService {
	return &ExerciseService{DB: db}
}

func (s *ExerciseService) CreateExercise(req *repository.Exercise) (*repository.Exercise, error) {
	exercise := &repository.Exercise{
		Name:         req.Name,
		MuscleGroups: req.MuscleGroups,
		Equipment:    req.Equipment,
		ImageURL:     req.ImageURL,
		Difficulty:   req.Difficulty,
	}
	if err := s.DB.Create(exercise).Error; err != nil {
		return nil, err
	}
	return exercise, nil
}

func (s *ExerciseService) GetExercises(page, limit int) ([]repository.Exercise, error) {
	exercise := []repository.Exercise{}
	offset := (page - 1) * limit

	if err := s.DB.Offset(offset).Limit(limit).Find(&exercise).Error; err != nil {
		return nil, err
	}
	return exercise, nil
}

func (s *ExerciseService) GetExercise(id string) (*repository.Exercise, error) {
	exercise := repository.Exercise{}
	if err := s.DB.First(&exercise, id).Error; err != nil {
		return nil, err
	}
	return &exercise, nil
}

func (s *ExerciseService) UpdateExercise(id string, req *repository.Exercise) (*repository.Exercise, error) {
	exercise := repository.Exercise{}
	if err := s.DB.First(&exercise, id).Error; err != nil {
		return nil, err
	}
	exercise.Name = req.Name
	exercise.MuscleGroups = req.MuscleGroups
	exercise.Equipment = req.Equipment
	exercise.ImageURL = req.ImageURL
	exercise.Difficulty = req.Difficulty

	if err := s.DB.Save(&exercise).Error; err != nil {
		return nil, err
	}
	return &exercise, nil
}

func (s *ExerciseService) DeleteExercise(id string) error {
	exercise := repository.Exercise{}
	if err := s.DB.First(&exercise, id).Error; err != nil {
		return err
	}
	if err := s.DB.Delete(&exercise).Error; err != nil {
		return err
	}
	return nil
}
