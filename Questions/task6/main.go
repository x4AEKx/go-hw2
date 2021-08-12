// Что выведет программа? Объяснить вывод программы.
// Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	fmt.Printf("len i = %d cap i = %d\n", len(i), cap(i))

	i = append(i, "4") /*аллоцируется память,
	/ i больше не ссылается на слайс i (который в свою очередь ссылается на массив),
	/ т.к. capacity удвоилось
	/ и i ссылается уже на другой массив(скопировав значения из первого массива),
	/ отличный от от того на который ссылался первоначальный i
	*/
	fmt.Printf("len i = %d cap i = %d\n", len(i), cap(i))
	fmt.Println(i)

	i[1] = "5"
	fmt.Printf("len i = %d cap i = %d\n", len(i), cap(i))
	fmt.Println(i)

	i = append(i, "6")
	fmt.Printf("len i = %d cap i = %d\n", len(i), cap(i))
	fmt.Println(i)
}
