package main

import (
	"context"
	"github.com/beego/beego/v2/client/cache"
	"github.com/beego/beego/v2/core/logs"
	"time"
)

func main() {
	cacheConfig := `{"CachePath":"./cache","FileSuffix":".cache", "EmbedExpiry": "120"}`
	bm, err := cache.NewCache("file", cacheConfig)
	if err != nil {
		logs.Error(err)
	}

	// put
	isPut := bm.Put(context.Background(), "name", "beego", time.Second*10)
	logs.Info(isPut)

	isPut = bm.Put(context.Background(), "hello", "golang", time.Second*10)
	logs.Info(isPut)

	isPut = bm.Put(context.Background(), "num", 1, time.Second*10)
	logs.Info(isPut)

	// get
	result, _ := bm.Get(context.Background(), "name")
	logs.Info(result)

	multiResult, _ := bm.GetMulti(context.Background(), []string{"name", "hello"})
	for _, v := range multiResult {
		logs.Info(v)
	}

	// isExists
	isExists, _ := bm.IsExist(context.Background(), "name")
	logs.Info(isExists)

	// delete
	isDelete := bm.Delete(context.Background(), "hello")
	logs.Info(isDelete)

	// incr decr
	err = bm.Incr(context.Background(), "num")
	logs.Info(err)

	err = bm.Decr(context.Background(), "num")
	logs.Info(err)

	result, _ = bm.Get(context.Background(), "num")
	logs.Info(result)

	bm.ClearAll(context.Background())

}
