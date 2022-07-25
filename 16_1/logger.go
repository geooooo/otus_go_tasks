package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var logFile *os.File

func InitLogger(pathToLogFile string, logLevel LogLevel) error {
	if error := openLogFile(pathToLogFile); error != nil {
		return error
	}

	log.SetPrefix(fmt.Sprintf("%s: ", logLevel))
	log.SetFlags(log.Ldate | log.Lmicroseconds)

	return nil
}

func openLogFile(path string) error {
	logFile, error := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if error != nil {
		return fmt.Errorf("cannot open log file: %s", path)
	}

	log.SetOutput(logFile)

	return nil
}

func CloseLogger() error {
	if error := logFile.Close(); error != nil {
		return errors.New("logger closing error")
	}

	return nil
}

func Log(text string) {
	log.Println(text)
}
