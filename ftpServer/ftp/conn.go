package ftp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// Conn represents a connection to FTP server
type Conn struct {
	conn         net.Conn
	binary       bool
	dataHostPort string
	rootDir      string
	workDir      string
}

// NewConn returns a FTP connection
func NewConn(conn net.Conn, rootDir string) *Conn {
	return &Conn{
		conn:    conn,
		rootDir: rootDir,
		workDir: "/",
	}
}

const (
	status150 = "150 File status okay; about to open data connection."
	status200 = "200 Command okay."
	status215 = "215 UNIX Type: L8"
	status220 = "220 Service ready for new user."
	status221 = "221 Service closing control connection."
	status226 = "226 Closing data connection. Requested file action successful."
	status230 = "230 User %s logged in, proceed."
	status425 = "425 Can't open data connection."
	status426 = "426 Connection closed; transfer aborted."
	status501 = "501 Syntax error in parameters or arguments."
	status502 = "502 Command not implemented."
	status504 = "504 Cammand not implemented for that parameter."
	status550 = "550 Requested action not taken. File unavailable."
)

// Serve scans for incoming commands and routes them to handler function
func (c *Conn) Serve() {
	c.writeln(status220)
	var cmd string
	var args []string
	s := bufio.NewScanner(c.conn)
	for s.Scan() {
		fields := strings.Fields(s.Text())
		if len(fields) == 0 {
			continue
		}
		cmd = strings.ToUpper(fields[0])
		args = nil
		if len(fields) > 1 {
			args = fields[1:]
		}
		log.Printf("<< %s %v", cmd, args)

		switch cmd {
		case "CWD":
			c.cwd(args)
		case "RETR":
			c.retr(args)
		case "LIST":
			c.list(args)
		case "TYPE":
			c.setDataType(args)
		case "USER":
			c.writeln(status230)
		case "SYST":
			c.writeln(status215)
		case "PORT":
			c.Port(args)
		case "QUIT":
			c.writeln(status221)
			return
		default:
			c.writeln(status502)
		}
	}
	if s.Err() != nil {
		log.Print(s.Err())
	}
}

// writeln writes to ftp connection
func (c *Conn) writeln(msg string) {
	log.Print(">> ", msg)
	_, err := fmt.Fprintf(c.conn, msg+c.EOL())
	if err != nil {
		log.Print(err)
	}
}

// EOL returns the line terminator matching the FTP standard for the datatype.
func (c *Conn) EOL() string {
	if c.binary {
		return "\n"
	}
	return "\r\n"
}
