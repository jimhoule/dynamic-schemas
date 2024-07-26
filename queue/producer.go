package queue

import (
	"github.com/IBM/sarama"
)

type Producer = sarama.SyncProducer
type ProducerMessage = sarama.ProducerMessage

type ProducerHandler struct {
	Producer Producer
}

func NewProducerHandler(addresses []string) (ProducerHandler, error) {
	producerHandler := ProducerHandler{}

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producer, err := sarama.NewSyncProducer(addresses, config)
	if err != nil {
		return producerHandler, err
	}

	producerHandler.Producer = producer

	return producerHandler, nil
}

func (ph *ProducerHandler) SendMessage(topic string, payload []byte) error {
	message := &ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(payload),
	}

	_, _, err := ph.Producer.SendMessage(message)
	if err != nil {
		return err
	}

	return nil
}
