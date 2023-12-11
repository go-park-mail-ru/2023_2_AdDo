package init_kafka

import (
	"github.com/IBM/sarama"
	"main/internal/common/logger"
	"time"
)

func NewDefaultConfig() *sarama.Config {
	return sarama.NewConfig()
}

func NewClient(connParam []string) (sarama.Client, error) {
	return sarama.NewClient(connParam, NewDefaultConfig())
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

func NewProducer(connParam []string) (sarama.SyncProducer, error) {
	producer, err := sarama.NewSyncProducer(connParam, getDefaultConfig())
	if err != nil {
		l := logger.Singleton{}
		l.GetLogger().Errorln(err)
		return nil, err
	}

	return producer, nil
}

func getDefaultConfig() *sarama.Config {
	config := sarama.NewConfig()

	// Set the required ACKs from brokers for producer requests.
	// 0 = No responses required, 1 = Wait for leader acknowledgement, -1 = Wait for all in-sync replicas acknowledgement.
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	// The maximum number of times to retry sending a message.
	config.Producer.Retry.Max = 5

	// The maximum number of messages allowed in the producer queue.
	config.Producer.Flush.MaxMessages = 1000

	// The maximum duration the broker will wait before answering a Produce request.
	config.Producer.Timeout = 5 * time.Second

	// The compression algorithm to use for compressing message sets.
	config.Producer.Compression = sarama.CompressionSnappy

	return config
}
