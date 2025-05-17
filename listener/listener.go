package listener

import (
	"fmt"
	"mqtt"
)

type ListenerMessage struct {
	MessageNames []string
}

type ListenerConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Handlers map[string]func()
}

type ListenerConnection struct {
	config ListenerConfig
	client mqtt.Client
}

// Here we will take the message from the worker and add it to list of available workers
func (connection ListenerConnection) handleMessage(message ListenerMessage) {

}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected on host: %s")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func initializeListeners() {

}

func connectToBroker() {

}
