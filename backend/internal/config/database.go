package config

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB() *sqlx.DB {
	// postgres://<username>:<password>@<host>:<port>/<nama_database>?sslmode=disable
	dsn := "postgres://admin:password_rahasia@localhost:5432/notes_db?sslmode=disable"

	// sqlx.Connect automatically calls db.Open() and db.Ping()
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		// log.Fatalf stops the application immediately if the connection fails
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Succesfully connected to PostgreSQL database!")

	return db
}
