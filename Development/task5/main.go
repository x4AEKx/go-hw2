package main

import (
	"flag"
	"fmt"
	"go-hw2/Development/task5/mygrep"
	"log"
	"os"
)

func usage() {
	log.Printf("Usage: [OPTION]... PATTERNS [FILE]... \n")
	flag.PrintDefaults()
}

func showUsageAndExit(exitcode int) {
	usage()
	os.Exit(exitcode)
}

func main() {
	var after = flag.Int("A", 0, "печатать +N строк после совпадения")
	var before = flag.Int("B", 0, "печатать +N строк до совпадения")
	var context = flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	var count = flag.Bool("c", false, "количество строк")
	var ignoreCase = flag.Bool("i", false, "игнорировать регистр")     // +
	var invert = flag.Bool("v", false, "вместо совпадения, исключать") // +
	var fixed = flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	var lineNum = flag.Bool("n", false, "печатать номер строки") // +

	var showHelp = flag.Bool("h", false, "Show help message") // +

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if *showHelp {
		showUsageAndExit(0)
	}

	var pattern = flag.Args()[0]
	var files = flag.Args()[1:]

	opts := mygrep.NewOpts(*after, *before, *context, *count, *ignoreCase, *invert, *fixed, *lineNum)

	result := mygrep.Search(pattern, *opts, files)

	for _, v := range result {
		fmt.Println(v)
	}
}

// grep -A 3 -i "example" demo_text
