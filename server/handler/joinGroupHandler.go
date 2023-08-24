package handler

import (
	"encoding/json"
	"fmt"
	"liveChatServer/models"
	"liveChatServer/publisher"
	"net"
)

func joinGroupHandler(conn net.Conn, groups *[]*models.Group, groupID int, client *models.Client) bool {
	group, found := models.GetItemByID(*groups, groupID)

	if !found {
		err := fmt.Errorf("Group with ID %d not found\n", groupID)
		fmt.Print(err)
		return false
	}

	group.Clients = append(group.Clients, client)
	client.GroupId = group.Id

	sendingMessage := publisher.InnerMessage{
		MessageType:    JoinGroup,
		GroupId:        group.Id,
		SenderUsername: client.Username,
		Context:        fmt.Sprintf("You've joined the group(id: %d)", groupID),
	}

	sendingMessageBytes, err := json.Marshal(sendingMessage)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return false
	}

	_, err = conn.Write(sendingMessageBytes)
	if err != nil {
		fmt.Println("Error sending data:", err)
		return false
	}
	return true
}
