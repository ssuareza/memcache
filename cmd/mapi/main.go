package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	port          = ":8090"
	tmpDir        = "/tmp/"
	memcachedAddr = "127.0.0.1:11211"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", Health).Methods("GET")
	router.HandleFunc("/flush", Flush).Methods("POST")
	router.HandleFunc("/key/{id}", Get).Methods("GET")
	router.HandleFunc("/key/{id}", Set).Methods("POST")
	router.HandleFunc("/key/{id}", Update).Methods("PUT")
	router.HandleFunc("/key/{id}", Delete).Methods("DELETE")
	router.HandleFunc("/file/{id}", SetFile).Methods("POST")
	router.HandleFunc("/file/{id}", UpdateFile).Methods("PUT")

	fmt.Println("starting")
	log.Fatal(http.ListenAndServe(port, router))
}
