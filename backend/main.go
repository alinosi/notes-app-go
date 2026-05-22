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

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
		var credentials Account

		switch r.Method {

		case http.MethodGet:
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(Message{"ini adalah halaman login"})

		case http.MethodPost:

			// read the data from http stream
			err := json.NewDecoder(r.Body).Decode(&credentials)

			guard := r.Header.Get("Origin")

			// csrf guardian simulation
			if guard != "localhost:8080/login" {
				http.Redirect(w, r, "login", http.StatusMovedPermanently)
				json.NewEncoder(w).Encode(Message{"you are not from allowed origin"})
				return
			}

			if err != nil {
				w.WriteHeader(http.StatusBadRequest) // Error 400

				w.Write([]byte("Format JSON kamu rusak\n"))
				return
			}

			w.Header().Set("content-type", "application/json")
			// database password simulation
			if credentials.Username == "admin" && credentials.Password == "admin123" {
				// w.Header().Set("Content-Type", "applicaton/json")

				w.WriteHeader(http.StatusAccepted)

				json.NewEncoder(w).Encode(Message{"login berhasil dilakukan"})
			} else {
				// w.Header().Set("Content-Type", "applicaton/json")
				w.WriteHeader(http.StatusOK)

				json.NewEncoder(w).Encode(Message{"password/username salah"})
			}

		default:
			w.WriteHeader(http.StatusBadRequest)
			message := Message{"status tidak diizinkan"}

			json.NewEncoder(w).Encode(message)
		}
	})

	address := "192.168.1.4:80"
	// address := ":80"
	// address := ":80"

	fmt.Println("server berhasil dijalankan pada %d", address)
	http.ListenAndServe(address, nil)
	// http.ListenAndServe(":80", nil)
}
