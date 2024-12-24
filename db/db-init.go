package db

import (
	"fmt"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var (
	db   *sqlx.DB
	once sync.Once
)

func NewClient() (*sqlx.DB, error) {
	var err error

	once.Do(func() {
		dbURL := os.Getenv("SUPABASE_DB_URL")
		if dbURL == "" {
			// Construct the connection string from individual parameters
			dbURL = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				os.Getenv("SUPABASE_URL"),
				os.Getenv("SUPABASE_PORT"),
				os.Getenv("SUPABASE_USER"),
				os.Getenv("SUPABASE_PASSWORD"),
				os.Getenv("SUPABASE_DB_NAME"),
			)
		}

		db, err = sqlx.Connect("postgres", dbURL)
		if err != nil {
			return
		}

		err = db.Ping()
	})

	if err != nil {
		return nil, fmt.Errorf("database initialization error: %w", err)
	}

	return db, nil
}
