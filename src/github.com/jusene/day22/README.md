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

