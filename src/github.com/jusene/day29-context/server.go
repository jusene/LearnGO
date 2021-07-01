package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	number := rand.Intn(2)
	if number == 0 {
		time.Sleep(10 * time.Second)
		fmt.Fprintf(w, "slow response")
		return
	}
	fmt.Fprintf(w, "quick response")
}

func main() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
