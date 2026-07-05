package main

import (
	"backend/internal/config"
	"backend/internal/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("Starting backend server....")

	db := config.ConnectDB()
	defer db.Close()

	app := fiber.New()

	routes.Setup(app, db)

	log.Println("Server is ready to accept request on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
