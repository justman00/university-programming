package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/justman00/tehnici-si-mecanisme-de-proiectare-software/db"
	"github.com/justman00/tehnici-si-mecanisme-de-proiectare-software/handlers"
	"github.com/justman00/tehnici-si-mecanisme-de-proiectare-software/messaging"
	"github.com/justman00/tehnici-si-mecanisme-de-proiectare-software/models"
	_ "github.com/lib/pq"
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

func main() {
	dbInstance, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbInstance.Close()

	if err := migrateUp(dbInstance.DB); err != nil {
		log.Fatal(err)
	}

	router := bunrouter.New(
		bunrouter.Use(reqlog.NewMiddleware()),
		bunrouter.Use(errorHandler),
	)

	clientModels := &models.ClientModels{dbInstance.DB}
	bookingModels := &models.BookingModels{dbInstance.DB}

	sender, err := messaging.NewSender("email")
	if err != nil {
		log.Fatal(err)
	}

	handlersCollections := handlers.NewHandler(clientModels, bookingModels, sender)

	router.GET("/", indexHandler)
	router.POST("/clients", handlersCollections.CreateClient)
	router.POST("/bookings", handlersCollections.CreateReservation)
	router.GET("/bookings", handlersCollections.GetReservations)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, req bunrouter.Request) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return bunrouter.JSON(w, "Hello, World!")
}

func errorHandler(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		// Call the next handler on the chain to get the error.
		err := next(w, req)

		if err != nil {
			fmt.Println(fmt.Printf("Error: %v", err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return nil
	}
}
