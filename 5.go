package main

import (
	"fmt"
	"time"
)

// Горутина которая записывает в себя поданный на нее массив чисел
func channel(n []int, c chan []int) {
	a := []int{}
	for i := 0; i < len(n); i++ {
		a = append(a, n[i])
	}
	c <- a
}

func main() {
	c := make(chan []int)
	n := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for range time.Tick(10 * time.Second) { // ограничение работы по времени
		go channel(n[:], c)
		fmt.Print(<-c)
		return // необходим для остановки горутины , чтобы не было зацикливания
	}
}
