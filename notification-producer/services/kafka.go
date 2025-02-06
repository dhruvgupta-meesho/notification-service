package services

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
)

func SendKafka (id string){
	brokers := []string{"localhost:29092"}

	topic := "notify"

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("Failed to close Kafka producer: %v", err)
		}
	}()

	res, _ := json.Marshal(struct{
		EmailMsg string
		Id string
	}{
		"Send Email to this", id,
	})

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(res),
	}

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	log.Printf("Message sent successfully to partition %d with offset %d\n", partition, offset)
}