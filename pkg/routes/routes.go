package routes

import "chat-app/pkg/handlers"

func InitializeRoutes() {
	go handlers.HandleMessages()
}
