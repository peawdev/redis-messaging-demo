package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Change this if Redis is on a different host/port
	})

	defer client.Close()

	channel := "chat"
	for i := 1; i <= 5; i++ {
		message := fmt.Sprintf("Hello %d from publisher!", i)
		err := client.Publish(ctx, channel, message).Err()
		if err != nil {
			log.Fatalf("Error publishing message: %v", err)
		}
		fmt.Printf("Published: %s\n", message)
		time.Sleep(1 * time.Second) // Simulate delay
	}
}
