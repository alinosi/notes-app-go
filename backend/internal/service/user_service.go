package service

import (
	"backend/internal/model"
	"backend/internal/repository"
	"errors"
	"unicode/utf8"
)

type UserService interface {
	CreateUser(email, username, password string) (*model.User, error)
	// CreateUser(email, username, password string) (*model.User, error)
	// CreateUser(email, username, password string) (*model.User, error)
}

type userServiceImpl struct {
	repo repository.UserRepository
}

func (s *userServiceImpl) CreateUser(email, username, passowrd string) (*model.User, error) {
	if utf8.RuneCountInString(passowrd) == 0 {
		return nil, errors.New("passowrd cannot be empty")
	}

	user := &model.User{
		Username: username,
		Email:    email,
		Passowrd: passowrd,
	}

	if mark, _ := s.repo.GetUserByEmail(user.Email); mark == true {
		return nil, errors.New("Email already used")
	}

	err := s.repo.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userServiceImpl{
		repo: repo,
	}
}
