package initialize

import (
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/twinbeard/goLearning/global"
)

var (
	// kafkaConsumer *kafka.Reader
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL   = "localhost:19092" // kafka server address trong docker là 9092 (luôn luôn là 9092)
	kafkaTopic = "otp-auth-topic"  // kafka topic name to push message into

)

// for kafka producer (write) - use to push message into kafka
func InitKafka(kafkaURL, topic string) {
	// after creating kafka.Writer object, return it address
	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL), // kafka server address
		Topic:    topic,               // kafka topic name to push message into
		Balancer: &kafka.LeastBytes{}, // default balancer
	}
}
func closeKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		log.Fatal("Failed to close kafka producer", err)
	}
}
