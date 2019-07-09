package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/jacobrec/kamal/server/logger"
)

func listen(res http.ResponseWriter, req *http.Request) {
	scheme := "http"
	if req.TLS != nil {
		scheme = "https"
	}
	send(getDestination(scheme, strings.Split(req.Host, ":")[0]), res, req)
}

func send(destination string, res http.ResponseWriter, req *http.Request) {
	url, err := url.Parse(destination)
	if err != nil {
		logger.Warn("Could not parse url: ", destination)
		return
	}

	logger.OCD("destination:", destination, " scheme:", url.Scheme,
		" host:", url.Host)
	proxy := httputil.NewSingleHostReverseProxy(url)

	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forward-Host", req.Header.Get("Host"))
	req.Host = url.Host

	proxy.ServeHTTP(res, req)
}
