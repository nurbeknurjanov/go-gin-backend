package kafka

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const (
	flushTimeout = 5000
)

var errUnknownType = errors.New("unknown event type")

type Producer struct {
	producer *kafka.Producer
}

func NewProducer(addrs []string) (*Producer, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers": strings.Join(addrs, ","),
		//"acks":              "0",
		"acks": "1",
		//"acks":              "all", or "-1"
	}
	p, err := kafka.NewProducer(config)
	if err != nil {
		return nil, fmt.Errorf("error creating kafka producer: %w", err)
	}
	return &Producer{producer: p}, nil
}

func (p *Producer) Produce(message, topic string, key *string) error {
	var Key []byte
	if key != nil {
		Key = []byte(*key)
	}

	kafkaMessage := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value:     []byte(message),
		Key:       Key,
		Timestamp: time.Now(),
	}

	kafkaChannel := make(chan kafka.Event)
	if err := p.producer.Produce(kafkaMessage, kafkaChannel); err != nil {
		return fmt.Errorf("error sending message to kafka: %w", err)
	}
	e := <-kafkaChannel
	switch ev := e.(type) {
	case *kafka.Message:
		return nil
	case kafka.Error:
		return ev
	default:
		return errUnknownType
	}
}
func (p *Producer) Close() {
	p.producer.Flush(flushTimeout)
	p.producer.Close()
}
