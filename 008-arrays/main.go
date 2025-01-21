package main

import (
	"fmt"
)

// https://gobyexample.com/arrays
func main() {
	fmt.Println("Arrays:")
	var someArr [5]int
	fmt.Println("-- A length of 5 empty array looks like this 'var someArr [5]int' ...", someArr)

	someArr[0] = 100
	fmt.Println("-- You can set values in an array like this 'someArr[0] = 100'...", someArr)

	fmt.Println("-- You can get the length of an array using the 'len' function 'len(someArr)'...", len(someArr))

	anotherArr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("-- Arrays dont have to be created empty 'anotherArr := [5]int{1, 2, 3, 4, 5}'...", anotherArr)

	yetAnotherArr := [...]int{9, 8, 7}
	fmt.Println("-- The compiler can also determine the length using the '...' syntax 'yetAnotherArr := [...]int{9, 8, 7}'...", yetAnotherArr)

	weirdArr := [...]int{25, 4: 35}
	fmt.Println("-- Using the ':' operator you can inject multiple zero valued items into an array at compile time 'weirdArr := [...]int{25, 4: 35}'...", weirdArr)

	var twoDimA [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoDimA[i][j] = j * 3
		}
	}
	fmt.Println("-- Hey check out this 2D array why not...", twoDimA)

	twoDimB := [3][4]int{
		{3, 4, 5, 6},
		{11, 22, 33, 44},
		{9, 7, 6, 4},
	}
	fmt.Println("-- Hey check out this TOTALLY different 2D array why not...", twoDimB)

}
