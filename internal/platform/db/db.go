package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect creates a *gorm.DB using environment variables:
// POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_HOST, POSTGRES_DB
// sslmode is disabled by default for local/dev usage.
func Connect() (*gorm.DB, error) {
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	dbname := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, pass, host, dbname)

	gormLogger := logger.Default.LogMode(logger.Warn)
	if os.Getenv("GORM_LOG_LEVEL") == "info" {
		gormLogger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, err
	}

	// Optionally configure connection pool
	if sqlDB, err := db.DB(); err == nil {
		sqlDB.SetMaxOpenConns(25)
		sqlDB.SetMaxIdleConns(25)
		sqlDB.SetConnMaxLifetime(5 * time.Minute)
	}

	return db, nil
}
