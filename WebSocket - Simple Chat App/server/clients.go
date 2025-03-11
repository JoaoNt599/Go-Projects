package server

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// Channel for messages and map of connected clients
var Clients = make(map[*websocket.Conn]bool)
var Broadcast = make(chan Message)
var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
