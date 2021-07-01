package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y)
}

type NamePoint struct {
	Point
	name string
}

func main() {
	n := &NamePoint{
		Point: Point{3, 4},
		name:  "gogogo",
	}
	fmt.Println(n.Abs(), n.name)
}
