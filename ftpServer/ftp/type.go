package ftp

func (c *Conn) setDataType(args []string) {
	if len(args) == 0 {
		c.writeln(status501)
	}

	switch args[0] {
	case "A":
		c.binary = false
	case "I": // image/binary
		c.binary = true
	default:
		c.writeln(status504)
		return
	}
	c.writeln(status200)
}
