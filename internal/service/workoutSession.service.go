package service

import (
	"github.com/aifia105/grind/internal/repository"
	"gorm.io/gorm"
)

type WorkoutSessionService struct {
	DB *gorm.DB
}

func NewWorkoutSessionService(db *gorm.DB) *WorkoutSessionService {
	return &WorkoutSessionService{DB: db}
}

func (s *WorkoutSessionService) CreateWorkoutSession(req *repository.WorkoutSession) (*repository.WorkoutSession, error) {
	workoutSession := &repository.WorkoutSession{
		ProgramID: req.ProgramID,
		Duration:  req.Duration,
		Notes:     req.Notes,
		Exercises: req.Exercises,
	}

	if err := s.DB.Create(workoutSession).Error; err != nil {
		return nil, err
	}
	return workoutSession, nil
}

func (s *WorkoutSessionService) GetWorkoutSessions(page, limit int) ([]repository.WorkoutSession, error) {
	workoutSessions := []repository.WorkoutSession{}
	offset := (page - 1) * limit

	if err := s.DB.Offset(offset).Limit(limit).Find(&workoutSessions).Error; err != nil {
		return nil, err
	}
	return workoutSessions, nil
}

func (s *WorkoutSessionService) GetWorkoutSession(id string) (*repository.WorkoutSession, error) {
	workoutSession := repository.WorkoutSession{}
	if err := s.DB.First(&workoutSession, id).Error; err != nil {
		return nil, err
	}
	return &workoutSession, nil
}

func (s *WorkoutSessionService) UpdateWorkoutSession(id string, req *repository.WorkoutSession) (*repository.WorkoutSession, error) {
	workoutSession := repository.WorkoutSession{}
	if err := s.DB.First(&workoutSession, id).Error; err != nil {
		return nil, err
	}
	workoutSession.ProgramID = req.ProgramID
	workoutSession.Duration = req.Duration
	workoutSession.Notes = req.Notes
	workoutSession.Exercises = req.Exercises

	if err := s.DB.Save(&workoutSession).Error; err != nil {
		return nil, err
	}
	return &workoutSession, nil
}

func (s *WorkoutSessionService) DeleteWorkoutSession(id string) error {
	workoutSession := repository.WorkoutSession{}
	if err := s.DB.First(&workoutSession, id).Error; err != nil {
		return err
	}
	if err := s.DB.Delete(&workoutSession).Error; err != nil {
		return err
	}
	return nil
}

func (s *WorkoutSessionService) AddExerciseSetToWorkoutSession(sessionId string, exerciseSet *repository.ExerciseSet) (*repository.WorkoutSession, error) {
	workoutSession := repository.WorkoutSession{}
	if err := s.DB.First(&workoutSession, sessionId).Error; err != nil {
		return nil, err
	}
	exerciseSet.WorkoutSessionID = workoutSession.ID
	if err := s.DB.Create(exerciseSet).Error; err != nil {
		return nil, err
	}
	workoutSession.Exercises = append(workoutSession.Exercises, *exerciseSet)
	if err := s.DB.Save(&workoutSession).Error; err != nil {
		return nil, err
	}
	return &workoutSession, nil
}
