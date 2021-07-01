package main

import (
	"context"
	"github.com/beego/beego/v2/client/cache"
	"github.com/beego/beego/v2/core/logs"
	"time"
)

func main() {
	// create memory
	bm, err := cache.NewCache("memory", `{"interval": 60}`)
	if err != nil {
		logs.Error(err)
	}

	// put
	isPut := bm.Put(context.Background(), "name", "beego", time.Second*10)
	logs.Info(isPut)

	// get
	isExists, _ := bm.IsExist(context.Background(), "name")
	logs.Info(isExists)

	result, _ := bm.Get(context.Background(), "name")
	logs.Info(result)
}
