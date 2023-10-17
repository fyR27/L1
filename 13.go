package main

import "fmt"

func Remove(nums []int, c chan []int) { // пример выполнен на самом простом массиве из 2 чисел , если надо будет заменить в более объемном массиве
	for i := 0; i < len(nums)-1; i++ { // то просто меняем i на позиции тех переменных которые нам необходиом поменять
		nums[i] = nums[i] + nums[i+1]
		nums[i+1] = nums[i] - nums[i+1]
		nums[i] = nums[i] - nums[i+1]
	}
	c <- nums
}
func main() {
	c := make(chan []int)
	x := []int{1, 2}
	go Remove(x[:], c) // запуск горутины на выходе получаем переставленные элементы
	a := <-c
	fmt.Println(a)
}
