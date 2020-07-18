// go run reverb1 first and then:
// go run gopl.io/ch8/exercise/ex8.03
// Hello
//          HELLO
//          Hello
//          hello
// Wahahaha
//          WAHAHAHA
// ^D       Wahahaha
//          wahahaha
// 2020/07/18 14:09:39 done.

// old version netcat3 is output:
// go run gopl.io/ch8/netcat3
// Hello
//          HELLO
//          Hello
//          hello
// Hello
//          HELLO // Ctrl+D then program come over.
// 2020/07/18 14:10:04 done.

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
	conn.(*net.TCPConn).CloseWrite()
	// if connn, ok := conn.(*net.TCPConn); ok {
	//         connn.CloseWrite()
	// }
	// conn.Close()
	<-done // wait for background goroutine to finish
	// log print, then exit.
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
