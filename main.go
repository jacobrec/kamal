package main

import (
	"net/http"

	"github.com/jacobrec/kamal/logger"
)

func main() {

	redirectMap = make(map[string]string)
	redirectMap["test.rac.reckhard.ca"] = "reckhard.ca"
	redirectMap["loop.rac.reckhard.ca"] = "reckhard.ca:8888"

	logger.Info("Starting Proxy Server")
	http.HandleFunc("/", listen)
	if err := http.ListenAndServe(":8888", nil); err != nil {
		panic(err)
	}
}
