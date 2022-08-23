package repositories

import (
	"database/sql"
	"os"
)

// DB holds a pointer to the opened database handle.
type DB struct {
	*sql.DB
}

// Connect initializes a new postgres connection.
func Connect() (*DB, error) {
	connStr := os.Getenv("DB_CONN")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
