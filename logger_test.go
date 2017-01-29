package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNewLogger(t *testing.T) {
	stdout := new(bytes.Buffer)
	logger := NewLogger(stdout)
	logger.Println("test")

	fmt.Println(stdout.Bytes())

	//expected := []byte("test")
	//if bytes.Compare(expected, stdout.Bytes()) != 0 {
	//	t.Fatal("not matched", expected, stdout.Bytes())
	//}
}
