package main

import "fmt"

// функция производит рачет сумы квадратов каждого элемента в массиве
func sqrtSlice(x []int, c chan int) {
	sum := 0
	for i := 0; i < len(x); i++ {
		x[i] *= x[i]
		sum += x[i]
	}
	c <- sum
}

func main() {
	c := make(chan int)
	x := []int{1, 2, 3, 4, 5, 6, 7}
	go sqrtSlice(x[:], c) // запсукаем горутину которая выведет нам сумму квардатов всех элементов в массиве
	a := <-c
	fmt.Println(a)
}
