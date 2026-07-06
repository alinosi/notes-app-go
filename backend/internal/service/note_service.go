package service

import (
	"backend/internal/model"
	"backend/internal/repository"
	"errors"
	"unicode/utf8"
)

type NoteService interface {
	CreateNote(userID, title, content string) (*model.Note, error)
	DeleteNote(userID, notesID string) error
}

type noteServiceImpl struct {
	repo repository.NoteRepository
}

func NewNoteService(repo repository.NoteRepository) NoteService {
	return &noteServiceImpl{
		repo: repo,
	}
}

// CreateNote handles the business logic before saving to the database
func (s *noteServiceImpl) CreateNote(userID, title, content string) (*model.Note, error) {
	// 1. Business Validation
	if utf8.RuneCountInString(title) == 0 {
		return nil, errors.New("title cannot be empty")
	}

	// 2. Assemble the Struct Blueprint
	note := &model.Note{
		UserID:  userID,
		Title:   title,
		Content: content,
	}

	// 3. Command the Repository to save the data
	err := s.repo.CreateNote(note)

	if err != nil {
		return nil, err
	}

	return note, nil
}

// CreateNote handles the business logic before saving to the database
func (s *noteServiceImpl) DeleteNote(noteID, userID string) error {
	// 1. Business Validation
	err := s.noteRightsValidation(noteID, userID)
	if err != nil {
		// Check if the user attempting to delete the note is indeed the note's owner.
		return errors.New("The user does not have access to this note.") // or perhaps the note is not found
	}

	// 2. Command the Repository to save the data
	err = s.repo.DeleteNote(noteID)

	if err != nil {
		return err
	}

	return nil
}

func (s *noteServiceImpl) noteRightsValidation(noteID, userID string) error {
	_, err := s.repo.ReadNoteByUserAndNoteID(noteID, userID)
	if err != nil {
		return err
	}

	return nil
}
