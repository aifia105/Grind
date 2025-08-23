package service

import (
	"github.com/aifia105/grind/internal/repository"
	"gorm.io/gorm"
)

type ProgramService struct {
	DB *gorm.DB
}

func NewProgramService(db *gorm.DB) *ProgramService {
	return &ProgramService{DB: db}
}

func (s *ProgramService) CreateProgram(req *repository.Program) (*repository.Program, error) {
	program := &repository.Program{
		StartDate:         req.StartDate,
		Description:       req.Description,
		Duration:          req.Duration,
		Split:             req.Split,
		WorkoutSessionIDs: req.WorkoutSessionIDs,
	}
	if err := s.DB.Create(program).Error; err != nil {
		return nil, err
	}
	return program, nil
}

func (s *ProgramService) GetPrograms(page, limit int) ([]repository.Program, error) {
	programs := []repository.Program{}
	offset := (page - 1) * limit

	if err := s.DB.Offset(offset).Limit(limit).Find(&programs).Error; err != nil {
		return nil, err
	}
	return programs, nil
}

func (s *ProgramService) GetProgramById(id string) (*repository.Program, error) {
	program := repository.Program{}
	if err := s.DB.First(&program, id).Error; err != nil {
		return nil, err
	}
	return &program, nil
}

func (s *ProgramService) GetProgramWithWorkoutSessions(id string) (*repository.Program, []repository.WorkoutSession, error) {
	program := repository.Program{}
	if err := s.DB.First(&program, id).Error; err != nil {
		return nil, nil, err
	}

	var workoutSessions []repository.WorkoutSession
	if len(program.WorkoutSessionIDs) > 0 {
		if err := s.DB.Where("id IN ?", program.WorkoutSessionIDs).Find(&workoutSessions).Error; err != nil {
			return &program, nil, err
		}
	}

	return &program, workoutSessions, nil
}

func (s *ProgramService) UpdateProgram(id string, req *repository.Program) (*repository.Program, error) {
	program := repository.Program{}
	if err := s.DB.First(&program, id).Error; err != nil {
		return nil, err
	}
	program.StartDate = req.StartDate
	program.Description = req.Description
	program.Duration = req.Duration
	program.Split = req.Split
	program.WorkoutSessionIDs = req.WorkoutSessionIDs
	if err := s.DB.Save(&program).Error; err != nil {
		return nil, err
	}
	return &program, nil
}

func (s *ProgramService) DeleteProgram(id string) error {
	program := repository.Program{}
	if err := s.DB.First(&program, id).Error; err != nil {
		return err
	}
	if err := s.DB.Delete(&program).Error; err != nil {
		return err
	}
	return nil
}

func (s *ProgramService) AddWorkoutSessionToProgram(programId string, workoutSessionId int) (*repository.Program, error) {
	var program repository.Program
	if err := s.DB.First(&program, programId).Error; err != nil {
		return nil, err
	}
	for _, id := range program.WorkoutSessionIDs {
		if id == workoutSessionId {
			return &program, nil
		}
	}
	program.WorkoutSessionIDs = append(program.WorkoutSessionIDs, workoutSessionId)

	if err := s.DB.Save(&program).Error; err != nil {
		return nil, err
	}

	return &program, nil
}
