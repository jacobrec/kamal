package main

import (
	"net/url"
	"os"
	"strconv"
	"strings"
)

// The help command was left in main.go

func addOrRemove(fn func(string, string)) bool {
	if len(os.Args) != 4 {
		return false
	}
	entry, ok := getEntrypoint()
	if !ok {
		return false
	}

	target, ok := getTarget()
	if !ok {
		badUsage()
		return false
	}

	fn(entry, target)
	return true
}

func add() {
	if !addOrRemove(addEntry) {
		badUsage()
	}
}

func remove() {
	if !addOrRemove(removeEntry) {
		badUsage()
	}
}

func run() {
	var parseArgs = func() (bool, string, string, []string) {
		if len(os.Args) <= 4 {
			return false, "", "", nil
		}
		entry, ok := getEntrypoint()
		if !ok {
			return false, "", "", nil
		}
		port, ok := getRunPort()
		if !ok {
			return false, "", "", nil
		}
		run_cmd := os.Args[4:]
		return true, entry, port, run_cmd
	}
	ok, entry, port, run_cmd := parseArgs()

	if ok {
		runProg(entry, port, run_cmd)
	} else {
		badUsage()
	}
}

func getEntrypoint() (string, bool) {
	path, err := url.Parse("http://" + os.Args[2])
	return path.Hostname(), err == nil && path.Hostname() != ""
}

func getRunPort() (string, bool) {
	port := os.Args[3]
	port_num, err := strconv.Atoi(port)
	return port, err == nil && port_num <= 65535
}

func getTarget() (string, bool) {
	raw := os.Args[3]
	if !strings.Contains(raw, "://") {
		raw = "http://" + raw
	}
	path, err := url.Parse(raw)
	return path.Hostname() + ":" + path.Port(), err == nil && path.Hostname() != ""
}
