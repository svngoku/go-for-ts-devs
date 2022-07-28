package main

import (
	"fmt"
	"strings"
)

func comparaison() {
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))
}


func main() {
    comparaison()
}
