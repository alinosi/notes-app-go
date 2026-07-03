package service

import (
	"backend/internal/model"
	"backend/internal/repository"
)

type NoteService interface {
	SaveNote(note *model.Note) error
	UpdateNote(note *model.Note) error
}

type noteServiceImpl struct {
	repo repository.NoteRepository
}

func NewNoteService(repo repository.NoteRepository) NoteService {
	return &noteServiceImpl{
		repo: repo,
	}
}

// dummy code for simulation
func (s *noteServiceImpl) SaveNote(note *model.Note) error {
	err := s.repo.CreateNote(note)
	if err != nil {
		return err
	}
	return nil
}

func (s *noteServiceImpl) UpdateNote(note *model.Note) error {
	// user credentials validantions

	err := s.repo.UpdateNote(note)
	if err != nil {
		return err
	}
	return nil
}
