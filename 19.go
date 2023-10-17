package main

import (
	"fmt"
)

func main() {
	s := "strings"
	new := []byte{}
	for ch := range s { // записываем всю строку в массив байт
		new = append(new, s[ch])
	}
	for i := len(new)/2 - 1; i >= 0; i-- { // идем по циклу с конца массива
		opp := len(new) - 1 - i
		new[i], new[opp] = new[opp], new[i] // меняем местами противоположностоящие элементы
	}
	fmt.Println(string(new))
}
