package main

import "fmt"

func Remove(nums []int, i int, c chan []int) {
	if i >= 0 && i < len(nums) {
		nums[i] = nums[len(nums)-1] // просто заменяем нужный нам элемент на последний и выводим новый массив без последнего чтобы не было повторения
		c <- nums[:len(nums)-1]
	} else {
		c <- nums
	}
}
func main() {
	c := make(chan []int)
	x := []int{1, 2, 3, 4, 5, 6, 7}
	i := 0
	go Remove(x[:], i, c)
	a := <-c
	fmt.Println(a)
}
