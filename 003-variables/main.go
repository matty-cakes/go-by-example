package main

import "fmt"

// https://gobyexample.com/variables
func main() {
	fmt.Println("Variables:")
	var first = "varibles can be strings"
	fmt.Println(first)

	var second, third int = 2, 3
	fmt.Println("or maybe ints like:", second, "and", third)

	var fourth = true
	fmt.Println("it is", fourth, "that variables dont have to be typed upon creation")

	var fifth int
	fmt.Println("all types have a zero value in go for instance the int zero value is:", fifth)

	sixth := "sixth"
	fmt.Println("the syntax 'sixth := \"sixth\"' is how variables can be declared and typed in shorthand... it could then be referenced in a Println and would produce the string:", sixth)
}
