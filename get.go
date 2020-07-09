package memcache

import (
	"bufio"
	"fmt"
	"strings"
)

// Get gets the item for the given key
func (c *Client) Get(key string) (*Item, error) {
	var i Item

	// send get to socket
	fmt.Fprintf(c.Conn, "get %s\n", key)

	// get response
	scanner := bufio.NewScanner(c.Conn)

	for scanner.Scan() {
		line := scanner.Bytes()
		if string(line) == "END" {
			break
		}
		if strings.HasPrefix(string(line), "VALUE") {
			continue
		}
		i.Key = key
		i.Value = []byte(line)
	}

	return &i, nil
}

// GetMulti is a batch version of Get
func (c *Client) GetMulti(keys []string) (map[string]*Item, error) {
	// create a map to save all values
	m := make(map[string]*Item)

	var i Item
	for _, key := range keys {
		// send get to socket
		fmt.Fprintf(c.Conn, "get %s\n", key)

		// get response
		scanner := bufio.NewScanner(c.Conn)
		for scanner.Scan() {
			line := scanner.Bytes()
			if string(line) == "END" {
				break
			}
			if strings.HasPrefix(string(line), "VALUE") {
				i.Key = string(strings.Fields(string(line))[1])
				continue
			}
			i.Key = key
			i.Value = []byte(line)
		}
		m[key] = &i
	}

	return m, nil
}
