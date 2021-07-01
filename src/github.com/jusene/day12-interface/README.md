## Go语言接口

接口（interface）定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节。

- 接口(interface)是一种类型
- 接口(interface)是方法的集合

### 接口的定义

```
type 接口类型名 interface {
    方法名1（参数列表1） 返回值列表1
    方法名2（参数列表2） 返回值列表2
}
```

```go
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
```

### 接口多态

```go
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
```

实现接口的条件，一个对象只要全部实现了接口中方法，那么久实现了这个接口，接口就是一个需要实现的方法列表。

### 接口值接受者和指针接受者的区别

```go
package main

import "fmt"

type Mover interface {
	move()
}

type dog struct {}

func (d dog) move() {
	fmt.Println("狗动了")
}

type cat struct {}

func (c *cat) move() {
	fmt.Println("猫动了")
}

func main() {
	wangcai := dog{} //值接受者实现接口
	fugui := &dog{}
	Mover(wangcai).move()
	Mover(fugui).move() // Go语言中有对指针类型变量求值的语法糖，dog指针fugui内部会自动求值*fugui。

	//miao := cat{} // 值接受者实现接口
	aiai := &cat{}
	//Mover(miao).move() // Mover不接受miao的类型
	Mover(aiai).move() // Mover可以接受*cat类型

}
```

### 一个类型实现多个接口

```go
package main

import "fmt"

type Sayer interface {
	say()
}

type Mover interface {
	move()
}

type dog struct {
	name string
}

func (d dog) say() {
	fmt.Printf("%s会叫\n", d.name)
}

func (d dog) move() {
	fmt.Printf("%s会动\n", d.name)
}

func main() {
	var x Sayer
	var y Mover

	a := dog{name: "旺财"}
	x = a
	y = a
	x.say()
	y.move()

	b := &dog{name: "来福"}
	x = b
	y = b
	x.say()
	y.move()

	c := dog{name: "葫芦"}
	Sayer(c).say()
	Mover(c).move()

	d := &dog{name: "琅琊"}
	Sayer(d).say()
	Mover(d).move()
}
```

### 多个类型实现统一接口

```go
package main

import "fmt"

type Mover interface {
	move()
}

type dog struct {
	name string
}

type car struct {
	name string
}

func (d dog) move() {
	fmt.Printf("%s会跑\n", d.name)
}

func (c car) move() {
	fmt.Printf("%s会跑\n", c.name)
}

func main() {
	var x Mover
	var a = dog{name: "旺财"}
	var b = car{name: "路虎"}
	x = a
	x.move()
	x = b
	x.move()
}
```

### 接口结构体嵌套

```go
package main

import "fmt"

type WashingMachine interface {
	wash()
	dry()
}

type dryer struct {
	name string
}

func (d dryer) dry() {
	fmt.Printf("%s甩甩\n", d.name)
}

type haier struct {
	dryer // 嵌入甩干机
}

func (h haier) wash() {
	fmt.Printf("%s洗洗\n", h.dryer.name)
}

func main() {
	var machine WashingMachine
	h := haier{dryer{name: "海尔"}}
	machine = h
	machine.dry()
	machine.wash()
}
```

### 接口嵌套

接口与接口间可以通过嵌套创造出新的接口。

```go
package main

import "fmt"

type Sayer interface {
	say()
}

type Mover interface {
	move()
}

// 嵌套得到的接口的使用与普通接口一样，这里我们让cat实现animal接口：
type animal interface {
	Sayer
	Mover
}

type cat struct {
	name string
}

func (c cat) say() {
	fmt.Println("喵喵")
}

func (c cat) move() {
	fmt.Println("动了")
}

func main() {
	var x animal
	x = cat{name: "花花"}
	x.move()
	x.say()
}
```

### 空接口

空接口是指没有定义任何方法的接口，空接口类型的变量可以存储任意类型的变量

```go
package main

import "fmt"

func main() {
	var x interface{}
	s := "hello world"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)
}
```

### 空接口应用

空接口实现可以接收任意类型的函数参数

```go
package main

import "fmt"

func show(a interface{}) (b interface{}){
	fmt.Printf("type:%T value:%v\n", a, a)
	return true
}
func main() {
	a := 21212
	x := show(a)
	fmt.Printf("type:%T value:%v\n", x, x)
	b := false
	show(b)
}
```

```go
import "fmt"

func main() {
	var studInfo = make(map[string]interface{})
	studInfo["name"] = "jusene"
	studInfo["age"] = 27
	studInfo["marr"] = false

	fmt.Printf("%v", studInfo)
}
```

### 类型断言

```go
package main

import "fmt"

func main() {
	var x interface{}
	x = "hello world"
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}
```

```go
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
```