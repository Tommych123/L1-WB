package main

import "fmt"

type TypeCConnector interface { // интерфейс клиента
	UseTypeC()
}

type JackConnector interface { // интерфейс устройства который несовместим с клиентом
	UseJack()
}

type Headphones struct{} // устройство

func (h *Headphones) UseJack() { // устройство реализует только метод для Jack разъема
	fmt.Println("Музыка через Jack connector")
}

type JackToTypeCAdapter struct { // адаптер между Jack и Type-C интерфейсами
	device JackConnector
}

func (a *JackToTypeCAdapter) UseTypeC() { // метод для подключения через адаптер
	fmt.Println("Подключаем адаптер")
	a.device.UseJack()
}

func main() {
	headphones := Headphones{}                         // создаем устройство с разъемом Jack
	adapter := JackToTypeCAdapter{device: &headphones} // подключаем адаптер Jack -> Type-C

	var phonePort TypeCConnector = &adapter // теперь к телефону можно подключить наушники с разъемом Jack через адаптер
	phonePort.UseTypeC()                    // вызов метода подключения
}
