package main

import (
	"log"
	"os"
	"smart-waste-system/internal/app/handlers"
	"smart-waste-system/internal/connections"

	"smart-waste-system/internal/routes"

	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
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

	// err = models.MigrateBooks(db)
	// if err != nil {
	// 	log.Fatal("could not migrate db")
	// }

	repo := &handlers.Repository{
		DB: db,
	}

	// print(db)
	app := fiber.New()

	//socket xem sau
	routes.UserRoutes(app, repo)
	routes.TrashBinRoutes(app, repo)
	routes.ReportRoutes(app, repo)
	routes.TransactionRoutes(app, repo)
	routes.AreaRoutes(app, repo)

	app.Listen(":3000")

}
