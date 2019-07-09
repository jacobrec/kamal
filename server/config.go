package main

import (
	"bufio"
	"os"
	"strings"
	"time"

	"github.com/jacobrec/kamal/server/logger"
)

var lastReadFile time.Time
var filePath = os.Getenv("CONFIG")
var redirectMap map[string]string = make(map[string]string)

func getDestination(scheme, from string) string {
	to := ""
	checkConfig()
	if val, ok := redirectMap[from]; ok {
		to = scheme + "://" + val

		logger.Debug("Redirect: ", from, " => ", to)
	} else {
		logger.Debug("Unknown origin: ", from)
	}
	return to
}

func loadConfigFile() {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		logger.Warn("Could not open the config file. ", err, "\nThe server will be useless")
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		loadConfigLine(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		logger.Warn("An error occured while reading the config file. ", err)
		return
	}

	logger.Info("Successfully loaded config file")
	lastReadFile = time.Now()
}

func loadConfigLine(line string) {
	data := strings.Split(line, " ")
	redirectMap[data[0]] = data[1]
}

func shouldReloadConfigFile() bool {
	info, err := os.Stat(filePath)
	if err != nil {
		logger.Warn("Error reading file", err)
		return false
	}
	return info.ModTime().After(lastReadFile)
}

func checkConfig() {
	if shouldReloadConfigFile() {
		loadConfigFile()
	}
}
