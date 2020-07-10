package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/ssuareza/filesplit"
)

func TestReceiveFile(t *testing.T) {
	// create test file
	file := "../../test/file.dat"
	size := int64(512) // 512 bytes
	if err := filesplit.CreateTestFile(file, size); err != nil {
		t.Errorf("Not able to create \"%s\" file", file)
	}

	// open the file
	data, err := os.Open(file)
	if err != nil {
		t.Error(err)
	}
	defer data.Close()

	// create the body
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(data.Name()))
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(part, data)
	writer.Close()

	// create the request
	req, err := http.NewRequest("POST", "", body)
	if err != nil {
		t.Error(err)
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	// and test it!
	tmpFile, err := ReceiveFile(req)
	if err != nil {
		t.Error(err)
	}

	expected := "/tmp/file.dat"
	if tmpFile != expected {
		t.Errorf("File retrieved wrong: got %s want %s\n", tmpFile, expected)
	}
}

func TestReceiveFileToStruct(t *testing.T) {
	// create test file
	file := "../../test/file.dat"
	size := int64(512) // 512 bytes
	if err := filesplit.CreateTestFile(file, size); err != nil {
		t.Errorf("Not able to create \"%s\" file", file)
	}

	// open the file
	data, err := os.Open(file)
	if err != nil {
		t.Error(err)
	}
	defer data.Close()

	// create the body
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(data.Name()))
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(part, data)
	writer.Close()

	// create the request
	req, err := http.NewRequest("POST", "", body)
	if err != nil {
		t.Error(err)
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	// and test it!
	tmpFile, err := ReceiveFileToStruct(req)
	if err != nil {
		t.Error(err)
	}

	expected := 512
	fileSize := len(tmpFile.Content)
	if fileSize != expected {
		t.Errorf("File size is %v and should be %v", fileSize, expected)
	}
}
