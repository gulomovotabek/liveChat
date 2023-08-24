package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"liveChatClient/models"
	"net"
)

func MessagesHandler(conn net.Conn, client *models.Client) {
	var (
		data = make([]byte, 1024)
		err  error
		n    int
	)
	for {
		n, err = conn.Read(data)

		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection lost!")
				return
			}
			fmt.Println(err)
			continue
		}

		var message InnerMessage

		err = json.Unmarshal(data[:n], &message)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			continue
		}

		switch message.MessageType {
		case SendMessage:
			fmt.Printf("%s: %s\n", message.SenderUsername, message.Context)
		case CreateGroup:
			client.GroupId = message.GroupId
			fmt.Printf("Group Id has been set %d\n", message.GroupId)
			fmt.Printf("%s: %s\n", message.SenderUsername, message.Context)
		case JoinGroup:
			client.GroupId = message.GroupId
			fmt.Printf("%s: %s\n", message.SenderUsername, message.Context)
		case ListGroup:
			fmt.Printf("%s: %s\n", message.SenderUsername, message.Context)
		default:
			fmt.Printf("%s: %s\n", message.SenderUsername, message.Context)
		}
	}
}
