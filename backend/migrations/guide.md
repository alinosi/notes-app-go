command running psql:
	psql -h localhost -U postgres -d db_name
	psql "postgres://postgres:password@localhost:5432/nama_database?sslmode=disable"


-p is flag for port, not password


psql basic command: - always user backslash as prefix (\)
	
	\l or \list (print all databases)
	\c db_name (connect/switch database)
	\dt (print all tables)
	\d table_name (print table describe)
	\q (quit)
