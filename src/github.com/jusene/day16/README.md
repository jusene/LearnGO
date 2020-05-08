## Go语言HTTP编程

Go语言标准库内置net/http包，涵盖HTTP客户端和服务端的具体实现。

### HTTP客户端

```
func (c *Client) Get(url string) (r *Response, err error)
func (c *Client) Post(url string, bodyType string, body io.Reader) (r *Response, err error)
func (c *Client) PostForm(url string, data url.Values) (r *Response, err error)
func (c *Client) Head(url string) (r *Response, err error)
func (c *Client) Do(req *Request) (resp *Response, err error)
```

#### Get

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:9090")
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	defer resp.Body.Close()
	// io.Copy(os.Stdout, resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read from resp.Body failed,err:", err)
		return
	}
	fmt.Println(string(body))
}
```

#### Get Params

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	data := url.Values{
		"name": []string{"jusene"},
		"age": []string{"27"},
	}
	data.Set("job", "enginer")
	u, err := url.ParseRequestURI("http://127.0.0.1:9090")
	if err != nil {
		fmt.Errorf("parse url requestUrl failed, err %v", err)
	}
	u.RawQuery = data.Encode()
	fmt.Printf("%s", u.String())

	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	defer resp.Body.Close()
	// io.Copy(os.Stdout, resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read from resp.Body failed,err:", err)
		return
	}
	fmt.Println(string(body))
}
```

#### Post

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "http://127.0.0.1:9090/post"

	// form
	 // contentType := "application/x-www-form-urlencoded"
	 // data := "name=jusene&age=27"
	// json
	contentType := "application/json"
	data := `{"name": "jusene", "age": 27}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed; err:%v\n", err)
		return
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))

}
```

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// 上传图片
func main() {
	data := url.Values{
		"name": []string{"jusene"},
		"age": []string{"27"},
	}
	resp, err := http.PostForm("http://127.0.0.1:9090", data)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
```

#### Head

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Head("https://www.baidu.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Header["Cache-Control"])
}
```

#### http.Client

```
type Client struct {
    // Transport 用于确定HTTP请求的创建机制
    // 如果为空，将会使用DefaultTransport
    Transport RoundTripper
    // CheckRedirect定义重定向策略
    // 如果CheckRedirect不为空，客户端将在跟踪HTTP重定向前调用该函数
    // 两个参数req和via分别为即将发起的请求和已经发起的所有请求，最早的已发送请求在最前面
    // 如果CheckRedirect返回错误，客户端将直接返回错误，不会再发起该请求
    // 如果CheckRedirect为空，Client将采用一种确认策略，将在10个连续请求后终止
    CheckRedirect func(req *Request, via []*Request) error
    // 如果jar为空，cookie将不会在请求中发送，并会在响应中被忽略
    Jar CookieJar
}
```