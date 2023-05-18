package db

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

// singleton pattern
// https://refactoring.guru/design-patterns/singleton/go/example#example-0
var dbInstance *DB
var syncOnce sync.Once

func NewDB() (*DB, error) {
	var err error
	syncOnce.Do(func() {
		host := os.Getenv("DB_HOST")         // localhost
		port := os.Getenv("DB_PORT")         // 5432
		user := os.Getenv("DB_USER")         // postgres
		password := os.Getenv("DB_PASSWORD") // postgres
		dbname := os.Getenv("DB_NAME")       // postgres

		fmt.Println(host, port, user, password, dbname)

		connString := fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

		db, err := sql.Open("postgres", connString)
		if err != nil {
			err = fmt.Errorf("failed to open db: %w", err)
			return
		}

		if err := db.Ping(); err != nil {
			err = fmt.Errorf("failed to ping db: %w", err)
		}

		dbInstance = &DB{db}
	})

	return dbInstance, err
}
