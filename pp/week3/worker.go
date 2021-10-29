package main

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/labstack/gommon/log"
)

type worker struct {
	period time.Duration
}

// should get the rabbit mq and db connection
func newWorker() *worker {
	duration, err := time.ParseDuration("1m")
	if err != nil {
		panic("Could not parse duration")
	}

	return &worker{
		period: duration,
	}
}

func (w *worker) work(db *DB, q *queue) {
	ticker := time.NewTicker(w.period)

	// infinte loop that runs every once in a while determined by period
	for {
		select {
		case <-ticker.C:
			messages, err := q.consumeMessages()
			for message := range messages {
				buf := bytes.NewBuffer(message.Body)
				dec := gob.NewDecoder(buf)

				o := order{}
				if err = dec.Decode(&o); err != nil {
					log.Error("could not decode message from RabbitMQ")
					panic(err)
				}

				log.Infof("consuming order: %v", o)
			}
		}
	}
}
