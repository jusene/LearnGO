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

func main() {
	sq := new(Square)
	sq.side = 5
	area := Shaper(sq)
	fmt.Println(area.Area())
}
