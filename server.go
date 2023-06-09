package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"wallester.com/assignment/models"
	"wallester.com/assignment/service"
	"wallester.com/assignment/storage"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {

		log.Fatal(err)

	}

	config := &storage.Config{

		Host: os.Getenv("DB_HOST"),

		Port: os.Getenv("DB_PORT"),

		Password: os.Getenv("DB_PASS"),

		User: os.Getenv("DB_USER"),

		SSLMode: os.Getenv("DB_SSLMODE"),

		DBName: os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {

		log.Fatal("could not load database")

	}

	err = models.MigrateBooks(db)

	if err != nil {

		log.Fatal("could not migrate db")

	}

	r := &service.Repository{

		DB: db,
	}

	app := fiber.New()

	r.SetupRoutes(app)

	app.Listen("localhost:3000")
}
