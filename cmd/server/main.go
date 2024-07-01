package main

import (
	"log"
	"net/http"

	"chat-app/pkg/config"
	"chat-app/pkg/handlers"
	"chat-app/pkg/routes"
)

func main() {
	http.HandleFunc("/ws", handlers.WSHandler)
	routes.InitializeRoutes()

	log.Printf("Server started on %s\n", config.WsAddr+config.WsPort)
	err := http.ListenAndServe(config.WsPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
