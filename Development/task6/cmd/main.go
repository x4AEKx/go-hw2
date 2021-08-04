package main

import (
	"flag"
	"fmt"
	"go-hw2/Development/task6/pkg"
	"log"
	"os"
	"strings"
)

func usage() {
	log.Printf("Usage: ./cut [OPTION]... STDIN... \n")
	flag.PrintDefaults()
}

func showUsageAndExit(exitcode int) {
	usage()
	os.Exit(exitcode)
}

func main() {
	var fields = flag.String("f", "", "выбрать поля (колонки)")             // +
	var delimiter = flag.String("d", "", "использовать другой разделитель") // +
	var separated = flag.Bool("s", false, "только строки с разделителем")   // ?не понял назначение флага

	var showHelp = flag.Bool("h", false, "Show help message") // +

	if *delimiter == "" {
		*delimiter = "\n"
	}

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if *showHelp {
		showUsageAndExit(0)
	}

	if len(*delimiter) > 1 {
		log.Println("cut: the delimiter must be a single character")
		showUsageAndExit(1)
	}

	if len(*fields) < 1 {
		log.Println("cut: you must specify a list of fields")
		showUsageAndExit(1)
	}

	fmt.Println([]byte(*delimiter))

	opts := pkg.NewOpts(*fields, *delimiter, *separated)

	result, err := pkg.Cut(*opts)
	if err != nil {
		log.Fatalf("cut: %s", err)
	}

	fmt.Println(strings.Join(result, " "))
}
