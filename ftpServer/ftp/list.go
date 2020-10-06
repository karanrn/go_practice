package ftp

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func (c *Conn) list(args []string) {
	var target string
	if len(args) > 0 {
		target = filepath.Join(c.rootDir, c.workDir, args[0])
	} else {
		target = filepath.Join(c.rootDir, c.workDir)
	}

	files, err := ioutil.ReadDir(target)
	if err != nil {
		log.Print(err)
		c.writeln(status550)
		return
	}
	c.writeln(status150)

	dataConn, err := c.dataConnect()
	if err != nil {
		log.Print(err)
		c.writeln(status425)
		return
	}
	defer dataConn.Close()

	for _, file := range files {
		_, err := fmt.Fprintf(dataConn, file.Name(), c.EOL())
		if err != nil {
			log.Print(err)
			c.writeln(status426)
		}
	}
	_, err = fmt.Fprintf(dataConn, c.EOL())
	if err != nil {
		log.Print(err)
		c.writeln(status426)
	}
	c.writeln(status226)
}
