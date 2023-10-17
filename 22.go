package main

import "fmt"

func main() {
	var a int64 = 2000000000
	var b int64 = 10000000
	umn := 0
	del := 0
	summ := 0
	minus := 0
	umn = int(a * b)   // умножение
	del = int(a / b)   // деление
	summ = int(a + b)  // сложение
	minus = int(a - b) // вычитание
	fmt.Printf(" Умножение: %d.\n Деление: %d.\n Сумма: %d.\n Вычитание: %d.\n", umn, del, summ, minus)

}
