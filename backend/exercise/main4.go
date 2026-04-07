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
		// fmt.Printf("Type: %T\n", r.Body) // Output: Type: int

		switch r.Method {

		case http.MethodGet:
			var sliceData []Notes

			n1 := Notes{ID: "abcd-edf", UserID: "user-1", Text: "this is your first note", CreatedAt: time.Now()}
			n2 := Notes{ID: "abcd-edf", UserID: "user-2", Text: "yesterdey i got new supercar", CreatedAt: time.Now()}

			sliceData = append(sliceData, n1, n2)

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(sliceData)
		case http.MethodPost:
			// 1. Siapkan wadah kosong berdasarkan Struct yang sudah kita buat
			var newNote Notes

			// 2. Ambil data dari r.Body, decode, dan masukkan ke wadah &newNote
			err := json.NewDecoder(r.Body).Decode(&newNote)

			// 3. Cek apakah proses bongkar paket berhasil?
			if err != nil {
				w.WriteHeader(http.StatusBadRequest) // Error 400

				w.Write([]byte("Format JSON kamu rusak\n"))
				return
			}
			// 4. Sekarang data sudah ada di variabel newNote
			fmt.Printf("Data masuk: %s\n", newNote.Text)
			json.NewEncoder(w).Encode(newNote)

			w.Write([]byte("ini adalah data yang anda masukkan:"))
			w.Write([]byte(newNote.Text))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)

			w.Write([]byte("metode tidak diizinkan\n"))
		}
	})

	// fmt.Println("server sudah berjalan di http://172.21.224.1:8080")
	// http.ListenAndServe("172.21.224.1:8080", nil)
	http.ListenAndServe(":8080", nil)
}
