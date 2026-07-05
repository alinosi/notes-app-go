<!-- dummy password -->

migrate -path migrations -database "postgres://admin:password_rahasia@localhost:5432/notes_db?sslmode=disable" up
make migrate-up

migrate -path migrations -database "postgres://admin:password_rahasia@localhost:5432/notes_db?sslmode=disable" down -all
make migrate-down