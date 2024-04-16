package main

import (
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
	"io"
	"log"
	"os"
)

const (
	graylogAddr = "localhost:12201"
)

func SetupLog() {

	gelfWriter, err := gelf.NewUDPWriter(graylogAddr)
	if err != nil {
		log.Fatalf("gelf.NewWriter: %s", err)
	}
	// log to both stderr and graylog2
	log.SetOutput(io.MultiWriter(os.Stderr, gelfWriter))
}
