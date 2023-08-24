package handler

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"log"
	"net"
)

func usernameHandler(conn net.Conn) (string, bool) {
	data := make([]byte, 1024)

	n, err := conn.Read(data)

	if err != nil {
		log.Println("Failed to read username from client: ", err)
		return "", false
	}

	var message Message

	err = json.Unmarshal(data[:n], &message)

	messageHeader := message.Header

	if messageHeader.MessageType != SetUsername {
		log.Println("Username didn't send properly:", err)
		return "", false
	}

	var messageBody SetUsernameMessage
	err = mapstructure.Decode(message.Body, &messageBody)
	if err != nil {
		return "", false
	}

	username := messageBody.Username

	// todo: send success message

	return username, true
}
