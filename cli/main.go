package main

import (
	"fmt"
	"os"
	"strings"
)

func badUsage() {
	fmt.Println("Usage:", os.Args[0], "--<ACTION> <ENTRY> <DATA>")
	fmt.Println("\tWhere <ENTRY> is the incoming host to proxy from")
	fmt.Println("\tWhere <ACTION> is one of the following:")
	fmt.Println("\t - add: <DATA> should be a single string for the outgoing target")
	fmt.Println("\t - rm:  <DATA> should be a single string for the outgoing target")
	fmt.Println("\t - run: <DATA> is the port followed by the command. See --help")
	fmt.Println("\t - help: <ENTRY> and <DATA> are ignored, and the help is shown")
}

func help() {
	fmt.Println("Example Usage:")
	fmt.Println()
	fmt.Println("Add a new rule:", os.Args[0], "--add test.example.com localhost:8080")
	fmt.Println("\tThis will forward all requests trying to reach test.example.com")
	fmt.Println("\tthat hit this server to localhost on port 8080")
	fmt.Println()
	fmt.Println("Remove a rule:", os.Args[0], "--rm test.example.com localhost:8080")
	fmt.Println("\tThis will remove the rule, such that requests to test.example.com")
	fmt.Println("\twill no longer go to localhost:8080")
	fmt.Println()
	fmt.Println("Run a program:", os.Args[0], "--run test.example.com 8080 ./my-binary my-binary-arg1 my-arg2")
	fmt.Println("\tThis will run the program my-binary, passing through the specifed")
	fmt.Println("\targs and for the duriation of the running program will proxy")
	fmt.Println("\tall traffic for test.example.com to localhost:8080")
	fmt.Println()
	fmt.Println("Basic load balancing:")
	fmt.Println("\t", os.Args[0], "--add test.example.com localhost:8000")
	fmt.Println("\t", os.Args[0], "--add test.example.com localhost:8001")
	fmt.Println("\tNow traffic will be split between these 2 servers")
}

func main() {
	if len(os.Args) < 2 {
		badUsage()
		return
	}
	var argFuncMap = map[string]func(){
		"add":  add,
		"rm":   remove,
		"run":  run,
		"help": help,
		"h":    help,
	}
	action := strings.Trim(os.Args[1], "-")
	if fn, ok := argFuncMap[action]; ok {
		fn()
	} else {
		badUsage()
	}
}
