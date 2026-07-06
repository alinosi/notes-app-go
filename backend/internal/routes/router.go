// File: internal/routes/router.go
package routes

import (
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

// Setup berfungsi sebagai terminal utama untuk menyambungkan semua kabel
func Setup(app *fiber.App, db *sqlx.DB) {
	// ---------------------------------------------------------
	// 1. Dependency Injection (Area Pemasangan Kabel)
	// ---------------------------------------------------------

	// Kabel Notes
	noteRepo := repository.NewNoteRepository(db)
	noteService := service.NewNoteService(noteRepo)
	noteHandler := handler.NewNoteHandler(noteService)

	// Kabel Users (Nantinya di sini)
	// userRepo := repository.NewUserRepository(db)
	// userService := service.NewUserService(userRepo)
	// userHandler := handler.NewUserHandler(userService)

	// ---------------------------------------------------------
	// 2. Routing (Area Penunjuk Jalan)
	// ---------------------------------------------------------

	// Kita bisa membuat prefix "/api/v1" agar URL lebih rapi dan profesional
	api := app.Group("/api/v1")

	// Grup khusus Notes -> URL akan menjadi "/api/v1/notes"
	notes := api.Group("/notes")
	notes.Post("/", noteHandler.CreateNote)
	notes.Post("/delete", noteHandler.DeleteNote)
	// notes.Get("/", noteHandler.GetAllNotes)       <-- Untuk tugas mandirimu nanti
	// notes.Get("/:id", noteHandler.GetNoteByID)    <-- Untuk tugas mandirimu nanti
	// notes.Put("/:id", noteHandler.UpdateNote)     <-- Untuk tugas mandirimu nanti
	// notes.Delete("/:id", noteHandler.DeleteNote)  <-- Untuk tugas mandirimu nanti

	// Grup khusus Users -> URL akan menjadi "/api/v1/users"
	// users := api.Group("/users")
	// users.Post("/register", userHandler.Register)
	// users.Post("/login", userHandler.Login)
}
