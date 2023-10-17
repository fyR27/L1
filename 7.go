package main

import (
	"fmt"
)

// маппим элементы поданые с массива строк
func Mapping(x []string, c chan map[int]string) {
	m := make(map[int]string)
	for i := 0; i < len(x); i++ {
		m[i+1] = x[i]
	}
	c <- m
}

func main() {
	c := make(chan map[int]string)
	x := []string{"first", "second", "third"}
	go Mapping(x[:], c) // запускаем горутины чтобы на выходе получить размапленную строку
	a := <-c
	fmt.Println(a)
}
