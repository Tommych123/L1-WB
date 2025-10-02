package main

import (
	"fmt"
)

// функция для разворота подмассива рун
func reverse(runes []rune, left, right int) {
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
}

func reverseWords(s string) string {
	runes := []rune(s)
	reverse(runes, 0, len(runes)-1) // перевернуть всю строку

	start := 0
	for i := 0; i <= len(runes); i++ { // перевернуть каждое слово используя цикл
		if i == len(runes) || runes[i] == ' ' {
			reverse(runes, start, i-1)
			start = i + 1
		}
	}
	return string(runes)
}

func main() {
	input := "snow dog sun" // исходные данные
	output := reverseWords(input) // вызов основной функции
	fmt.Println(output)
}
