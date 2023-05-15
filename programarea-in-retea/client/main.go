package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"

	"github.com/google/uuid"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {
	username := os.Args[1]
	c, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Welcome to your messages ", username)
	fmt.Println("---------------------")

	var lastMSG *message

	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)

			msg := message{
				Message: text,
				ID:      uuid.New().String(),
			}

			marshalledMsg, err := json.Marshal(msg)
			if err != nil {
				fmt.Println("marshal message")
				log.Fatal(err)
			}
			fmt.Fprintf(c, string(marshalledMsg)+"\n")

			lastMSG = &msg
		}
	}()

	go func() {
		for {
			m, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				fmt.Println("read message")
				log.Fatal(err)
			}
			
			var msg message
			err = json.Unmarshal([]byte(m), &msg)
			if err != nil {
				fmt.Println("unmarshal message", string(m))
				log.Fatal(err)
			}
			
			if lastMSG != nil && lastMSG.ID == msg.ID {
				continue
			} else {
				fmt.Println(msg.Message)
			}

		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	for {
		select {
		case <-ch:
			fmt.Println("Exiting")
			os.Exit(1)
		}
	}
}

type message struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}
