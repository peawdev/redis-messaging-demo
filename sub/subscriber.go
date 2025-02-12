package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	defer client.Close()

	// Subscribe to a channel
	channel := "chat"
	pubsub := client.Subscribe(ctx, channel)
	defer pubsub.Close()

	fmt.Println("Subscribed to channel:", channel)

	// Listen for messages
	for msg := range pubsub.Channel() {
		fmt.Printf("Received message: %s\n", msg.Payload)
	}
}
