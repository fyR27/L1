package main

import "fmt"

func main() {
	fmt.Println(unicLatters("abcd"))
	fmt.Println(unicLatters("abCdefAaf"))
	fmt.Println(unicLatters("aabcd"))
}

func unicLatters(str string) (string, bool) {
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] { // побуквенно проверяем переменные в массиве
				return str, false
			}
		}
	}

	return str, true
}
