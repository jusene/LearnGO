package main

import "fmt"

// 简单工厂
type Shape interface {
	Draw()
}

type Rectangle struct {
}

func (s *Rectangle) Draw() {
	fmt.Println("draw Rectangle")
}

type Square struct {
}

func (s *Square) Draw() {
	fmt.Println("draw Square")
}

type SimplenessFactory struct {
}

func (s *SimplenessFactory) GetShape(shapeType string) (Shape, bool) {
	if shapeType == "" {
		return nil, false
	}

	switch shapeType {
	case "Rectangle":
		return new(Rectangle), true
	case "Square":
		return new(Square), true
	default:
		return nil, false
	}
}

func TestSimplenessFactory() {
	f := new(SimplenessFactory)
	var s Shape
	s, ok := f.GetShape("Rectangle")
	if ok {
		s.Draw()
	}
}

func main() {
	TestSimplenessFactory()

	/*
		var s Shape
		s = new(Rectangle)
		s.Draw()
	*/

	/*
		r := &Rectangle{}
		Shape(r).Draw()
	*/
}
