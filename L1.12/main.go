package main

import "fmt"

func main() {
	arr1 := []string{"cat", "cat", "dog", "cat", "tree"} // исходный массив
	m := make(map[string]bool)                           // мапа для определения какие уникальные строки есть в массиве
	for _, v := range arr1 {                             // перебираем все строки в массиве, все уникальные заносим в мапу
		m[v] = true
	}
	var result []string // итоговый массив
	for v := range m {
		result = append(result, v) // добавляем уникальные строки в итоговый массив через мапу
	}
	fmt.Println(result) // вывод результата
}
