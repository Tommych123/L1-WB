package main

import "fmt"

func quickSort(arr []int) []int { // функция сортировки
	if len(arr) == 1 || len(arr) == 0 { // если длина массива 0 или 1 возвращаем его
		return arr
	}
	pivot := arr[0]             // берем первый элемент как опорный
	var left []int              // массив в котором элементы меньше либо равны опорному
	var right []int             // массив в котором элементы больше опорного
	for _, v := range arr[1:] { // перебираем весь массив кроме первого элемента
		if v <= pivot { // сортировка оставшегося массива на left и right
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	left = quickSort(left)                       // рекурсивный вызов сортировки
	right = quickSort(right)                     // рекурсивный вызов сортировки
	return append(append(left, pivot), right...) // возвращаем склееный массив в виде отсортированный left + pivot + отсортированный right
}

func main() {
	arr := []int{13, 8, -1, 7, 50, 20}          // исходный массив
	fmt.Println("Массив до сортировки", arr)    // печатаем до функции
	arr = quickSort(arr)                        // вызов функции сортировки
	fmt.Println("Массив после сортировки", arr) // печатаем после функции
}
