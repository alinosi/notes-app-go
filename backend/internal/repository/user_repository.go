package repository

import (
	"backend/internal/model"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	ReadUser(user *model.User) error
}

type userRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (r *userRepositoryImpl) CreateUser(user *model.User) error { return nil }
func (r *userRepositoryImpl) UpdateUser(user *model.User) error { return nil }
func (r *userRepositoryImpl) ReadUser(user *model.User) error   { return nil }
