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

#### Parameters in path

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

#### Querystring parameters

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

#### Multipart/Urlencoded Form

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