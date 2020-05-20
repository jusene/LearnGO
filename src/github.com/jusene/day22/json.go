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
