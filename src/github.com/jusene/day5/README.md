## Go语言复合数据类型

### 切片（slice）

切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。

切片是一个引用类型，它的内部结构包含地址、长度和容量。

#### 切片的定义

make和切面字面量

字面量
```go
var name []T
```

make()函数构造切片
```go
make([]T, size, cap)
```

```go
// 其长度和容量都是5个元素切片
slice := make([]string, 5)
fmt.Println(slice)

// 创建一个长度为3个元素，容量为5个元素的切片
slice := make([]int, 3, 5)

// 使用空字符串初始化第100个元素
slice := []string{99: ""}

//[]运算符里指定一个值，那么创建的就是数组而不是切片，只有不指定值的时候才会创建切片
array := [3]int{10, 20, 30} // 数组
slice := []int{10, 20, 30} //切片 
```

#### 空切片和nil

```go
package main

import "fmt"

func main() {
	var a []string   // 声明一个字符串切片
	var b = []string{}  // 声明一个整型切片并初始化
	slice := make([]string, 0)
	fmt.Println(a == nil) // true
	fmt.Println(b == nil) // false
	fmt.Println(slice == nil) // false
}
```

#### 切片的长度和容量

切片拥有自己的长度和容量，我们可以通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量。

#### 切片的使用

基于数组定义切片

```go
package main

import "fmt"

func main() {
	a := [5]int{55, 56, 57, 58, 59}
	b := a[1:4]
	fmt.Printf("%T %v\n", a, a)
	fmt.Printf("%T %v\n", b, b)
}
```

使用三个索引创建切片

```go
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
    // slice[i:j:k] 对底层容量k的切片，长度为j-i，容量为k-i
	newSlice := slice[2:3:4]
	fmt.Printf("%v", newSlice)
	fmt.Println(len(newSlice), cap(newSlice))
}
```

赋值分割

```go
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	slice1 := slice[2:5]
	slice2 := slice1[1:3]
	slice2[0] = 100
    // 切片共享同一个底层数组，如果一个切片修改底层数组的共享部分，另一个切片也受影响
    // 对切片进行再切片时，索引不能超过原数组的长度，否则会出现索引越界的错误
	fmt.Println(slice)
	fmt.Println(slice1)
	fmt.Println(slice2)
}
```

切片扩容

```go
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
    // append() 只能增加新切片的长度，而容量有可能会改变，也有可能不会改变，这取决于被操作的切片的可用容量。
	newslice := append(slice, 100, 200)
	fmt.Println(newslice)

    slice1 := []int{10, 20, 30}
    newslice1 := append(slice, slice1...)
    fmt.Println(newslice1)
}
```

长度和容量自动扩容

```go
package main

import "fmt"

func main() {
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v len:%d cap:%d ptr:%p \n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
}

[0] len:1 cap:1 ptr:0xc0000b4008 
[0 1] len:2 cap:2 ptr:0xc0000b4040 
[0 1 2] len:3 cap:4 ptr:0xc0000b8020 
[0 1 2 3] len:4 cap:4 ptr:0xc0000b8020 
[0 1 2 3 4] len:5 cap:8 ptr:0xc0000ac080 
[0 1 2 3 4 5] len:6 cap:8 ptr:0xc0000ac080 
[0 1 2 3 4 5 6] len:7 cap:8 ptr:0xc0000ac080 
[0 1 2 3 4 5 6 7] len:8 cap:8 ptr:0xc0000ac080 
[0 1 2 3 4 5 6 7 8] len:9 cap:16 ptr:0xc0000ba000 
[0 1 2 3 4 5 6 7 8 9] len:10 cap:16 ptr:0xc0000ba000 
```

从长度与容量的变化和指针的变化，append会在容量不足的时候，在底层数组上自动扩容

#### 切片本质

切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）

切片就是一个框，框住一块连续的内存

切片是应用类型，真正的数据都是保存在底层数组中

#### 判断切片是否为空

```go
package main

import "fmt"

func main() {
	a := []int{}
	fmt.Println(len(a), a == nil)

	var b []int
	fmt.Println(len(b), b == nil)

	c := make([]int, 0, 5)
	fmt.Println(len(c), c == nil)
}
```

#### 切片不能直接比较

切片之间是不能比较的，我们不能使用==操作符来判断两个切片是否含有全部相等元素。 切片唯一合法的比较操作是和nil比较。 一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。但是我们不能说一个长度和容量都是0的切片一定是nil

```go
var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
```

#### 切片遍历

```go
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	for i := 0; i < len(s); i++ {
		fmt.Println(i)
	}

	for index, value := range s {
		fmt.Println(index, value)
	}
}
```

#### 使用copy()函数复制切片

切片是引用类型，切片共享同一个底层数组，如果一个切片修改底层数组的共享部分，另一个切片也受影响

```go
copy(destSlice, srcSlice []T)
```

```go
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	c := make([]int, 5, 6)
	copy(c, a)
	fmt.Println(a)
	fmt.Println(c)
	c[0] = 1000
	fmt.Println(a)
	fmt.Println(c)
}
```

#### 从切片中删除元素

```go
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	newNums := append(nums[:2], nums[3:]...) // 通过切片重新组成新的切片*
	fmt.Println(newNums)
}
```

#### 限制容量

```go
package main

import "fmt"

func main() {
source := []int{10, 20, 30, 40, 50}
slice := source[2:3]
fmt.Println(len(slice), cap(slice))
// 容量足够，append会修改source的值
slice = append(slice, 11)
fmt.Println(slice, source)

slice1 := source[2:3:3]
// 容量不足， append会新创建一个新的底层数组
fmt.Println(len(slice1), cap(slice1))
slice2 := append(slice1, 11)
fmt.Println(slice2, source)
}
```

#### 多维切片

```go
slice := [][]int{{10}, {100, 200}}
slice[0] = append(slice[0], 20)
fmt.Println(slice)
```

#### 将切片传递给函数

```go
package main

import "fmt"

func main() {
	slice := make([]int, 1e6)
	fmt.Println(foo(&slice))
}


func foo(slice *[]int) *[]int {
	return slice
}
```