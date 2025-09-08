package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup           // группа ожидания для синхронизации горутин
	arr := []int{1, 2, 3, 4, 5}     // массив с исходными данными
	ch := make(chan int, len(arr))  // канал для записи чисел из массива
	ch2 := make(chan int, len(arr)) // канал для записи чисел*2 из ch
	wg.Add(1)
	go func(arr []int) {
		defer wg.Done()
		for i := 0; i < len(arr); i++ {
			ch <- arr[i] // записываем в канал ch числа из массива
		}
		close(ch)
	}(arr)
	wg.Add(1)
	go func(arr []int) {
		defer wg.Done()
		for v := range ch {
			ch2 <- v * 2 // записываем в канал ch2 числа из ch
		}
		close(ch2)
	}(arr)
	for v := range ch2 {
		fmt.Println(v) // вывод в stdout числе из ch2
	}
	wg.Wait()
}
