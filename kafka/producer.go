package kafka

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var producer *kafka.Producer

// Initialize Kafka producer
func InitKafka() {
	var err error
	producer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BROKERS"),
	})

	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}

	log.Println("Kafka producer initialized successfully")
}

// Publish a message to Kafka
func Publish(topic string, message string) error {
	if producer == nil {
		log.Println("Kafka producer is not initialized")
		return nil // Avoid nil pointer dereference
	}

	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

	if err != nil {
		log.Printf("Failed to send Kafka message: %v", err)
		return err
	}

	log.Printf("Kafka message sent to topic %s: %s", topic, message)
	return nil
}
