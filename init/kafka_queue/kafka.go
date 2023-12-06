package kafka

import (
	"github.com/IBM/sarama"
)

func NewDefaultConfig() *sarama.Config {
	return sarama.NewConfig()
}

// []string{"kafka:9092"}
func NewClient(connParam []string) (sarama.Client, error) {
	client, err := sarama.NewClient(connParam, NewDefaultConfig())
	if err != nil {
		defer client.Close()
		return nil, err
	}

	return client, nil
}

func NewConsumer(connParam []string) (sarama.Consumer, error) {
	client, err := NewClient(connParam)
	if err != nil {
		return nil, err
	}

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		defer consumer.Close()
		return nil, err
	}

	return consumer, nil
}

func NewProducer(connParam []string) (sarama.AsyncProducer, error) {
	client, err := NewClient(connParam)
	if err != nil {
		return nil, err
	}

	producer, err := sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		defer producer.Close()
		return nil, err
	}

	return producer, nil
}
