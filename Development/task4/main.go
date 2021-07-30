package main

import (
	"fmt"

	"github.com/x4aekx/go-hw2/Development/task4/find"
)

func main() {
	arr := []string{"пятак", "пятак", "пятка", "Яптка", "апятк", "тяпка", "листок", "ислток", "слиток", "столик"}

	result := find.SetsOfAnagramInDict(arr)

	fmt.Println(result)
}
