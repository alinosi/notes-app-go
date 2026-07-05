DB_URL="postgres://admin:password_rahasia@localhost:5432/notes_db?sslmode=disable"

migrate-up:
	migrate -path migrations -database $(DB_URL) up

migrate-down:
	migrate -path migrations -database $(DB_URL) down -all