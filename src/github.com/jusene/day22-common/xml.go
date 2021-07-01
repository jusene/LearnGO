package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type RecurlyServers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
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
