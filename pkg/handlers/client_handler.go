package handlers

import (
	"chat-app/pkg/chat"
	"chat-app/pkg/config"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

func Connect(username string) *websocket.Conn {
	conn, _, err := websocket.DefaultDialer.Dial(
		fmt.Sprintf("ws://%s%s/ws", config.WsAddr, config.WsPort), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	return conn
}

func HandleIncomingMessages(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		fmt.Printf("Received: %s\n", message)
	}
}

func HandleCliCommand(cmd, username string, conn *websocket.Conn) {
	parts := strings.SplitN(cmd, " ", 3)

	switch parts[0] {
	case "create":
		if len(parts) < 3 {
			fmt.Println("Usage: create [room] [password]")
			return
		}
		createRoom(parts[1], parts[2], conn)
	case "join":
		if len(parts) < 3 {
			fmt.Println("Usage: join [room] [password]")
			return
		}
		joinRoom(parts[1], parts[2], username, conn)
	case "leave":
		if len(parts) < 2 {
			fmt.Println("Usage: leave [room]")
			return
		}
		leaveRoom(parts[1], username, conn)
	case "delete":
		if len(parts) < 3 {
			fmt.Println("Usage: delete [room] [password]")
			return
		}
		deleteRoom(parts[1], parts[2], conn)
	case "send":
		if len(parts) < 3 {
			fmt.Println("Usage: send [room] [message]")
			return
		}
		sendMessage(parts[1], parts[2], username, conn)
	case "dm":
		if len(parts) < 3 {
			fmt.Println("Usage: dm [user] [message]")
			return
		}
		directMessage(parts[1], parts[2], username, conn)
	case "exit":
		conn.Close()
		fmt.Println("Disconnected from server")
		os.Exit(0)
	default:
		fmt.Println("Unknown command")
	}
}

func createRoom(room, password string, conn *websocket.Conn) {
	msg := chat.Message{Type: "create", Room: room, Password: password}
	conn.WriteJSON(msg)
}

func joinRoom(room, password, username string, conn *websocket.Conn) {
	msg := chat.Message{Type: "join", Room: room, Sender: username, Password: password}
	conn.WriteJSON(msg)
}

func leaveRoom(room, username string, conn *websocket.Conn) {
	msg := chat.Message{Type: "leave", Room: room, Sender: username}
	conn.WriteJSON(msg)
}

func deleteRoom(room, password string, conn *websocket.Conn) {
	msg := chat.Message{Type: "delete", Room: room, Password: password}
	conn.WriteJSON(msg)
}

func sendMessage(room, content, username string, conn *websocket.Conn) {
	msg := chat.Message{Type: "message", Room: room, Content: content, Sender: username}
	conn.WriteJSON(msg)
}

func directMessage(receiver, content, username string, conn *websocket.Conn) {
	msg := chat.Message{Type: "dm", Receiver: receiver, Content: content, Sender: username}
	conn.WriteJSON(msg)
}
