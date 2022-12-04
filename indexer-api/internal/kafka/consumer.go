package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/soa/indexer-api/internal/config"
)

type consumer struct {
	reader         *kafka.Reader
	productService productService
	searchService  searchService
	handleMap      handleMap
}

func (c *consumer) Start(ctx context.Context) {
	for {
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Println(err)
			continue
		}
		if handle, ok := c.handleMap[string(m.Value)]; ok {
			log.Println(handle(ctx))
		}
	}
}

func (c *consumer) productAPIIvent(ctx context.Context) error {
	info, err := c.productService.GetNewData(ctx)
	if err != nil {
		return err
	}

	return c.searchService.SendNewInfo(ctx, info)
}

func NewConsumer(cfg *config.Config, productService productService, searchService searchService) Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{cfg.Kafka.Address},
		Topic:     cfg.Kafka.Topic,
		Partition: cfg.Kafka.Partition,
		GroupID:   "1",
	})
	c := &consumer{
		reader:         r,
		productService: productService,
		searchService:  searchService,
	}
	hm := handleMap{
		"productAPI": c.productAPIIvent,
	}
	c.handleMap = hm
	return c
}
