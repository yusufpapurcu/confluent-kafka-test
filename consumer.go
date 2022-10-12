package kafka

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	zapctx "github.com/saltpay/go-zap-ctx"
	"go.uber.org/zap"
)

type Consumer struct {
	consumer *kafka.Consumer
}

type Processor func(context.Context, *kafka.Message) error

func NewConsumer(conf *kafka.ConfigMap) (*Consumer, error) {
	c, err := kafka.NewConsumer(conf)
	return &Consumer{consumer: c}, err
}

func (c *Consumer) Listen(ctx context.Context, prc Processor) {
	for {
		m, err := c.consumer.ReadMessage(-1)
		if err != nil {
			zapctx.Error(ctx, "error reading message, trying again")

			continue
		}

		err = prc(ctx, m)
		if err != nil {
			zapctx.Error(ctx, "error processing message", zap.Error(err))

			continue
		}

		zapctx.Info(ctx, "message processed")

		_, err = c.consumer.CommitMessage(m)
		if err != nil {
			zapctx.Error(ctx, "error commiting message", zap.Error(err))

			continue
		}

		zapctx.Info(ctx, "message committed")
	}
}

func (c *Consumer) Subscribe(s ...string) error {
	return c.consumer.SubscribeTopics(s, nil)
}
