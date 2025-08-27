package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
	kafkaGo "github.com/segmentio/kafka-go"
)

type Producer interface {
	Publish(topic, msg string) error
	Close() error
}

type KafkaProducer struct {
	writer *kafkaGo.Writer
}

func NewProducer(broker string) Producer {
	writer := &kafkaGo.Writer{
		Addr:     kafkaGo.TCP(broker),
		Balancer: &kafkaGo.LeastBytes{},
	}
	return &KafkaProducer{writer: writer}
}

func (p *KafkaProducer) Publish(topic, message string) error {
	err := p.writer.WriteMessages(context.Background(),
		kafka.Message{Topic: topic, Value: []byte(message)},
	)
	if err != nil {
		log.Printf("Kafka publish error: %v", err)
	}
	return err
}

func (p *KafkaProducer) Close() error {
	return p.writer.Close()
}
