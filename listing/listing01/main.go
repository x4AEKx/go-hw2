package main

import (
	"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] // B слайс ссылается на массив A
	b[0] = 1
	b[1] = 2
	b[2] = 3
	fmt.Printf("len a = %d cap a = %d\n", len(a), cap(a))
	fmt.Printf("len b = %d cap b = %d\n", len(b), cap(b))
	fmt.Println("a = ", a)
	fmt.Printf("b = %v\n\n", b)

	b = append(b, 4) // B слайс все ещё ссылается на массив A, т.к. capacity позволяет
	fmt.Printf("len a = %d cap a = %d\n", len(a), cap(a))
	fmt.Printf("len b = %d cap b = %d\n", len(b), cap(b))
	fmt.Println("a = ", a)
	fmt.Printf("b = %v\n\n", b)

	b = append(b, 5) // B не ссылается на массив A, т.к. capacity удвоилось и B ссылается уже на другой массив, отличный от A
	b[0] = 42
	fmt.Printf("len a = %d cap a = %d\n", len(a), cap(a))
	fmt.Printf("len b = %d cap b = %d\n", len(b), cap(b))
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
}
