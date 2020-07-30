package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp4", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string // an outgoing message channel
var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		//listens for events on the global messages channel,
		// to which each client sends all its incoming messages.
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			// Broadcasts the message to every connected client.
			for cli := range clients {
				cli <- msg
			}
		// The broadcaster listens on the global entering
		// and leaving channels for announcements of
		// arriving and departing clients.
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {

}
