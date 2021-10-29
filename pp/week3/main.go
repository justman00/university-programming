package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

// salvarea datelor din request in baza de date.
// Setarea unui queue, cum ar fi rabbitmq si
// citirea din el

func main() {
	log.Println("starting project")

	db := newDB()
	e := echo.New()
	q := newQueue()
	defer q.channel.Close()
	defer q.conn.Close()

	e.GET("/", helloHandler)
	e.PUT("/", createOrderHandler(db, q))

	w := newWorker()
	go w.work(db, q)

	e.Logger.Fatal(e.Start(":8080"))
}
