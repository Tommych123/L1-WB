package main

import (
	"fmt"
	"math/big"
)

func main() {
	first := new(big.Int)
	second := new(big.Int)
	result := new(big.Int) // переменная для результата

	fmt.Println("Введите первое число")
	fmt.Scan(first) // получаем первое число
	fmt.Println("Введите второе число")
	fmt.Scan(second) // получаем второе число
	var operand string
	fmt.Println("Введите операцию( + - / *)")
	fmt.Scan(&operand) // получаем нужную операцию
	switch operand {
	case "+":
		result.Add(first, second) // сложение
	case "-":
		result.Sub(first, second) // вычитание
	case "/":
		if second.Sign() == 0 { // проверка на деление на 0
			fmt.Println("Ошибка: деление на 0")
			return
		}
		result.Div(first, second) // деление
	case "*":
		result.Mul(first, second) // умножение
	default:
		fmt.Println("Неизвестная операция") // в случае неверно введенного знака
		return
	}
	fmt.Println("Результат:", result) // печатаем результат
}
