package main

import "fmt"

// первая горутина отвечает за наполнение нового массива данным из массива старого
func conv1(x []int, c chan []int) {
	arr1 := []int{}
	for i := 0; i < len(x); i++ {
		arr1 = append(arr1, x[i])
	}
	c <- arr1
}

// вторая горутина отвечает за наполнение нового массива квадратами переменных старого массива
func conv2(x []int, k chan []int) {
	arr2 := []int{}
	for i := 0; i < len(x); i++ {
		x[i] *= x[i]
		arr2 = append(arr2, x[i])
	}
	k <- arr2
}

func main() {
	c := make(chan []int)
	k := make(chan []int)
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	go conv2(x, k) // заупск в таком порядке по причине того что идет изменение массива
	go conv1(x, c)
	_, b := <-c, <-k
	fmt.Printf("Array 2 is contain:%v.", b)

}
