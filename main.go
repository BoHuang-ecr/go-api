package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/BoHuang-ecr/go-api/database"
	"github.com/BoHuang-ecr/go-api/models"
	"github.com/BoHuang-ecr/go-api/repos"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := database.NewConnection(config)

	if err != nil {
		log.Fatal("could not load the database")
	}
	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	repo := repos.Repository{
		DB: db,
	}
	app := fiber.New()
	repo.SetupRoutes(app)
	app.Listen(":8080")

}
