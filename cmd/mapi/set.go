package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ssuareza/filesplit"
	"github.com/ssuareza/memcache"
)

// Set sets a new memcache key/value.
func Set(w http.ResponseWriter, r *http.Request) {
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
	if err := client.Set(item); err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}

	fmt.Fprintf(w, "STORED\n")
}

// SetFile sets a new memcache key/value from a file.
func SetFile(w http.ResponseWriter, r *http.Request) {
	// vars from uri
	vars := mux.Vars(r)
	key := vars["id"]

	// retrieve file
	file, err := ReceiveFileToStruct(r)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}

	// get md5
	fileMD5 := MD5(file.Content)

	// split it
	chunks, err := filesplit.SplitFromBytes(file.Content)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}

	// create index
	index := Index(key, fileMD5, chunks)

	// set memcache client
	client, err := memcache.New(memcachedAddr)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}

	// insert index
	indexItem := &memcache.Item{Key: key, Value: []byte(fmt.Sprint(index))}
	if err := client.Set(indexItem); err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}

	// insert chunks
	for _, v := range chunks {
		item := &memcache.Item{Key: key + v.Name, Value: v.Content}
		if err := client.Set(item); err != nil {
			fmt.Fprintf(w, "%s\n", err)
			return
		}
	}

	fmt.Fprintf(w, "STORED\n")

}

// Index creates a index of the file.
func Index(key, md5 string, chunks []filesplit.Chunk) map[string]interface{} {
	index := make(map[string]interface{}, 0)

	list := make([]string, 0)
	for _, v := range chunks {
		list = append(list, key+v.Name)
	}

	index["md5"] = md5
	index["chunks"] = list

	return index
}

// MD5 hashes using md5 algorithm.
func MD5(text []byte) string {
	algorithm := md5.New()
	algorithm.Write(text)
	return hex.EncodeToString(algorithm.Sum(nil))
}
