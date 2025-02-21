package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDatabase() (*sql.DB, error) {
	connStr := "user=postgres dbname=arquitectura sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
