package main

import (
	"bytes"
	"encoding/gob"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// should receive rabbit mq connection and push events to it
func createOrderHandler(db *DB, q *queue) echo.HandlerFunc {
	return func(c echo.Context) error {
		var o order

		if err := c.Bind(&o); err != nil {
			log.Errorf("binding the request to a struct: %v", err)
			return echo.NewHTTPError(http.StatusBadRequest, "wrong data structure")
		}

		if err := db.createOrder(o); err != nil {
			log.Errorf("creating order in DB: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "could not create order")
		}

		// publish event to the queue
		var buf bytes.Buffer

		encoder := gob.NewEncoder(&buf)
		if err := encoder.Encode(o); err != nil {
			log.Errorf("encoding the order: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "could not encode the order")
		}

		if err := q.publishEvent(buf.Bytes()); err != nil {
			log.Errorf("publishing event to Rabbit MQ: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "could not publish event to Rabbit MQ")
		}

		return c.JSON(http.StatusOK, "OK")
	}
}
