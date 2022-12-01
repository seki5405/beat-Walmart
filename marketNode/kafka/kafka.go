package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	// KafkaTopic is the name of the Kafka topic
	KafkaTopic = "market-kafka-test"
	// KafkaBroker is the address of the Kafka broker
	KafkaBroker = "sw-kafka.kafka-ns.svc.cluster.local:9092"
	// KafkaBroker = "localhost:9092"
	// KafkaBroker = "kafka:9092"
)

// Publish a message to a Kafka topic
func Publish(topic string, message string) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", KafkaBroker, KafkaTopic, 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte(message),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
