package main

import "fmt"

// реализация простейшего счетчика
func counter(c int, k chan int) {
	for c < 10 {
		c++
	}
	k <- c
}

func main() {
	k := make(chan int)
	c := 0
	go counter(c, k)
	a := <-k
	fmt.Printf("Counter is %d", a)
}
