package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"gopl.io/ch8/exercise/ex8.01/kdama/client/console"
)

// Server is the server that reports the time.
// Has the server name, address, and latest output.
type Server struct {
	name, address, output string
}

func main() {
	servers, err := parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		conn, err := net.Dial("tcp", server.address)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go mustCopy(conn, server)
	}

	for {
		var data [][]string
		for _, server := range servers {
			data = append(data, []string{server.name, server.output})
		}

		table := console.SprintTable(data)
		console.Clear()
		fmt.Fprintf(os.Stdout, table)

		time.Sleep(time.Second)
	}
}

// mustCopy updates the value as the latest output of the server each time it reads the value from Reader.
func mustCopy(src io.Reader, server *Server) {
	sc := bufio.NewScanner(src)
	for sc.Scan() {
		server.output = sc.Text()
		if err := sc.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

// parse takes a string of the form "name=address" and returns a server that reports the time.
func parse(args []string) (servers []*Server, err error) {
	for _, arg := range args {
		s := strings.SplitN(arg, "=", 2)

		if len(s) != 2 {
			return nil, fmt.Errorf("failed to parse 'name=address': %s", arg)
		}

		name, address := s[0], s[1]
		servers = append(servers, &Server{name, address, ""})
	}
	return
}
