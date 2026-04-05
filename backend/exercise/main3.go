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
	http.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			var sliceData []Notes

			n1 := Notes{ID: "abcd-edf", UserID: "user-1", Text: "this is your first note", CreatedAt: time.Now()}
			n2 := Notes{ID: "abcd-edf", UserID: "user-2", Text: "yesterdey i got new supercar", CreatedAt: time.Now()}

			sliceData = append(sliceData, n1, n2)

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(sliceData)
		case http.MethodPost:
			sliceData := []Message{Message{"Rencana simpan data berhasil dideteksi"}}

			// n1 := Message{"Rencana simpan data berhasil dideteksi"}
			// header
			// sliceData = append(sliceData, n1)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			// body
			json.NewEncoder(w).Encode(sliceData)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("metode tidak diizinkan\n"))
		}
	})

	fmt.Println("server sudah berjalan di http://172.21.224.1:8080")
	http.ListenAndServe("172.21.224.1:8080", nil)
}
