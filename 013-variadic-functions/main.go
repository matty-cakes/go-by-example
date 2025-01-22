package main

import (
	"fmt"
)

func sumAll(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// https://gobyexample.com/variadic-functions
func main() {
	fmt.Println("Variadic Functions:")
	fmt.Println("-- This is how variadic functions work. See the source to understand.")
	fmt.Println("-- Here is an example of a variadic function declaration:")
	fmt.Println(`func sumAll(nums ...int) int {
	var total int
	for i := range nums {
	   total += i
	}
	return total
}`)
	fmt.Println("-- Calling 's := sumAll(1, 3, 5, 7)' results in:", sumAll(1, 3, 5, 7))
	fmt.Println("-- EFFECTIVE GO SIDE NOTE: The '...' syntax works both ways. For variadic function it combines all args INTO a slice (Ex: 'nums ...int', nums == [1,3,5,7]), but when placed on the otherside the elements in a slice are spread out INTO individual items.")
	nums := []int{1, 2, 3, 4}
	fmt.Println("---- Example: 'nums := []int{1, 2, 3, 4}' & 'sumAll(nums...)'", sumAll(nums...))
	fmt.Println("-- FURTHER NOTE: I realize this way of demonstrating is a little wasteful and mind-bendy but hopefuly you get it.")

}
