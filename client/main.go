package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"liveChatClient/handler"
	"liveChatClient/models"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

var client models.Client

func main() {
	host := flag.String("host", "localhost", "host domain or IP address")
	port := flag.String("port", "1337", "port number")
	flag.Parse()
	conn, err := net.Dial("tcp", net.JoinHostPort(*host, *port))
	if err != nil {
		log.Fatal("Failed to connect to server: ", err)
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	sendingMessage := handler.Message{
		Header: handler.MessageHeader{MessageType: handler.SetUsername},
		Body:   handler.SetUsernameMessage{Username: username},
	}

	sendingMessageBytes, err := json.Marshal(sendingMessage)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	_, err = conn.Write(sendingMessageBytes)
	if err != nil {
		log.Fatal("Failed to send username to server:", err)
	}

	client = models.Client{Username: username, GroupId: 0}

	go handler.MessagesHandler(conn, &client)

	fmt.Println("You are registered as:", username)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Failed to read input:", err)
			continue
		}
		input = input[:len(input)-1]

		splittedInput := strings.Split(input, " ")

		switch splittedInput[0] {
		case "/create_group":
			sendingMessage = handler.Message{Header: handler.MessageHeader{MessageType: handler.CreateGroup}}

			sendingMessageBytes, err = json.Marshal(sendingMessage)
			if err != nil {
				fmt.Println("Error marshaling JSON:", err)
				continue
			}

			_, err = conn.Write(sendingMessageBytes)
			if err != nil {
				fmt.Println("Error sending data:", err)
				continue
			}
		case "/join_group":
			groupId, err := strconv.Atoi(splittedInput[1])
			if err != nil {
				fmt.Println("Can't parse group ID, please send correct group ID as integer")
				continue
			}

			sendingMessage := handler.Message{
				Header: handler.MessageHeader{MessageType: handler.JoinGroup},
				Body:   handler.JoinGroupDTO{GroupId: groupId},
			}

			sendingMessageBytes, err := json.Marshal(sendingMessage)
			if err != nil {
				fmt.Println("Error marshaling JSON:", err)
				continue
			}

			_, err = conn.Write(sendingMessageBytes)
			if err != nil {
				fmt.Println("Error sending data:", err)
				continue
			}

		case "/list_group":
			sendingMessage := handler.Message{Header: handler.MessageHeader{MessageType: handler.ListGroup}}

			sendingMessageBytes, err := json.Marshal(sendingMessage)
			if err != nil {
				fmt.Println("Error marshaling JSON:", err)
				continue
			}

			_, err = conn.Write(sendingMessageBytes)
			if err != nil {
				fmt.Println("Error sending data:", err)
				continue
			}
		case "/leave_group":
			// leave
		default:
			if client.GroupId == 0 {
				fmt.Println("please first join group or create group")
				continue
			}

			sendingMessage := handler.Message{
				Header: handler.MessageHeader{MessageType: handler.SendMessage},
				Body:   handler.SendMessageDTO{Message: input},
			}

			sendingMessageBytes, err := json.Marshal(sendingMessage)
			if err != nil {
				fmt.Println("Error marshaling JSON:", err)
				continue
			}

			_, err = conn.Write(sendingMessageBytes)
			if err != nil {
				fmt.Println("Error sending data:", err)
				continue
			}
		}
	}
}
