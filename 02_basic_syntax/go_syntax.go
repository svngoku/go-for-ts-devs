package main

import "fmt"

// This example shows basic control structures and variables
func main() {
	for i := 0; i < 3; i++ {
		if i%2 == 0 {
			fmt.Printf("%d is even\n", i)
		} else {
			fmt.Printf("%d is odd\n", i)
		}
	}
}
