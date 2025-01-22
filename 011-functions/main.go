package main

import (
	"fmt"
)

func addTwoInts(a int, b int) int {
	return a + b
}

func addThreeInts(a, b, c int) int {
	return a + b + c
}

// https://gobyexample.com/functions
func main() {
	fmt.Println("Functions:")
	fmt.Println("-- This is how functions work. See the source to understand.")
	fmt.Println("-- Here is an example of a function declaration:")
	fmt.Println(`func addTwoInts(a int, b int) int {
  return a + b
}`)
	fmt.Println("-- Calling 'addTwoInts(1,2)' results in:", addTwoInts(1, 2))
	fmt.Println("-- Calling 'addThreeInts(1,2,3)' results in:", addThreeInts(1, 2, 3))
}
