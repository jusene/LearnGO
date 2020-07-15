## Go语音  设计模式

### 简单工厂模式

主要解决接口选择的问题，在不通的条件下使用不同的实例，让子类实现工厂接口，返回一个抽象产品，创建过程在子类中完成。

```go
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
```

## 抽象工厂

主要解决接口选择的问题，系统的产品有多于一个的产品族，而系统只需要消费某一族的产品，在一个工厂中聚合多个同类产品。

```go
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
```

