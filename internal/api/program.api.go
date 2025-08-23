package api

import (
	"net/http"
	"strconv"

	"github.com/aifia105/grind/internal/repository"
	"github.com/aifia105/grind/internal/service"
	"github.com/gin-gonic/gin"
)

type ProgramHandler struct {
	programService *service.ProgramService
}

func NewProgramHandler(programService *service.ProgramService) *ProgramHandler {
	return &ProgramHandler{programService: programService}
}

func (h *ProgramHandler) CreateProgram(ctx *gin.Context) {
	req := repository.Program{}
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	program, err := h.programService.CreateProgram(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, program)
}

func (h *ProgramHandler) GetPrograms(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)

	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)

	if err != nil || limit < 1 {
		limit = 1
	}

	users, err := h.programService.GetPrograms(page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (h *ProgramHandler) GetProgramById(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid program ID"})
		return
	}

	program, err := h.programService.GetProgramById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, program)
}

func (h *ProgramHandler) GetProgramWithWorkoutSessions(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid program ID"})
		return
	}

	program, workoutSessions, err := h.programService.GetProgramWithWorkoutSessions(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := gin.H{
		"program":         program,
		"workoutSessions": workoutSessions,
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *ProgramHandler) UpdateProgram(ctx *gin.Context) {
	req := repository.Program{}
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid program ID"})
		return
	}

	program, err := h.programService.UpdateProgram(id, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, program)
}

func (h *ProgramHandler) DeleteProgram(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid program ID"})
		return
	}

	if err := h.programService.DeleteProgram(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
func (h *ProgramHandler) AddWorkoutSessionToProgram(ctx *gin.Context) {
	programId := ctx.Param("id")
	if programId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid program ID"})
		return
	}

	var req struct {
		WorkoutSessionID string `json:"workoutSessionId" validate:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	workoutSessionID, err := strconv.Atoi(req.WorkoutSessionID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	program, err := h.programService.AddWorkoutSessionToProgram(programId, workoutSessionID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, program)
}
