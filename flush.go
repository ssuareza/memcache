package memcache

import (
	"bufio"
	"fmt"
)

// FlushAll flush all the contents in memcache.
func (c *Client) FlushAll() error {
	// send flush_all to socket
	fmt.Fprintf(c.Conn, "flush_all\r\n")

	// get response
	scanner := bufio.NewScanner(c.Conn)
	if scanner.Scan() {
		if scanner.Text() == "OK" {
			return nil
		}
		return fmt.Errorf(scanner.Text())
	}

	return nil
}
