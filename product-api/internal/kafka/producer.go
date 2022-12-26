package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/soa/product-api/internal/config"
)

type (
	Msg struct {
		Service string  `json:"service,omitempty"`
		Type    string  `json:"type"`
		IDs     []int64 `json:"ids,omitempty"`
	}

	Producer interface {
		Write(ctx context.Context, msg []byte) error
	}

	producer struct {
		partition int
		writer    *kafka.Writer
	}
)

func NewProducer(cfg *config.Config) Producer {
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    cfg.Kafka.Topic,
		Balancer: &kafka.LeastBytes{},
	}

	return &producer{
		writer:    w,
		partition: cfg.Kafka.Partition,
	}
}

func (p *producer) Write(ctx context.Context, msg []byte) error {
	return p.writer.WriteMessages(ctx, kafka.Message{
		Partition: p.partition,
		Value:     msg,
	})
}
