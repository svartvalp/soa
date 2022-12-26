package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/soa/indexer-api/internal/config"
)

type (
	consumer struct {
		reader         *kafka.Reader
		productService productService
		handleMap      handleMap
	}

	msg struct {
		Service string  `json:"service,omitempty"`
		Type    string  `json:"type"`
		IDs     []int64 `json:"ids,omitempty"`
	}
)

func (c *consumer) Start(ctx context.Context) {
	for {
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Println(err)
			continue
		}

		var msg msg
		err = json.Unmarshal(m.Value, &msg)
		if err != nil {
			log.Println(err)
			continue
		}

		if handle, ok := c.handleMap[fmt.Sprintf("%s-%s", msg.Service, msg.Type)]; ok {
			log.Println(handle(ctx, msg.IDs))
		}
	}
}

func NewConsumer(cfg *config.Config, productService productService) Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{cfg.Kafka.Address},
		Topic:     cfg.Kafka.Topic,
		Partition: cfg.Kafka.Partition,
		GroupID:   "1",
	})
	c := &consumer{
		reader:         r,
		productService: productService,
	}
	hm := handleMap{
		"productAPI-UPDATE": c.productService.ProductAPIUpdateIvent,
		"productAPI-DELETE": c.productService.ProductAPIDeleteIvent,
		"productAPI-CREATE": c.productService.ProductAPICreateIvent,
	}
	c.handleMap = hm
	return c
}
