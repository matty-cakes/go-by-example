package main

import (
	"fmt"
)

func sumAndDiff(a int, b int) (int, int) {
	return a + b, a - b
}

// https://gobyexample.com/multiple-return-values
func main() {
	fmt.Println("Functions w/ Multiple Return Values:")
	fmt.Println("-- This is how functions with multiple return values work. See the source to understand.")
	fmt.Println("-- Here is an example of a function declaration with two return values:")
	fmt.Println(`func sumAndDiff(a int, b int) (int, int) {
  return a + b, a - b
}`)
	s, d := sumAndDiff(1, 2)
	fmt.Println("-- Calling 's, d := sumAndDiff(1, 2)' results in:", s, d)
	fmt.Println("-- We can exclude a return value from creating a variable like so '_, d := sumAndDiff(1, 2)':", d)
	fmt.Println("-- EFFECTIVE GO NOTE: The most common iditomatic use of mutliple return values is Go's idiom of returning an 'err' value as the last return of any function that might error out in its operation.")
}
