## Go语音 JSON技巧

### 基本序列化

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string
	Age int64
	Weight float64
}

func main() {
	p1 := Person{
		Name:   "jusene",
		Age:    18,
		Weight: 71.5,
	}

	// struct -> json string
	b, err := json.Marshal(p1)
	if err != nil {
		log.Fatalf("json marshal failed, err:%v\n", err)
	}
	fmt.Printf("str:%s\n", b)

	// json string -> struct
	var p2 Person
	err = json.Unmarshal(b, &p2)
	if err != nil {
		log.Fatalf("json unmarshal failed, err:%v\n", err)
	}
	fmt.Printf("p2:%v\n", p2)
}
```

### 结构体tag介绍

`tag`是结构体的元数据，可以在运行的时候通过反射的机制读取出来。 。

```go
// 使用json tag指定序列化与反序列化时的行为
type Person struct {
	Name   string `json:"name"` // 指定json序列化/反序列化时使用小写name
	Age    int64
	Weight float64
}
```
```go
// 使用json tag指定json序列化与反序列化时的行为
type Person struct {
	Name   string `json:"name"` // 指定json序列化/反序列化时使用小写name
	Age    int64
	Weight float64 `json:"-"` // 指定json序列化/反序列化时忽略此字段
}
```

- 忽略空值字段
```go
package main

import (
	"encoding/json"
	"fmt"
)

type User1 struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Hobby []string `json:"hobby"`
}

// 在tag中添加omitempty忽略空值
type User2 struct {
	Name string `json:"name"`
	Email string `json:"email,omitempty"`
	Hobby []string `json:"hobby,omitempty"`
}

func main() {
	u1 := User1{
		Name:  "jusene",
	}

	u2 := User2{
		Name:  "jusene",
	}

	b1, err := json.Marshal(u1)
	if err != nil {
		panic(err)
	}
	b2, err := json.Marshal(u2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b1))
	fmt.Println(string(b2))
}
```

### 忽略嵌套结构体空值字段

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Email string `json:"email,omitempty"`
	Hobby []string `json:"hobby,omitempty"`
	*Profile `json:"profile,omitempty"`
}

type Profile struct {
	Website string `json:"site"`
	Blog string `json:"blog"`
}

func main() {
	u1 := User{
		Name:    "jusene",
	}

	b1, err := json.Marshal(u1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b1))
}
```