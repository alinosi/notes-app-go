package service

import (
	"backend/internal/model"
	"backend/internal/repository"
)

type NoteService interface {
	CreateNote(note *model.Note) error
}

type noteServiceImpl struct {
	repo repository.NoteRepository
}

// dummy code for simulation
func (s *noteServiceImpl) CreateNote(note *model.Note) error {
	return nil
}

func NewNoteService(repo repository.NoteRepository) NoteService {
	return &noteServiceImpl{
		repo: repo,
	}
}
