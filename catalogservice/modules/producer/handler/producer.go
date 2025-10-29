package handler

import (
	"encoding/json"
	"fmt"
	"catalogservice/modules/entities"
	"catalogservice/modules/logs"

	"github.com/IBM/sarama"
)

type eventProducer struct {
	producer sarama.SyncProducer
}

func NewEventProducer(producer sarama.SyncProducer) entities.EventProducer {
	return &eventProducer{producer}
}

func (obj eventProducer) Produce(event entities.Event) error {
	// topic := reflect.TypeOf(event).Name()
	topic := event.Name()

	value, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(value),
	}

	_, _, err = obj.producer.SendMessage(&msg)
	if err != nil {
		logs.Error(err)
		return err
	}
	logs.Debug(fmt.Sprintf("%s, %s", msg.Topic, msg.Value))

	return nil
}
