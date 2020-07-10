package main

import (
	"io/ioutil"
	"net/http"
)

// File contains the file received.
type File struct {
	Name    string
	Content []byte
}

// ReceiveFile retrieves the files sent from the client and save it.
func ReceiveFile(r *http.Request) (string, error) {
	var tmpFile string

	// get file
	r.ParseForm()
	file, handler, err := r.FormFile("file")
	if err != nil {
		return tmpFile, err
	}
	defer file.Close()

	// write temporary file on our server
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return tmpFile, err
	}

	tmpFile = tmpDir + handler.Filename
	err = ioutil.WriteFile(tmpFile, fileBytes, 0644)
	if err != nil {
		return tmpFile, err
	}

	return tmpFile, err
}

// ReceiveFileToStruct retrieves the files sent from the client and returns a struct.
func ReceiveFileToStruct(r *http.Request) (File, error) {
	var file File

	// get file
	r.ParseForm()
	f, handler, err := r.FormFile("file")
	if err != nil {
		return file, err
	}
	defer f.Close()

	// convert into bytes
	fileBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return file, err
	}

	file.Name = handler.Filename
	file.Content = fileBytes

	return file, nil
}
