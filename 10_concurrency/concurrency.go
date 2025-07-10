package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan<- int) {
	time.Sleep(time.Millisecond * 50)
	ch <- id * 2
}

func main() {
	ch := make(chan int)
	for i := 1; i <= 3; i++ {
		go worker(i, ch)
	}

	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
}
