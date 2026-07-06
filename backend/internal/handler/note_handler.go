package handler

import (
	"backend/internal/service"

	"github.com/gofiber/fiber/v2"
)

type NoteHandler struct {
	service service.NoteService
}

func NewNoteHandler(service service.NoteService) *NoteHandler {
	return &NoteHandler{
		service: service,
	}
}

// CreateNote handles the HTTP POST request for creating a new note
func (h *NoteHandler) CreateNote(c *fiber.Ctx) error {
	// 1. Prepare an empty struct to read the incoming JSON request
	var req struct {
		UserID  string `json:"user_id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	// 2. Parse the JSON body into the 'req' struct
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body format",
		})
	}

	// 3. Pass the extracted data to the Service Layer (Business Logic)
	note, err := h.service.CreateNote(req.UserID, req.Title, req.Content)
	if err != nil {
		// If the Service throws an error (e.g., empty title), return it to the user
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// 4. Return a success response to the Frontend
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Note created successfully",
		"data":    note,
	})
}

func (h *NoteHandler) DeleteNote(c *fiber.Ctx) error {
	// 1. Prepare an empty struct to read the incoming JSON request
	var req struct {
		ID     string `json:"id"`
		NoteID string `json:"note_id"`
	}

	// 2. Parse the JSON body into the 'req' struct
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body format",
		})
	}

	// 3. Pass the extracted data to the Service Layer (Business Logic)
	err := h.service.DeleteNote(req.ID, req.NoteID)
	if err != nil {
		// If the Service throws an error (e.g., notefailure), return it to the user
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// 4. Return a success response to the Frontend
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Note deleted successfully",
	})
}
