package websockets

import (
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients   = make(map[uint]*websocket.Conn)
	clientsMu sync.Mutex
)

type WebSocketMessage struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type ClientMessage struct {
	Action string `json:"action"`
	Data   string `json:"data"`
}
