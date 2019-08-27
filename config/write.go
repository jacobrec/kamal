package config

import (
	"os"
	"path/filepath"
	"strings"
)

func SaveConfig(config Config) bool {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0755)
	if err == nil {
		f.Truncate(0)
		f.WriteString(stringConfig(config))
		return true
	}
	return false
}

func stringConfig(config Config) string {
	var sb strings.Builder
	for k, v := range config {
		if v {
			sb.WriteString(k.src)
			sb.WriteString(" ")
			sb.WriteString(k.dest)
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

func CreateIfNotExist() {
	os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0755)
}
