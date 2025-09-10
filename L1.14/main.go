package main

import "fmt"

func detectType(value interface{}) {
	switch value := value.(type) { // определение типа переменной
	case int:
		fmt.Println("Тип переменной - int") // переменная типа int
	case string:
		fmt.Println("Тип переменной - string") // переменная типа string
	case bool:
		fmt.Println("Тип переменной - bool") // переменная типа bool
	case chan int:
		fmt.Println("Тип переменной - chan int") // переменная типа chan int
	default:
		fmt.Printf("Неизвестный тип переменной %T ", value) // переменная неизвестного типа
	}
}

func main() {
	// объявляем все необходимые переменные
	numberValue := 5
	stringValue := "12345"
	boolValue := true
	channelValue := make(chan int)
	undefinedValue := 10.0
	detectType(numberValue)    // передаем int
	detectType(stringValue)    // передаем string
	detectType(boolValue)      // передаем bool
	detectType(channelValue)   // передаем chan
	detectType(undefinedValue) // передаем float64
}
