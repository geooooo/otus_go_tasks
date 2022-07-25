package main

import (
	"fmt"
	"strings"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
)

func LogLevelFromString(value string) (LogLevel, error) {
	switch strings.ToLower(value) {
	case "debug":
		return Debug, nil
	case "info":
		return Info, nil
	case "warn":
		return Warn, nil
	case "error":
		return Error, nil
	default:
		return Debug, fmt.Errorf("undefined LogLevel: '%s'", value)
	}
}

func (logLevel LogLevel) String() string {
	switch logLevel {
	case Debug:
		return "Debug"
	case Info:
		return "Info"
	case Warn:
		return "Warn"
	case Error:
		return "Error"
	default:
		panic(fmt.Sprintf("undefined LogLevel: %d", logLevel))
	}
}
