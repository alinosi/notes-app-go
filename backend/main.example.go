package main

import (
	"encoding/json"
	"net/http"
)

// users dummy data
type User struct {
	User_id string `json:"id"`
	Email   string `json:"email"`
	// Username   string `json:"username"`
	// Password   string `json:"password"`
	// Role       string `json:"role"`
	// Created_at string `json:"created_at"`
	// Updated_at string `json:"updated_at"`
	// Deleted_at string `json:"deleted_at"`
}

// notes dummy data
type Note struct {
	Note_Id    string `json:"id"`
	User_id    string `json:"email"`
	Text       string `json:"text"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Deleted_at string `json:"deleted_at"`
}

func main() {

	// 1. ROUTING: Mendaftarkan alamat
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {

		// 2. LOGIKA API: Membuat data pura-pura (Dummy)
		data := User{User_id: "abcd-efgh-123", Email: "admin123@gmail.com"}

		// 3. HEADER: Kasih tahu browser kalau ini adalah JSON, bukan HTML
		w.Header().Set("Content-Type", "application/json")

		// 4. ENCODING: Mengubah Struct Go menjadi teks JSON
		json.NewEncoder(w).Encode(data)
	})

	// 5. START SERVER
	http.ListenAndServe(":8080", nil)
}
