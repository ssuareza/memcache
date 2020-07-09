package memcache

import (
	"net"
	"time"
)

const (
	defaultTimeout    = 10 * time.Millisecond
	defaultExpiration = 900
)

// Client is a memcache client.
type Client struct {
	Addr string
	Conn net.Conn
}

// Item is an item to be got or stored in a memcached server.
type Item struct {
	Key   string
	Value []byte
}

// New returns a memcache client
func New(addr string) (*Client, error) {
	d := net.Dialer{Timeout: defaultTimeout}
	conn, err := d.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Client{
		Addr: addr,
		Conn: conn,
	}, nil
}
