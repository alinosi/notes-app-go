package main

import (
	"backend/internal/config"
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("Starting backend server....")

	db := config.ConnectDB()
	defer db.Close()

	// repo
	noteRepo := repository.NewNoteRepository(db)
	// service
	noteService := service.NewNoteService(noteRepo)
	// handler
	noteHandler := handler.NewNoteHandler(noteService)

	app := fiber.New()

	app.Post("/Notes", noteHandler.CreateNote)

	log.Println("Server is ready to accept request on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
