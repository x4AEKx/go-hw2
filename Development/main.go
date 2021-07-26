package main

import (
	"fmt"

	"github.com/x4AEKx/go-hw2/Development/ext"
)

func main() {
	time, err := ext.GetTime("")

	fmt.Println(time)
	fmt.Println(err)
}
