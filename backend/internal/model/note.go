package model

import "time"

type Note struct {
	ID        string     `json:"id"`
	UserID    string     `json:"user_id" db:"user_id"`
	Title     string     `json:"title" db:"title"`
	Content   string     `json:"content" db:"content"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"` // allowed NULL
}
