package server

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/xande/go-lang-task-queue/queue"
)

var serverQueue = queue.CreateQueue()

type ServerConfig struct {
	Address string
	Port    string
	Handler http.Handler
}

func createServer(s ServerConfig) *http.Server {
	return &http.Server{
		Addr:           s.Address + ":" + s.Port,
		Handler:        s.Handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func enqueueMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	queable, _ := queue.DeserializeJsonIntoQueable(body)
	serverQueue.Enqueue(*queable)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is alive"))
}

func initializeHandlers() {
	// Enqueue message route
	http.HandleFunc("/enqueue", enqueueMessage)
	http.HandleFunc("/health", healthCheck)
}

func StartServer(config ServerConfig) {
	initializeHandlers()

	server := createServer(config)

	fmt.Println("Starting server on", config.Address+":"+config.Port)
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server error:", err)
	}
}
