package ftp

import "fmt"

func hostPortFromFTP(address string) (string, error) {
	var a, b, c, d byte
	var p1, p2 int
	_, err := fmt.Sscanf(address, "%d,%d,%d,%d,%d,%d", &a, &b, &c, &d, &p1, &p2)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d.%d.%d.%d:%d", a, b, c, d, 256*p1+p2), nil
}

// Port sets port for ftp connection
func (c *Conn) Port(args []string) {
	if len(args) != 1 {
		c.writeln("501 Usage: PORT a,b,c,d,p1,p2")
		return
	}
	var err error
	c.dataHostPort, err = hostPortFromFTP(args[0])
	if err != nil {
		c.writeln("501 Can't parse address.")
		return
	}
	c.writeln(status200)
}
