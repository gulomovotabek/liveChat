package handler

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io"
	"liveChatServer/models"
	"liveChatServer/publisher"
	"log"
	"net"
)

func HandleClient(conn net.Conn, messageQueue chan publisher.InnerMessage, clients *[]models.Client, groups *[]*models.Group) {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println("Connection didn't close properly:", err)
		}
	}()

	username, isSuccess := usernameHandler(conn)

	if !isSuccess {
		return
	}

	client := models.Client{Username: username, Conn: conn}
	*clients = append(*clients, client)

	var (
		data    []byte
		n       int
		message Message
		err     error
	)

	for {
		data = make([]byte, 1024)

		n, err = conn.Read(data)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println(err)
			continue
		}

		message = Message{}
		err = json.Unmarshal(data[:n], &message)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			continue
		}

		switch message.Header.MessageType {
		case SendMessage:
			var messageBody SendMessageDTO
			err = mapstructure.Decode(message.Body, &messageBody)
			if err != nil {
				fmt.Println("Can't decode message body: ", err)
				continue
			}

			messageQueue <- publisher.InnerMessage{
				MessageType:    SendMessage,
				GroupId:        client.GroupId,
				SenderUsername: client.Username,
				Context:        messageBody.Message,
			}

		case ListGroup:
			listGroupHandler(conn, *groups)

		case CreateGroup:
			group, isSuccess := createGroupHandler(conn, client)
			if !isSuccess {
				continue
			}
			client.GroupId = group.Id
			*groups = append(*groups, &group)

		case JoinGroup:
			var messageBody JoinGroupDTO
			err = mapstructure.Decode(message.Body, &messageBody)
			if err != nil {
				fmt.Println("Couldn't decode message body: ", err)
				return
			}

			isSuccess = joinGroupHandler(conn, groups, messageBody.GroupId, &client)
			if !isSuccess {
				continue
			}

			messageQueue <- publisher.InnerMessage{
				MessageType:    SendMessage,
				GroupId:        messageBody.GroupId,
				SenderUsername: "System",
				Context:        fmt.Sprintf("%s has joined group", client.Username),
			}
		case LeaveGroup:
			// todo: not implemented

		case SetUsername:
			// todo: make validation for username
			username, isSuccess = usernameHandler(conn)
			if !isSuccess {
				continue
			}
			client.Username = username

		default:
			// todo: handle unexpected error
			break
		}
	}

	// Todo: remove user from group and send user left chat message

	err = conn.Close()
	if err != nil {
		return
	}
}
