package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	// Ввод N из аргументов командной строки при запуске программы
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <N>")
		return
	}
	// Считывание N и проверка на валидность
	N, err := strconv.Atoi(os.Args[1])
	if err != nil || N < 0 {
		fmt.Println("Invalid N")
		return
	}
	// Создание канала
	ch := make(chan int)
	// Запуск горутины для записи в канал
	go func() {
		i := 0
		for {
			ch <- i
			i++
			time.Sleep(1 * time.Second)
		}
	}()
	// Создание таймера который сработает через N секунд
	timeout := time.After(time.Duration(N) * time.Second)
	for {
		select {
		// Считываем из канала и выводим
		case v := <-ch:
			fmt.Println(v)
		// Завершаем программу по истечению времени
		case <-timeout:
			fmt.Println("Time is over")
			return
		}
	}
}
