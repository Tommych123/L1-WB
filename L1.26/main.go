package main

import (
	"fmt"
	"strings"
)

func isUnique(str string) bool {
	str = strings.ToLower(str)     // переводим строку в нижний регистр
	symbols := make(map[rune]bool) // мапа для проверки на уникальность символа в строке
	for _, k := range str {
		if symbols[k] { // если символ уже встречался возвращем false
			return false
		}
		symbols[k] = true // если не встречался записываем его по руне
	}
	return true
}

func main() {
	var str string
	fmt.Println("Введите строку для проверки на уникальность")
	fmt.Scan(&str)     // ввод исходных данных
	if isUnique(str) { // вызов функции и вывод в зависимости от исхода функции
		fmt.Println("Строка из уникальных символов")
	} else {
		fmt.Println("Строка не из уникальных символов")
	}
}
