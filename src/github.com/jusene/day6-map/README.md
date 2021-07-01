## Go语言复合数据类型

### 映射（map）

Go语言中提供的映射关系容器为map，其内部使用散列表（hash）实现。

map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用

#### map定义

```
map[KeyType]ValueType
```

map类型的变量默认初始值为nil，需要使用make()函数来分配内存

```
make(map[KeyType]ValueType, [cap])
```

#### 映射的创建

```go
package main

import "fmt"

func main() {
	dictMap := make(map[string]int)
	dictMap["张"] = 10
	dictMap["三"] = 20
	fmt.Println(dictMap)
	fmt.Println(dictMap["张"])

	dictMap1 := map[string]string{
		"username": "jusene",
		"password": "123rt",
	}
	fmt.Println(dictMap1["username"])

	dictMap2 := map[string][]int{}
	dictMap2["a"] = []int{1, 2, 3, 4}
	fmt.Println(dictMap2)
	fmt.Println(dictMap2["a"])

	// 不可以使用切片，函数以及包含切片结构类型由于具有引用语义，均不能作为映射的键
	// dictMap3 := map[[]string]int{}

	// value为map的切片
	dictMap4 := make([]map[string]int, 3)
	dictMap4[0] = make(map[string]int, 10)
	dictMap4[0]["name"] = 1
	dictMap4[0]["pass"] = 0
	fmt.Println(dictMap4)
}
```

#### 判断某个键是否存在

```
value , ok := map[key]
```

```go
package main

import "fmt"

func main() {
	dictMap := make(map[string]int)
	dictMap["张三"] = 90
	dictMap["小明"] = 100

	v, ok := dictMap["张三"]
	if ok {
		fmt.Println(v)
	}else {
		fmt.Println("没这个人")
	}
}
```

#### 映射遍历

```go
package main

import "fmt"

func main() {
	dictMap := make(map[string]int)
	dictMap["jusene"] = 90
	dictMap["zgx"] = 200

	for k, v := range dictMap {
		fmt.Println(k, v)
	}
}
```

`遍历map时的元素顺序与添加键值对的顺序无关。`

#### 删除键对值

```
delete(map, key)
```

```go
package main

import "fmt"

func main() {
	dictMap := make(map[string]int)
	dictMap["jusene"] = 90
	dictMap["zgx"] = 100

	delete(dictMap, "jusene")
	fmt.Println(dictMap)
}
```

#### 指定顺序遍历map

```go
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())
	
	dictMap  := make(map[string]int, 200)
	for i := 0; i< 100; i++ {
		key := fmt.Sprintf("stu%02d", i) // 生成stu开头的字符串
		value := rand.Intn(100)

		dictMap[key] = value
	}

	keys := make([]string, 0, 200)
	for key := range dictMap {
		keys = append(keys, key)
	}

	// 对切片进行排序
	sort.Strings(keys)
	// 按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, dictMap[key])
	}
}
```









