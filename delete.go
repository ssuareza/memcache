package memcache

import (
	"bufio"
	"fmt"
)

// Delete deletes the given key.
func (c *Client) Delete(key string) error {
	// send set to socket
	fmt.Fprintf(c.Conn, "delete %s\r\n", key)

	// get response
	scanner := bufio.NewScanner(c.Conn)
	if scanner.Scan() {
		if scanner.Text() == "DELETED" {
			return nil
		}
		return fmt.Errorf(scanner.Text())
	}

	return nil
}
