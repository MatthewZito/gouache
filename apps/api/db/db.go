package db

import (
	"database/sql"
	"os"
)

type DB struct {
	*sql.DB
}

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
