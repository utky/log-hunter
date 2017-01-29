package main

import (
	"io"
	"log"
)

func NewLogger(writer io.Writer) *log.Logger {
	logger := log.New(writer, "", log.LstdFlags)
	return logger
}
