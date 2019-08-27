package config

import (
	"os"
	"time"
)

var filePath = getEnv("KAMAL_CONFIG", "/etc/kamal/config")

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

type Entry struct {
	src  string
	dest string
}

type Config = map[Entry]bool

func NewEntry(src, dest string) Entry {
	return Entry{src, dest}
}

func (e Entry) String() string {
	return e.src + " => " + e.dest
}

func LastModTime() (time.Time, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return time.Now(), err
	}
	return info.ModTime(), nil
}
