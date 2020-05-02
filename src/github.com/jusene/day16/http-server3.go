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
