package main

import (
	"fmt"
	"net/http"
	"ws/server"
)

func main() {
	fs := http.FileServer(http.Dir("cmd/src"))
	http.Handle("/", fs)

	// Endpoint WebSocket
	http.HandleFunc("/ws", server.HandleConnections)

	// Goroutine to process messages
	go server.HandleMessages()

	fmt.Println("App running on port: 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
