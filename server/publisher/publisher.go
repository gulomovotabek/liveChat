package publisher

import (
	"encoding/json"
	"fmt"
	"liveChatServer/models"
	"log"
)

func BroadcastMessenger(messageQueue chan InnerMessage, groups *[]*models.Group) {
	for message := range messageQueue {
		group, found := models.GetItemByID(*groups, message.GroupId)

		if !found {
			fmt.Printf("Group with ID %s not found\n", message.GroupId)
			continue
		}

		data, err := json.Marshal(message)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}

		for _, client := range group.Clients {
			if client.Username == message.SenderUsername {
				continue
			}
			_, err = client.Conn.Write(data)
			if err != nil {
				log.Println("Failed to send message to client:", err)
			}
		}
	}
}
