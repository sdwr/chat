package main

import (
	"fmt"
	"strings"
	"encoding/json"
)

type Message struct {
	Type string `json:"Type"`
	Text string `json:"Text"`
}

type MessageError struct {
	Text string
}

type Command struct {
	Type string
	Args []string
}

// func ParseJSON(data []byte) interface{} {
// 	var msg Message
// 	json.Unmarshal(data, &msg)
// 	switch msg.Type {
// 	case "COMMAND":
		
// 	}
// }

func Parse(msg *Message) (interface{}, error) {
	if msg == nil  {
		return nil, MessageError{"No message recieved"}
	}
	switch msg.Type {
	case "COMMAND":
		command := strings.Split(msg.Text, " ")
		return &Command{msgAr[0], msgAr[1:]}, nil		
	}

	return nil, MessageError{"Unknown message type"}
	
}

func (e MessageError) Error() string  {
	return fmt.Sprintf("%s", e.Text)
}

