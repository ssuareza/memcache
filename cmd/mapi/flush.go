package main

import (
	"fmt"
	"net/http"

	"github.com/ssuareza/memcache"
)

// Flush flush memcache content.
func Flush(w http.ResponseWriter, r *http.Request) {
	client, err := memcache.New(memcachedAddr)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
	}

	if err := client.FlushAll(); err != nil {
		fmt.Fprintf(w, "%s\n", err)
	}

	fmt.Fprintf(w, "FLUSHED\n")
}
