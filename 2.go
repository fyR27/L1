package main

import (
	"fmt"
)

// функция которвя возводит каждый элемент массива в квадрат
func sqrtSlice(x []int, c chan []int) {
	for i := 0; i < len(x); i++ {
		x[i] *= x[i]
	}
	c <- x
}

func main() {
	c := make(chan []int)
	x := []int{1, 2, 3, 4, 5, 6, 7}
	go sqrtSlice(x[:], c) // запускаем горутину которая по завершеию должна вывести квадрат каждого числа в массиве
	a := <-c
	fmt.Println(a)
}
