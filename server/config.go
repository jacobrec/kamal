package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/jacobrec/kamal/config"
	"github.com/jacobrec/kamal/server/logger"
)

var lastReadFile time.Time
var filePath = os.Getenv("CONFIG")
var redirectMap map[string][]string = make(map[string][]string)

func getDestination(scheme, from string) string {
	to := ""
	checkConfig()
	if val, ok := redirectMap[from]; ok {
		to = scheme + "://" + val[rand.Intn(len(val))]

		logger.Debug("Redirect: ", from, " => ", to)
	} else {
		logger.Debug("Unknown origin: ", from)
	}
	return to
}

func loadConfigFile() {
	redirects, err := config.ParseConfigAsMap()
	if err == nil {
		logger.Info("Successfully loaded config file.")
		lastReadFile = time.Now()
		redirectMap = redirects
	} else if err.Error() == "Read" {
		logger.Warn("An error occured while reading the config file.")
	} else if err.Error() == "Open" {
		logger.Warn("Could not open the config file. \nThe server will be useless.")
	} else {
		logger.Warn("An unexpected error occured with the config file.")
	}
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
