package main

import (
	"fmt"
)

// https://gobyexample.com/if-else
func main() {
	fmt.Println("Conditional Branching:")

	fmt.Print("did you you know that... ")
	if 7%2 == 0 {
		fmt.Println("7 is even?")
	} else {
		fmt.Println("7 is odd?")
	}

	eightIsDivisivbleByFour := 8%4 == 0
	fmt.Print("did you you know that... ")
	if eightIsDivisivbleByFour {
		fmt.Println("8 is divisble by 4?")
	}

	fmt.Print("did you you know that... ")
	if 11%2 == 0 || 10%2 == 0 {
		fmt.Println("either 11 or 10 is even?")
	}

	if mattCoolFactor := 99; mattCoolFactor < 20 {
		fmt.Println("Matt isn't very cool?")
	} else if mattCoolFactor >= 20 && mattCoolFactor < 90 {
		fmt.Println("Matt is cool enough")
	} else {
		fmt.Println("Matt is so so so cool")
	}
}
