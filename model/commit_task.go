package model

import (
	"log"

	"github.com/Sar-Kerson/img_task_server/dal/kafka"
	"github.com/Shopify/sarama"
)

const (
	MessageTopic = "gan"
	MessageKey   = "commit"
)

type MsgMeta struct {
	Topic string
	Key   string
	Value string
}

func NewMsgMeta(taskId string) *MsgMeta {
	return &MsgMeta{
		Topic: MessageTopic,
		Key:   MessageKey,
		Value: taskId,
	}
}

func CommitTask(taskId string) error {
	// We are not setting a message key, which means that all messages will
	// be distributed randomly over the different partitions.
	log.Printf("[CommitTask] req taskID: %s", taskId)

	meta := NewMsgMeta(taskId)

	partition, offset, err := kafka.KafkaCli.SendMessage(&sarama.ProducerMessage{
		Topic: meta.Topic,
		Key:   sarama.StringEncoder(meta.Key),
		Value: sarama.StringEncoder(meta.Value),
	})

	if err != nil {
		log.Printf("[CommitTask] SendMessage failed, taskID: %s, err: %s", taskId, err.Error())
		return err
	}

	log.Printf("[CommitTask] SendMessage success, taskID: %s, partition: %d, offset: %d", taskId, partition, offset)
	return nil
}
