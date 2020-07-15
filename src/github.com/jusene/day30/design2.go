package main

import "fmt"

type AbstractFactory interface {
	GetColor(colorType string) Color
	GetShape(shapeType string) Shape
}

type AbsFactory struct {
}

func (s *AbsFactory) GetShape(shapeType string) Shape {
	if shapeType == "" {
		return nil
	}

	switch shapeType {
	case "Rectangle":
		return new(Rectangle)
	case "Square":
		return new(Square)
	default:
		return nil
	}
}

func (s *AbsFactory) GetColor(colorType string) Color {
	if colorType == "" {
		return nil
	}

	switch colorType {
	case "Red":
		return new(Red)
	case "Green":
		return new(Green)
	case "Blue":
		return new(Blue)
	default:
		return nil
	}
}

type Color interface {
	Fill()
}

type Red struct {
}

func (s *Red) Fill() {
	fmt.Println("Red Fill")
}

type Green struct {
}

func (s *Green) Fill() {
	fmt.Println("Green Fill")
}

type Blue struct {
}

func (s *Blue) Fill() {
	fmt.Println("Blue Fill")
}

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

func testAbstractFactory() {
	f := new(AbsFactory)
	color := f.GetColor("Red")
	color.Fill()
	S := f.GetShape("Rectangle")
	S.Draw()
}

func main() {
	testAbstractFactory()
}
