package main

import "fmt"

func main() {
	a := 10
	b := 7
	fmt.Printf("Начальные числа a=%d | b=%d\n", a, b)
	a = a + b // a=10+7=17
	b = a - b // b=17-7=10(начальное а)
	a = a - b // a=17-10=7(начальное b)
	fmt.Printf("Числа после обмена значениями a=%d | b=%d", a, b)
}
