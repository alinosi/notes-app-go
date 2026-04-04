package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Notes struct {
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	Text      string     `json:"text"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func main() {
	http.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {
		var listNotes []Notes

		// Note 1
		n1 := Notes{ID: "1", UserID: "user_a", Text: "Beli kopi", CreatedAt: time.Now()}
		// Note 2
		n2 := Notes{ID: "2", UserID: "user_a", Text: "Belajar Go", CreatedAt: time.Now()}

		listNotes = append(listNotes, n1, n2)

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(listNotes)
	})

	fmt.Println("server sudah berjalan di http://localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
