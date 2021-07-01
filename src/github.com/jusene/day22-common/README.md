## Go语言 数据结构

### xml

```xml
<?xml version="1.0" encoding="utf-8" ?>
<servers version="1">
    <server>
        <serverName>Local_web</serverName>
        <serverIP>192.168.66.100</serverIP>
    </server> 
    <server>
        <serverName>Local_DB</serverName>
        <serverIP>192.168.66.101</serverIP>
    </server>      
</servers>
```

- 解析xml
```go
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type RecurlyServers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs []server `xml:"server"`
	Description string `xml:",innerxml"`
}

type server struct {
	XMLName xml.Name `xml:"server"`
	ServerName string `xml:"serverName"`
	ServerIP string `xml:"serverIP"`
}

func main() {
	file, err := os.Open("src/github.com/jusene/day22/server.xml")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	v := RecurlyServers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v.Description)
}
```
- 生成xml
```go
package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs []server `xml:"server"`
}

type server struct {
	ServerName string `xml:"serverName"`
	ServerIP string `xml:"serverIP"`
}

func main() {
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{"Remote_WEB", "192.168.66.122"})
	v.Svs = append(v.Svs, server{"Remote_DB", "192.168.66.144"})
	output, err := xml.MarshalIndent(v, " ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}
```

### JSON

```json
{"servers": [
  {"serverName": "Local_WEB", "serverIP": "192.168.66.100"},
  {"serverName": "Local_DB", "serverIP": "192.168.66.101"}
]}
```

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Server struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}

type ServerSlice struct {
	Servers []Server
}

func main() {
	s := new(ServerSlice)
	file, err := os.Open("src/github.com/jusene/day22/server.json")
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	json.Unmarshal(buf[:n], s)
	fmt.Println(s.Servers[0].ServerIP)

}
```

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}

type ServerSlice struct {
	Servers []Server
}

func main() {
	s := new(ServerSlice)
	s.Servers = append(s.Servers, Server{
		ServerName: "LOCAL",
		ServerIP:   "192.168.66.100",
	})

	s.Servers = append(s.Servers, Server{
		ServerName: "REMOTE",
		ServerIP:   "192.168.66.101",
	})

	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
```

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
)

type MyDATA struct {
	Name  string  `json:"item"`
	Other float64 `json:"amount"`
}

func main() {
	s := new(MyDATA)
	s.Name = "jusene"
	s.Other = 0.23
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	js, err := simplejson.NewJson(b)
	if err != nil {
		panic(err)
	}
	fmt.Println(js.Get("item").String())
}
```

### 日志记录

```go
package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	log.SetOutput(os.Stdout)

	log.SetLevel(log.WarnLevel)
}

func main() {
	log.WithFields(log.Fields{
		"tool": "pen",
		"price": 10,
	}).Warn("This is a 10 dollars pen")

	log.WithFields(log.Fields{
		"tool": "pen",
		"price": 10,
	}).Info("This is a 10 dollars pen")

	contextLogger := log.WithFields(log.Fields{
		"common": "一个字段",
	})

	contextLogger.Fatal("TEST")
}
```

### 打包解包

- 打包
```go
package main

import (
	"archive/zip"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	// 创建一个打包文件
	outfile, err := os.Create("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	// 使用zip包创建函数zipWriter
	zipWriter := zip.NewWriter(outfile)

	// 往打包文件中写入文件
	var filesToArchive = []struct{
		name, body string
	} {
		{"test.txt", "String contents of file"},
		{"test2.txt", "\x61\x62\x63\n"},
	}

	for _, file := range filesToArchive {
		fileWriter, err := zipWriter.Create(file.name)
		if err != nil {
			log.Fatal(err)
		}

		_, err = fileWriter.Write([]byte(file.body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// 清理
	err = zipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}
}
```

- 解包
```go
package main

import (
	"archive/zip"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)

func main() {
	zipReader, err := zip.OpenReader("test.zip")
	if err != nil {
		panic(err)
	}
	defer zipReader.Close()

	// 遍历打包文件中的每一个文件和文件夹
	for _, file := range zipReader.Reader.File {
		zipFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zipFile.Close()

		targetDir := "./"
		extraFile := filepath.Join(targetDir, file.Name)

		if file.FileInfo().IsDir() {
			log.Println("正在创建目录: ", extraFile)
			os.MkdirAll(extraFile, file.Mode())
		} else {
			log.Println("正在提取文件: ", file.Name)

			outFile, err := os.OpenFile(
				extraFile,
				os.O_WRONLY | os.O_CREATE | os.O_TRUNC,
				file.Mode())
			if err != nil {
				panic(err)
			}

			defer outFile.Close()

			_, err = io.Copy(outFile, zipFile)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
```

### 压缩解压

- 压缩
```go
package main

import (
	"compress/gzip"
	"os"
)

func main() {
	outFile, err := os.Create("test.gz")
	if err != nil {
		panic(err)
	}

	gzipWriter := gzip.NewWriter(outFile)
	defer gzipWriter.Close()

	_, err = gzipWriter.Write([]byte("Gopher!\n"))
	if err != nil {
		panic(err)
	}
}
```

- 解压
```go
package main

import (
	"compress/gzip"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	gzipFile, err := os.Open("test.gz")
	if err != nil {
		log.Fatal(err)
	}

	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		panic(err)
	}
	defer gzipReader.Close()

	// 解压缩到一个writer，它是一个file writer
	outfileWriter, err := os.Create("unzipped.txt")
	if err != nil {
		panic(err)
	}
	defer outfileWriter.Close()

	// 复制内容
	_, err = io.Copy(outfileWriter, gzipReader)
	if err != nil {
		panic(err)
	}
}
```