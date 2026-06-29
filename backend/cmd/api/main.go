package main

import (
	"backend/internal/config"
	"log"
)

func main() {
	log.Println("Starting App-Notes Backend Server...")

	// 1. Initialize Database Connection
	db := config.ConnectDB()

	// defer ensures the database connection is closed gracefully when the app stops
	defer db.Close()

	log.Println("Server is ready to accept requests!")

	// (Nanti kode untuk menyalakan Fiber / Routing HTTP akan kita taruh di bawah sini)
}
