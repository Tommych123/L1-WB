package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func worker(id int, ch <-chan int) { // функция воркера
	for info := range ch { // цикл по информации в канале
		fmt.Printf("Worker %d read %d\n", id, info) // работа с информацией из канала
	}
}

func main() {
	ch := make(chan int, 10) // создаем канал
	var N int                // переменная для количества воркеров
	if len(os.Args) < 2 {    // проверка на правильный запуск программы
		fmt.Println("Usage: go run main.go <num_workers>")
		return
	}

	N, err := strconv.Atoi(os.Args[1]) // считываем количество воркеров из запуска программы
	if err != nil || N <= 0 {          // проверка на валидность введенного количества
		fmt.Println("Invalid number of workers")
		return
	}
	for i := 1; i <= N; i++ { // запуск воркеров каждого в своей горутине
		go worker(i, ch)
	}
	info := 0 // информация которая будет положена в канал
	for {     // бесконечный цикл для записи
		ch <- info              // запись в канал
		info++                  // изменение информации
		time.Sleep(time.Second) // небольшая задержка между записями для удобного отслеживания работы программы
	}
}
