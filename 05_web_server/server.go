package main

import (
    "fmt",
    "net/http"
)


func home(w http.ResponseWriter, *http.Request) {
    fmt.Println("Home path")
}



func main() {
    
}
