package memcache

import (
	"bufio"
	"fmt"
)

// Set writes the given item.
func (c *Client) Set(item *Item) error {
	// send set to socket
	fmt.Fprintf(c.Conn, "add %s 0 %v %v\r\n%s\r\n", item.Key, defaultExpiration, len(item.Value), item.Value)

	// get response
	scanner := bufio.NewScanner(c.Conn)
	if scanner.Scan() {
		if scanner.Text() == "STORED" {
			return nil
		}
		return fmt.Errorf(scanner.Text())
	}

	return nil
}
