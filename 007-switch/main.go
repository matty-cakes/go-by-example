package main

import (
	"fmt"
	"time"
)

// https://gobyexample.com/switch
func main() {
	fmt.Println("Switch:")

	i := 2
	fmt.Print("did you you know that the variable 'i' was set to... ")
	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	}

	fmt.Print("did you know that today was... ")
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("a wonderful weekend day!")
	default:
		fmt.Println("a lovely weekday!")
	}

	whatAmI := func(i interface{}) {
		fmt.Print("what is t? ")
		switch t := i.(type) {
		case bool:
			fmt.Println("t is a bool")
		case int:
			fmt.Println("t is an int")
		default:
			fmt.Println("not sure what t is: ", t)
		}
	}

	whatAmI(true)
	whatAmI("hello")
	whatAmI(14)
}
