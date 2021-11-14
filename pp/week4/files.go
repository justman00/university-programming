package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/streadway/amqp"
)

func saveFile(db *DB, orders []order) error {
	csvFile, err := ordersToCSV(orders)
	if err != nil {
		return fmt.Errorf("transorming orders to CSV: %v", err)
	}

	f := file{
		File:     csvFile,
		FileName: time.Now().Format(time.Layout),
	}

	if err = db.createFile(f); err != nil {
		return fmt.Errorf("saving file to DB: %v", err)
	}

	return nil
}

func composeOrders(messages <-chan amqp.Delivery) []order {
	orders := []order{}
	start := time.Now()

	for message := range messages {
		buf := bytes.NewBuffer(message.Body)
		dec := gob.NewDecoder(buf)

		o := order{}
		if err := dec.Decode(&o); err != nil {
			log.Print("could not decode message from RabbitMQ")
			panic(err)
		}

		orders = append(orders, o)

		log.Printf("consuming order: %v", o)

		// runs for 3 seconds
		if time.Since(start) < time.Second*3 {
			break
		}
	}

	log.Printf("Finished consuming, got %d orders", len(orders))

	return orders
}

func ordersToCSV(orders []order) ([]byte, error) {
	file, err := gocsv.MarshalBytes(orders)
	if err != nil {
		return []byte{}, fmt.Errorf("marshalling orders to csv bytes")
	}

	return file, nil
}
