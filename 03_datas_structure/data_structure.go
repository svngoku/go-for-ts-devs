package main

import (
	"fmt"
)

// Write a function that takes a string as input and returns the string reversed.
func reversed(f string) string {
	n := 0
	rune := make([]rune, len(f))
	for _, r := range f {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	output := string(rune)
	fmt.Println(output)

	return output
}

func main() {
	var name string = "The quick brown 狐 jumped over the lazy 犬"
	reversed(name)
}
