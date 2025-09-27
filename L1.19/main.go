package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseString(s string) string { // функция разворота строки
	r := []rune(s)                                    // преобразуем строку в срез рун
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 { // i — индекс с начала, j — индекс с конца
		r[i], r[j] = r[j], r[i] // меняем местами символы на позициях i и j
	}
	return string(r) // преобразуем []rune обратно в строку
}

func main() {
	fmt.Print("Введите строку: ")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n') // читаем строку до символа перевода строки (\n)
	input = strings.TrimRight(input, "\r\n")               // убираем символы перевода строки \r\n
	fmt.Println(reverseString(input))
}
