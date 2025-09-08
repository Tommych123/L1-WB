package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3} // исходный массив
	arr2 := []int{2, 3, 4} // исходный массив

	m := make(map[int]bool) // мапа для проверки пересечений
	var result []int        // итоговый массив

	for _, v := range arr1 {
		m[v] = true // проверка на то какие числа есть в первом массиве
	}
	for _, v := range arr2 {
		if m[v] { // проверка есть ли такие же числа во втором массиве
			result = append(result, v) // добавляем это число в итоговый массив
			delete(m, v)               // убираем дубликаты из мапы
		}
	}
	fmt.Println(result) // вывод результата
}
