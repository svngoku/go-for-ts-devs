package main

import "fmt"

// update changes the value pointed by p
func update(p *string) {
	*p = "gopher"
}

func main() {
	name := "initial"
	namePointer := &name

	fmt.Println(name)
	fmt.Println(namePointer) // prints the address

	update(namePointer)

	fmt.Println(name)
}
