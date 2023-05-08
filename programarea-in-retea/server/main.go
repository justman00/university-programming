package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	port := 1234
	fmt.Println(fmt.Sprintf("starting chat server on PORT %v", port))
	chat := &chat{
		users: make([]*user, 0),
		emit:  make(chan message),
		event: make(chan string),
	}
	if err := chat.serve(port); err != nil {
		log.Fatal(err)
	}
}

type chat struct {
	users []*user
	emit  chan message
	event chan string
}

type message struct {
	Message string `json:"message"`
	ID	  string `json:"id"`
}

type user struct {
	username  string
	Connection *websocket.Conn
}

func (chat *chat) serve(port int) error {
	srv := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: chat.mux(),
	}
	go chat.run()
	return srv.ListenAndServe()
}

func (chat *chat) mux() http.Handler {
	mux := http.NewServeMux()
	// Use websocket.Server because we want to accept non-browser clients,
	// which do not send an Origin header. websocket.Handler does check
	// the Origin header by default.
	mux.Handle("/", websocket.Server{
		Handler: chat.handler(),
		// Set a Server.Handshake to nil - does not check the origin.
		// We can always provide a custom handshake method to access
		// the handshake http request and implement origin check or
		// other custom logic before the connection is established.
		Handshake: nil,
	})

	return mux
}

func (chat *chat) handler() func(*websocket.Conn) {
	return func(connection *websocket.Conn) {
		username := connection.Request().Header.Get("Username")
		chat.event <- fmt.Sprintf("%v connected", username)
		user := &user{
			username:  username,
			Connection: connection,
		}
		chat.users = append(chat.users, user)

		for {
			message := message{}
			err := websocket.JSON.Receive(connection, &message)
			if err != nil {
				fmt.Println("the error is: ", err.Error())
				// EOF connection closed by the client
				chat.event <- fmt.Sprintf("%v disconnected", username)
				return
			}
			chat.emit <- message
		}
	}
}

func (chat *chat) broadcast(message message) {
	for _, user := range chat.users {
		err := websocket.JSON.Send(user.Connection, message)
		if err != nil {
			log.Println(err)
		}
	}
}

func (chat *chat) run() {
	for {
		select {
		case message := <-chat.emit:
			fmt.Println("emitting message", message)
			chat.broadcast(message)
		case event := <-chat.event:
			fmt.Println(fmt.Sprintf("event: %s", event))
		}
	}
}