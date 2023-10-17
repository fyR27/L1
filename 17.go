package main

import "fmt"

func binarySearch(nums []int, key int) int {
	var i, j = 0, len(nums) - 1
	for i <= j { // если i элемент меньше последненго элемента
		var mid = i + (j-i)/2 // обозначаем середину между i и j
		if nums[mid] == key { // проверяем являяется ли число по середине нашим числом
			return mid // если да то выводим его
		}
		if key > nums[mid] { // если наше число больше тогда сдвигаем i на +1 от середины
			i = mid + 1
		} else { // иначе сдвигаем j на -1 от середины
			j = mid - 1
		}
	}
	return -1
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	index := binarySearch(a, 13)
	fmt.Print(index)
}
