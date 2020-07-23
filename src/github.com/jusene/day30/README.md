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

## 原型设计模式

### Golang深拷贝浅拷贝

- 浅拷贝同根
```go
package main

import "fmt"

func main() {
	// 数组
	nums := [5]int{}
	nums[0] = 1
	fmt.Printf("nums: %v, len: %d, cap: %d\n", nums, len(nums), cap(nums))

	// 切片
	dnums := nums[1:2]
	dnums[0] = 3
	fmt.Printf("nums: %v, len: %d, cap: %d\n", nums, len(nums), cap(nums))
	fmt.Printf("dnums: %v, len: %d, cap: %d\n", dnums, len(dnums), cap(dnums))
}
```

- 扩容摆脱同根
slice与array最大的区别在于slice不需要指定大小会自动扩容等一些特性,满足扩容策略，这时候内部就会重新申请一块内存空间，将原本的元素拷贝一份到新的内存空间上。此时其与原本的数组就没有任何关联关系了，再进行修改值也不会变动到原始数组。
```go
package main

import "fmt"

func main() {
	nums := [3]int{}
	nums[0] = 1

	fmt.Printf("nums: %v, len: %d, cap: %d\n", nums, len(nums), cap(nums))

	dnums := nums[0:2]
	dnums = append(dnums, []int{2, 3}...)
	dnums[0] = 100
	fmt.Printf("nums: %v, len: %d, cap: %d\n", nums, len(nums), cap(nums))
	fmt.Printf("nums: %v, len: %d, cap: %d\n", dnums, len(dnums), cap(dnums))
}
```

- Empty and nil
empty
一个有分配空间（Empty）一个没有分配空间（nil）
```go
package main

import "fmt"

func main() {
	nums := []int{}
	renums := make([]int, 0)
	var anums []int
	fmt.Printf("nums: %v, len: %d, cap: %d\n", nums, len(nums), cap(nums))
	fmt.Printf("renums: %v, len: %d, cap: %d\n", renums, len(renums), cap(renums))
	fmt.Printf("anums: %v, len: %d, cap: %d\n", anums, len(anums), cap(anums))

	if nums == nil {
		fmt.Println("nums is nil")
	}
	if renums == nil {
		fmt.Println("renums is nil")
	}
	if anums == nil {
		fmt.Println("anums is nil")
	}
}
```
- 浅拷贝
```go
package main

import "fmt"

//速度速值
type Speed int

//风扇转速
type FanSpeed struct {
	Speed Speed
}

//售价
type Money struct {
	Length float64
}

// 内存数量以及大小
type Memory struct {
	Count int
	MemorySize []int
}

// 电脑信息
type Computer struct {
	SystemName string
	UseNumber int
	Memory Memory
	Fan   map[string]FanSpeed
	Money Money
}


func main() {
	pc1 := Computer{
		SystemName: "Windows",
		UseNumber:  1000,
		Memory:     Memory{
			Count:      4,
			MemorySize: []int{32, 32, 32, 32},
		},
		Fan: map[string]FanSpeed{"left": {2500}, "right": {2000}},
		Money:      Money{123.45},
	}

	// 浅拷贝
	pc2 := pc1
	fmt.Printf("PcInfo Pc1:%v, Pc2:%v\n", pc1, pc2)
	// 修改切片内容以及map信息影响pc1
	pc2.SystemName = "MacOS"
	pc2.UseNumber = 100
	pc2.Memory.Count = 3
	pc2.Memory.MemorySize[0] = 8
	pc2.Fan["left"] = FanSpeed{200}
	fmt.Printf("PcInfo Pc1:%v, Pc2:%v\n", pc1, pc2)
}
```

- 深拷贝

```go
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

//速度速值
type Speed int

//风扇转速
type FanSpeed struct {
	Speed Speed
}

//售价
type Money struct {
	Length float64
}

// 内存数量以及大小
type Memory struct {
	Count int
	MemorySize []int
}

// 电脑信息
type Computer struct {
	SystemName string
	UseNumber int
	Memory Memory
	Fan   map[string]FanSpeed
	Money Money
}

// 基于序列化和反序列化来实现对象的深度拷贝
// 需要深拷贝的变量必须首字母大写才可以被拷贝
func deepcopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func main() {
	pc1 := Computer{
		SystemName: "Windows",
		UseNumber:  1000,
		Memory:     Memory{
			Count:      4,
			MemorySize: []int{32, 32, 32, 32},
		},
		Fan: map[string]FanSpeed{"left": {2500}, "right": {2000}},
		Money:      Money{123.45},
	}

	// 深拷贝
	pc2 := new(Computer)
	if err := deepcopy(pc2, pc1); err != nil {
		panic(err.Error())
	}
	fmt.Printf("PcInfo Pc1:%v, Pc2:%v\n", pc1, pc2)

	pc2.SystemName ="MacOs"
	pc2.UseNumber =100
	pc2.Memory.Count =2
	pc2.Memory.MemorySize[0]=8
	pc2.Memory.MemorySize[1]=8
	pc2.Memory.MemorySize[2]=0
	pc2.Memory.MemorySize[3]=0
	pc2.Fan["left"]=FanSpeed{2000}
	pc2.Fan["right"]=FanSpeed{1500}
	fmt.Printf("PcInfo Pc1:%v, Pc2:%v\n", pc1, pc2)
}
```

### 原型模式设计

这种模式是实现了一个原型接口，该接口用于创建当前对象的克隆。当直接创建对象的代价比较大时，则采用这种模式。

```go
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

//速度速值
type Speed int

//风扇转速
type FanSpeed struct {
	Speed Speed
}

//售价
type Money struct {
	Length float64
}

//内存数量以及大小
type Memory struct {
	Count      int
	MemorySize []int
}

//电脑信息
type Computer struct {
	SystemName string              //系统名字
	UseNumber  int                 //使用次数
	Memory     Memory              //存储
	Fan        map[string]FanSpeed //风扇
	Money      Money               //售价
}

func (s *Computer) Clone() *Computer {
	resume := *s
	return &resume
}

func (s *Computer) BackUp() *Computer {
	pc := new(Computer)
	if err := deepCopy(pc, s); err != nil {
		panic(err.Error())
	}
	return pc
}

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func main() {
	pc1 := Computer{
		SystemName: "Windows",
		UseNumber:  1000,
		Memory:     Memory{
			Count:      4,
			MemorySize: []int{32, 32, 32, 32},
		},
		Fan: map[string]FanSpeed{"left": {2500}, "right": {2000}},
		Money:      Money{123.45},
	}

	// 浅拷贝
	pc2 := pc1.Clone()
	pc2.Memory.MemorySize[0] = 9999
	fmt.Printf("pcinfo pc1:%v, pc2:%v\n", pc1, pc2)

	// 深拷贝
	pc2 = pc1.BackUp()
	pc2.Memory.MemorySize[0] = 8888
	fmt.Printf("pcinfo pc1:%v, pc2:%v\n", pc1, pc2)
}
```

### 适配器模式

```go
package main

import "fmt"

type MediaPlayer interface {
	Play(audioType string, filename string)
}

type AdvanceMediaPlayer interface {
	PlayVlc(filename string)
	PlayMp4(filename string)
}

type VlcPlayer struct {

}

func (s *VlcPlayer) PlayVlc(filename string) {
	fmt.Println("Playing vlc file. Name: "+filename)
}

func (s *VlcPlayer) PlayMp4(filename string) {

}

type Mp4Player struct {

}

func (s *Mp4Player) PlayVlc(filename string) {
}

func (s *Mp4Player) PlayMp4(filename string) {
	fmt.Println("Playing mp4 file. Name: "+filename)
}

type MediaAdapter struct {
	advanceMediaPlayer AdvanceMediaPlayer
}

func (s *MediaAdapter) MediaAdapter(audioType string) {
	if audioType == "vlc" {
		s.advanceMediaPlayer = new(VlcPlayer)
	} else if audioType == "mp4" {
		s.advanceMediaPlayer = new(Mp4Player)
	}
}

func (s *MediaAdapter) Play(audioType string, filename string) {
	if audioType == "vlc" {
		s.advanceMediaPlayer.PlayVlc(filename)
	} else if audioType == "mp4" {
		s.advanceMediaPlayer.PlayMp4(filename)
	}
}

type AudioPlayer struct {
	mediaAdapter MediaAdapter
}

func (s *AudioPlayer) Play(audioType string, filename string) {
	if audioType == "mp3" {
		fmt.Println("Playing mp3 file. Name: " + filename)
	} else if audioType == "vlc" || audioType == "mp4" {
		s.mediaAdapter = MediaAdapter{}
		s.mediaAdapter.MediaAdapter(audioType)
		s.mediaAdapter.Play(audioType, filename)
	} else {
		fmt.Println("Invalid media. " + audioType + " format not supported")
	}
}

func main() {
	audioPlayer := AudioPlayer{}

	audioPlayer.Play("mp3", "beyond.mp3")
	audioPlayer.Play("mp4", "alone.mp4")
	audioPlayer.Play("vlc", "away.vlc")
	audioPlayer.Play("avi", "mind.avi")
}
```

### 策略模式

```go
package main

import "fmt"

// 策略模式
type Strategy interface {
	DoOperation(num1, num2 int) int
}

type OperationAdd struct {

}

func (o *OperationAdd) DoOperation(num1, num2 int) int {
	return  num1 + num2
}

type OperationSub struct {

}

func (o *OperationSub) DoOperation(num1, num2 int) int {
	return num1 - num2
}

type OperationMul struct {

}

func (o *OperationMul) DoOperation(num1, num2 int) int {
	return num1 * num2
}

type Context struct {
	strategy Strategy
}

func (c *Context) Context(strategy Strategy) *Context {
	c.strategy = strategy
	return c
}

func (c *Context) Execute(num1, num2 int) int {
	return c.strategy.DoOperation(num1, num2)
}

func main() {
	context := new(Context).Context(new(OperationAdd))
	fmt.Println(context.Execute(10, 5))

	context = new(Context).Context(new(OperationSub))
	fmt.Println(context.Execute(10, 5))

	context = new(Context).Context(new(OperationMul))
	fmt.Println(context.Execute(10, 5))
}
```

### 过滤器模式

