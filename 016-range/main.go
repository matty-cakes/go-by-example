package main

import (
	"fmt"
)

// https://gobyexample.com/range-over-built-in-types
func main() {
	fmt.Println("Ranging over built-In types:")
	fmt.Println("-- In Go you can use 'range' to iterate over slices, arrays, maps (keys and values), and strings...")
	s := []int{5, 7, 9}
	fmt.Println("-- Let's take a basic slice and demonstrate... 's := []int{5, 7, 9}'")
	fmt.Println(`for index, value := range s {
	fmt.Println("--- The value:", value, "is at index:", index)
}`)
	for index, value := range s {
		fmt.Println("--- The value:", value, "is at index:", index)
	}

	m := map[string]int{"a": 4, "b": 6, "c": 8}
	fmt.Println("-- We can also range over the keys and values (either keys or values OR both) 'm := map[string]int{\"a\":4, \"b\":6, \"c\":8}'")

	fmt.Println(`for k, v := range m {
	fmt.Println("--- The value:", v, "is at key:", k)
}`)
	for k, v := range m {
		fmt.Println("--- The value:", v, "is at key:", k)
	}

	str := "cool"
	fmt.Println("-- Lastly, we can range over BYTES in strings 'str := \"cool\"'")

	fmt.Println(`for _, l := range str {
	fmt.Printf("--- LETTER: %d (char repr: %c)\n", l, l)
}`)
	for _, l := range str {
		fmt.Printf("--- LETTER: %d (char repr: %c)\n", l, l)
	}

}
