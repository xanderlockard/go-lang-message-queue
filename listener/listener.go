package listener

import (
	"fmt"
	"mqtt"
)

type ListenableMessage struct {
	MessageName   string
	MessageParams string
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

func (connection ListenerConnection) handleMessage() {

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
