# Introduction 1: Understand golang

## Printing

```go
// Print
fmt.Print()  //Prints output to the stdout console 
fmt.Fprint() // Returns number of bytes and an error 
fmt.Printf() // (The error is generally not worried about

// Fprint
fmt.Fprint() // Prints the output to an external source (file, browser) 
fmt.Fprintln() // Does not print to the stdout console
fmt.Fprintf() // Returns number of bytes, and any write errors

// Sprint
fmt.Sprint() // Stores output on a character buffer 
fmt.Sprintln() // Does not print to stdout console
fmt.Sprintf() // Returns the string you want to print
``` 