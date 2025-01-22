package main

import (
	"fmt"
)

func intSeq() func() int {
	num := 0
	return func() int {
		num = num + 1
		return num
	}
}

// https://gobyexample.com/closures
func main() {
	fmt.Println("Function Closures:")
	fmt.Println("-- See the source to understand...")

	incr := intSeq()
	one := incr()
	two := incr()
	three := incr()

	fmt.Println(one, two, three)

	fmt.Println("-- Here is an example of a function that retains data within a closure:")
	fmt.Println(`func intSeq() func() int {
  num := 0
  return func() int {
	num = num + 1
	return num
  }
}`)
	fmt.Printf("-- In this example the intSeq CREATES and RETURNS a function whose data is enclosed within the body of 'intSeq' so that 'closure' remains AFTER 'intSeq' has returned so call intSeq will give a function WHICH is a value: %T\n", intSeq())
	f1 := intSeq()
	fmt.Println("-- So if we ASSIGN intSeq's return value to a variable like so 'f1 := intSeq()' and then call it like so 'f1()' what do we get...", f1())
	fmt.Println("-- What if we call it again 'f1()' ???", f1())
	fmt.Println("-- And again! 'f1()'", f1())
	fmt.Println("-- You get the idea...")
	f2 := intSeq()
	fmt.Println("-- As a final note, the closures are one per call meaning that you can create as many independent counters using multiple calls to 'intSeq' as you want! 'f1()' and 'f2()'", f1(), f2())
}
