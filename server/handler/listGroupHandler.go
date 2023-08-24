package handler

import (
	"encoding/json"
	"fmt"
	"liveChatServer/models"
	"liveChatServer/publisher"
	"net"
)

func listGroupHandler(conn net.Conn, groups []*models.Group) {
	messageContext := "Here is all groups:\n"
	for _, group := range groups {
		messageContext += fmt.Sprintf("ID: %d   -   %d active users\n", group.Id, len(group.Clients))
	}

	sendingMessage, err := json.Marshal(publisher.InnerMessage{
		MessageType:    ListGroup,
		GroupId:        0,
		SenderUsername: "System",
		Context:        messageContext,
	})
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	_, err = conn.Write(sendingMessage)
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}
}
