package handler

import (
	"encoding/json"
	"fmt"
	"liveChatServer/models"
	"liveChatServer/publisher"
	"net"
)

func createGroupHandler(conn net.Conn, client models.Client) (models.Group, bool) {
	group := models.CreateGroup(client, true)
	sendingMessage := publisher.InnerMessage{
		MessageType:    CreateGroup,
		GroupId:        group.Id,
		SenderUsername: client.Username,
		Context:        fmt.Sprintf("You've created group: %d", group.Id),
	}

	sendingMessageBytes, err := json.Marshal(sendingMessage)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return models.Group{}, false
	}

	_, err = conn.Write(sendingMessageBytes)
	if err != nil {
		fmt.Println("Error sending data:", err)
		return models.Group{}, false
	}

	return group, true
}
