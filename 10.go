package main

import "fmt"

func main() {
	m := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	minusTwenty := []float64{} // создае новые массивы в которых будем хранить все данные по условию
	minusTen := []float64{}
	zero := []float64{}
	plusTen := []float64{}
	plusTwenty := []float64{}
	plusThirdty := []float64{}
	for i := 0; i < len(m); i++ {
		if m[i] > float64(-30.0) && m[i] < float64(-19.9) { // смотрим удовлетворение i переменной на принадлежность одному из условий
			minusTwenty = append(minusTwenty, m[i]) // при удовлетворении добавляем его
		} else if m[i] > float64(-20.0) && m[i] < float64(-9.9) {
			minusTen = append(minusTen, m[i])
		} else if m[i] > float64(-10.0) && m[i] < float64(0) {
			zero = append(zero, m[i])
		} else if m[i] > float64(9.9) && m[i] < float64(20) {
			plusTen = append(plusTen, m[i])
		} else if m[i] > float64(19.9) && m[i] < float64(30) {
			plusTwenty = append(plusTwenty, m[i])
		} else if m[i] > float64(29.9) && m[i] < float64(40) {
			plusThirdty = append(plusThirdty, m[i])
		}
	}
	fmt.Printf("-20:%v;\n -10:%v;\n 0:%v;\n 10:%v;\n 20:%v;\n 30:%v;\n", minusTwenty, minusTen, zero, plusTen, plusTwenty, plusThirdty)
}
