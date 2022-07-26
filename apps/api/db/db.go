package db

import (
	"database/sql"
)

type DB struct {
	*sql.DB
}

// func Connect() (*DB, error) {
// 	connStr := os.Getenv("DB_CONN")
// }
