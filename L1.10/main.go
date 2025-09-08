package main

import "fmt"

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5} // исходный массив температур
	m := make(map[int][]float64)                                        // мапа для распределения их
	for i := 0; i < len(arr); i++ {
		if arr[i] < -10 || arr[i] > 10 { // температуры которые вне диапазона [-10;10]
			m[int(arr[i]/10.0)*10] = append(m[int(arr[i]/10.0)*10], arr[i])
		} else { // температуры относящиеся к группе 0
			m[0] = append(m[0], arr[i])
		}
	}
	fmt.Println(m) // вывод мапы
}
