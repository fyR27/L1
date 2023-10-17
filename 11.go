package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{20, 45, 68, 70}
	y := []int{10, 19, 24, 59, 100}
	new_arr := []int{}
	for i := 0; i < len(x); i++ { // добавляем переменные из 1 массива в новый массив
		new_arr = append(new_arr, x[i])
	}
	for i := 0; i < len(y); i++ { // добавляем переменные из 1 массива в новый массив
		new_arr = append(new_arr, y[i])
	}
	sort.Ints(new_arr) // сотрируем поучившийся массив
	fmt.Println(new_arr)
}
