package main

import (
	"log"
	"os"

	"github.com/aifia105/grind/internal/config"
	"github.com/aifia105/grind/internal/repository"
	"github.com/aifia105/grind/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env vars")
	}

	port := os.Getenv("PORT")

	config.ConnectDB()

	config.DB.AutoMigrate(
		&repository.BodyMeasurement{},
		&repository.Exercise{},
		&repository.ExerciseSet{},
		&repository.Food{},
		&repository.Meal{},
		&repository.MealFood{},
		&repository.Nutrition{},
		&repository.Program{},
		&repository.Sleep{},
		&repository.Supplement{},
		&repository.WorkoutSession{})

	r := routes.SetupRoutes(config.DB)

	if err := r.Run(":" + port); err != nil {
		panic(err)
	}

	log.Println("Server starting on :" + port)

}
