package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"path/filepath"

	"github.com/karanrn/go_practice/ftpServer/ftp"
)

func main() {
	port := flag.Int("port", 8000, "listen port")
	flag.Parse()

	ln, err := net.Listen("tcp4", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal("opening main listener:", err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			log.Print("Accepting new connection:", err)
		}
		go ftp.NewConn(c, "/home/karan/Downloads/Books").Serve()
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	absPath, err := filepath.Abs("/home/karan/Downloads/Books")
	if err != nil {
		log.Fatal(err)
	}
	ftp.NewConn(c, absPath).Serve()
}
