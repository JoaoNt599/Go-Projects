package main

import (
	"fmt"
	"log"
	"net/http"

	"ws/rabbitmq"
	"ws/server"
)

func main() {
	// Connect to RabbitMQ
	conn, ch := rabbitmq.ConnectRabbitMQ()
	defer conn.Close()
	defer ch.Close()

	// Start consuming messages from RabbitMQ
	go server.ConsumeMessages(ch)

	// WebSocket endpoint
	http.HandleFunc("/ws", server.HandleConnections)

	// Serve static files (HTML, CSS, JS)
	fs := http.FileServer(http.Dir("src"))
	http.Handle("/", http.StripPrefix("/", fs))

	fmt.Println("Server running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
