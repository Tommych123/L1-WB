package main

import "fmt"

// Дана структура Human
type Human struct{
	Name string
	Age int
}

// Дана структура Action
type Action struct{
	Human
}

// Метод Say для структуры Human
func (H *Human)Say(){
	fmt.Println("I am Human")
}

// Метод Say для структуры Action
func (A *Action)Say(){
	fmt.Println("I am Action")
}


func main() {
	A:=Action{Human: Human{Name: "Вася", Age: 18}} // Создание переменной Action
	A.Say() // Происходит вызов метода для Action (I am Action)
	A.Human.Say() // Происходит вызов метода для Human (I am Human)
}