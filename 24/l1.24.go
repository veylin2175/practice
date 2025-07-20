package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func newPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) Distance(other Point) float64 {
	return math.Sqrt((p.x-other.x)*(p.x-other.x) + (p.y-other.y)*(p.y-other.y))
}

func main() {
	a := newPoint(1.0, 2.0)
	b := newPoint(-1.0, 1.0)

	dist := a.Distance(*b)
	fmt.Println(dist)
}
