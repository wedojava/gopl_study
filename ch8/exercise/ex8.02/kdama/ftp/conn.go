package ftp

import (
	"log"
	"net"
)

type Conn struct {
	conn     net.Conn
	dataport *dataport
	datatype datatype
	rootDir  string
	workDir  string
}

func NewConn(conn net.Conn, rootDir string) Conn {
	return Conn{
		conn:    conn,
		rootDir: rootDir,
		workDir: "/",
	}
}

func (c *Conn) Close() error {
	err := c.conn.Close()
	if err != nil {
		log.Print(err)
	}
	return err
}
