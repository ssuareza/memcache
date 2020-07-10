package main

import (
	"fmt"
	"net/http"
)

// Health returns status of the server.
func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
