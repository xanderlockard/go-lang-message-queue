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

func deserializeJsonIntoQueable(jsonStr string) (*Queable, error) {
	var queableJSON QueableJSON
	err := json.Unmarshal([]byte(jsonStr), &queableJSON)
	if err != nil {
		return nil, errors.New("failed to parse JSON: " + err.Error())
	}

	queable := &Queable{
		queueName:     queableJSON.QueueName,
		processId:     0,
		processStatus: "pending",
	}

	return queable, nil
}
