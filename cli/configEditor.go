package main

import "fmt"

type Entry struct {
	src  string
	dest string
}

type Config = map[Entry]bool

func editEntry(entry, target string, val bool) {
	e := Entry{src: entry, dest: target}
	config := ParseConfig()
	if config != nil {
		config[e] = val
		SaveConfig(config)
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

func ParseConfig() Config {
	config := make(map[Entry]bool)

	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), " ")
		config[Entry{data[0], data[1]}] = true
	}

	if err := scanner.Err(); err != nil {
		return nil
	}

	return config
}

func SaveConfig(config Config) {
}
