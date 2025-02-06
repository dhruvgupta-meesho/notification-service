package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/dhruvgupta7733/notification-consumer/services"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	sc := services.MakeContainer()
	sc.Kafkainit()

	for {
		select {
		case <-signals:
			log.Println("Interrupt detected, shutting down...")
			return
		}
	}
}
