package main

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

var wg sync.WaitGroup

func main() {
	consumer, err := sarama.NewConsumer([]string{"35.240.232.168:9092"}, nil)
	if err != nil {
		fmt.Println("consumer connect error:", err)
		return
	}
	fmt.Println("connnect success...")
	defer consumer.Close()
	partitions, err := consumer.Partitions("test")
	if err != nil {
		fmt.Println("geet partitions failed, err:", err)
		return
	}

	for _, p := range partitions {
		partitionConsumer, err := consumer.ConsumePartition("test", p, sarama.OffsetOldest)
		if err != nil {
			fmt.Println("partitionConsumer err:", err)
			continue
		}
		wg.Add(1)
		go func() {
			for m := range partitionConsumer.Messages() {
				fmt.Printf("key: %s, text: %s, offset: %d\n", string(m.Key), string(m.Value), m.Offset)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
