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

### 不修改原结构体忽略空值字段

需要序列化User，但是不想把密码也序列化，又不想修改User结构体，这个时候我们就可以使用创建另外一个结构体PublicUser匿名嵌套原User，同时指定Password字段为匿名结构体指针类型，并添加omitemptytag，

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name string `json:"name"`
	Password string `json:"password"`
}

type PublicUser struct {
	*User //匿名嵌套
	Password *struct{} `json:"password,omitempty"`
}

func main() {
	u1 := User{
		Name:     "jusene",
		Password: "1234567",
	}

	b, err := json.Marshal(PublicUser{
		User:     &u1,
	})
	if err != nil {
		log.Fatalf("json.Marshal u1 failed, err:%v\n", err)
	}

	fmt.Printf("str:%s\n", b)
}
```

### 优雅处理字符串格式的数字

前端在传递来的json数据中可能会使用字符串类型的数字，这个时候可以在结构体tag中添加string来告诉json包从字符串中解析相应字段的数据

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Card struct {
	ID int64 `json:"id,string"`  // 添加string tag
	Score float64 `json:"score,string"`
}

func main() {
	jsonStr1 := `{"id": "123456", "score": "98.5"}`
	var c1 Card
	if err := json.Unmarshal([]byte(jsonStr1), &c1); err != nil {
		log.Fatalf("json.Unmarsha jsonStr1 failed, err:%v\n", err)
	}
	fmt.Printf("c1:%#v\n", c1)
}
```

### 整数变浮点数

在 JSON 协议中是没有整型和浮点型之分的，它们统称为number。json字符串中的数字经过Go语言中的json包反序列化之后都会成为float64类型。

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	// map[string]interface{} -> json string
	var m = make(map[string]interface{}, 1)
	m["count"] = 1 // int
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
	}
	fmt.Printf("string:%v\n", string(b))

	// json string -> map[string]interface{}
	var m2 map[string]interface{}
	err = json.Unmarshal(b, &m2)
	if err != nil {
		fmt.Printf("unmarshl failed, err:%v\n", err)
	}
	fmt.Printf("value:%v\n", m2["count"]) // 1
	fmt.Printf("value:%T\n", m2["count"]) // float64

	// 整型变成了浮点型，解决
	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.UseNumber()
	err = decoder.Decode(&m2)
	if err != nil {
		fmt.Printf("unmarshl failed, err:%v\n", err)
	}
	fmt.Printf("value:%v\n", m2["count"]) // 1
	fmt.Printf("value:%T\n", m2["count"]) // json.Number
	// 将m2["count"]装换为json.Number之后调用Int64()方法获得int64类型的值
	count, err := m2["count"].(json.Number).Int64()
	if err != nil {
		fmt.Printf("parse to int64 failed, err:%v\n", err)
		return
	}
	fmt.Printf("value:%v\n", count) // 1
	fmt.Printf("value:%T\n", count) // int64
}
```

### 自定义解析时间字段

