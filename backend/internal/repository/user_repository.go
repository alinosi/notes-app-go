package repository

import (
	"backend/internal/model"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	GetUserByEmail(email string) (bool, error)
	// ReadUser(user *model.User) error
	// ReadUsers(user *model.User) error
}

type userRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (r *userRepositoryImpl) CreateUser(user *model.User) error {
	query := `
		INSERT INTO user (email, username, passowrd)
		VALUES (:email, :username, :password)
		RETURNING id, created_at, updated_at
	`

	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(user)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *userRepositoryImpl) UpdateUser(user *model.User) error { return nil }
func (r *userRepositoryImpl) GetUserByEmail(email string) (bool, error) {
	query := `SELECT id FROM user WHERE email = $1`

	var id string

	err := r.db.Get(&id, query, email)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *userRepositoryImpl) ReadUsers(user *model.User) (bool, error) {
	query := `SELECT FROM user WHERE email = :email`

	err := r.db.Get(&user, query, user.Email)
	if err != nil {
		return false, err
	}

	return true, nil
}
