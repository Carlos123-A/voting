package config

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func KafkaConfig() kafka.ConfigMap {
	configMap := kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_SERVER"),
		"group.id":          "vote-consumer",
		"auto.offset.reset": "earliest",
	}

	return configMap
}
