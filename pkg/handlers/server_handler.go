package handlers

import (
	"chat-app/pkg/chat"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]string)
var broadcast = make(chan chat.Message)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WSHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	for {
		var msg chat.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		clients[ws] = msg.Sender
		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast
		chatService := chat.NewChat()

		switch msg.Type {
		case "message":
			// Check if the user is a member of the room
			isMember, err := chatService.IsUserInRoom(msg.Room, msg.Sender)
			if err != nil {
				log.Printf("error: checking: %v", err)
				msg.Content = "Error checking if user is in room"
			}
			if !isMember {
				log.Printf("error: not exist: %v", err)
				msg.Content = "You are not a member of this room"
			}
			chatService.SaveMessage(msg.Room, msg.Content)
		case "create":
			err := chatService.SetRoomPassword(msg.Room, msg.Password)
			msg.Content = fmt.Sprintf("Created room: %s", msg.Room)
			if err != nil {
				msg.Content = err.Error()
				log.Printf("error: %v", err)
			}
		case "join":
			err := chatService.CheckRoomPassword(msg.Room, msg.Password)
			msg.Content = fmt.Sprintf("Joined room: %s", msg.Room)
			if err != nil {
				msg.Content = err.Error()
				log.Printf("error: %v", err)
			}
			err = chatService.JoinRoom(msg.Room, msg.Sender)
			if err != nil {
				msg.Content = err.Error()
				log.Printf("error: %v", err)
			}
		case "leave":
			err := chatService.LeaveRoom(msg.Room, msg.Sender)
			msg.Content = fmt.Sprintf("Left room: %s", msg.Room)
			if err != nil {
				log.Printf("error: %v", err)
				msg.Content = err.Error()
			}
		case "delete":
			err := chatService.DeleteRoom(msg.Room, msg.Password)
			msg.Content = "Deleted room"
			if err != nil {
				msg.Content = err.Error()
				log.Printf("error: %v", err)
			}
		}

		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
