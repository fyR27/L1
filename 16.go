package main

import (
	"fmt"
	"sort"
)

func sort_arr(x []int, c chan []int) {
	sort.Ints(x) // запускаем быструю сортировку
	c <- x
}

func main() {
	x := []int{20, 20, 548, 12, 543, 575, 1000, 54, 76}
	c := make(chan []int)
	go sort_arr(x, c) // подключаем горутину
	x = <-c
	fmt.Print(x)
}
