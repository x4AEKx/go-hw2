package main

import (
	"fmt"
	"sort"

	"github.com/x4aekx/go-hw2/Development/task3/mySort"
	"github.com/x4aekx/go-hw2/Development/task3/scan"
)

func main() {

	str, err := scan.ReadFile("file.txt")
	if err != nil {
		fmt.Println(err)
	}

	result := mySort.SortStr(str)
	fmt.Println(result)
	fmt.Println("Sorted: ", sort.StringsAreSorted(result))
}
