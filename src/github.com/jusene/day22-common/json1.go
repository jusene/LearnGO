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
