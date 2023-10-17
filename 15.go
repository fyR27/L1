package main

var justString string

func someFunc() {
	v := createHugeString(1 << 10) // делим строку на 2 10 раз
	justString = v[:]              // как я понял нам просто надо брать значение этой строки изменном виде по всей ее длине
}

func main() {
	someFunc()
}
