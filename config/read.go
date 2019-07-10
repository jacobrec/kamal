package config

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func ParseConfig(onLine func(string)) error {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return errors.New("Open")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		onLine(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return errors.New("Read")
	}
	return nil
}

func ParseConfigAsMap() (map[string][]string, error) {
	var redirectMap = make(map[string][]string)
	var onLine = func(line string) {
		data := strings.Split(line, " ")
		if arr, ok := redirectMap[data[0]]; ok {
			redirectMap[data[0]] = append(arr, data[1])
		} else {
			redirectMap[data[0]] = []string{data[1]}
		}
	}
	return redirectMap, ParseConfig(onLine)
}

func ParseConfigAsSet() (Config, error) {
	config := make(map[Entry]bool)

	var onLine = func(line string) {
		data := strings.Split(line, " ")
		config[Entry{data[0], data[1]}] = true
	}
	return config, ParseConfig(onLine)
}
