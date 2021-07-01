package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// 带有超时机制的client
type respData struct {
	resp *http.Response
	err  error
}

func doCall(ctx context.Context) {
	transport := http.Transport{
		DisableKeepAlives: true,
	}

	client := http.Client{
		Transport: &transport,
	}

	respChan := make(chan *respData, 1)
	req, err := http.NewRequestWithContext(ctx, "GET", "http://127.0.0.1:8000", nil)
	if err != nil {
		fmt.Printf("new request failed, err: %v\n", err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go func() {
		resp, err := client.Do(req)
		rd := &respData{
			resp: resp,
			err:  err,
		}
		respChan <- rd
		wg.Done()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("call api timeout")
	case result := <-respChan:
		fmt.Println("call api success")
		if result.err != nil {
			fmt.Println("failed")
			return
		}
		defer result.resp.Body.Close()
		data, _ := ioutil.ReadAll(result.resp.Body)
		fmt.Println(string(data))
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	doCall(ctx)
}
