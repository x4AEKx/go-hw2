package main

import (
	"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}

	var b []int = a[1:4] // слайс, который ссылается на массив

	b[0] = 1

	b = append(b, 3)
	b = append(b, 4)

	a[2] = 2

	fmt.Println(b, a)
}
