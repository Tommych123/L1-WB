package main

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) {
	done := make(chan bool, 1) // создаем канал
	go func() {
		<-time.After(d) // создаем таймер
		done <- true
	}()
	<-done // блокируем пока горутина не отправит сигнал
}

func main() {
	fmt.Println("Стар программы: ", time.Now())
	Sleep(2 * time.Second) // вызов нашей фунции Sleep
	fmt.Println("Конец программы: ", time.Now())
}
