package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

//func ChatHandler(response http.ResponseWriter, request *http.Request) {
//
//	username := request.URL.Query().Get("user")
//	fmt.Println(username)
//	connection, err := upgrader.Upgrade(response, request, nil)
//	if err != nil {
//		return
//	}
//	defer connection.Close()
//
//	for {
//
//		messageType, message, err := connection.ReadMessage()
//		if err != nil {
//			return
//		}
//
//		err = connection.WriteMessage(messageType, message)
//		if err != nil {
//			return
//		}
//
//	}
//}

func NewChatHandler(hub *Hub) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		username := request.URL.Query().Get("user")

		if username == "" {
			http.Error(response, "username is required", http.StatusBadRequest)
			return
		}

		connection, err := upgrader.Upgrade(response, request, nil)

		if err != nil {
			return
		}
		hub.Register(username, connection)

		defer connection.Close()

		defer hub.Unregister(username, connection)

		for {

			_, data, err := connection.ReadMessage()
			if err != nil {
				return
			}
			var message ChatMessage
			if err := json.Unmarshal(data, &message); err != nil {
				continue
			}

			if message.RecipientUsername == "" || message.Text == "" {
				continue
			}

			if err := hub.Send(message.RecipientUsername, []byte(message.Text)); err != nil {
				fmt.Println(err)
			}

		}

	}

}
func (hub *Hub) Register(username string, connection *websocket.Conn) {

	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	hub.connections[username] = connection
}

func (hub *Hub) Unregister(username string, connection *websocket.Conn) {

	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	if hub.connections[username] == connection {
		delete(hub.connections, username)
	}
}

func (hub *Hub) Send(username string, message []byte) error {

	hub.mutex.Lock()
	connection := hub.connections[username]
	hub.mutex.Unlock()

	if connection == nil {
		return fmt.Errorf(" username %q  not found", username)
	}

	return connection.WriteMessage(websocket.TextMessage, message)

}

var _ sync.Locker
