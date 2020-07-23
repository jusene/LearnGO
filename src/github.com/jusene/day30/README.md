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

## 单例设计模式

单例对象的类必须保证只有一个实例存在。许多时候整个系统只需要拥有一个的全局对象，这样有利于我们协调系统整体的行为。比如在某个服务器程序中，该服务器的配置信息存放在一个文件中，这些配置数据由一个单例对象统一读取

### 懒汉模式

```go
package main

import (
	"fmt"
)

// 单例设计模式
// 懒汉模式
type SingleObject struct {
	Count int
}

var singleObj *SingleObject

// 存在线程安全问题，多线程时，会创建多个对象，所有出现饿汉模式
func GetInstance1() *SingleObject {
	if singleObj == nil {
		singleObj = new(SingleObject)
	}
	return singleObj
}

func testSingleton() {
	s1 := GetInstance1()
	s1.Count = 5
	fmt.Printf("%v, %v", &s1, s1)
	s2 := GetInstance1()
	fmt.Printf("%v, %v", &s2, s2)
}

func main() {
	testSingleton()
}
```

### 饿汉模式

```go
package main

import "fmt"

// 单例设计模式
// 饿汉模式

type SingleObject struct {
	Count int
}

var singleObj *SingleObject

// 饿汉模式将在包加载的时候就创建单例对象，当程序中用不到该对象时，浪费了一部分空间
func init() {
	singleObj = new(SingleObject)
}

func GetInstance2() *SingleObject {
	return singleObj
}

func main() {
	s1 := GetInstance2()
	s1.Count = 5
	fmt.Printf("%v, %v", &s1, s1)
	s2 := GetInstance2()
	fmt.Printf("%v, %v", &s2, s2)
}
```

### 双重检查机制

```go
package main

import (
	"fmt"
	"sync"
)

// 单例设计模式
// 双重检查机制
type SingleObject struct {
	Count int
}

var singleObj *SingleObject

var lock *sync.Mutex = &sync.Mutex{}
func GetInstance3() *SingleObject {
	if singleObj == nil {
		lock.Lock()
		defer lock.Unlock()
		singleObj = new(SingleObject)
	}
	return singleObj
}

func testSingleton1() {
	s1 := GetInstance3()
	s1.Count = 5
	fmt.Printf("%v, %v", &s1, s1)
	s2 := GetInstance3()
	fmt.Printf("%v, %v", &s2, s2)
}

func main() {
	testSingleton1()
}
```

## 建造者设计模式

造者模式（Builder Pattern）使用多个简单的对象一步一步构建成一个复杂的对象。这种类型的设计模式属于创建型模式，它提供了一种创建对象的最佳方式

一个 Builder 类会一步一步构造最终的对象。该 Builder 类是独立于其他对象的

1. Product需要创建的复杂对象
2. Builder用来规范建造者
3. ConcreteBuilder具体的Builder实现，主要用来根据不用的业务来创建对象的所有组件
4. Director 用来规范复杂对象的创建流程

```go
package main

import "fmt"

type Computer struct {
	CPU string
	Memory string
	HardDisk string
}

func (c *Computer) SetCPU(cpu string) {
	c.CPU = cpu
}

func (c *Computer) GetCPU() string {
	return c.CPU
}

func (c *Computer) SetMemory(memory string) {
	c.Memory = memory
}

func (c *Computer) GetMemory() string {
	return c.Memory
}

func (c *Computer) SetHardDisk(hardDisk string) {
	c.HardDisk = hardDisk
}

func (c *Computer) GetHardDisk() string {
	return c.HardDisk
}

// Builder规范建造者
type Builder interface {
	SetCPU(cpu string) Builder
	SetMemory(memory string) Builder
	SetHardDisk(hardDisk string) Builder
	Build() *Computer
}

// 根据不同的业务完成创建对象的组建
type ComputerBuilder struct {
	computer *Computer
}

func (c *ComputerBuilder) SetCPU(cpu string) Builder {
	if c.computer == nil {
		c.computer = new(Computer)
	}
	c.computer.SetCPU(cpu)
	return c
}

func (c *ComputerBuilder) SetMemory(memory string) Builder {
	if c.computer == nil {
		c.computer = new(Computer)
	}
	c.computer.SetMemory(memory)
	return c
}

func (c *ComputerBuilder) SetHardDisk(hardDisk string) Builder {
	if c.computer == nil {
		c.computer = new(Computer)
	}
	c.computer.SetHardDisk(hardDisk)
	return c
}

func (c *ComputerBuilder) Build() *Computer {
	return c.computer
}

type Director struct {
	Builder Builder
}

func (d Director) Create(cpu string, memory string, hardDisk string) *Computer {
	return d.Builder.SetCPU(cpu).SetMemory(memory).SetHardDisk(hardDisk).Build()
}

func main() {
	builder := new(ComputerBuilder)
	director := Director{Builder: builder}
	computer := director.Create("17", "32G", "4T")
	fmt.Println(*computer)
}
```