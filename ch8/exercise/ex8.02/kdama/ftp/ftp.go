package ftp

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func Serve(c Conn) {
	c.respond("220 Service ready for new user.")
	s := bufio.NewScanner(c.conn)
	for s.Scan() {
		input := strings.Fields(s.Text())
		if len(input) == 0 {
			continue
		}
	}
}

func (c *Conn) respond(s string) {
	log.Print(">>", s)
	_, err := fmt.Fprint(c.conn, s, c.eol())
	if err != nil {
		log.Print(err)
	}
}
