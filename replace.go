package memcache

import (
	"bufio"
	"fmt"
)

// Replace replaces the given item.
func (c *Client) Replace(item *Item) error {
	// send set to socket
	fmt.Fprintf(c.Conn, "replace %s 0 %v %v\r\n%s\r\n", item.Key, defaultExpiration, len(item.Value), item.Value)

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
