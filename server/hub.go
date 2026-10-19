package server

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	connections map[string]*websocket.Conn
	mutex       sync.Mutex
}

func NewHub() *Hub {

	return &Hub{
		connections: make(map[string]*websocket.Conn),
	}
}
