package main

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

// Redis setup
var ctx = context.Background()
var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379", // Redis address
})

// WebSocket setup
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections (change for production security)
	},
}

var clients = make(map[*websocket.Conn]bool) // Track active WebSocket clients
var mutex = sync.Mutex{}                     // Prevent race conditions

// Handle WebSocket connections
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	// Add client to list
	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	log.Println("New WebSocket client connected")

	// Keep the connection open
	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			// Remove client on disconnect
			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()
			log.Println("WebSocket client disconnected")
			break
		}
	}
}

// Subscribe to Redis and forward messages to WebSocket clients
func redisSubscriber() {
	pubsub := redisClient.Subscribe(ctx, "chat")
	defer pubsub.Close()

	ch := pubsub.Channel()
	for msg := range ch {
		log.Println("Received from Redis:", msg.Payload)

		// Broadcast to all WebSocket clients
		mutex.Lock()
		for client := range clients {
			if err := client.WriteMessage(websocket.TextMessage, []byte(msg.Payload)); err != nil {
				log.Println("WebSocket send error:", err)
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}

func main() {
	// Start Redis subscriber in a goroutine
	go redisSubscriber()

	// Start WebSocket server
	http.HandleFunc("/ws", handleWebSocket)

	port := ":3000"
	log.Println("WebSocket server listening on ws://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
