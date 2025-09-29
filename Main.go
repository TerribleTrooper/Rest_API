package main

import (
	"Rest_API/http"
	"Rest_API/todo"
	"fmt"
)

func main() {
	todoList := todo.NewList()
	httpHandlers := http.NewHTTPHandlers(todoList)
	httpServer := http.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("Failed to start HTTP server", err)
	}
}
