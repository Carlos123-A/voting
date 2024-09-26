package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jsusmachaca/voting/internal/config"
	"github.com/jsusmachaca/voting/pkg/model"
)

func worker() {
	var kafkaData model.KafkaData
	kafkaConfig := config.KafkaConfig()

	consumer, err := kafka.NewConsumer(&kafkaConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	err = consumer.Subscribe("voting", nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}
		json.Unmarshal(msg.Value, &kafkaData)
		log.Printf("message received: %s\n", msg.Value)
	}
}

func main() {
	go worker()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hola")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Printf("Server listen on http://localhost%s\n", server.Addr)
	server.ListenAndServe()
}
