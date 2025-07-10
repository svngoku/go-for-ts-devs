package main

import (
	"fmt"
	"os"
)

// mayPanic intentionally panics
func mayPanic() {
	panic("a problem")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	mayPanic()
	_, err := os.Create("file")
	if err != nil {
		panic(err)
	}
}
