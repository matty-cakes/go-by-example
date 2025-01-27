package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func (p person) greet() { fmt.Printf("Hi I'm %s and I am %d years old...\n", p.name, p.age) }

// https://gobyexample.com/methods
func main() {
	fmt.Println("Methods:")
	fmt.Println("-- In Go methods can be attached to struct types")
	fmt.Println("-- Let's use our 'person' struct from the previous lesson:")
	fmt.Println(`type person struct {
  name string
  age int
}`)
	fmt.Println("-- Let's attach a method to this struct")
	fmt.Println("func (p person) greet() { fmt.Printf(\"Hi I'm %s and I am %d years old...\\n\", p.name, p.age) }")
	fmt.Println("-- Notice after the word 'func' we have the following '(p person)' this is known as a method reciever or the TYPE that is recieving the method. In this case values of type person can run the 'greet' func and have access to the person value they are used on.")
	p := person{"John", 33}
	fmt.Println("-- For example: 'p := person{\"John\", 33}'. p.greet():")
	p.greet()

}
