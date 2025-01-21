package main

import (
	"fmt"
)

// https://gobyexample.com/for
func main() {
	fmt.Println("Loops:")

	i := 1
	for i <= 3 {
		fmt.Println("Basic conditional loop iteration:", i)
		i = i + 1
	}

	for j := 0; j < 4; j++ {
		fmt.Println("Classic three part loop iteration:", j)
	}

	for k := range 5 {
		fmt.Println("Range loop iteration:", k)
	}

	for {
		fmt.Println("You can just use 'for {code...}' to produce an infinite loop")
		break
	}

	fmt.Println("The keyword 'continue' can be used to move on to the next iteration of a loop")
	for l := range 6 {
		if l%2 == 0 {
			continue
		}
		fmt.Println("This iteration of a range loop will only print and odd number like:", l)
	}

}
