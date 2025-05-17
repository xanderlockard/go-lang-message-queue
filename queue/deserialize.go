package queue

import (
	"encoding/json"
	"errors"

	"github.com/xande/go-lang-task-queue/listener"
)

type QueableJSON struct {
	MessageName   string            `json:"messageName"`   // this is the name of the event to queue
	MessageParams map[string]string `json:"messageParams"` // these are the parameters sent with the server request for the event.
}

type ListenerMessageJson struct {
	EventNames []string `json:"messageNames"` // this is the list of the events the worker can handle
}

func DeserializeJsonIntoQueable(requestBody []byte) (*Queable, error) {
	var queableJson QueableJSON
	err := json.Unmarshal(requestBody, &queableJson)
	if err != nil {
		return nil, errors.New("failed to parse Json: " + err.Error())
	}

	queable := &Queable{
		MessageName:   queableJson.MessageName,
		MessageParams: queableJson.MessageParams,
	}

	return queable, nil
}

func DeserializeMessageIntoMqttParam(messageBody []byte) (*listener.ListenerMessage, error) {
	var listenerMessageJson ListenerMessageJson
	err := json.Unmarshal(messageBody, &listenerMessageJson)
	if err != nil {
		return nil, errors.New("failed to parse Json: " + err.Error())
	}

	listenerMessage := &listener.ListenerMessage{
		MessageNames: listenerMessageJson.EventNames,
	}

	return listenerMessage, nil
}
