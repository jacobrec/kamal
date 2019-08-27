package main

import (
	"net/http"
	"os"

	"github.com/jacobrec/kamal/server/logger"
)

func main() {
	logger.Info("Starting Proxy Server on Port", port())
	http.HandleFunc("/", listen)
	if err := http.ListenAndServe(port(), nil); err != nil {
		panic(err)
	}
}

func port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	return ":" + port
}
