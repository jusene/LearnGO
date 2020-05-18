## Go语言 数据结构和算法

- math/big 大整数的高精度计算实现
- math/cmplx 复数基本函数操作
- math/rand 伪随机数生成器
- sort 包含基本的排序方法，支持切片数据排序以及用户自定义数据集合排序
- index/suffixary 该包实现了后缀数组相关算法以支持许多常见的字符串操作
- container 实现三个复杂的数据结构，堆、链表、环

### 数据集合排序

数据集合（包括自定义数据类型的集合）排序需要实现sort.Interface接口的三个方法：
```
type Interface interface {
    // 获取数据集合元素的个数
    Len() int
    // 如果i索引的数据小于j所有的数据，返回true，不会调用
    // 下面的Swap()，即数据升序排序
    Less(i, j int) bool
    // 交换i和j索引的两个元素的位置
    Swap(i, j int)
}
```
```go
package main

import (
	"fmt"
	"sort"
)

// 学生成绩结构体
type StuScore struct {
	name string
	score int
}

type StuScores []StuScore

// Len() 人数
func (s StuScores) Len() int {
	return len(s)
}

// Less() 成绩将由低到高排序
func (s StuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

// Swap() 排序
func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	stus := StuScores{
		{"张三", 95},
		{"李四", 91},
		{"赵六", 96},
		{"王六", 90},
	}

	fmt.Println("=======默认=======")
	// 原始顺序
	for _, v := range stus {
		fmt.Println(v.name, ":", v.score)
	}

	fmt.Println()
	// stuScores已经实现了sort.Interface接口
	sort.Sort(stus)
	fmt.Println("=======排序之后=======")
	for _, v := range stus {
		fmt.Println(v.name, v.score)
	}

	fmt.Println(sort.IsSorted(stus))
}
```

自定义类型实现了StuScore实现了sort.Interface接口，可以将对象作为sort.Sort和sort.IsSorted的参数传入。

实现降序
```go
package main

import (
	"fmt"
	"sort"
)

// 学生成绩结构体
type StuScore struct {
	name string
	score int
}

type StuScores []StuScore

// Len() 人数
func (s StuScores) Len() int {
	return len(s)
}

// Less() 成绩将由高到低排序
func (s StuScores) Less(i, j int) bool {
	return s[i].score > s[j].score
}

// Swap() 排序
func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	stus := StuScores{
		{"张三", 95},
		{"李四", 91},
		{"赵六", 96},
		{"王六", 90},
	}

	fmt.Println("=======默认=======")
	// 原始顺序
	for _, v := range stus {
		fmt.Println(v.name, ":", v.score)
	}

	fmt.Println()
	// stuScores已经实现了sort.Interface接口
	sort.Sort(stus)
	fmt.Println("=======排序之后=======")
	for _, v := range stus {
		fmt.Println(v.name, v.score)
	}

	fmt.Println(sort.IsSorted(stus))
}

    sort.Sort(sort.Reverse(stus)) // Reverse只是重新了Less方法
	for _, v := range stus {
		fmt.Println(v.name, v.score)
	}
```

sort包提供Reverse()方法，不必修改Less()代码
```
func Reverse(data interface) Interface
```

内部实现:
```
// 定义一个reverse结构类型，嵌入Interface接口
type reverse struct {
    Interface
}

// reverse结构类型的Less()方法拥有嵌入的Less()方法相反的行文
// Len() 和 Swap方法则保持嵌入类型的方法
func (r reverse) Less(i, j int) bool {
    return r.Interface.Less(j, i)
}

// 返回新的实现Interface接口的数据类型
func Reverse(data Interface) Interface {
    return &reverse(data)
}
```

```go
sort.Sort(sort.Reverse(stus)) // Reverse只是重新了Less方法
	for _, v := range stus {
		fmt.Println(v.name, v.score)
	}
```

sort.Search()
```go
package main

import (
	"fmt"
	"sort"
)

type num []int

func (s num) Len() int {
	return len(s)
}

func (s num) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s num) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	x := 11
	n := num{12, 22, 31, 413, 11}

	sort.Sort(n)
	pos := sort.Search(len(n), func(i int) bool {
		return n[i] >= x
	})

	if pos < len(n) && n[pos] == x {
		fmt.Println(x, "在n的位置为：", pos)
	} else {
		fmt.Println("n不包含元素", x)
	}
}
```

### 切片排序

sort包原生支持[]int,[]string,[]float64三种内建数据切片的排序操作

#### []int排序

```
type IntSlice []int

func (p IntSlice) Len() int {return len(p)}
func (p IntSlice) Less(i, j int) bool {return p[i] < p[j]}
func (p IntSlice) Swap(i, j int) {return p[i], p[j] = p[j], p[i]}
// IntSlice类型定义了Sort()方法，包装了sort.Sort()函数
func (p IntSlice) Sort() {Sort(p)}
// InitSlice类型定义了SearchInts()方法，包装了SearchInts()函数
func (p IntSlice) Search(x int) int (return SearchInts(p, x))
```

#### []float64排序

```
type Float64Slice []float64

func (p Float64Slice) Len() int {return len(p)}
func (p Float64Slice) Less(i, j int) bool {return p[i] < p[j] || isNaN(p[i])}
func (p Float64Slice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p Float64Slice) Sort() { Sort(p) }
func (p Float64Slice) Search(x float64) int {return SearchFloat64s(p, x) }
```

- func Float64s(a []float64)
- func Float64AreSorted(a []float64) bool
- func SearchFloat64s(a {}float64, x float64) int

#### []string 排序

```
type StringSlice []string

func (p StringSlice) Len() int {return len(p)}
func (p StringSlice) Less(i, j int) bool {return p[i] < p[j]}
func (p StringSlice) Swap(i, j int) {p[i], p[j] = p[j], p[i]}
func (p StringSlice) Sort() {Sort(p)}
func (p StringSlice) Search(x string) int {return SearchStrings(p, x)}
```

### container

#### 堆

```
type Interface interface {
    sort.Interface
    Push(x interface{}) // 添加一个元素x并返回Len()
    Pop(x interface{}) // 移除并返回元素长度Len() - 1
}
```

实现一个堆需要实现5个方法:

```go
package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func(h IntHeap) Len() int {
	return len(h)
}

func(h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func(h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func(h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func(h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Println(h)
	heap.Pop(h)
	fmt.Println(h)
}
```

### 链表

链表就是一个有prev和next指针的数组了，它维护两个type

```
type Element struct {
    next, prev *Element // 上一个元素和下一个元素
    list *List // 元素所在链表
    Value interface{} // 元素
}

type List struct {
    root Element // 链表的根元素
    len int // 链表的长度
}
```

链表是先创建list，然后往list中插入值，list就在内部创建一个Element，并在内部设置好Element的next、prev

```go
package main

import (
	list2 "container/list"
	"fmt"
)

func main() {
	list := list2.New()
	list.PushBack(1)
	list.PushBack(2)

	fmt.Printf("长度: %v\n", list.Len())
	fmt.Printf("第一个元素: %#v, %#v\n", list.Front(), list.Front().Value)
	fmt.Printf("第二个元素: %#v, %#v\n", list.Front().Next(), list.Front().Next().Value)

}
```

```
type Element
    func(e *Element) Next() *Element
    func(e *Element) Prev() *Element
type List
    func New() *List
    func (l *List) Back() *Element  // 最后一个元素
    func (l *List) Front() *Element // 第一个元素
    func (l *List) Init() *List // 链表初始化
    // 往某个元素后插入
    func (l *List) InsertAfter(v interface{}, mark *Element) *Element
    // 往某个元素前插入
    func (l *List) InsertBefore(v interface{}, mark *Element) *Element
    func (l *List) Len() int // 链表长度
    func (l *List) MoveAfter(e, mark *Element) // 把e元素移到mark之后
    func (l *List) MoveBefore(e, mark *Element) // 把e元素移到mark之前
    func (l *List) MoveToBack(e *Element) // 把e元素移到队列最后
    func (l *List) MoveToFront(e *Element) // 把e元素移到队列最头部
    func (l *List) PushBack(v interface{}) *Element // 在队列最后插入元素
    func (l *List) PushBackList(other *List) // 在队列最后插入新队列
    func (l *List) PushFront(v interface{}) *Element // 在队列头部插入元素
    func (l *List) PushFrontList(other *List) // 在队列头部插入新的队列
    func (l *List) Remove(e *Element) interface{} //删除某个元素
```

### 环

环的结构有点特殊，环的尾部就是头部
```
type Ring struct {
    next, prev *Ring
    Value interface{}
}
```

