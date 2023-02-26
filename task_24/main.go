package main

import (
	"fmt"
	"math"
)

// Объявляем структуру Point с двумя координатами x и y
type Point struct {
	x float64
	y float64
}

// Конструктор для структуры Point
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Считаем расстояние между двумя точками по формуле sqrt((x2 - x1)^2 + (y2 - y1)^2)
func (p *Point) DistanceTo(q *Point) float64 {
	dx := q.x - p.x
	dy := q.y - p.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p := NewPoint(1.0, 2.0)
	q := NewPoint(4.0, 6.0)

	distance := p.DistanceTo(q)

	fmt.Printf("Distance between point%v and point%v = %v", *p, *q, distance)
}
