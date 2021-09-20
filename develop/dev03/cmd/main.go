package main

import (
	"flag"
	"go-hw2/develop/dev03/pkg/sort"
	"log"
	"os"
)

func usage() {
	log.Printf("Usage: ./sort [OPTION]... [FILE]... \n")
	flag.PrintDefaults()
}

func showUsageAndExit(exitcode int) {
	usage()
	os.Exit(exitcode)
}

func main() {
	var fields = flag.String("k", "", "указание колонки для сортировки") // +
	var number = flag.Bool("n", false, "сортировать по числовому значению")
	var reverse = flag.Bool("r", false, "сортировать в обратном порядке")     // +
	var duplicate = flag.Bool("u", false, "не выводить повторяющиеся строки") // +

	var showHelp = flag.Bool("h", false, "Show help message")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if *showHelp {
		showUsageAndExit(0)
	}

	args := flag.Args()

	if len(args) < 1 {
		showUsageAndExit(1)
	}

	var files = args[0:]

	sort := sort.New(*fields, *number, *reverse, *duplicate, files)

	err := sort.Execute()
	if err != nil {
		log.Fatalf("sort: %s", err)
	}

	err = sort.Output()
	if err != nil {
		log.Fatalf("sort: %s", err)
	}
}
