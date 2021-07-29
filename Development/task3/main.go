package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"

	"github.com/x4aekx/go-hw2/Development/task3/mySort"
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
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	defer file.Close()

	str, _ := fileScanner(file)
	result := mySort.SortStr(str)
	fmt.Println(result)
	fmt.Println("Sorted: ", sort.StringsAreSorted(result))
}
