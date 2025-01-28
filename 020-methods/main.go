package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func (p person) greet()     { fmt.Printf("Hi I'm %s and I am %d years old...\n", p.name, p.age) }
func (p *person) stateAge() { fmt.Printf("Hi I'm %d years old...\n", p.age) }

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
	fmt.Println("-- We can also attach methods to POINTERS to VALUES of TYPES instead of VALUES of TYPES. By doing this we can mutate the value in question (as opposed to creating copies within the function)")
	fmt.Println("-- Recall as well that Go will automatically handle pointer dereferncing for us so we can call it just like it were attached to a value.")
	fmt.Println("-- Let's show an example of using a pointer receiver:")
	fmt.Println("func (p *person) stateAge() { fmt.Printf(\"Hi I'm %d years old...\\n\", p.age) }")
	fmt.Println("-- 'p.stateAge()':")
	p.stateAge()
	fmt.Println("NOTE: In the majority of cases it LIKELY makes sense to attach your methods to pointers...")

}
