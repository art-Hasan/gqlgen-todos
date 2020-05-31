package main

import (
	"context"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

const (
	defaultHost = "127.0.0.1"
	defaultPort = 8080
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(defaultPort)
	}

	service, err := Service()
	if err != nil {
		panic(err)
	}

	if err := service.Migrate(context.Background()); err != nil {
		panic(err)
	}

	if err := service.ListenAndServe(); err != nil {
		panic(err)
	}
}
