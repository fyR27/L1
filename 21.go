package main

import "fmt"

type Target interface {
	Operation()
}

// адаптируемый класс
type Adaptee struct {
}

// Метод адаптируемого класса, который нужно вызвать где-то
func (adaptee *Adaptee) AdaptedOperation() {
	fmt.Println("\nI am AdaptedOperation()")
}

// класс конкретного адаптера
type ConcreteAdapter struct {
	*Adaptee
}

// реализация метода интерфейса, реализующего вызов адаптируемого класса
func (adapter *ConcreteAdapter) Operation() {
	adapter.AdaptedOperation()
}

// конструктор адаптера
func NewAdapter(adaptee *Adaptee) Target {
	return &ConcreteAdapter{adaptee}
}

// основной метод для демонстрации
func main() {
	fmt.Println("\nAdapter demo:")
	adapter := NewAdapter(&Adaptee{})
	adapter.Operation()
}
