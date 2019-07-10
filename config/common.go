package config

import "os"

var filePath = os.Getenv("CONFIG")

type Entry struct {
	src  string
	dest string
}

type Config = map[Entry]bool

func NewEntry(src, dest string) Entry {
	return Entry{src, dest}
}
