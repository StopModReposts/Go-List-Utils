package main

import (
	"log"
	"os"
)

func GetCommand() string {
	return os.Args[1]
}

func EnsureArgs() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: " + os.Args[0] + " <command> <list-directory>")
		os.Exit(1)
	}
}

func GetListFile() string {
	return os.Args[2]
}
