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
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodPost:

			var err *int
			// simulation if the server has run correctly (password is true)
			err = nil

			if err != nil {
				message := Message{"Kesalahan server"}

				w.WriteHeader(http.StatusBadGateway)
				w.Header().Set("Content-Type", "application/json")

				json.NewEncoder(w).Encode(message)
			}

			// credentials simulation
			accountValidations := true

			if accountValidations == true {
				message := Message{"Login berhasil"}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				json.NewEncoder(w).Encode(message)
			} else {
				message := Message{"Username/Password salah"}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				json.NewEncoder(w).Encode(message)
			}
		default:
			w.WriteHeader(http.StatusBadRequest)
			message := Message{"status tidak diizinkan"}

			json.NewEncoder(w).Encode(message)
		}
	})

	fmt.Println("server berhasil dijalankan pada localhost:8080")
	http.ListenAndServe(":8080", nil)
}
