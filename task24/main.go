package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p Point) Distance(p2 Point) float64 {
	x := p2.x - p.x
	y := p2.y - p.y

	return math.Sqrt(x*x + y*y)
}

func main() {
	point1 := NewPoint(3, 4)
	point2 := NewPoint(6, 8)
	point3 := NewPoint(0, 0)
	point4 := NewPoint(-3, -4)

	fmt.Printf("Расстояние между точками (%.1f, %.1f) и (%.1f, %.1f): %.2f\n",
		point1.x, point1.y, point2.x, point2.y, point1.Distance(point2))
	fmt.Printf("Расстояние между точками (%.1f, %.1f) и (%.1f, %.1f): %.2f\n",
		point2.x, point2.y, point3.x, point3.y, point2.Distance(point3))
	fmt.Printf("Расстояние между точками (%.1f, %.1f) и (%.1f, %.1f): %.2f\n",
		point3.x, point3.y, point4.x, point4.y, point3.Distance(point4))
}
