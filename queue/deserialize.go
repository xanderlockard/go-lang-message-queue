package queue

import (
	"encoding/json"
	"errors"
)

// JSON structure expected for deserialization
type QueableJSON struct {
	QueueName     string   `json:"queueName"`
	MessageParams []string `json:"messageParams"`
}

type HandleableMessage struct {
	MessageName   string   `json:"messageName"`
	MessageParams []string `json:"messageParams"`
}

func DeserializeJsonIntoQueable(requestBody []byte) (*Queable, error) {
	var queableJSON QueableJSON
	err := json.Unmarshal(requestBody, &queableJSON)
	if err != nil {
		return nil, errors.New("failed to parse JSON: " + err.Error())
	}

	queable := &Queable{
		QueueName:     queableJSON.QueueName,
		ProcessId:     0,
		ProcessStatus: "pending",
	}

	return queable, nil
}

func DeserializeMessageIntoMqttParam(messageBody []byte) (*ListenableMessage, error)
