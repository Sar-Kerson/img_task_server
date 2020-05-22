package kafka

import (
	"log"

	"github.com/Shopify/sarama"
)

var (
	KafkaCli sarama.SyncProducer
)

func init() {
	KafkaCli = newProducer([]string{"35.240.232.168:9092"})
}

func Close() {
	_ = KafkaCli.Close()
}

func newProducer(brokers []string) sarama.SyncProducer {
	// For the data collector, we are looking for strong consistency semantics.
	// Because we don't change the flush settings, sarama will try to produce messages
	// as fast as possible to keep latency low.
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Retry.Max = 10 // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	return producer
}
