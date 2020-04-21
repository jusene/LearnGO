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


