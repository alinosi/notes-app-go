package service

import (
	"backend/internal/model"
	"backend/internal/repository"
	"errors"
	"unicode/utf8"
)

type NoteService interface {
	CreateNote(userID, title, content string) (*model.Note, error)
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
