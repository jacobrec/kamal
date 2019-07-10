package main

import (
	"fmt"

	"github.com/jacobrec/kamal/config"
)

func editEntry(entry, target string, val bool) {
	e := config.NewEntry(entry, target)
	config.CreateIfNotExist()
	configFile, err := config.ParseConfigAsSet()
	if err == nil {
		configFile[e] = val
		if !config.SaveConfig(configFile) {
			fmt.Println("Error saving config file")
		}
	} else {
		fmt.Println("Config file error:", err)
	}
}
func removeEntry(entry, target string) {
	fmt.Println("Removing:", entry, " => ", target)
	editEntry(entry, target, false)
}

func addEntry(entry, target string) {
	fmt.Println("Adding:", entry, " => ", target)
	editEntry(entry, target, true)
}
