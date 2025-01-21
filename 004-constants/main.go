package main

import (
	"fmt"
	"math"
)

const first string = "CONSTANTS can be strings"

// https://gobyexample.com/constants
func main() {
	fmt.Println("Constants:")
	fmt.Println(first)

	const second, third int = 78, 99
	fmt.Println("CONSTANTS may alsobe ints like:", second, "and", third)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println("CONSTANT expressions perform arithmetic with arbitrary precision. Ex: 'const d = 3e20 / n' ==", d)

	fmt.Println("A numeric CONSTANT has no type until it’s given one, such as by an explicit conversion. Ex: 'int64(d)'", int64(d))

	fmt.Println("CONSTANT types MAY also be infered through their context like how math.Sin requires a float: 'math.Sin(n)':", math.Sin(n))
}
