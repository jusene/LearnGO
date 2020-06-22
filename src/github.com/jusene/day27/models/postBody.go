package models

type User struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Hobbys []Hobby `json:"hobbys,omitempty"`
}

type Hobby struct {
	Name string `json:"name"`
	year int    `json:"year"`
}

type File struct {
	Msg      string `json:"msg"`
	Filename string `json:"filename"`
}
