package main

import (
	"fmt"
)

func zeroVal(n int)     { n = 0 }
func zeroValPtr(n *int) { *n = 0 }

// https://gobyexample.com/pointers
func main() {
	fmt.Println("Pointers:")
	fmt.Println("-- In Go functions will either make copies of values provided to them or you can pass a reference to a value using a pointer")
	fmt.Println("-- Read over the following two functions:")
	fmt.Println("---- func zeroVal(n int)     { n = 0 }")
	fmt.Println("---- func zeroValPtr(n *int) { *n = 0 }")
	v := 1
	fmt.Println("-- initial value:", v)
	zeroVal(v)
	fmt.Println("-- value after call to 'zeroVal':", v)
	zeroValPtr(&v)
	fmt.Println("-- value after call to 'zeroValPtr':", v)
}
