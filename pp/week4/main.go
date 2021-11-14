package main

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

// salvarea datelor din request in baza de date.
// Setarea unui queue, cum ar fi rabbitmq si
// citirea din el

func main() {
	log.Println("starting project")

	db := newDB()
	log.Println("DB is initialized")

	e := echo.New()
	q := newQueue()
	defer q.channel.Close()
	defer q.conn.Close()

	log.Println("The queue is initialized successfully")

	e.GET("/", helloHandler)
	e.PUT("/", createOrderHandler(db, q))
	e.GET("/files/:id", getFileHandler(db, q))

	w := newWorker()
	go w.work(db, q)
	log.Println("The worker has been initialized")

	err := e.Start(fmt.Sprintf(":%s", portOrDefault()))
	e.Logger.Fatal(err)
}

func portOrDefault() string {
	if port := os.Getenv("PORT"); port != "" {
		return port
	}

	return "8080"
}
