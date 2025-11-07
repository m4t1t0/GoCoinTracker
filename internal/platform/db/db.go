package db

import (
	"database/sql"
	"fmt"
	"os"

	// pgx stdlib driver
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Connect creates a *sql.DB using environment variables:
// POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_HOST, POSTGRES_DB
// sslmode is disabled by default for local/dev usage.
func Connect() (*sql.DB, error) {
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	dbname := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, pass, host, dbname)
	return sql.Open("pgx", dsn)
}
