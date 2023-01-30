package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	conn *sqlx.DB
}

const schema = `
CREATE TABLE IF NOT EXISTS orders (
	id UUID PRIMARY KEY,
	SKU text NOT NULL,
	price int NOT NULL
);
CREATE TABLE IF NOT EXISTS files (
	id SERIAL PRIMARY KEY,
	file bytea NOT NULL,
	file_name text NOT NULL
);
`

func newDB() *DB {
	log.Println("Database URL: ", os.Getenv("DATABASE_URL"))
	conn, err := sqlx.Connect("postgres", newConnStr())
	if err != nil {
		panic(err)
	}

	if err = conn.Ping(); err != nil {
		panic("DB is not live")
	}

	conn.MustExec(schema)

	return &DB{
		conn,
	}
}

func newConnStr() string {
	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		return dbURL
	}

	return fmt.Sprintf("host=%s port=%d user=%s password=%s connect_timeout=%d sslmode=disable",
		"localhost",
		1234,
		"postgres",
		"postgres",
		1000,
	)
}

// TODO: introduce transactions
func (db *DB) createOrder(o order) error {
	sql := `
		INSERT INTO orders (id, sku, price) VALUES ($1, $2, $3)
	`

	_, err := db.conn.Exec(sql, o.ID, o.SKU, o.Price)
	if err != nil {
		return fmt.Errorf("inserting a new order: %v", err)
	}

	return nil
}

func (db *DB) createFile(f file) error {
	sql := `
		INSERT INTO files (file, file_name) VALUES ($1, $2)
	`

	_, err := db.conn.Exec(sql, f.File, f.FileName)
	if err != nil {
		return fmt.Errorf("inserting a new file: %v", err)
	}

	return nil
}

func (db *DB) getFile(id string) (file, error) {
	var f file

	sql := `
		SELECT * FROM files WHERE id = $1
	`

	err := db.conn.Get(&f, sql, id)
	if err != nil {
		return file{}, fmt.Errorf("getting a new file: %v", err)
	}

	return f, nil
}
