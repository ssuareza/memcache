package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ssuareza/memcache"
)

// Delete deletes a memcache key.
func Delete(w http.ResponseWriter, r *http.Request) {
	// get vars from request
	vars := mux.Vars(r)
	key := vars["id"]

	// set memcache client
	client, err := memcache.New(memcachedAddr)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}

	// delete key
	if err := client.Delete(key); err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}

	fmt.Fprintf(w, "DELETED\n")
}
