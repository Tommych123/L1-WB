package main

import "fmt"

func binarySearch(arr []int, target int) int { // функция бинарного поиска
	left := 0             // левая граница массива
	right := len(arr) - 1 // правая граница массива

	for left <= right { // цикл идет пока границы не сомкнутся
		mid := (left + right) / 2 // середина массива
		if arr[mid] == target {   // если средний элемент массива равен искомому возвращаем индекс
			return mid
		} else if arr[mid] < target { // если средний элемент массива меньше искомого то отбрасываем левую часть от середины
			left = mid + 1
		} else { // если средний элемент массива больше искомого то отбрасываем правую часть от середины
			right = mid - 1
		}
	}
	return -1 // если не нашли элемент
}

func main() {
	arr := []int{1, 2, 3, 5, 9, 10, 21}   // исходный массив
	target := 9                           // искомый элемент
	position := binarySearch(arr, target) // вызов функции бинарного поиска
	fmt.Println(position)                 // печать результата функции
}
