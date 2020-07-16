package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"io"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8000, "port number")
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	flag.Parse()
	server := fmt.Sprintf("localhost:%d", port)
	listener, err := net.Listen("tcp", server)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Listening at localhost:%d\n", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
