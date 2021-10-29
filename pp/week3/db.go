package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	conn *sqlx.DB
}

const schema = `
CREATE TABLE IF NOT EXISTS orders (
    id    UUID        PRIMARY KEY,
    SKU   text        NOT NULL,
    price int         NOT NULL
)
`

func newDB() *DB {
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
	return fmt.Sprintf("host=%s port=%d user=%s password=%s connect_timeout=%d sslmode=disable",
		"localhost",
		5432,
		"postgres",
		"postgres",
		1000,
	)
}

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
