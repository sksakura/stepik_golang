package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "test", 0)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = conn.WriteMessages(
		kafka.Message{
			Key:   []byte("Some Message"),
			Value: []byte("Hello World!"),
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("done")
}
