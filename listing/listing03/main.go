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

	fmt.Println(err == nil) // пустой интерфейс не равен nil, т.к. содержит внутри себя тип(*os.PathError) и значение(равное nil)

	fmt.Printf("%T\n", err) // тип интерфейса
	fmt.Printf("%v\n", err) // значение интерфейса

	fmt.Printf("%T\n", nil)
}
