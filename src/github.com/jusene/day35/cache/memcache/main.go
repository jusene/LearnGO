package main

import (
	"context"
	"github.com/beego/beego/v2/client/cache"
	"github.com/beego/beego/v2/core/logs"
	"time"

	_ "github.com/beego/beego/v2/client/cache/memcache"
)

func main() {
	bm, err := cache.NewCache("memcache", `{"conn": "127.0.0.1:11211"}`)
	if err != nil {
		logs.Error(err)
	}

	isPut := bm.Put(context.Background(), "name", "beego", time.Second*10)
	logs.Info(isPut)
}
