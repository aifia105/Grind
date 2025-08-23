package routes

import (
	"github.com/aifia105/grind/internal/api"
	"github.com/aifia105/grind/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.SetTrustedProxies(nil)

	programService := service.NewProgramService(db)
	programHandler := api.NewProgramHandler(programService)

	workoutSessionService := service.NewWorkoutSessionService(db)
	workoutSessionHandler := api.NewWorkoutSessionHandler(workoutSessionService)

	exerciseService := service.NewExerciseService(db)
	exerciseHandler := api.NewExerciseHandler(exerciseService)

	api := router.Group("/api/v1")
	{
		programs := api.Group("/programs")
		{
			programs.POST("/", programHandler.CreateProgram)
			programs.GET("/", programHandler.GetPrograms)
			programs.GET("/:id", programHandler.GetProgramById)
			programs.GET("/:id/full", programHandler.GetProgramWithWorkoutSessions)
			programs.POST("/:id/workoutSessions", programHandler.AddWorkoutSessionToProgram)
			programs.PUT("/:id", programHandler.UpdateProgram)
			programs.DELETE("/:id", programHandler.DeleteProgram)
		}
		workoutSession := api.Group("/workoutSession")
		{
			workoutSession.POST("/", workoutSessionHandler.CreateWorkoutSession)
			workoutSession.GET("/", workoutSessionHandler.GetWorkoutSessions)
			workoutSession.GET("/:id", workoutSessionHandler.GetWorkoutSession)
			workoutSession.PUT("/:id", workoutSessionHandler.UpdateWorkoutSession)
			workoutSession.DELETE("/:id", workoutSessionHandler.DeleteWorkoutSession)
			workoutSession.POST("/:id/exerciseSets", workoutSessionHandler.AddExerciseSetToWorkoutSession)
		}
		exercises := api.Group("/exercises")
		{
			exercises.POST("/", exerciseHandler.CreateExercise)
			exercises.GET("/", exerciseHandler.GetExercises)
			exercises.GET("/:id", exerciseHandler.GetExercise)
			exercises.PUT("/:id", exerciseHandler.UpdateExercise)
			exercises.DELETE("/:id", exerciseHandler.DeleteExercise)
		}

	}
	return router
}
