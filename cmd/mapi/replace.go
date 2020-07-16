package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ssuareza/memcache"
)

// Replace replaces a key.
func Replace(w http.ResponseWriter, r *http.Request) {
	// get vars from request
	vars := mux.Vars(r)
	key := vars["id"]
	value := r.FormValue("value")

	if len(value) == 0 {
		fmt.Fprintf(w, "you should specify \"value\"\n")
		return
	}

	// set memcache client
	client, err := memcache.New(memcachedAddr)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}

	// insert key+value
	item := &memcache.Item{Key: key, Value: []byte(value)}
	if err := client.Replace(item); err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}

	fmt.Fprintf(w, "STORED\n")
}

// ReplaceFile updates a memcache key from a file.
func ReplaceFile(w http.ResponseWriter, r *http.Request) {
	Delete(w, r)
	SetFile(w, r)
}
