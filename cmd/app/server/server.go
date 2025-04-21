package server

import (
	"io/ioutil"
	"net/http"
	"time"
)

var serverQueue = LinkedListQueue{}

type ServerConfig struct {
	address string
	port    string
	handler http.Handler
}

func createServer(s ServerConfig) *http.Server {
	return &http.Server{
		Addr:           s.address + ":" + s.port,
		Handler:        s.handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func enqueueMessage(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	serverQueue.Enqueue()
}

func initializeHandlers() {
	// Enqueue message route
	http.HandleFunc("/enqueue", enqueueMessage)
}
