## Go语言 gin框架

### 快速开始

第一个实例：
```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
        // 创建一个默认的路由引擎
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
	// 返回JSON格式数据
        context.JSON(200, gin.H{
			"message": "pong",
		})
	})
    // 启动HTTP服务，默认0.0.0.0:8080启动服务
	r.Run()
}
```

### RESTful API

#### 路径参数

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func gettinng(ctx *gin.Context) {
	ctx.String(http.StatusOK, "GET")
}

func main() {
	router := gin.Default()
	// http://127.0.0.1:8080/get GET
	router.GET("/get", gettinng)
	// http://127.0.0.1:8080/get/jusene HELLO, jusene
	router.GET("/get/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "HELLO, %s", name)
	})
	// http://127.0.0.1:8080/get/jusene/29 {"age":"29","code":200,"name":"jusene"}
	router.GET("/get/:name/:age", func(context *gin.Context) {
		name := context.Param("name")
		age := context.Param("age")
		context.JSON(200, gin.H{
			"name": name,
			"age": age,
			"code": 200,
		})
	})
	// http://127.0.0.1:8080/get/jusene/29/send   jusene is /send
	router.GET("/get/:name/:age/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		message := name + " is " + action
		context.String(http.StatusOK, message)
	})
	// http://127.0.0.1:8080/post/jusene/29/send /post/:name/:age/*action
	router.POST("/post/:name/:age/*action", func(context *gin.Context) {
		context.String(http.StatusOK, context.FullPath())
	})
	router.Run()
}
```

#### 参数

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	// http://127.0.0.1:8080/welcome?firstname=zhang&lastname=jusene  hello zhang jusene
	router.GET("/welcome", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname", "Guest") // 默认query param
		lastname := context.Query("lastname")

		context.String(http.StatusOK, "hello %s %s", firstname, lastname)
	})

	router.Run()
}
```

#### 表单

```go
package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	// curl -d "message=you are work" http://127.0.0.1:8080/form  {"message":"you are work","nick":"jusene","status":"post"}
	router.POST("/form", func(context *gin.Context) {
		message := context.PostForm("message")
		nick := context.DefaultPostForm("nick", "jusene")

		context.JSON(200, gin.H{
			"status": "post",
			"message": message,
			"nick": nick,
		})
	})
	router.Run()
}
```

#### query + post form
```go
package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	// curl -d "name=jusene" http://127.0.0.1:8080/post?id=1 {"id":"1","message":"ok","name":"jusene","page":"0"}
	router.POST("/post", func(context *gin.Context) {
		id := context.Query("id")
		page := context.DefaultQuery("page", "0")
		name := context.PostForm("name")
		message := context.DefaultPostForm("message", "ok")

		context.JSON(200, gin.H{
			"id": id,
			"page": page,
			"name": name,
			"message": message,
		})
	})

	router.Run()
}
```

#### Map as querystring

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// curl -X POST -d "names[c]=jusne" http://127.0.0.1:8080/post?ids[a]=123&ids[b]=456  {"ids":{"a":"123","b":"456"},"name":{"c":"jusene"}}
	router.POST("/post", func(context *gin.Context) {
		ids := context.QueryMap("ids")
		name := context.PostFormMap("names")

		context.JSON(200, gin.H{
			"ids": ids,
			"name": name,
		})
	})

	router.Run()
}
```

#### 上传文件

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	// set a lower memory limit for multipart forms(default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8MiB
        // curl -X POST http://127.0.0.1:8080/upload -F "upload[]=@gin1.go" -H "Content-Typ e: multipart/form-data"
	router.POST("/upload", func(context *gin.Context) {
		// Multipart form
		form, _ := context.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			context.SaveUploadedFile(file, "./test.go")
		}
		context.String(http.StatusOK, fmt.Sprintf("%d files uploads!", len(files)))
	})
	router.Run()
}
```

#### 分组路由

```go
package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// GROUP v1
	v1 := router.Group("/api/v1")
	{
        // curl http://127.0.0.1:8080/api/v1/get GET
		v1.GET("/get", func(context *gin.Context) {
			context.String(200, "GET")
		})
        // curl -X POST http://127.0.0.1:8080/api/v1/post POST
		v1.POST("/post", func(context *gin.Context) {
			context.String(200, "POST")
		})
	}
	router.Run()
}
```

### 修改默认中间件

写入日志文件
```go
package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	gin.DisableConsoleColor()

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}
```

改变日志格式
```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}
```

修改响应头信息
```go
package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.New()

	router.Use(Cors)
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.Run()
}

func Cors(ctx *gin.Context) {
	ctx.Header("BY-X-SERVER", "GIN")
	ctx.Next()
}

curl  -vv http://127.0.0.1:8080/ping
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> GET /ping HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.55.1
> Accept: */*
>
< HTTP/1.1 200 OK
< By-X-Server: GIN
< Content-Type: text/plain; charset=utf-8
< Date: Fri, 19 Jun 2020 15:35:37 GMT
< Content-Length: 4
<
pong* Connection #0 to host 127.0.0.1 left intact
```

### model绑定和验证

要将请求主体主体绑定类型，使用model绑定，目前支持JSON,XML,YAML和标准表单值得绑定。

需要在绑定的所有字段上设置相应的绑定标签。

GIN提供了两种绑定方法:
- Type - 必须绑定
- - 方法: Bind,BindJSON,BindXML,BindQuery,BindYAML,BindHeader
- - 行为: 这些方法在后台使用`MustBindWith`,如果存在绑定错误，则使用终止请求`c.AbortWithError(400, err).SetType(ErrorTypeBind)`。这会将响应的状态码设置为400，并将`Content-Type`设置为`text/plain; charset=utf-8`

- Type - 应该绑定
- - 方法: ShouldBind,ShouldBindJSON,ShouldBindXML,ShouldBindQuery,ShouldBindYAML,ShouldBindHeader
- - 行为: 这些方法`ShouldBindWith`在后台使用，存在绑定错误，则将返回错误，开发人员应适当处理请求和错误。

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	// Example for binding JSON ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(context *gin.Context) {
		var json Login
		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if json.User != "jusene" || json.Password != "123" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": "unauthorized",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status": "you are login in",
		})
	})

	// Example for binding XML (
	//	<?xml version="1.0" encoding="UTF-8"?>
	//	<root>
	//		<user>user</user>
	//		<password>123</password>
	//	</root>)
	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if xml.User != "jusene" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Example for binding a HTML form (user=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// This will infer what binder to use depending on the content-type header.
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "jusene" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
```

跳过验证，将`binding: "required"`修改为`binding: "-"`,将不会返回error

### 通常验证

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func main() {
	route := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	route.GET("/bookable", getBookable)
	route.Run()
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindQuery(&b); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
```

绑定字符查询
```go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Person struct {
	Name string `form:"name"`
	Address string `form:"address"`
}

func main() {
	route := gin.Default()
	route.Any("/testing", startPage)
	route.Run()
}

func startPage(ctx *gin.Context) {
	var person Person
	if ctx.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	ctx.String(200, "Success")
}
```

绑定字符查询或者post数据
```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type Person1 struct {
	Name string `form:"name"`
	Address string `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	CreateTime time.Time `form:"createTime" time_format:"unixNano"`
	UnixTime time.Time `form:"unixTime" time_format:"unix"`
}

func main() {
	route := gin.Default()
	route.GET("/testing", startPage1)
	route.Run()
}

func startPage1(ctx *gin.Context) {
	var person Person1
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	fmt.Println(ctx.ShouldBindQuery(&person))
	if ctx.ShouldBindQuery(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
		log.Println(person.CreateTime)
		log.Println(person.UnixTime)
	}
	ctx.String(200, "Success")
}
```

绑定uri
```go
package main

import "github.com/gin-gonic/gin"

type Person struct {
	ID string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	route := gin.Default()
	route.GET("/:name/:id", func(context *gin.Context) {
		var person Person
		if err := context.ShouldBindUri(&person); err != nil {
			context.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		context.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})
	route.Run()
}
```
- curl -v localhost:8080/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3
- curl -v localhost:8080/thinkerou/not-uuid // error

绑定头部
```go
package main

import "github.com/gin-gonic/gin"

type Header struct {
	Rate int `header:"Rate"`
	Domain string `header:"Domain"`
}

func main() {
	r := gin.Default()

    // curl -H "rate:300" -H "domain:http" 127.0.0.1:8080 {"Domain":"http","Rate":300}
	r.GET("/", func(context *gin.Context) {
		H := Header{}
		if err := context.ShouldBindHeader(&H); err != nil {
			context.JSON(200, err)
		}

		context.JSON(200, gin.H{"Rate": H.Rate, "Domain": H.Domain})
	})

	r.Run()
}
```

绑定html checkbox
```go
package main

import "github.com/gin-gonic/gin"

type myForm struct {
	Colors []string `form:"colors[]"`
}

func main() {
	r := gin.Default()

	r.POST("/", func(context *gin.Context) {
		var fakeForm myForm
		context.ShouldBind(&fakeForm)
		context.JSON(200, gin.H{
			"color": fakeForm.Colors,
		})
	})
	r.Run()
}
```
```html
<html>
<form action="http://127.0.0.1:8080/" method="post">
<p>Check some colors</p>
    <label for="red">Red</label>
    <input type="checkbox" name="colors[]" value="red" id="red">
    <label for="green">Green</label>
    <input type="checkbox" name="colors[]" value="green" id="green">
    <label for="blue">Blue</label>
    <input type="checkbox" name="colors[]" value="blue" id="blue">
    <input type="submit">
</form>
</html>
```

Multipart/Urlencoded binding

```go
package main

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

type ProfileForm struct {
	Name   string                `form:"name" binding:"required"`
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`

	// or for multiple files
	// Avatars []*multipart.FileHeader `form:"avatar" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/profile", func(c *gin.Context) {
		// you can bind multipart form with explicit binding declaration:
		// c.ShouldBindWith(&form, binding.Form)
		// or you can simply use autobinding with ShouldBind method:
		var form ProfileForm
		// in this case proper binding will be automatically selected
		if err := c.ShouldBind(&form); err != nil {
			c.String(http.StatusBadRequest, "bad request")
			return
		}

		err := c.SaveUploadedFile(form.Avatar, form.Avatar.Filename)
		if err != nil {
			c.String(http.StatusInternalServerError, "unknown error")
			return
		}

		// db.Save(&form)

		c.String(http.StatusOK, "ok")
	})
	router.Run(":8080")
}
```

### XML, JSON, YAML and ProtoBuf rendering

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/someJSON", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "key", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(context *gin.Context) {
		var msg struct{
			Name string `json:"user"`
			Message string
			Number int
		}

		msg.Name = "jusnee"
		msg.Message = "hey"
		msg.Number = 123

		context.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// The specific definition of protobuf is written in the testdata/protoexample file.
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// Note that data becomes binary data in the response
		// Will output protoexample.Test protobuf serialized data
		c.ProtoBuf(http.StatusOK, data)
	})

	r.Run()
}
```

### SecureJSON

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/someJSON", func(context *gin.Context) {
		names := []string{"name", "ga", "age"}

		context.SecureJSON(http.StatusOK, names)
	})
	// while(1);["name","ga","age"]
	r.Run()
}
```

### JSONP

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/JSONP", func(c *gin.Context) {
		data := gin.H{
			"foo": "bar",
		}

		//callback is x
		// Will output  :   x({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")

	// client
	// curl http://127.0.0.1:8080/JSONP?callback=x
	//callback is x
	// Will output  :   x({\"foo\":\"bar\"})
}
```

### AsciiJSON

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		data := gin.H{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
```

### PureJSON

```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// Serves unicode entities
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// Serves literal characters
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
```

### 静态文件服务

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Static("/mod", "./")
	r.StaticFS("/more", http.Dir("./"))
	r.StaticFile("/favicon.ico", "./favicon.ico")

	r.Run()
}
```

### 文件数据服务

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/local/file", func(context *gin.Context) {
		context.File("./gin.log")
	})

	var fs http.FileSystem 
	router.GET("/fs/file", func(context *gin.Context) {
		context.FileFromFS("gin1.go", fs)
	})

	router.Run()
}
```

### Serving data from reader

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
	router.Run(":8080")
}
```

## html 渲染

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	// r.LoadHTMLFiles("templates/template1.html")
	r.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
			"body": "hello world",
		})
	})

	r.Run()
}
```

```html
<html>
	<h1>
		{{ .title }}
	</h1>

	<p> {{ .body }}</p>
</html>
```

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*/*")
	r.GET("/posts/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/index.tmpl", gin.H{
			"title": "Users",
		})
	})
	r.Run(":8080")
}
```

```html
{{ define "posts/index.tmpl" }}
<html><h1>
	{{ .title }}
</h1>
<p>Using posts/index.tmpl</p>
</html>
{{ end }}
```

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
	r := gin.Default()
	// r.Delims("{[{", "}]}")
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	r.LoadHTMLFiles("templates/raw.tmpl")

	r.GET("/raw", func(context *gin.Context) {
		context.HTML(http.StatusOK, "raw.tmpl", gin.H{
			"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run()
}
```

```html
Date: {{.now | formatAsDate }}
```

### 重定向

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/redic", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	r.GET("/test", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/redic")
	})

	r.GET("/test1", func(context *gin.Context) {
		context.Request.URL.Path = "/test2"
		r.HandleContext(context)
	})

	r.GET("/test2", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})
	r.Run()
}
```

### 通常中间件

```go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin.New()
	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// it would print: "12345"
		log.Println(example)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
```

### basic auth 中间件

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var secrets = gin.H{
	"foo": gin.H{"email": "foo@bar.com", "phone": "123433"},
}

func main() {
	r := gin.Default()

	authorizd := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo": "bar",
	}))

	authorizd.GET("/secrets", func(context *gin.Context) {
		user := context.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			context.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			context.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET"})
		}
	})

	r.Run()
}
```

### 中间件内goroutines

```go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/long_async", func(context *gin.Context) {
		cCp := context.Copy()
		go func() {
			time.Sleep(5 * time.Second)

			log.Println("Done! "+ cCp.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(context *gin.Context) {
		time.Sleep(5 * time.Second)

		log.Println("Done! "+ context.Request.URL.Path)
	})

	r.Run()
}
```

### gin中启动多服务

```go
package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

var (
	g errgroup.Group
)

func router1() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "Welcome server 1",
		})
	})
	return e
}

func router2() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "Welcome server 2",
		})
	})
	return e
}

func main() {
	server1 := &http.Server{
		Addr:         ":8080",
		Handler:      router1(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server2 := &http.Server{
		Addr:         ":8081",
		Handler:      router2(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		err := server1.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	g.Go(func() error {
		err := server2.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
```

### 正常关闭和重启

```go
package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		time.Sleep(5 * time.Second)
		context.String(http.StatusOK, "Welcome gin server")
	})

	srv := &http.Server{
		Addr: ":8080",
		Handler: r,
	}

	// Initializing the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待信号
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 syscall.SIGINI
	// kill -9 syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<- quit
	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exiting")
}
```

### 构建包含模板的二进制文件

```go
package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/html/index.tmpl", gin.H{
			"Foo": "World",
		})
	})
	r.GET("/bar", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/html/bar.tmpl", gin.H{
			"Bar": "World",
		})
	})
	r.Run(":8080")
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		fmt.Println(name, file)
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
```
```bash
go-assets-builder.exe .\html -o assets.go
```
```go
package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsbfa8d115ce0617d89507412d5393a462f8e9b003 = "<!doctype html>\r\n<body>\r\n  <p>Can you see this? → {{.Bar}}</p>\r\n</body>"
var _Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2 = "<!doctype html>\r\n<body>\r\n  <p>Hello, {{.Foo}}</p>\r\n</body>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{".\\html"}, "/html": []string{"bar.tmpl", "index.tmpl"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1592791609, 1592791609649929800),
		Data:     nil,
	}, "/html": &assets.File{
		Path:     "/html",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1592791669, 1592791669877880000),
		Data:     nil,
	}, "/html/bar.tmpl": &assets.File{
		Path:     "/html/bar.tmpl",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1592791669, 1592791669876892100),
		Data:     []byte(_Assetsbfa8d115ce0617d89507412d5393a462f8e9b003),
	}, "/html/index.tmpl": &assets.File{
		Path:     "/html/index.tmpl",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1592791641, 1592791641333471000),
		Data:     []byte(_Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2),
	}}, "")
```

### http2.0 server push

```go
package main

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html)

	r.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// use pusher.Push() to do server push
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})

	// Listen and Server in https://127.0.0.1:8080
	r.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}
```

### cookie

```go
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/cookie", func(c *gin.Context) {

		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}

		fmt.Printf("Cookie value: %s \n", cookie)
	})

	router.Run()
}
```