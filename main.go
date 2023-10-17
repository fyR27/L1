package main

import (
	"fmt"
)

// объявляем структуру
type Action struct {
	A, B string
}

// объявляем интерфейс в котором будет храниться метод Name
type Human interface {
	Name() string
}

// объявляем функцию a с наследованием метода Name
func (a *Action) Name() string {
	return a.A + a.B
}

func main() {
	var c Human
	sh := Action{A: "2", B: "1"} // Присваеваем переменной значения в фунции
	c = &sh
	fmt.Printf(c.Name())
}
