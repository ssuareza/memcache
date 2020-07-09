package main

import (
	"fmt"
	"log"

	"github.com/ssuareza/memcache"
)

func main() {
	client, err := memcache.New("localhost:11211")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Conn.Close()

	// set and get
	item := &memcache.Item{Key: "dog", Value: []byte("dago")}
	err = client.Set(item)
	if err != nil {
		log.Fatal(err)
	}

	value, err := client.Get("dog")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
}
