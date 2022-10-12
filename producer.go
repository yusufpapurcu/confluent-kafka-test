package kafka

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	producer *kafka.Producer
}

func NewProducer(conf *kafka.ConfigMap) (*Producer, error) {
	p, err := kafka.NewProducer(conf)
	return &Producer{producer: p}, err
}

func (p *Producer) WriteMessage(ctx context.Context, m *kafka.Message) error {
	return p.producer.Produce(m, nil)
}

func (p *Producer) Flush(a int) int {
	return p.producer.Flush(a)
}
