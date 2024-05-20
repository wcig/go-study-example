package ch101_kafka

import (
	"fmt"
	"testing"

	"github.com/Shopify/sarama"
)

func TestKafkaProducer(t *testing.T) {
	config := sarama.NewConfig()
	producer, err := sarama.NewSyncProducer([]string{brokerServer}, config)
	if err != nil {
		t.Fatal(err)
	}

	msg := &sarama.ProducerMessage{
		Topic: "test",
		Key:   sarama.ByteEncoder("ok"),
		Value: sarama.ByteEncoder("ok"),
	}
	partition, offset, err := producer.SendMessage(msg)
	fmt.Printf("partition=%d, offset=%d, err=%s", partition, offset, err)
}
