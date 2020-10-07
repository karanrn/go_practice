package ftp

import (
	"io"
	"log"
	"net"
)

func (c *Conn) dataConnect() (conn io.ReadWriteCloser, err error) {
	conn, err = net.Dial("tcp", c.dataHostPort)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return conn, nil
}
