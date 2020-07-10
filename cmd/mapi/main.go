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
	router.HandleFunc("/get/{key}", Get).Methods("GET")
	router.HandleFunc("/set/{key}", Set).Methods("POST")
	router.HandleFunc("/file/{key}", SetFile).Methods("POST")

	fmt.Println("starting")
	log.Fatal(http.ListenAndServe(port, router))
}
