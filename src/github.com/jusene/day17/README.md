## Go语言 RPC

GO语言支持三个级别的RPC：TCP, HTTP, JSONRPC, Go语言的RPC与传统的RPC不同，只支持Go语言开发的服务器和客户端之间交互，因为内部采用Gob来编码。

Go RPC的函数只有符合下列条件才能被远程访问：
- 函数必须是导出（首字母大写）
- 必须有两个导出类型的参数
- 第一个是接受的参数，第二个是返回给客户端的参数，第二个参数必须是指针类型
- 函数还要有一个返回值error

```
func (t *T) MethodName(argType T1, replyType *T2) error
```

#### HTTP RPC

- server
```go
package main

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("除数不能为零")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	log.Fatal(http.ListenAndServe(":1234", nil))
}
```

- client
```go
package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	//serverAddress := os.Args[1]
	serverAddress := "127.0.0.1"

	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("Dial错误", err)
	}

	args := Args{17, 1}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith错误: ", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith 错误: ", err)
	}
	fmt.Printf("Arith: %d/%d=%d 余 %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
```

#### TCP RPC

- server
```go
package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("除数不能为零")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
```

- client
```go
package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("Dial 错误: ", err)
	}

	args := Args{
		A: 15,
		B: 8,
	}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith 错误: ", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith 错误: ", err)
	}
	fmt.Printf("Arith: %d/%d=%d 余 %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
```

#### JSON RPC

json rpc是数据编码采用json，而不是Gob

```go
package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("除数不能为零")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go jsonrpc.ServeConn(conn)
	}
}
```

- client
```go
package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}

	args := Args{15, 7}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith 错误: ", err)
	}
	fmt.Printf("Arith: %d/%d=%d 余 %d\n", args.A, args.B, quot.Quo, quot.Rem)

}
```