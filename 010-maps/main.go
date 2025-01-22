package main

import (
	"fmt"
	"maps"
)

// https://gobyexample.com/maps
func main() {
	fmt.Println("Maps:")

	m := make(map[string]int)
	fmt.Println("-- Maps can be made using the following syntax 'm := make(map[string]int)':", m)

	m["a"] = 5
	fmt.Println("-- We can set map keys like so 'm[\"a\"]=5' and then retrive map values by 'm[\"a\"]':", m["a"])

	m["b"] = 10
	m["c"] = 20

	fmt.Println("-- We added a few more keys to the map:", m)

	fmt.Println("-- If a key doesn't have a value the value's type's zero value will be returned 'm[\"x\"]'", m["x"])

	delete(m, "b")
	fmt.Println("-- If we wish to delete a key from a map we can use the delete keyword 'delete(m, \"b\")'", m)

	clear(m)
	fmt.Println("-- If we wish to clear the whole map we can use the clear keyword 'clear(m)'", m)

	_, present := m["a"]
	fmt.Println("-- We can check to see if a key is present in a map by using the second return value of a map lookup like so: _, present := m[\"a\"]:", present)

	am := map[string]int{"a": 1}
	fmt.Println("-- We can delcare and initialize a map on the same line like so 'am := map[string]int{\"a\": 1}'", am)

	yam := maps.Clone(am)
	fmt.Println("-- We can use the `maps` pkg for several helpful map functions like 'yam := maps.Clone(am)':", yam)
}
