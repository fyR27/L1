package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "sun dog flower"
	words := strings.Fields(name)
	for i := len(words)/2 - 1; i >= 0; i-- { // идем с конца массива
		opp := len(words) - 1 - i
		words[i], words[opp] = words[opp], words[i] // меняем противоположностоящие элементы в массиве
	}
	fmt.Print(words)
}
