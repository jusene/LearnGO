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

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("User-Agent", "test agent")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	d, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(d))
}
```

自定义transport

```go
type Transport struct {
    // Proxy 指定用于针对特定请求返回代理的函数
    // 如果该函数返回一个非空的错误，请求将终止并返回该错误
    // 如果Proxy为空或者返回一个空的URL指针，将不使用代理
    Proxy func(*Request) (*url.URL, error)
    // Dial 指定用于创建TCP连接的dial函数
    // 如果Dial为空，将默认使用net.Dial()函数
    Dial func(net, addr string) (c net.Conn, err error)
    // TLSClientConfig指定用于tls.Client的TLS配置
    // 如果为空则使用默认配置
    TLSConfig *tls.Config
    DisableKeepAlives bool
    DisableCompression bool
    // 如果MaxIdleConnsPerHost为非零值，它用于控制每个host所需要保持的最大空闲连接数
    // 如果该值为空，则使用DefaultMaxIdleConnPerHost
    MaxIdleConnsPerHost int
}
```

### HTTP服务端

#### GET

```go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(os.Stdout, r.RemoteAddr)
	fmt.Fprintln(w, "hello jusene")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err%v\n", err)
		return
	}
}
```

#### GET Param

```go
package main

import (
	"fmt"
	"net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))
	fmt.Println(data.Get("job"))
	ans := `{"status": "ok"}`
	w.Write([]byte(ans))
}

func main() {
	http.HandleFunc("/", getHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err%v\n", err)
		return
	}
}
```

#### POST

```go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// form
	r.ParseForm()
	fmt.Println(r.PostForm)
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
	// json
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read request.Body failed")
	}
	type info struct {
		Name string
		Age int
	}
	person := new(info)
	err1 := json.Unmarshal(b, person)
	fmt.Println(string(b))
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(1111, person.Name, person.Age)
	answ := `{"status": "ok"}`
	w.Write([]byte(answ))
}

func main() {
	http.HandleFunc("/", postHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err%v\n", err)
		return
	}
}
```

更多控制服务器端的行为
```
s := &http.Server{
    Addr: ":8080",
    Handler: myHandler,
    ReadTimeout: 10 * time.Second,
    WriteTimeout: 10 * time.Second,
    MaxHeaderBytes: 1 << 20,
}

log.Fatal(s.ListenAndServe())
```

#### HTTPS

```
func ListenAndServeTLS(addr string, certFile string, keyFile string, handler Handler) error
```

```
http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))	
})

log.Fatal(http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil))
```

```
ss := &http.Server{
    Addr: ":10443",
    Handler: myHandler,
    ReadTimeout: 10 * time.Second,
    WriteTimeout: 10 * time.Second,
    MaxHeaderBytes: 1 << 20,
}

log.Fatal(ss.ListenAndServeTLS("cert.pem", "key.pem"))
```

#### Cookie

```
http.SetCookie(w ResponseWriter, cookie *Cookie)
```

```
type Cookie struct {
    Name string
    Value string
    Path string
    Domain string
    Expires time.Time
    RawExpires string

    // MaxAge=0 意味着没有指定Max-Age值
    // MaxAge<0 意味着现在就删除Cookie，等价于Max-Age=0
    // MaxAge>0 意味着Max-Age属性存在并以秒为单位存在

    MaxAge int
    Secure bool
    HttpOnly bool
    Raw string
    Unparsed []string // 未解析attribute-value属性值对
}
```

```
expiration := time.Time()
expiration = expiration.AddDate(1, 0, 0)
cookie := http.Cookie{Name: "username", Value: "test", Expires: expiration}
http.SetCookie(w, &cookie)
```

读取Cookie
```
cookie, _ := r.Cookie("username")
fmt.Fprint(w, cookie)
```
