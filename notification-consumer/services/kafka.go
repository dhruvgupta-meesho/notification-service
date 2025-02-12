package services

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/IBM/sarama"
)

type kafkamsg struct{
	EmailMsg string
	Id int64
}

func (sc *ServiceContainer)Kafkainit(){
	brokers := "localhost:29092"
	topic := "notify"
	consumer, err := sarama.NewConsumer(strings.Split(brokers, ","), nil)

	if err != nil {
		log.Fatalf("Error creating Kafka consumer: %v", err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalf("Error closing Kafka consumer: %v", err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error creating partition consumer: %v", err)
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalf("Error closing partition consumer: %v", err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	log.Printf("Consuming messages from topic: %s", topic)

	done := make(chan struct{})

	func() {
		for {
			select {
			case msg := <-partitionConsumer.Messages():
				var v kafkamsg
				json.Unmarshal(msg.Value, &v)
				log.Printf("Message received: for id = %d value = %s", v.Id, string(v.EmailMsg))
				sc.SendMail(string(msg.Value), v.Id)
			case <-signals:
				log.Println("Interrupt detected, shutting down...")
				close(done)
				return
			}
		}
	}()
	<-done
}