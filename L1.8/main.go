package main

import "fmt"

func setBit(number int64, i int, value int) int64 {
	mask := int64(1 << i) // создаем маску 1 сдвинуто на i позиций влево
	if value == 1 {
		return number | mask // установка в 1
	}
	return number &^ mask // установка в 0
}

func main() {
	var number int64 = 5              // 0101₂
	fmt.Println(setBit(number, 0, 0)) // пример сброса i-того бита в 0
	fmt.Println(setBit(number, 1, 1)) // пример установки 1 в i-тый бит
}
