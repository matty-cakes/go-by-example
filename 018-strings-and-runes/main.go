package main

import (
	"fmt"
	"unicode/utf8"
)

func examineRune(r rune) {
	// Notice the single quotes here
	if r == 't' {
		fmt.Println("🚨 ALERT found a letter \"t\"")
	}
}

// https://gobyexample.com/strings-and-runes
func main() {
	fmt.Println("String and Runes:")
	s := "stuff"
	fmt.Println("-- I have a string 's' whose value is stuff... see...", s)
	fmt.Println("-- I can use the builtin 'len' to see the length of 's' like so len(s):", len(s))

	fmt.Println("-- Remember in the 'range' lesson we mentioned that strings are slices of bytes?")
	fmt.Println("-- What do you think will happen when we loop over 's'?")
	for i := 0; i < len(s); i++ {
		fmt.Printf("---- The value of s at index %d is %v\n", i, s[i])
		fmt.Printf("---- The value of s at index %d CAST AS A HEX VALUE is %x\n", i, s[i])
	}

	fmt.Println("-- The bytes in strings in Go are known as 'runes' we can inspect the runes in a string using the utf8 pkg. Ex: 'utf8.RuneCountInString(s))'")

	fmt.Println("-- Rune count for 's':", utf8.RuneCountInString(s))

	fmt.Println("-- Note: some runes will span 2 bytes so the results of rune count and len might not a;ways be the same... Cool!")

	fmt.Println("-- There is some magic in the way range works for Go strings. Check this out...")
	for index, val := range s {
		fmt.Printf("--- %U starts at %d\n", val, index)
	}

	for i, w := 0, 0; i < len(s); i += w {
		fmt.Println(s[i:])
		r, widthOfRune := utf8.DecodeRuneInString(s[i:])
		fmt.Println(r)
		fmt.Println(widthOfRune)
		examineRune(r)
		fmt.Println("-----")
		w = widthOfRune
	}

	fmt.Println("-- Note: This lesson is mostly dealing with grokking more text encoding and I wouldn't worry too much about it to be honest...")
}
