package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	port := 1234
	fmt.Println(fmt.Sprintf("starting chat server on PORT %v", port))
	chat := &chat{}

	err := chat.serve(port)
	if err != nil {
		log.Fatal(err)
	}
}

type chat struct {
	conns []net.Conn
}

type message struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}

func (chat *chat) serve(port int) error {
	l, err := net.Listen("tcp4", fmt.Sprintf(":%v", port))
	if err != nil {
		return fmt.Errorf("could not listen on port %v: %v", port, err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			return fmt.Errorf("could not accept connection: %v", err)
		}
		chat.conns = append(chat.conns, c)
		defer c.Close()
		go chat.handler(c)
	}
}

func (chat *chat) handler(c net.Conn) error {
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			return fmt.Errorf("could not read from connection: %v", err)
		}

		fmt.Println("received: ", netData)
		chat.broadcast(netData)
	}
}

func (chat *chat) broadcast(message string) {
	for _, c := range chat.conns {
		c.Write([]byte(message))
	}
}