package main

import (
	"fmt"
	"net/http"
)

// Update updates a memcache key.
func Update(w http.ResponseWriter, r *http.Request) {
	value := r.FormValue("value")

	if len(value) == 0 {
		fmt.Fprintf(w, "you should specify \"value\"\n")
		return
	}

	Delete(w, r)
	Set(w, r)
}

// UpdateFile updates a memcache key from a file.
func UpdateFile(w http.ResponseWriter, r *http.Request) {
	Delete(w, r)
	SetFile(w, r)
}
