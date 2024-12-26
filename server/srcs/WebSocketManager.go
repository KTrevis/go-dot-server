package main

import (
	"server/database"
	"sync"

	"github.com/gorilla/websocket"
)

type WebSocketManager struct {
	Clients map[*websocket.Conn]*Client
	db		*database.DB
	mutex   sync.Mutex
}

func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		Clients: make(map[*websocket.Conn]*Client),
	}
}

func (manager *WebSocketManager) AddClient(socket *websocket.Conn) {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()
	manager.Clients[socket] = &Client{
		manager: manager,
		socket:  socket,
	}
}

func (manager *WebSocketManager) RemoveClient(socket *websocket.Conn) {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()
	delete(manager.Clients, socket)
	socket.Close()
}
