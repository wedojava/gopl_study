package main

import (
	"log"
	"net"
	"os"

	"io"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		// Closing the read half network connection causes
		// the background goroutine’s call to io.Copy
		// to return a ‘read from closed connection’ error,
		// which is why we’ve removed the error logging
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		// now this background goroutine logs a message
		log.Println("done.")
		// send a value n the done channel
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // wait for background goroutine to finish
	// log print, then exit.
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
