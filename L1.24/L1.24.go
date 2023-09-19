package main

import (
	"fmt"
	"math"
)

// Point - структура для представления точки
type Point struct {
	x, y float64
}

// NewPoint - конструктор для создания новой точки
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// DistanceTo - метод для вычисления расстояния между текущей точкой и другой точкой
func (p1 Point) DistanceTo(p2 Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаем две точки
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)

	// Вычисляем расстояние между точками
	distance := point1.DistanceTo(point2)

	fmt.Printf("Расстояние между точкой 1 и точкой 2: %.2f\n", distance)
}
