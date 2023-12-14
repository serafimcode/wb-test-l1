package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

func newPoint(x, y float64) point {
	return point{x: x, y: y}
}

/*
Расстояние между 2 мя точками равняетя квадратному корню
суммы разностей координат. Используем везде float64, чтобы не мучиться
с приведением типов
*/
func getDistance(p1, p2 point) float64 {
	squareX := (p1.x - p2.x) * (p1.x - p2.x)
	squareY := (p1.y - p2.y) * (p1.y - p2.y)

	return math.Sqrt(squareX + squareY)
}
func main() {
	p1 := newPoint(5, 2)
	p2 := newPoint(10, 2)

	fmt.Println(getDistance(p1, p2))
}
