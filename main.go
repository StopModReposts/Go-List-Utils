package main

import (
	"log"
	"os"
)

type IllegalSite struct {
	Domain string
	Notes  string
	Path   string
	Reason string
}

func main() {
	EnsureArgs()
	f := GetListFile()

	log.Println("Using list: " + f)
	command := GetCommand()

	switch command {
	case "issue":
		IssueCommand()
	case "sort":
		SortCommand()
	}
}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
