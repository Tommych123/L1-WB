package main

import (
	"fmt"
	"math"
)

type Point struct { // объявление структуры
	x float64
	y float64
}

func NewPoint(x, y float64) Point { // конструктор для Point
	return Point{x, y}
}

func (p *Point) Distance(other Point) float64 { // метод для расчета расстояния между 2 точками
	return math.Sqrt((math.Pow((p.x - other.x), 2)) + (math.Pow((p.y - other.y), 2))) // формула для вычисления расстояния между 2 точками на плоскости
}

func main() {
	A := NewPoint(1, 2)        // первая точка
	B := NewPoint(4, 5)        // вторая точка
	fmt.Println(A.Distance(B)) // вызов метода для расчета расстояния между ними
}
