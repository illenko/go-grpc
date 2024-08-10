package main

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
)

func createKafkaWriter(brokers []string, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    topic,
		Balancer: &kafka.Hash{},
	}
}

func sendPaymentStatusUpdate(writer *kafka.Writer, event PaymentEvent) error {
	jsonResponse, err := toJson(event)
	if err != nil {
		return err
	}
	msg := kafka.Message{
		Key:   []byte(event.PaymentId),
		Value: []byte(jsonResponse),
	}
	return writer.WriteMessages(context.Background(), msg)
}

func toJson(event PaymentEvent) (string, error) {
	jsonData, err := json.Marshal(event)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
