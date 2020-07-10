package main

import (
	"testing"

	"github.com/ssuareza/filesplit"
)

func TestMD5(t *testing.T) {
	testString := []byte("random text")
	md5 := MD5(testString)

	expected := "d9b9bec3f4cc5482e7c5ef43143e563a"
	if md5 != expected {
		t.Errorf("MD5 has is %s and should be %v", md5, expected)
	}
}

func TestIndex(t *testing.T) {
	key := "mykey"
	md5 := "d9b9bec3f4cc5482e7c5ef43143e563a"
	chunks := []filesplit.Chunk{
		{Name: "chunk0", Content: []byte("random")},
		{Name: "chunk1", Content: []byte("random")},
	}

	index := Index(key, md5, chunks)

	expected := 2
	if len(index) != expected {
		t.Errorf("Index number is %v and should be %v", len(index), expected)
	}
}
