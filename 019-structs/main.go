package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

// https://gobyexample.com/strings-and-runes
func main() {
	fmt.Println("Structs:")
	fmt.Println("In Go we can group logical entities together using structs. For example below is a basic struct for creating a person:")
	type person struct {
		name string
		age  int
	}
	fmt.Println(`type person struct {
  name string
  age int
}`)

	fmt.Println("-- We can create a new person like so:")
	walter := person{name: "walter", age: 22}
	fmt.Println("-- walter := person{name: \"walter\", age: 22} --", walter)
	fmt.Println("-- We can use the dot char '.' to access attributes of structs. Ex: 'walter.age':", walter.age)

	fmt.Println("-- You can create a struct, get it's pointer value AND use the same dot syntax on the pointer. The pointer is magically derefenced for you:")
	sam := &person{name: "sam", age: 44}
	fmt.Println("-- sam := &person{name: \"sam\", age: 44} --", sam)
	fmt.Println("-- What is sam's age: 'same.age'", sam.age)

	fmt.Println("-- Lastly, you can create inline/anonymous structs for one time uses:")
	dog := struct {
		name   string
		isGood bool
	}{
		"rufus",
		true,
	}
	fmt.Println(`dog := struct {
  name   string
  isGood bool
}{
  "rufus",
  true,
}`)
	boyStatus := "boy"
	if dog.isGood {
		boyStatus = "good boy"
	} else {
		boyStatus = "bad boy"
	}
	fmt.Printf("-- %s is a %s\n", dog.name, boyStatus)

}
