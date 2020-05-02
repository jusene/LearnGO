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
