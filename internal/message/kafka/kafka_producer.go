package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2"
	"os"
)

func ProduceKafkaSender() *kafka_sarama.Sender {
	kafkaHost := os.Getenv("KAFKA_HOST")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = sarama.V2_0_0_0
	sender, _ := kafka_sarama.NewSender([]string{kafkaHost}, saramaConfig, kafkaTopic)
	return sender
}
