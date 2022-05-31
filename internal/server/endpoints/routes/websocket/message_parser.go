package websocket

import "encoding/json"

const (
	CreatePost = iota
	SavePost
)

type Message struct {
	Type int         "json:type"
	Data interface{} "json:data"
}

func ParseMessage(message []byte) (error, *Message) {
	var messageParsed Message
	err := json.Unmarshal(message, messageParsed)
	if err != nil {
		return err, nil
	}
	return nil, &messageParsed
}
