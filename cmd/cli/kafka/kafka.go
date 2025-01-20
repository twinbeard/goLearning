package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

// ! File này chỉ dùng để đọc tham khảo không có sử dụng dự án này
var (
	// kafkaConsumer *kafka.Reader
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL   = "localhost:19092" // kafka server address trong docker là 9092 (luôn luôn là 9092)
	kafkaTopic = "user_topic_vip"  // kafka topic name to push message into

)

// for kafka producer (write) - use to push message into kafka
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	// after creating kafka.Writer object, return it address
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL), // kafka server address
		Topic:    topic,               // kafka topic name to push message into
		Balancer: &kafka.LeastBytes{}, // default balancer
	}
}

// for kafka consumer (reader) - use to read message from kafka
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,     // kafka server address (broker) - can be multiple [localhost:9092,localhost:9093] -> use to manage topic
		GroupID:        groupID,     // consumer group id -> Manage consumer
		Topic:          topic,       // kafka topic name to read message from
		MinBytes:       10e3,        // 10KB -> minimum number of bytes to fetch from kafka
		MaxBytes:       10e6,        // 10MB
		CommitInterval: time.Second, // Khoảng thời gian giữa các lần commit offset là 1s
		// StartOffset:    kafka.FirstOffset, // offset (vị trí của message) to start reading from -> FirstOffset: read from the oldest message
		StartOffset: kafka.LastOffset, // offset (vị trí của message) to start reading from -> LastOffset: read from the newest message
	})
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func newStock(
	msg string,
	msgType string,
) *StockInfo {
	s := StockInfo{}
	s.Message = msg
	s.Type = "stock"

	return &s
}

// actionStock is a function to push message into kafka -> producer action stock
func actionStock(c *gin.Context) {
	// Extract query params from http request
	s := newStock(c.Query("msg"), c.Query("type"))

	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = s

	jsonBody, _ := json.Marshal(body)

	// Create messages as byte type in array
	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}
	// kafkaProducer object push message into kafka queue
	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"err": "",
		"msg": "action Successfully",
	})
}

// ConsumerATC is a function to read message from kafka -> consumer buy stock
func RegisterConsumerATC(id int) {
	// group consumer??
	kafkaGroupId := "consumer-group"
	// init kafka reader object with kafka server address, topic name, group id để thằng reader biết đọc message từ đâu
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)
	defer reader.Close() // close reader after finish this function - RegisterConsumerATC

	fmt.Printf("Consumer (%d) Hong Phien ATC:: ", id)
	// loop to read message from kafka to see what consumer buy stock
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Consumer (%d) error: %v", id, err)
		}
		fmt.Printf("Consumer (%d), hong topic: %v, partition: %v, offset: %v, time: %d %s = %s\n", id, m.Topic, m.Partition, m.Offset, m.Time.Unix(), string(m.Key), string(m.Value))
	}
}

func main() {
	r := gin.Default()

	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic) // create kafka producer object
	defer kafkaProducer.Close()

	r.POST("action/stock", actionStock) // create a router in order for producer push message into kafka queue

	// đăng ký 2 consumer để đọc message từ kafka queue - đăng ký 2 consumer để mua stock
	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)

	r.Run(":8999") // run server on port 8999
}
