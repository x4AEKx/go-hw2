package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func fileScanner(r io.Reader) ([]string, error) {
	var strs []string

	fileScanner := bufio.NewScanner(r)

	for fileScanner.Scan() {
		strs = append(strs, fileScanner.Text())
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	return strs, nil
}

func main() {
	file, err := os.Open("file.txt")
	defer file.Close()
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	str, _ := fileScanner(file)
	result := mySort.sortStr(str)
	fmt.Println(result)
	fmt.Println("Sorted: ", sort.StringsAreSorted(result))
}
