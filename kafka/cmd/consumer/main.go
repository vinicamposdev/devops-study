package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka_kafka_1:9092",
		"client.id":         "goapp-consumer",
		"group.id":          "goapp-group",
		"auto.offset.reset": "earliest",
	}
	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		fmt.Println("error consumer", err.Error())
	}
	topics := []string{"test"}
	c.SubscribeTopics(topics, nil)
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)
		}
	}
}
