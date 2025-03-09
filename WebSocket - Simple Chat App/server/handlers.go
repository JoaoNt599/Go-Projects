package server

import (
	"fmt"
	"net/http"
)

// HandleConnections manages WebSocket connections
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer ws.Close()

	Clients[ws] = true

	for {
		var msg Message
		if err := ws.ReadJSON(&msg); err != nil {
			delete(Clients, ws)
			break
		}
		Broadcast <- msg
	}
}

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
