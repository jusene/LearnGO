## Go语言复合数据类型

### 数组（array）

数组定义

```go
var 数组变量名 [元素数量]T
```

声明数组

```go
package main

import "fmt"

func main() {
	var array [5]int
	fmt.Println(array)
	array1 := [5]int{10, 20, 30, 40, 50}
	fmt.Println(array1)
    // 编译器根据初始值的个数自行推断数组的长度
	array2 := [...]int{10, 20, 30, 40, 50}
    fmt.Printf("%T\n", array2)
	fmt.Println(array2)
	array3 := [5]int{1: 10, 3: 20}
	fmt.Println(array3)
}
```

数组遍历

```go
package main

import "fmt"

func main() {
	var a = [...]string{"北京", "上海", "广州", "深圳", "杭州"}
	// for循环遍历
	for i:=0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// for range遍历
	for index, value := range a {
		fmt.Println(index, value)
	}
}
```

指针数组

```go
package main

import "fmt"

func main() {
	array := [5]*int{1: new(int)}
	*array[1] = 10
	fmt.Println(array)
	fmt.Println(*array[1])
}
```

多维数组

```go
package main

import "fmt"

func main() {
	var array [4][2]int
	fmt.Println(array)

	// 声明并初始化外层数组中索引为1和3的元素
	array1 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println(array1)

	// 声明并初始化外层数组和内层数组的单个元素
	array2 := [4][2]int{1: {0: 21}, 3: {1: 41}}
	fmt.Println(array2)
}
```

二维数组遍历

```go
package main

import "fmt"

func main() {
	a := [3][2]string{
		{"a", "b"},
		{"c", "d"},
		{"e", "f"},
	}
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}

	// 多维数组只有第一层可以使用...来让编译器推导数组长度
	b := [...][2]string{
		{"a", "b"},
		{"c", "d"},
		{"e", "f"},
	}
	for _, v1 := range b {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
	// 不支持多维数组的内层使用
	/*
	c := [3][...]string{
		{"a", "b"},
		{"c", "d"},
		{"e", "f"},
	}
	for _, v1 := range c {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
	*/

}
```

数组是值的类型

数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。

```go
package main

import "fmt"

func main() {
	i := [3]int{10, 20, 30}
	A(i)
	fmt.Println(i)

	j := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	B(j)
	fmt.Println(j)

	k := [2]int{10, 20}
	C(&k)
	fmt.Println(k)
}

func A(a [3]int) {
	a[0] = 100
}

func B(b [3][2]int) {
	b[2][0] = 100
}

func C(c *[2]int) { // 指针数据传递
	c[0] = 100
}
```

- 数组支持 “==“、”!=” 操作符，因为内存总是被初始化过的。
- `[n]*T`表示指针数组，`*[n]T`表示数组指针 。

将数组传递给函数

需要传递100万int元素数组，仅8m内存消耗
```go
package main

import "fmt"

var array [1e6]int

func main() {
	foo(array)
}

func foo(array [1e6]int) {
	fmt.Println(array)
}
```
指针优化
```go
package main

import "fmt"

var array [1e6]int

func main() {
	foo(&array)
}

func foo(array *[1e6]int) {
	fmt.Println(array)
}
```
