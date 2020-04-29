## Go语言网络编程

### 检查ip示例

```go
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "用法：%s IP地址\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("无效IP地址")
	} else {
		fmt.Println("IP地址是", addr.String())
	}
}
```

### Dial() 函数

```
func Dial(net, addr string) (Conn, err)
```

- TCP连接
```
conn, err := net.Dial("tcp", "127.0.0.1:2100")
```

- UDP连接
```
conn, err := net.Dial("udp", "127.0.0.1:2100")
```

- ICMP连接
```
conn, err := net.Dial("ip4:icmp", "127.0.0.1")
```
在建立连接后，就可以使用数据发送和接受，使用conn的Write()发送，Read()接受。

### TCP Socket

```
func (c *TCPConn) Write(b []byte) (n int, err os.Error)
func (c *TCPConn) Read(b []byte) (n int, err os.Error)
```

需要TCPAddr类型
```
type TCPAddr struct {
   IP IP
   Port int
   Zone string // ipv6范围选址
}

// 在go语言中通过resolvetcpaddr获取一个tcpaddr
func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error)
```

#### TCP客户端

```
func DailTCP(net string, laddr, raddr *TCPAddr)(c *TCPConn, err os.Error)
```
- net 参数tcp、tcp4、tcp6中任意一个
- laddr 表示本机地址，一般为nil
- raddr 表示远程地址

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误：%s", err.Error())
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "用法: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]
	TCPAddr, err := net.ResolveTCPAddr("tcp4", service) // 生成地址结构体
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, TCPAddr)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	fmt.Println(string(result))
}
```

#### TCP服务端

```
func ListenTCP(net string, laddr *TCPAddr) (l *TCPListener, err os.Error)
func (l *TCPListener) Accept() (c conn, err os.Error)
```

设置建立连接的超时时间，客户端和服务器端都适用，当超时设置时间后，连接自动关机
```
func DialTimeout(net, addr string, timeout time.Duration) (Conn, err)
```

设置写入/读取一个连接的超时时间，当超过设置时间时，连接自动关闭
```
func (c *TCPConn) SetReadDeadline(t time.Time) error
func (c *TCPConn) SetWriteDeadline(t time.Time) error
```

间隔性地发送keepalive包，操作系统可以通过该包来判断一个tcp连接是否已经断开
```
func (c *TCPConn) SetKeepAlive(keepalive bool) os.Error
```

```go
package main

import (
	"log"
	"net"
	"time"
)

// TCP 服务端

func echo(conn *net.TCPConn) {
	tick := time.Tick(5 * time.Second) // 5秒请求一次
	for now := range tick {
		n, err := conn.Write([]byte(now.String()))
		if err != nil {
			log.Println(err)
			conn.Close()
			return
		}
		log.Printf("send %d bytes to %s\n", n, conn.RemoteAddr())
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	address, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8000")
	checkError(err)
	listener, err := net.ListenTCP("tcp4", address)
	checkError(err)

	for {
		conn, err := listener.AcceptTCP()
		checkError(err)
		log.Println("远程地址：", conn.RemoteAddr())
		go echo(conn)
	}
}
```

```go
package main

// TCP 客户端

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "用法： %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	TCPAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, TCPAddr)
	checkError(err)

	conn.SetKeepAlive(true) // 会间隔性地发送keepalive包，操作系统可以通过该包来判断一个tcp连接是否已经断开

	for {
		data := make([]byte, 256)
		n, err := conn.Read(data)
		checkError(err)
		log.Println(strings.TrimSpace(string(data[0:n])))
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误：%s", err.Error())
		os.Exit(1)
	}
}
```

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// tcp 服务端

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [512]byte
		n, err := reader.Read(buf[0:])
		checkError(err)
		recvStr := string(buf[:n])
		fmt.Println("接受client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr))
	}
}

func checkError(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9000")
	checkError(err)

	for {
		conn, err := listen.Accept()
		checkError(err)
		go process(conn)
	}
}
```

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// tcp 客户端

func checkError(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	checkError(err)
	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		inputinfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputinfo) == "Q" {
			return
		}

		_, err = conn.Write([]byte(inputinfo)) // 发送数据
		checkError(err)

		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		checkError(err)

		fmt.Println(string(buf[:n]))
	}
}
```

- TCP黏包

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

// tcp 服务端

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [1024]byte
	for {
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed, err: ", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("接受client的数据: ", recvStr)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("listen failed, err: ", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err: ", err)
			continue
		}
		go process(conn)
	}
}
```

```go
package main

import (
	"fmt"
	"net"
)

// TCP 客户端

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("dial failed, err: ", err)
		return
	}

	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, hello`
		conn.Write([]byte(msg))
	}
}
```

```
接受client的数据:  Hello, helloHello, helloHello, helloHello, hello
接受client的数据:  Hello, helloHello, helloHello, helloHello, helloHello, hello
接受client的数据:  Hello, hello
接受client的数据:  Hello, hello
接受client的数据:  Hello, hello
接受client的数据:  Hello, hello
接受client的数据:  Hello, hello
接受client的数据:  Hello, hello
接受client的数据:  Hello, hello
接受client的数据:  Hello, hello
接受client的数据:  Hello, hello
接受client的数据:  Hello, hello
接受client的数据:  Hello, hello
```

- TCP 黏包解决

```go
package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// Encode 将消息编码
func Encode(message string) ([]byte, error) {
	// 读取消息的长度，转换成int32类型（占4个字符）
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}

	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// Decode 解码消息
func Decode(reader *bufio.Reader) (string, error) {
	// 读取消息长度
	lengthByte, _ := reader.Peek(4) // 读取前4个字节的数据
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}

	// Buffered 返回缓冲中现有的可读取的字节数
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	// 读取真正的数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
```

```go
package main

import (
	"bufio"
	"fmt"
	"github.com/jusene/day15/proto"
	"io"
	"net"
)

// TCP 服务端

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("收到client发来的数据：", msg)
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
```

```go
package main

import (
	"fmt"
	"github.com/jusene/day15/proto"
	"net"
)

// TCP 客户端

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}
}
```
### UDP Socket

```
func ResolveUDPAddr(net, addr string)(*UDPAddr, os.Error)
func DialUDP(net, string, laddr, raddr *UDPAddr)(c *UDPConn, err os.Error)
func ListenUDP(net string, laddr *UDPAddr)(c *UDPConn, err os.Error)
func (c *UDPConn) ReadFromUDP(b []byte)(n int, addr *UDPAddr, err os.Error)
func (c *UDPConn) WriteToUDP(b []byte)(n int, err os.Error)
```

```go
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

// UDP 服务端

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误", err.Error())
		os.Exit(2)
	}
}

func main() {
	service := "127.0.0.1:1200"
	udpaddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)

	conn, err := net.ListenUDP("udp4", udpaddr)
	checkError(err)

	for {
		buf := make([]byte, 512)
		n, addr, err := conn.ReadFromUDP(buf)
		checkError(err)
		fmt.Fprintf(os.Stdout, "接受%s\n", addr.IP)
		fmt.Fprintf(os.Stdout, "%s\n", string(buf[0:n]))
		daytime := time.Now().String()
		conn.WriteToUDP([]byte(daytime), addr)
	}
}
```

```go
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

// UDP 客户端

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误", err.Error())
		os.Exit(2)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "用法：%s host:port", os.Args[0])
		os.Exit(2)
	}

	service := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	tick := time.Tick(5 * time.Second)
	for now := range tick {
		_ = now
		conn.Write([]byte("anything"))
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		checkError(err)
		fmt.Fprintf(os.Stdout, "%s\n", string(buf[0:n]))
	}
}
```


