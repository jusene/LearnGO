package main

import (
	"context"
	"github.com/beego/beego/v2/client/cache"
	_ "github.com/beego/beego/v2/client/cache/redis"
	"github.com/beego/beego/v2/core/logs"
	"time"
)

func main() {
	bm, err := cache.NewCache("redis", `{"conn":"127.0.0.1:6379", "dbNum":"0"}`)
	if err != nil {
		logs.Error(err)
	}

	// put
	isPut := bm.Put(context.Background(), "hello", 1, time.Second*100)
	logs.Info(isPut)

	// get
	result, _ := bm.Get(context.Background(), "hello")
	logs.Info(string(result.([]byte)))
}
