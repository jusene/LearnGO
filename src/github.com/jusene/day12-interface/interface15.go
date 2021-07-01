package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

type Shaper interface {
	Area() float64
}

func (sq *Square) Area() float64 {
	return sq.side * sq.side
}

func (cl *Circle) Area() float64 {
	return cl.radius * math.Pi
}

func main() {
	var areaInf Shaper
	sq := new(Square)
	sq.side = 5

	areaInf = sq
	switch t := areaInf.(type) {
	case *Square:
		fmt.Printf("%T, %v", t, t)
	case *Circle:
		fmt.Printf("%t, %v", t, t)
	case nil:
		fmt.Printf("%T, %v", t, t)
	default:
		fmt.Printf("%T", t)
	}
}
