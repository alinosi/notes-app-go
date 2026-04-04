package main

import (
	"net/http"
	"time"
)

type Notes struct {
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	Text      string     `json:"text"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func main() {
	http.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {
		data := Notes{ID: "abcd-efg", UserID: "user_a", Text: "this is my first note", CreatedAt: time.Now()}
	})
}
