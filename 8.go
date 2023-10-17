package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int64 = 10
	var a string
	var c string
	b := 3
	a = strconv.FormatInt(x, 2)
	for i := 0; i < len(a); i++ {
		if i == b {
			if a[i] == uint8(48) { // проверяем если наша переменная равна 0 в уникоде тогда меняем ее на 1
				c += "1"

			} else { // в противном случае меняем на 0
				c += "0"

			}
		} else if i != b { // если это не наша переенная просто записываем ее в новую строку
			c += string(a[i])
		}
	}
	fmt.Print(c)
}
