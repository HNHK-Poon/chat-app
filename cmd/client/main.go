package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"chat-app/pkg/handlers"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	conn := handlers.Connect(username)
	defer conn.Close()

	go handlers.HandleIncomingMessages(conn)

	fmt.Println("\nCommands: \njoin [room] [password]\ncreate [room] [password]\nsend [room] [message]\ndm [user] [message]\ndelete [room] [password]\nexit")

	for {
		time.Sleep(500 * time.Millisecond)
		fmt.Print("Enter command: ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		if strings.ToLower(cmd) == "exit" {
			fmt.Println("Exiting chat application...")
			break
		}

		handlers.HandleCliCommand(cmd, username, conn)
	}
}
