package main

import (
	"go-hw2/Development/task9/wget"
	"log"
	"os"
)

// wget test.txt http://example.org/

func usage() {
	log.Printf("Usage: wget [FILE] [URL] \n")
}

func showUsageAndExit(exitcode int) {
	usage()
	os.Exit(exitcode)
}

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		showUsageAndExit(1)
	}

	fileName := args[0]
	urlPath := args[1]

	if err := wget.Wget(fileName, urlPath); err != nil {
		log.Printf("Error: %s", err)
	}
}
