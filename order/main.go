package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	topic := "orders"

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"client.id":         "orderProducer",
		"acks":              "all"},
	)
	if err != nil {
		logger.Fatalf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	delivery_chan := make(chan kafka.Event, 10000)
	run := 1
	for run < 100 {
		err = p.Produce(
			&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          []byte(fmt.Sprintf("Order %d", run)),
			},
			delivery_chan,
		)
		if err != nil {
			logger.Fatalf("Failed to produce message: %s\n", err)
			os.Exit(1)
		}
		<-delivery_chan

		// simulating creating a new order into the topic once every 3 seconds
		logger.Printf("Order %d created - %v", run, time.Now())
		time.Sleep(time.Second * 3)
		run = run + 1
	}

}
