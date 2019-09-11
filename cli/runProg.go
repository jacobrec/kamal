package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
)

func runProg(entry, port string, cmd []string) {
	addEntry(entry, "localhost:"+port)

	// Catch interrupts and remove rule when they occur
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			removeEntry(entry, "localhost:"+port)
			os.Exit(0)
		}
	}()

	prgm := exec.Command(cmd[0], cmd[1:]...)
	prgm.Env = append(os.Environ(), "PORT="+port)
	fmt.Println("Run:", entry, ":", port, " ==> ", cmd)
	prgm.Run()

	removeEntry(entry, "localhost:"+port)
}
