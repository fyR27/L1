package main

import (
	"fmt"
)

// создаем структуру
type Point struct {
	X int
	Y int
}

func main() {
	var size Point
	size.X = 10 // присваиваем значение переменным в структуре
	size.Y = 5
	fmt.Print(size.X - size.Y)
}
