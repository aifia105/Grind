package api

import (
	"net/http"
	"strconv"

	"github.com/aifia105/grind/internal/repository"
	"github.com/aifia105/grind/internal/service"
	"github.com/gin-gonic/gin"
)

type WorkoutSessionHandler struct {
	workoutSessionService *service.WorkoutSessionService
}

func NewWorkoutSessionHandler(workoutSessionService *service.WorkoutSessionService) *WorkoutSessionHandler {
	return &WorkoutSessionHandler{workoutSessionService: workoutSessionService}
}

func (h *WorkoutSessionHandler) CreateWorkoutSession(ctx *gin.Context) {
	req := repository.WorkoutSession{}
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	workoutSession, err := h.workoutSessionService.CreateWorkoutSession(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, workoutSession)
}

func (h *WorkoutSessionHandler) GetWorkoutSessions(ctx *gin.Context) {
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

	workoutSessions, err := h.workoutSessionService.GetWorkoutSessions(page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, workoutSessions)
}

func (h *WorkoutSessionHandler) GetWorkoutSession(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}
	workoutSession, err := h.workoutSessionService.GetWorkoutSession(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, workoutSession)
}

func (h *WorkoutSessionHandler) UpdateWorkoutSession(ctx *gin.Context) {
	req := repository.WorkoutSession{}
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	workoutSession, err := h.workoutSessionService.UpdateWorkoutSession(id, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, workoutSession)
}

func (h *WorkoutSessionHandler) DeleteWorkoutSession(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}
	if err := h.workoutSessionService.DeleteWorkoutSession(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func (h *WorkoutSessionHandler) AddExerciseSetToWorkoutSession(ctx *gin.Context) {
	workoutSessionId := ctx.Param("id")
	if workoutSessionId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid workout session ID"})
		return
	}

	req := repository.ExerciseSet{}
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exerciseSet, err := h.workoutSessionService.AddExerciseSetToWorkoutSession(workoutSessionId, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, exerciseSet)
}
