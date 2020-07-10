package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ssuareza/memcache"
)

// Get memcache key/value.
func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	client, err := memcache.New(memcachedAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	item, err := client.Get(vars["key"])
	if err != nil {
		fmt.Println(err)
		return
	}

	// if key not exists
	if len(item.Value) == 0 {
		fmt.Fprintf(w, "%s\n", "key not found")
		return
	}

	fmt.Fprintf(w, "%s\n", string(item.Value))
}
