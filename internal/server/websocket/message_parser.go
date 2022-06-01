package websocket

import "encoding/json"

const (
	CreatePost = iota
	SavePost
)

type Message struct {
	Type int         `json:"type"`
	Data interface{} `json:"data"`
}

func ParseMessage(message []byte) (*Message, error) {
	var messageParsed Message
	err := json.Unmarshal(message, &messageParsed)
	if err != nil {
		return nil, err
	}
	return &messageParsed, nil
}
