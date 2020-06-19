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

修改相应头信息
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