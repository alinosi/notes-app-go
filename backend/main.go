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

type Message struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var sliceData []Notes

		n1 := Notes{ID: "abcd-edf", UserID: "user-1", Text: "this is your first note", CreatedAt: time.Now()}
		n2 := Notes{ID: "abcd-edf", UserID: "user-1", Text: "yesterdey i got new supercar", CreatedAt: time.Now()}

		sliceData = append(sliceData, n1, n2)

		w.Header().Set("Content-type", "application/json")

		json.NewEncoder(w).Encode(sliceData)
		w.Write([]byte("\nSHA256{this-is-flag}"))
	})
	http.HandleFunc("/detail", func(w http.ResponseWriter, r *http.Request) {
		var sliceData []Notes

		n1 := Notes{ID: "abcd-edf", UserID: "user-1", Text: "this is your first note", CreatedAt: time.Now()}

		sliceData = append(sliceData, n1)

		w.Header().Set("Content-Type", "application/json")

		// err simulation before using database
		var err *int
		err = nil

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(sliceData)
		}
		w.Write([]byte("\ndetail page"))
	})

	fmt.Println("server berhasil dijalankan pada localhost:8080")
	http.ListenAndServe(":8080", nil)
}
