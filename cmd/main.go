package main

import (
	_ "github.com/miguoliang/arch-go/configs"
	"github.com/miguoliang/arch-go/internal/resource"
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
	"io"
	"log"
	"os"
)

const (
	graylogAddr = "localhost:12201"
)

// @title Arch-Go API
// @description This is the API for Arch-Go
// @version 1.0
// @host localhost:8080
// @BasePath /api
// @schemes http
// @schemes https
// @contact.name Guoliang Mi
// @contact.email boymgl@qq.com
// @contact.url https://miguoliang.com
func main() {

	setupLog()

	r := resource.SetupRoutes()

	err := r.Run("0.0.0.0:8081")
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	log.Println("Started!")
}

func setupLog() {

	gelfWriter, err := gelf.NewUDPWriter(graylogAddr)
	if err != nil {
		log.Fatalf("gelf.NewWriter: %s", err)
	}
	// log to both stderr and graylog2
	log.SetOutput(io.MultiWriter(os.Stderr, gelfWriter))
}
