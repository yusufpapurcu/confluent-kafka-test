package main

import (
	"context"
	"fmt"

	kafkas "github.com/yusufpapurcu/confluent-kafka-test"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	client, err := kafkas.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "kafka:9093"})
	if err != nil {
		panic(err)
	}

	topic := `aaaa`
	err = client.WriteMessage(context.TODO(), &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(`dasdasdasdasda`),
	})
	if err != nil {
		fmt.Println(err)
	}

	client.Flush(15 * 1000)
}
