package kafka

import (
	"context"
	"log"

	kafkaGo "github.com/segmentio/kafka-go"
)

type Consumer interface {
	Consume(topic, groupID string, handler func(msg string)) error
}

type KafkaConsumer struct {
	Borker string
}

func NewConsumer(broker string) Consumer {
	return &KafkaConsumer{
		Borker: broker,
	}
}

func (c *KafkaConsumer) Consume(topic, groupID string, handler func(msg string)) error {
	r := kafkaGo.NewReader(kafkaGo.ReaderConfig{
		Brokers:  []string{c.Borker},
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("error in kafka reading", err)
		}
		handler(string(m.Value))
	}
}
