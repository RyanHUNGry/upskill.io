package producer

import (
	"log"
	"time"

	"github.com/IBM/sarama"
)

type AsyncLogGenerator struct {
	Producer sarama.AsyncProducer
}

func InitializeAsyncLogGenerator(brokerList []string, version sarama.KafkaVersion) (AsyncLogGenerator, error) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = version
	saramaConfig.Producer.Return.Successes = true
	// Only wait for the leader to acknowledge
	saramaConfig.Producer.RequiredAcks = sarama.WaitForLocal
	// Compress messages
	saramaConfig.Producer.Compression = sarama.CompressionSnappy
	// Flush batches every 500ms as Kafka buffers messages to encourage larger I/O over bursty, smaller I/O
	saramaConfig.Producer.Flush.Frequency = 500 * time.Millisecond
	// Retry up to 5 times to produce the message
	saramaConfig.Producer.Retry.Max = 5

	producer, err := sarama.NewAsyncProducer(brokerList, saramaConfig)

	if err != nil {
		return AsyncLogGenerator{}, err
	}

	// Log producer errors after all retry attempts are exhausted
	go func() {
		for err := range producer.Errors() {
			log.Println("Failed to write access log entry:", err)
		}
	}()

	return AsyncLogGenerator{Producer: producer}, nil
}

func (asyncLogGenerator *AsyncLogGenerator) Close() {
	if asyncLogGenerator.Producer != nil {
		if err := asyncLogGenerator.Producer.Close(); err != nil {
			log.Printf("Failed to close producer: %v", err)
		}
	}
}
