package main

import (
	"context"

	zapctx "github.com/saltpay/go-zap-ctx"
	kafkas "github.com/yusufpapurcu/confluent-kafka-test"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	consumer, err := kafkas.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9093",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}

	consumer.Subscribe("aaaa")

	consumer.Listen(context.TODO(), func(ctx context.Context, m *kafka.Message) error {
		zapctx.Info(ctx, string(m.Value))
		return nil
	})
}
