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
	var fields = flag.String("f", "", "выбрать поля (колонки)")               // +
	var delimiter = flag.String("d", "\t", "использовать другой разделитель") // +
	var separated = flag.Bool("s", false, "только строки с разделителем")     // ?не понял назначение флага

	var showHelp = flag.Bool("h", false, "Show help message") // +

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if *showHelp {
		showUsageAndExit(0)
	}

	if len(*fields) < 1 {
		log.Fatalf("cut: you must specify a list of fields")
	}

	opts := pkg.NewOpts(*fields, *delimiter, *separated)

	result, err := pkg.Cut(*opts)
	if err != nil {
		log.Fatalf("cut: %s", err)
	}

	fmt.Println(strings.Join(result, " "))
}
