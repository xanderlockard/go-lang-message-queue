package main

import (
	"net/http"

	"github.com/xande/go-lang-task-queue/server"
)

func main() {
	config := server.ServerConfig{
		Address: "localhost",
		Port:    "8080",
		Handler: http.DefaultServeMux,
	}

	server.StartServer(config)
}
