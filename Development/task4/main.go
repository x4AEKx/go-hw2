package main

import (
	"fmt"

	"go-hw2/Development/task4/pkg/find"
)

func main() {
	arr := []string{"пятак", "пятак", "пятка", "Яптка", "апятк", "тяпка", "листок", "ислток", "слиток", "столик"}

	result := find.SetsOfAnagram(arr)

	fmt.Println(result)
}
