package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

// This script will try and connect to the topic leader within the timeout, else return an os.Exit(1).
// We connect to the topic leader to ensure a leadership election succeeds and the topic is ready to be consumed.
func main() {
	kafkaEndpoint := os.Getenv("KAFKA_ENDPOINT")
	if kafkaEndpoint == "" {
		log.Fatal("missing KAFKA_ENDPOINT environment variable")
	}

	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		log.Fatal("missing KAFKA_TOPIC environment variable")
	}

	timeout := time.After(2 * time.Minute)
	for {
		select {
		case <-timeout:
			log.Fatal("timeout trying to connect to kafka")
		default:
			_, err := kafka.DialLeader(context.Background(), "tcp", kafkaEndpoint, topic, 0)
			if err == nil {
				fmt.Println("connected successfully to kafka")
				return
			}
			log.Println("failed to dial leader:", err)
			time.Sleep(10 * time.Second)
		}
	}
}
