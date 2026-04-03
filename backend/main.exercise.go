package main

import (
	"fmt"
	"net/http"
)

type Notes struct {
	notes_id string `json:"id"`
	notes_id string `json:"id"`
	notes_id string `json:"id"`
	notes_id string `json:"id"`
	notes_id string `json:"id"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("ini adalah halaman utama")
	})
	http.HandleFunc("/detail", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("ini adalah halaman utama")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("ini adalah halaman utama")
	})
}
