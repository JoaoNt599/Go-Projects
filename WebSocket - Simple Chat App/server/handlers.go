package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

// Handle Messages distributes messages to connected clients
func HandleMessages() {
	for {
		msg := <-Broadcast
		for client := range Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				delete(Clients, client)
				client.Close()
			}
		}
	}
}

// Handle new WebSocket connections
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer ws.Close()

	Clients[ws] = true
	log.Println("New client connected")

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Println("Read error:", err)
			delete(Clients, ws)
			break
		}
		Broadcast <- msg
	}
}

// Publish messages to RabbitMQ
func PublishMessage(ch *amqp.Channel, msg Message) {
	body, _ := json.Marshal(msg)
	err := ch.Publish("", "chat_messages", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
	if err != nil {
		log.Println("Failed to publish message:", err)
	}
}

// Consume messages from RabbitMQ
func ConsumeMessages(ch *amqp.Channel) {
	msgs, err := ch.Consume("chat_messages", "", true, false, false, false, nil)
	if err != nil {
		log.Fatal("Failed to consume messages:", err)
	}

	for msg := range msgs {
		var receivedMsg Message
		json.Unmarshal(msg.Body, &receivedMsg)

		// Broadcast to all connected clients
		for client := range Clients {
			err := client.WriteJSON(receivedMsg)
			if err != nil {
				log.Println("Failed to send message:", err)
				client.Close()
				delete(Clients, client)
			}
		}
	}
}
