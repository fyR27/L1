package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "cat, cat, god, cat, tree"
	n := strings.Fields(s)
	c := ""
	for i := 0; i < len(n)-1; i++ {
		for j := i + 1; j < len(n); j++ { // делаем двумерный цикл для сравнения всех элементов в массива
			if n[i] == n[j] { // если элемент повторяется то заменяем 2 из них на 1 , чтобы не словить ошибку о длине массива
			}
			if i == len(n)-2 { // провверка на последний элемент который может быть не изменен и удовлетворяет условию
				if n[j] != "1" {
					c += n[j] + " "
				}
			}

		}
		c += n[i] + " "
	}
	c = strings.Replace(c, "1", "", -1)   // замена всех 1 на пробелы
	c = strings.Replace(c, "  ", " ", -1) // удаление 2-х пробелов
	fmt.Println(c)
}