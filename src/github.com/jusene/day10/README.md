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

- 接受者变量：接受者中的参数变量名在命名时，建议使用接受者类型名称首字母小写， 而不是this、self之类的命名
- 接受者类型：接受者类型和参数类似，可以是指针类型和非指针类型
- 方法名、参数列表、返回参数，具体格式与函数定义相同

```go
package main

import "fmt"

// Person 结构体
type Person struct {
	name string
	age int
}

// NewPerson 构造函数
func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Dream Person 的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言\n", p.name)
}

func main() {
	p1 := NewPerson("jusene", 27)
	p1.Dream()
}
```
方法与函数的区别， 函数不属于任何类型，方法属于特定的类型

```go
package main

import "fmt"

type TwoInt struct {
	a int
	b int
}

func (t *TwoInt) AddThem() int {
	return t.a + t.b
}

func (t *TwoInt) AddToParam(param int) int {
	return t.a + t.b + param
}

func (_ *TwoInt) Echo() {
	fmt.Println("hello world")
}

func main() {
	two := new(TwoInt)
	two.a = 10
	two.b = 20

	fmt.Println(two.AddThem())
	fmt.Println(two.AddToParam(30))
	two.Echo()
}
```

非结构体的例子

```go
package main

import "fmt"

type IntVector []int

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

func main() {
	fmt.Println(IntVector{1, 2, 3}.Sum())
}
```

- 值类型的接受者
- 指针类型的接受者

```go
package main

import "fmt"

type People struct {
	name string
	age int
}

// 指针类型接受者
func(p *People) SetAge(newAge int) {
	p.age = newAge
}

// 值类型的接受者
func(p People) SetAge2(newAge int) int {
	p.age = newAge
	return p.age
}


func main() {
	per := People{"jusene", 27}
	fmt.Println(per.age) // 27
	per.SetAge(30)
	fmt.Println(per.age) // 30

	newAge := per.SetAge2(40)
	fmt.Println(per.age) // 30
	fmt.Println(newAge) // 40
}
```

### 函数与方法的区别

- 函数将变量作为参数： Function(recv);方法在变量上被调用： recv.Method()
- 当接受者是指针时，方法可以改变接受者的值和状态，这一点函数也能实现，当参数作为指针传递，即通过应用调用
- 接受者必须是一个显式名字，这个名字必须在方法中被使用，接受者类型必须在和方法同样在包中被声明。

### 结构体匿名字段

结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。

```go
package main

import "fmt"

type Person struct {
	string
	int
}

func main() {
	p1 := Person{"jusene", 27}
	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.string, p1.int)
}
```
匿名字段默认采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。

### 嵌套结构体

一个结构体中可以嵌套包含另一个结构体或结构体指针。

```go
package main

import "fmt"

type Address struct {
	Province string
	City string
}

type User struct {
	Name string
	Gender string
	Address Address
}

func main() {
	user1 := User{
		Name:    "jusene",
		Gender:  "男",
		Address: Address{
			Province: "浙江",
			City: "杭州",
		},
	}

	fmt.Printf("%#v", user1)
}
```

### 嵌套匿名结构体

```go
package main

import "fmt"

type Address struct {
	Province string
	City string
}

type User struct {
	Name string
	Gender string
	Address // 匿名结构体
}

func main() {
	var user2 User
	user2.Name = "JUSENE"
	user2.Gender = "男"
	user2.Address.Province = "浙江" // 通过匿名结构体，字段名访问
	user2.City = "杭州" // 直接访问匿名结构体的字段名

	fmt.Printf("%#v", user2)
}
```

当访问结构体成员时会先在结构体中查找该字段，找不到再去匿名结构体中查找。

### 结构体的字段名冲突

嵌套结构体内部可能存在相同的字段名。这个时候为了避免歧义需要指定具体的内嵌结构体的字段。

```go
package main

import (
	"fmt"
	"time"
)

type Address struct {
	Province string
	City string
	CreateTime time.Time
}

type Email struct {
	Account string
	CreateTime time.Time
}

type User struct {
	Name string
	Gender string
	Address
	Email
}

func main() {
	var user3 User
	user3.Name = "jusene"
	user3.Gender = "man"
	user3.Address.CreateTime = time.Now()
	user3.Email.CreateTime = time.Now()
	
	fmt.Printf("%#v", user3)
}
```

### 结构体的继承

Go语言中使用结构体也可以实现其他编程语言中面向对象的继承。

```go
package main

import "fmt"

type Aniaml struct {
	name string
}

func (a *Aniaml) move() {
	fmt.Printf("%s动了\n", a.name)
}

type Dog struct {
	Feet int8
	*Aniaml // 通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪\n", d.name)
}

func main() {
	d1 := &Dog{
		Feet:   4,
		Aniaml: &Aniaml{
			name: "葫芦",
		},
	}

	d1.move()
	d1.wang()
}
```


### 结构体与json序列化

package main

import (
	"encoding/json"
	"fmt"
)

// Student
type Student struct {
	ID int
	Gender string
	Name string
}

// Class
type Class struct {
	Title string
	Students []*Student
}

func main() {
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}

	for i := 0; i <= 10; i++ {
		stu := &Student{
			ID:     i,
			Gender: "man",
			Name:   fmt.Sprintf("Stu%02d", i),
		}
		c.Students = append(c.Students, stu)
	}

	// Json 序列化：结构体 --> JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("%s\n", data)

	// Json 反序列化：JSON格式字符串 --> 结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := new(Class)
	err1 := json.Unmarshal([]byte(str), c1)
	if err1 != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	
	fmt.Printf("%#v\n", c1)
}

### 结构体标签

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	StuID int `json:"id"` // 通过指定tag实现json序列化该字段时的key
	Gender string
	name string
}

func main() {
	s1 := Student{
		StuID:     1,
		Gender: "man",
		name:   "jusene",
	}

	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Println(string(data)) // {"id":1,"Gender":"man"} name 小写字母开头，小写表示私有，只能在结构体中使用
}
```

### 嵌入类型的方法和继承

```go
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y)
}

type NamePoint struct {
	Point
	name string
}

func main() {
	n := &NamePoint{
		Point: Point{3, 4},
		name:  "gogogo",
	}
	fmt.Println(n.Abs(), n.name)
}
```

### 多重继承

```go
package main

import "fmt"

type Camera struct {}

func (c *Camera) TakePicture() string {
	return "拍照"
}

type Phone struct {}

func (c * Camera) Call() string {
	return "响铃"
}

type CameraPhone struct {
	Camera
	Phone
}

func main() {
	cp := new(CameraPhone)
	fmt.Println(cp.TakePicture())
	fmt.Println(cp.Call())
}
```