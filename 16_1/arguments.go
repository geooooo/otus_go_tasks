package main

import "flag"

func InitArgs() string {
	pathToConfig := flag.String("config", "", "path/to/config")

	flag.Parse()

	return *pathToConfig
}
