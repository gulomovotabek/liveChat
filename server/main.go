package main

import (
	"fmt"
	"liveChatServer/handler"
	"liveChatServer/models"
	"liveChatServer/publisher"
	"log"
	"net"
)

func main() {
	runServer()
}

func runServer() {
	listener, err := net.Listen("tcp", ":1337")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
	defer listener.Close()

	fmt.Println("Server is running. Accepting connections...")

	groups := []*models.Group{}
	messageQueue := make(chan publisher.InnerMessage)

	go publisher.BroadcastMessenger(messageQueue, &groups)

	clients := []models.Client{}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection:", err)
			continue
		}

		go handler.HandleClient(conn, messageQueue, &clients, &groups)
	}
}
