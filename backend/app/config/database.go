package config

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB() *sqlx.DB {
	// postgres://<username>:<password>@<host>:<port>/<nama_database>?sslmode=disable
	dsn := "postgres://admin:password_rahasia@localhost:5432/notes_db?sslmode=disable"

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	log.Println("Berhasil terhubung ke database PostgreSQL!")

	return db
}
