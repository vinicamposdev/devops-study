package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChan := make(chan kafka.Event)
	producer := NewKafkaProducer()
	Publish("message", "test", producer, []byte("transference"), deliveryChan)
	go DeliveryReport(deliveryChan) // async
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "kafka_kafka_1:9092",
		"delivery.timeout.ms": "0",
		"ack":                 "all",
		"enable.idempotence":  "true",
	}
	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}
	err := producer.Produce(message, deliveryChan)
	return err
}

func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Error on sent message")
			} else {
				fmt.Println("Message sent:", ev.TopicPartition)
			}
		}
	}
}
