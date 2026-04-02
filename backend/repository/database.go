package repository

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/notes_app?parseTime=true")
	if err != nil {
		panic(err)
	} // sql.open return a database polling(DB struct), not make a connection to the database

	// db struct has many methods which are useful for database management
	db.SetMaxIdleConns(10) // idle connecton
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func OpenConnection() {
	db := GetConnection() // open the connection

	defer db.Close() // avoid the connection leak

}

/**
Dengan pooling:

- Sekali buka koneksi, disimpan di memori.
- Saat ada query baru, ambil koneksi yang idle (menganggur).
- Setelah query selesai, koneksi dikembalikan ke pool.

benar-benar mirip dengan konsep pooling di goroutines dimana data bisa
digunakan dan dikembalikan lalu digunakan ulang

**/
