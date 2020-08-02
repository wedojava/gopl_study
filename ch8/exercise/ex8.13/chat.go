package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

const timeout = 10 * time.Second // kick out client out of this time

type Client struct {
	Out  chan<- string // an outgoing message channel
	Name string
}

var (
	entering = make(chan Client)
	leaving  = make(chan Client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[Client]bool) // all connected clients
	for {
		select {
		//listens for events on the global messages channel,
		// to which each client sends all its incoming messages.
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			// Broadcasts the message to every connected client.
			for cli := range clients {
				cli.Out <- msg
			}
		// The broadcaster listens on the global entering
		// and leaving channels for announcements of
		// arriving and departing clients.
		case cli := <-entering:
			clients[cli] = true
			cli.Out <- "Current Presents:"
			for c := range clients {
				cli.Out <- c.Name
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.Out)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	cli := Client{ch, who}
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	timer := time.NewTimer(timeout)
	go func() {
		<-timer.C
		conn.Close()
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
		timer.Reset(timeout)
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

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
