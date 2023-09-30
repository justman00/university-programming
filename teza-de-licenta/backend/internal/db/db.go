package db

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Import the postgres driver
)

type DB interface {
	// Define your query methods here
}

type dbImpl struct {
	db *sqlx.DB
}

func (d *dbImpl) Query() error {
	// Implement your query method here
	return nil
}

func New() (DB, error) {
	var err error

	// Connect to the database
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Connect to the database
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName))
	if err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}

	// Run the database migrations
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf("create migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		return nil, fmt.Errorf("create migration instance: %w", err)
	}

	if err := m.Up(); !errors.Is(err, migrate.ErrNoChange) && err != nil {
		return nil, fmt.Errorf("run migrations: %w", err)
	}

	return &dbImpl{db: db}, nil
}
