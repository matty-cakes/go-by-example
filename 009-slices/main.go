package main

import (
	"fmt"
)

// https://gobyexample.com/slices
func main() {
	fmt.Println("Slices:")
	var someSlice []int
	fmt.Println("-- Check out this empty slice why not...", someSlice)

	tinkerSlice := make([]string, 4)
	fmt.Println("-- Check out this slice that was intiialized with some capacity added to it...", tinkerSlice, "\n---- I know that doesn't look like much but check out it's length and capacity", len(tinkerSlice), cap(tinkerSlice))

	tinkerSlice[0] = "t"
	tinkerSlice[1] = "n"
	tinkerSlice[2] = "k"
	tinkerSlice[3] = "r"

	fmt.Println("-- Look I filled in the 'tinkerSlice'...", tinkerSlice)

	tinkerSlice = append(tinkerSlice, "!", "!")
	fmt.Println("-- Also we can use 'append' to increase the size of slice! Check out our new slice: ", tinkerSlice, "and it's new len and cap:", len(tinkerSlice), cap(tinkerSlice))

	tinkerCopy := make([]string, len(tinkerSlice))
	copy(tinkerCopy, tinkerSlice)
	fmt.Println("-- We can copy slices like this: 'copy(tinkerCopy, tinkerSlice)':", tinkerCopy)

	exclamations := tinkerCopy[4:6]
	fmt.Println("-- We can creates slices of slices like this 'exclamations := tinkerCopy[4:6]' :", exclamations)

	exclamation := tinkerCopy[5:]
	fmt.Println("-- We can slice to the end from a known index like this 'exclamation := tinkerCopy[5:]':", exclamation)

	tnkr := tinkerCopy[:4]
	fmt.Println("-- We can slice to a known index from the start of a slice like this 'tnkr := tinkerCopy[:5]':", tnkr)

}
