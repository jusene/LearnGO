package main

import (
	"fmt"
	"net/http"
	"os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(os.Stdout, r.RemoteAddr)
	fmt.Fprintln(w, "hello jusene")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err%v\n", err)
		return
	}
}
