package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {
	username := os.Args[1]
	config, err := websocket.NewConfig(fmt.Sprintf("ws://:%v/", 1234), "http://")
	config.Header.Set("Username", username)
	if err != nil {
		log.Fatal(err)
	}
	connection, err := websocket.DialConfig(config)
	if err != nil {
		fmt.Println("Failed to establish connection")
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to your messages ", username)
	fmt.Println("---------------------")

	var lastMSG *message

	go func() {
		for {
			text, _ := reader.ReadString('\n')
			// convert CRLF to LF
			text = strings.Replace(text, "\n", "", -1)
			
			msg := message{
				Message: text,
				ID: uuid.New().String(),
			}
			err = websocket.JSON.Send(connection, msg)
			if err != nil {
				fmt.Println("Failed to send message")
				log.Fatal(err)
			}
			lastMSG = &msg
		}
	}()

	go func() {
		for {
			var response message
			err = websocket.JSON.Receive(connection, &response)
			if err != nil {
				log.Fatal(err)
			}
			if lastMSG != nil && lastMSG.ID == response.ID {
				continue
			} else {
				fmt.Println(response.Message)
			}
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for {
		select {
		case <-c:
			fmt.Println("Exiting")
			os.Exit(1)
		}
	}
}

type message struct {
	Message string `json:"message"`
	ID	  string `json:"id"`
}
