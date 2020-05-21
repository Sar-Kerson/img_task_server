package kafka

import (
	"testing"

	"github.com/Shopify/sarama"
)

func TestProducer(t *testing.T) {
	// We are not setting a message key, which means that all messages will
	// be distributed randomly over the different partitions.
	partition, offset, err := KafkaCli.SendMessage(&sarama.ProducerMessage{
		Topic: "important",
		Value: sarama.StringEncoder("hello world from req"),
	})
	t.Log(partition, offset, err)
}
