package server

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader for WebSocket
var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}
