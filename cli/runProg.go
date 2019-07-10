package main

import "fmt"

func runProg(entry, port string, cmd []string) {
	fmt.Println("Run:", entry, ":", port, " ==> ", cmd)
}
