## Go语言指针

任何数据载入内存后，在内存都有他们的地址，这就是指针。而为了保存一个数据在内存中的地址，需要指针变量。

Go语言中的指针：`&`(取指针)和`*`(根据地址取值)

### 指针地址和指针类型

```
ptr := &v // v的类型为T
```

```go
package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a)
	fmt.Printf("b:%p type:%T\n", b, b)
	fmt.Println(&b) // 指针的指针
}

a:10 ptr:0xc00000a0d8
b:0xc00000a0d8 type:*int
0xc000006028
```

### 指针取值

```go
package main

import "fmt"

func main() {
	// 指针取值
	a := 10
	b := &a
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
}
```

```go
package main

import "fmt"

func main() {
	a := 10
	modify1(10)
	fmt.Println(a)
	modify2(&a) // 指针操作可以改变原来的值
	fmt.Println(a)
}

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}
```

### make和new

```go
package main

import "fmt"

func main() {
	/*
	var a *int
	*a = 100 // 空指针 panic 只定义了应用类型，没有初始化
	fmt.Println(*a)
	*/
	a := new(int)
	*a = 100
	fmt.Println(*a)

	/*
	var b map[string]int // 空指针 panic
	b["a"] = 100
	fmt.Println(b)
	 */
	b := make(map[string]int)
	b["a"] = 100
	fmt.Println(b)
}
```

 在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。Go语言中new和make是内建的两个函数，主要用来分配内存。
 
 #### new

```
func new(Type) *Type
```

```go
package main

import "fmt"

func main() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Println(*a)
	fmt.Println(*b)
}
```

#### make

make也是用来分配内存的，区别于new，它只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。

```
func make(t Type, size ...IntegerType) Type
```

### new和make的区别

- 二者都是用来分配内存的
- make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身
- new用于类型的内存分配，而且内存对应的值为类型零值，返回的是指向类型的指针

