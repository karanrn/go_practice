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
	conn     net.Conn
	binary   bool
	dataPort *dataPort
	rootDir  string
	workDir  string
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

	s := bufio.NewScanner(c.conn)
	for s.Scan() {
		input := strings.Fields(s.Text())
		if len(input) == 0 {
			continue
		}

		command, args := input[0], input[1:]
		log.Printf("<< %s %v", command, args)

		switch command {
		case "CWD": // cd
			//c.cwd(args)
		case "RETR": // get
			//c.retr(args)
		case "LIST": // ls
			c.list(args)
		case "TYPE":
			//c.setDataType(args)
		case "USER":
			c.writeln(status230)
		case "SYST":
			c.writeln(status215)
		case "PORT":
			c.Port(args)
		case "LPRT":
			// TODO: Implement LPRT
			c.writeln(c.dataPort.toAddress())
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
	_, err := fmt.Fprintf(c.conn, msg, c.EOL())
	if err != nil {
		log.Print(err)
	}
}

// Port sets port for ftp connection
func (c *Conn) Port(args []string) {
	if len(args) != 1 {
		c.writeln(status501)
		return
	}
	var err error
	c.dataPort, err = dataPortFromHostPort(args[0])
	if err != nil {
		log.Printf("cmd: PORT, err: %v", err)
		c.writeln(status501)
		return
	}
	c.writeln(status200)
}

// EOL returns the line terminator matching the FTP standard for the datatype.
func (c *Conn) EOL() string {
	if c.binary {
		return "\n"
	}
	return "\r\n"
}
