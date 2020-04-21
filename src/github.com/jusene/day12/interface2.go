package main

import "fmt"

type Shaper interface {
	Area() float64
}

type Square struct {
	side float64
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

type Rectangle struct {
	length, width float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func main() {
	sq := new(Square)
	sq.side = 5
	r := Rectangle{
		length: 5,
		width:  3,
	}

	shapes := []Shaper{Shaper(sq), Shaper(r)}
	for _, n := range shapes {
		fmt.Println(n.Area())
	}
}
