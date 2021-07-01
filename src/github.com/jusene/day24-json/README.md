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

Go语言内置的 json 包使用 RFC3339 标准中定义的时间格式，对我们序列化时间字段的时候有很多限制

```go
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Post struct {
	CreateTime time.Time `json:"create_time"`
}

func main() {
	p1 := Post{CreateTime: time.Now()}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json.Marshal p1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	//jsonStr := `{"create_time":"2020-07-10T09:57:46.0743201+08:00"}`
	jsonStr := `{"create_time":"2020-07-10 09:57:46"}`
	var p2 Post
	if err := json.Unmarshal([]byte(jsonStr), &p2); err != nil {
		fmt.Printf("json.Unmarshl failed, err:%v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}
```

实现`json.Marshler`/`json.Unmarshaler`接口实现自定义的事件格式解析.

```go
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Order struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	CreatedTime time.Time `json:"created_time"`
}

const layout = "2006-01-02 15:04:05"

// MarshalJSON 为Order类型实现自定义的MarshalJSON方法
func (o *Order) MarshalJSON() ([]byte, error) {
	type TempOrder Order // 定义与Order字段一致的新类型
	return json.Marshal(struct {
		CreatedTime string `json:"created_time"`
		*TempOrder         // 避免直接嵌套Order进入死循环
	}{
		CreatedTime: o.CreatedTime.Format(layout),
		TempOrder:   (*TempOrder)(o),
	})
}

// UnmarshalJSON 为Order类型实现自定义的UnmarshalJSON方法
func (o *Order) UnmarshalJSON(data []byte) error {
	type TempOrder Order // 定义与Order字段一致的新类型
	ot := struct {
		CreatedTime string `json:"created_time"`
		*TempOrder         // 避免直接嵌套Order进入死循环
	}{
		TempOrder: (*TempOrder)(o),
	}
	if err := json.Unmarshal(data, &ot); err != nil {
		return err
	}
	var err error
	o.CreatedTime, err = time.Parse(layout, ot.CreatedTime)
	if err != nil {
		return err
	}
	return nil
}

// 自定义序列化方法
func main() {
	o1 := Order{
		ID:          123456,
		Title:       "《七米的Go学习笔记》",
		CreatedTime: time.Now(),
	}
	// 通过自定义的MarshalJSON方法实现struct -> json string
	b, err := json.Marshal(&o1)
	if err != nil {
		fmt.Printf("json.Marshal o1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	// 通过自定义的UnmarshalJSON方法实现json string -> struct
	jsonStr := `{"created_time":"2020-04-05 10:18:20","id":123456,"title":"《七米的Go学习笔记》"}`
	var o2 Order
	if err := json.Unmarshal([]byte(jsonStr), &o2); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("o2:%#v\n", o2)
}
```

### 使用匿名结构体添加字段

```go
package main

import (
	"encoding/json"
	"fmt"
)

type UserInfo struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

func main() {
	u1 := UserInfo{
		ID:   123456,
		Name: "JUSENE",
	}

	b, err := json.Marshal(struct {
		*UserInfo
		Token string `json:"token"`
	}{
		&u1,
		"1212121212",
	})
	if err != nil {
		fmt.Printf("json.Marsha failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
}
```

### 使用匿名结构体组合多个结构体

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Comment struct {
	Content string
}

type Image struct {
	Title string `json:"title"`
	URL string `json:"url"`
}

func main() {
	c1 := Comment{Content: "永远保持谦逊"}
	i1 := Image{
		Title: "jusnee",
		URL:   "http://www.baidu.com",
	}

	// struct -> json string
	b, err := json.Marshal(struct {
		*Comment
		*Image
	}{
		&c1,
		&i1,
	})
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)

	// json string -> struct
	jsonStr := `{"Content":"永远保持谦逊","title":"jusnee","url":"http://www.baidu.com"}`
	var (
		c2 Comment
		i2 Image
	)

	if err := json.Unmarshal([]byte(jsonStr), &struct {
		*Comment
		*Image
	}{&c2, &i2}); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("c2:%#v i2:%#v\n", c2, i2)
}
```

### 处理不确定层级的json

如果json串没有固定的格式导致不好定义与其相对应的结构体时，我们可以使用json.RawMessage原始字节数据保存下来。

```go
package main

import (
	"encoding/json"
	"fmt"
)

type sendMsg struct {
	User string `json:"user"`
	Msg string `json:"msg"`
}

func main() {
	jsonStr := `{"sendMsg":{"user":"q1mi","msg":"永远不要高估自己"},"say":"Hello"}`

	// 定义一个map，value类型为json.RawMessage，方便后续更灵活地处理
	var data map[string]json.RawMessage

	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		fmt.Printf("json.Unmarshal jsonStr failed, err:%v\n", err)
		return
	}
	fmt.Printf("data:%#v\n", data)
	var msg sendMsg
	if err := json.Unmarshal(data["sendMsg"], &msg); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("msg:%#v\n", msg)
}
```