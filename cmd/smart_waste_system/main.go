package main

import (
	"fmt"
	"log"
	"os"
	"smart-waste-system/internal/app/handlers"
	"smart-waste-system/internal/app/models"
	"smart-waste-system/internal/app/repository"
	"smart-waste-system/internal/routes"

	"smart-waste-system/internal/connections"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading env file")
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("DB_PORT must be a valid integer")
	}
	sql := &connections.Sql{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		Password: os.Getenv("DB_PASS"),
		UserName: os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DbName:   os.Getenv("DB_NAME"),
	}

	db, _ := sql.Connect()
	defer sql.Close()

	err = models.MigrateUser(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	fmt.Println("NGUYEN CONG MANH")

	repo := repository.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(repo)

	app := fiber.New()
	routes.UserRoutes(app, userHandler)
	log.Fatal(app.Listen(":1234"))
}
