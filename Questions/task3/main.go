package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil // fs.PathError - интерфейс, присваиваем значение nil
	return err
}

func main() {
	err := Foo()

	fmt.Println(err)

	fmt.Println(err == nil) // пустой интерфейс не равен nil, т.к. содержит внутри себя...

	fmt.Printf("%T\n", err)
	fmt.Printf("%T\n", nil)
}
