package main

import (
	"fmt"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

// https://gobyexample.com/recursion
func main() {
	fmt.Println("Function Recursion:")
	fmt.Println("-- See the source to understand...")
	fmt.Println("-- Below is an example of a recursive function:")
	fmt.Println(`func fact(n int) int {
  if n == 0 {
    return 1
  }
  return n * fact(n-1)
}`)

	fmt.Println("-- As you can see this function calls ITSELF until the value for n is 1. Then each successive (or recuirsive) return allows the function instance above it to use the result of more deeply nested results...")
	fmt.Println("-- Recursion isn't easy for anyone to just get, but with patience it will click.")
	fmt.Println("-- Let's try our recursive function out 'fact(4)':", fact(4))
	// Note this actually looks like 1 * 2 ==> 2 * 3 ==> 6 * 4 === 24

	fmt.Println("-- Not sure why this is surprising but recursive function can be assigned to var to be used anonymously. Like so...")
	fmt.Println("'var diffFact func(n int) int'")
	fmt.Println("'diffFact = func(n) int {...}'... you get the idea...")
}
