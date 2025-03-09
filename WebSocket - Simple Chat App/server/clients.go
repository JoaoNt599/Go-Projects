package server

import "github.com/gorilla/websocket"

// Channel for messages and map of connected clients
var Broadcast = make(chan Message)
var Clients = make(map[*websocket.Conn]bool)
