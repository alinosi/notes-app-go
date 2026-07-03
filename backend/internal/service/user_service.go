package service

import (
	"backend/internal/repository"
)

type UserService interface{}

type userServiceImpl struct {
	repo repository.UserRepository
}
