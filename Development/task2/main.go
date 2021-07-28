package main

import (
	"fmt"

	"github.com/x4aekx/go-hw2/Development/task2/unpack"
)

func main() {
	str1, _ := unpack.UnpStr("a4bc2d5e")
	fmt.Printf("a4bc2d5e = %s\n\n", str1)

	str2, _ := unpack.UnpStr("abcd")
	fmt.Printf("abcd = %s\n\n", str2)

	str3, err := unpack.UnpStr("45")
	fmt.Printf("45 = %s\n", str3)
	fmt.Printf("%s\n\n", err)

	str4, _ := unpack.UnpStr("")
	fmt.Printf("''= %s\n\n", str4)

	str5, _ := unpack.UnpStr("da4bc2d5e")
	fmt.Printf("da4bc2d5e = %s\n\n", str5)

	str6, _ := unpack.UnpStr(" ")
	fmt.Printf("' ' = %s\n", str6)

}
