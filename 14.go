package main

import (
	"fmt"
	"reflect"
)

// объявляем структуру в которой буем оперировать строками
type structHere struct {
	A, B string
}

// объявляем интерфейс в котором будет храниться метод Name()
type Core interface {
	Name() string
}

// объявляем функцию s в которой используется метод Name
func (s *structHere) Name() string {
	return s.A + s.B
}

func main() {
	var c Core
	sh := structHere{A: "2", B: "1"} // присваиваем значения переменным
	c = &sh
	fmt.Printf("Method Name() has type of %s", reflect.TypeOf(c.Name())) // печатаем тип метода
}
