## Go语言结构体

Go语言没有类的概念，也不支持类的继承等面向对象的概念，Go语言的解决方法是通过结构体内嵌接口等方式实现面向对象编程。

### 自定义类型和类型别名

```go
package main

import "fmt"

type myInt int64 // 自定义类型
type yourInt = int64 // 类型别名

func main() {
	//var a int64
	var n myInt
	n = 100
	//a = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	// fmt.Println(n == a) // 自定义类型是一种全新的类型，类型不同不能比较

	var b int64
	var m yourInt
	m = 100
	b = 100
	fmt.Println(m)
	fmt.Printf("%T\n", m)
	fmt.Println(m == b)
}
```

### 结构体的定义

结构体定义的一般方式，组成结构体类型的那些数据称为字段，每个字段都有一个类型和一个名字，在一个结构体中，字段名字必须是唯一。

```
type author struct {
    field1 type1
    field2 type2
    ...
}
```

```go
package main

import "fmt"

type author struct {
	name string
	age int
	tag []string
}

func main() {
	var s author
	// s := author

	s.name = "jusene"
	s.age = 27
	s.tag = []string{"devops", "ops"}

	fmt.Println(s)
}
```

### 结构体实例化

只有当结构体实例化时，才会真正分配内存，只有分配内存才能使用结构体。

```
var 结构体实例 结构体类型
```

使用new创建结构体，返回指向分配内存的指针

```go
package main

import "fmt"

type person struct {
	name string
	age int
	city string
}

func s1() person {
	var p person
	p.name = "jusene"
	p.age = 27
	p.city = "hangzhou"

	fmt.Printf("%#v\n", p)
	return p
}

func s2() *person {
	s := new(person)
	s.name = "jusene"
	s.age = 27
	s.city = "hangzhou"

	fmt.Printf("%#v\n", s)
	return s
}

func s3() *person {
	t := &person{
		name: "jusene",
		age:  27,
		city: "hangzhou",
	}

	fmt.Printf("%#v\n", t)
	return t
}


func main() {
	s1()
	s2()
	s3()
}
```

### 匿名结构体

定义一些临时数据结构等场景下还可以使用匿名结构体

```go
package main

import "fmt"

func main() {
	var user struct{Name string; Age int}
	user.Name = "JUSENE"
	user.Age = 27
	fmt.Printf("%#v\n", user)
}
```

### 结构体初始化

```go
package main

import "fmt"

type person struct {
	name string
	age int
	city string
}

func s1() person {
	p := person{ // 使用键值对初始化
		name: "jusene", 
		age:  27,
		city: "hangzhou",
	}
	return p
}

func s2() *person {
	p := &person{
		name: "jusene",
		age:  27,
		city: "hangzhou",
	}
	return p
}

func s3() *person {
    p := &person{  // 使用列表初始化，可以不写键对值
       "jusene",
       27,
       "hangzhou",
    }
    return p
}

func main() {
	p := s1()
	p1 := s2()
    p2 := s3()
	fmt.Printf("%#v\n", p)
	fmt.Printf("%#v\n", p1)
    fmt.Printf("%#v\n", p2)
}
```

### 结构体内存布局

结构体占用一块连续的内存

```go
package main

import "fmt"

type test struct {
	a int8
	b int8
	c int8
	d int8
}

func main() {
	n := test{1, 2, 3, 4}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)
}

n.a 0xc00000a0c8
n.b 0xc00000a0c9
n.c 0xc00000a0ca
n.d 0xc00000a0cb
```

### 空结构体

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a struct{}
	var b struct{User string; Age int}
	
	fmt.Println(unsafe.Sizeof(a)) // 空结构体不占用空间
	fmt.Println(unsafe.Sizeof(b))
}
```

### 结构体进阶

```go
package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{ // 结构体列表
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		fmt.Printf("%#v", stu)
		m[stu.name] = &stu // &stu  指针指向最后循环的数据
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name) // v.name "大王八"
	}
}
```

### 递归结构体

结构体类型可以通过引用自身来定义

```
// data字段用于存放有效数据，su指针指向后续节点
type Node struct {
    data float64
    su *Node
}
```

### 带标签的结构体

```go
package main

import (
	"fmt"
	"reflect"
)

type tagType struct {
	goods bool "是否有存货"
	name string "商品名称"
	price float64 "商品价格"
}

func main() {
	tt := tagType{
		goods: true,
		name:  "IPHONE SE",
		price: 1000,
	}

	for i := 0; i < 3 ; i++ {
		refTag(tt, i)
	}
}

/*
在一个变量上调用reflect.TypeOf()可以获取变量的正确类型
如果变量是一个结构体类型，就可以通过Field来索引结构体的字段，然后就可以使用Tag属性
*/

func refTag(tt tagType, ix int) {
	ttType := reflect.TypeOf(tt)
	iField := ttType.Field(ix)
	fmt.Printf("%v\n", iField.Tag)
}
```

### 结构体构造函数

```go
package main

import "fmt"

type people struct {
	name, city string
	age int
}

func newPerson(name, city string, age int) *people {
	return &people{
		name: name,
		city: city,
		age: age,
	}
}

func main() {
	p := newPerson("jusene", "hangzhou", 27)
	fmt.Printf("%#v", p)
}
```

### 方法和接受者

Go语言中的方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）。接收者的概念就类似于其他语言中的this或者 self。

```
func (接受者变量  接受者类型) 方法名(参数列表) (返回参数) {
    函数体
}
```

