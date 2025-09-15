package main

import "fmt"

var justString string

func someFunc() {
	v := createHugeString(1 << 10)       // Вызов функции
	justString = string([]byte(v[:100])) // берем последние 100 символов строки
}

func createHugeString(size int) string {
	return string(make([]byte, size)) //Создание строки
}

func main() {
	someFunc()                   // Вызов функции
	fmt.Println(len(justString)) // Печать длины обрезанной строки
}
