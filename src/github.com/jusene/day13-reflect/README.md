## Go语言反射

### 变量的内在机制

- 类型信息：预先定义好的元数据 
- 值信息：程序运行过程可动态变化

### 反射介绍

反射是指在程序运行期对程序本身进行访问和修改的能力。程序在编译时，变量被转换为内存地址，变量名不会被编译器写入可执行部分，在运行城西的时候，程序无法获取自身的信息。

#### reflect包

任何接口值都是`一个具体类型`和`具体类型的值`，任何接口值在反射中都可以理解由`reflect.Type`和`reflect.Value`两部分组成，reflect包提供了reflect.TypeOf和reflect.ValueOf两个函数来获取任意对象的Value和Type。

```go
package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	t := reflect.ValueOf(x)
	fmt.Printf("type:%v value:%v\n", v, t)
}

func main() {
	var a float64 
	reflectType(a)
	var b int64 = 100
	reflectType(b)
}
```

### type name和type kind

在反射中关于类型还划分为两种：类型（Type）和种类（Kind）。因为在Go语言中我们可以使用type关键字构造很多自定义类型，而种类（Kind）就是指底层的类型。

```go
package main

import (
	"fmt"
	"reflect"
)

type myInt int64

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

func main() {
	var a *float64
	var b myInt
	var c rune
	reflectType(a)
	reflectType(b)
	reflectType(c)

	type person struct {
		name string
		age int
	}

	type book struct {
		title string
	}

	var d = person{
		name: "jusene",
		age:  27,
	}
	var e = book{title: "golang"}
	reflectType(d)
	reflectType(e)
}
```

Kind类型
```go
type Kind uint
const (
    Invalid Kind = iota  // 非法类型
    Bool                 // 布尔型
    Int                  // 有符号整型
    Int8                 // 有符号8位整型
    Int16                // 有符号16位整型
    Int32                // 有符号32位整型
    Int64                // 有符号64位整型
    Uint                 // 无符号整型
    Uint8                // 无符号8位整型
    Uint16               // 无符号16位整型
    Uint32               // 无符号32位整型
    Uint64               // 无符号64位整型
    Uintptr              // 指针
    Float32              // 单精度浮点数
    Float64              // 双精度浮点数
    Complex64            // 64位复数类型
    Complex128           // 128位复数类型
    Array                // 数组
    Chan                 // 通道
    Func                 // 函数
    Interface            // 接口
    Map                  // 映射
    Ptr                  // 指针
    Slice                // 切片
    String               // 字符串
    Struct               // 结构体
    UnsafePointer        // 底层指针
)
```

### 通过反射获得接口值

```go
package main

import (
	"fmt"
	"reflect"
)

func refelectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%T,%v\n", v, v)
	k := v.Kind() // 获取底层的数据类型
	switch k {
	case reflect.Float32:
		fmt.Printf("%T\n", v.Float())
		ret := float32(v.Float())
		fmt.Printf("%T, %v", ret, ret)
	case reflect.Int32:
		fmt.Printf("%T\n", v.Int())
		ret := int32(v.Int())
		fmt.Printf("%T, %v", ret, ret)
	}
}

func main() {
	var aa int32 = 100
	refelectValue(aa) // 通过反射获得接口的值
}
```
- Interface() 返回interface{}
- Int() 返回int64
- Uint() 返回uint64
- Float() 返回float64
- Bool() 返回bool
- Bytes() 返回[]bytes
- String() 返回string

### 通过反射设置变量的值

```go
package main

import (
	"fmt"
	"reflect"
)

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) // 修改的是副本，reflect包会panic
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 放射中使用Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	var a int64 = 100
	//reflectSetValue1(a) // panic: reflect: reflect.Value.SetInt using unaddressable value
	reflectSetValue2(&a)
	fmt.Println(a)
}
```

### isNil和isValid

isNil
> func (v Value) IsNil() bool

IsNil()报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic。

isValid
> func (v Value) IsValid() bool

IsValid()返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic。

```go
// IsNil()常被用于判断指针是否为空；IsValid()常被用于判定返回值是否有效。
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// *int类型空指针
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// 实例化一个匿名结构体
	b := struct {}{}
	// 尝试从结构体中查找abc字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	// 禅师从结构体中查找abc方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())

	// map
	c := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map不存在的键,", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")))
}
```

## 结构体反射

### 与结构体相关的方法

任意值通过`reflect.TypeOf()`获取反射的对象，如果是结构体，可以通过反射值对象`reflect.Type`的`NumField()`和`Field()`方法获得结构体的详细信息。

|  方法   | 说明  |
|  ----  | ----  |
| Field(i int) StructField | 根据索引，返回索引对应的结构体字段的信息 |
| NumField() int | 返回结构体成员字段数量 |
| FieldByName(name string) (StructField, bool) | 根据给定字符串返回字符串对应的结构体字段的信息 |
| FieldByIndex(index []int) StructField | 多层成员访问时，根据[]int根据的每个结构体的字段索引，返回字段的信息 |
| FieldByNameFunc(match func(string) bool) (StructField,bool) | 根据传入的匹配函数匹配需要的字段 |
| NumMethod() int | 返回该类型的方法集中方法的数目 |
| Method(int) Method | 返回该类型的方法集中方法的数目 |
| MethodByName(string) (Method, bool) | 根据方法名返回该类型方法集中的方法 |

```
type StructField struct {
    Name string  // 字段的名字
    PkgPath string  // 非导出字段的包路径，对导出的为空
    Type Type  // 字段的类型
    Tag StructTag  // 字段的标签
    Offset uintptr  // 字段在结构体中的字节偏移量
    Index []int  // 用于Type.FieldByIndex时的索引切片
    Anonymous bool  // 是否匿名字段
}

```

### 结构体反射类型

```go
package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int `json:"score"`
}

func main()  {
	stu1 := student{
		Name:  "jusene",
		Score: 80,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}

	field := t.FieldByIndex([]int{0})
	fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
}
```

```go
package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func (s student) Study() string {
	msg := "good study"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "good night"
	fmt.Println(msg)
	return msg
}

func printInterface(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	fmt.Println(v.NumMethod())
	fmt.Println(v.Method(0).Type())
	fmt.Println(t.Method(0).Type)
	fmt.Println(t.Method(0).Name)
	fmt.Println(t.MethodByName("Sleep"))
	fmt.Println(v.MethodByName("Sleep").IsNil())
	for i := 0; i < v.NumMethod(); i ++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
		v.MethodByName("Sleep").Call(args)
		v.MethodByName("Study").Call(args)
	}
}


func main() {
	printInterface(student{
		Name:  "jusene",
		Score: 27,
	})
}
```
