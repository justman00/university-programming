package main

import (
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
			log.Info("Running worker")
			messages, err := q.consumeMessages()
			if err != nil {
				panic(err)
			}

			orders := composeOrders(messages)

			if err = saveFile(db, orders); err != nil {
				panic(err)
			}
		}
	}
}
