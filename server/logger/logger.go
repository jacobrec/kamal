package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	ocd = iota
	debug
	info
	warn
	err
)

var levelMap = map[string]int{
	"OCD":   ocd,
	"DEBUG": debug,
	"INFO":  info,
	"WARN":  warn,
	"ERROR": err,
	"":      debug, // default if env isn't set
}

var level = levelMap[os.Getenv("LOG_LEVEL")]

func jLog(level string, v ...interface{}) {
	log.Print("[", level, "] ", fmt.Sprint(v...))
}

// OCD is the most verbose logging level, use this if your printing information
// in like a tight loop, or something else that would similaraly clog logs
func OCD(v ...interface{}) {
	if level <= ocd {
		jLog("OCD", v...)
	}
}

// Debug can be used everywhere, this should be your goto log level, its for
// things you want to see while developing, but you don't really care about
// once things start working
func Debug(v ...interface{}) {
	if level <= debug {
		jLog("DEBUG", v...)
	}
}

// Info is for important information, it should be used sparingly, like for
// saying server shutdown, or sql database reset
func Info(v ...interface{}) {
	if level <= info {
		jLog("INFO", v...)
	}
}

// Warn is for when things go wrong but it's recoverable
func Warn(v ...interface{}) {
	if level <= warn {
		jLog("WARN", v...)
	}
}

// Error prints things out at the highest priority, only things going horribly
// wrong that will lead to crashes should use this
func Error(v ...interface{}) {
	if level <= err {
		jLog("ERROR", v...)
	}
}
